package ratelimit

import (
	"context"
	"encoding/binary"
	"fmt"
	"log"
	"os"
	"runtime/debug"
	"strconv"
	"strings"
	"time"

	"github.com/coocood/freecache"
	ratelimitExt "github.com/envoyproxy/go-control-plane/envoy/extensions/common/ratelimit/v3"
	ratelimit "github.com/envoyproxy/go-control-plane/envoy/service/ratelimit/v3"
	config "github.com/in4it/roxprox-ratelimit/pkg/config/ratelimit"
)

const cacheMbSizeDefault = 512
const expiry = 60
const limit = 5

var (
	DebugLogger   *log.Logger
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func initLoggers() {
	DebugLogger = log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	InfoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func newRateLimitService(cacheMbSize int, rules []config.RateLimitRule) *Service {
	initLoggers()

	rulesIndex := make(map[string]int)
	for k, rule := range rules {
		rulesIndex[rule.Name] = k
	}

	r := &Service{
		rules:      rules,
		rulesIndex: rulesIndex,
	}

	cacheSize := cacheMbSize * 1024 * 1024
	r.cache = freecache.NewCache(cacheSize)
	debug.SetGCPercent(20)
	r.startvalue = make([]byte, 8)
	binary.LittleEndian.PutUint64(r.startvalue, 1)
	return r
}

//Service is a Rate Limit Service implementing ShouldRateLimit
type Service struct {
	cache      *freecache.Cache
	rules      []config.RateLimitRule
	rulesIndex map[string]int
	startvalue []byte
}

//ShouldRateLimit is triggered for every request. This function determines whether to rate limit the request or not
func (r *Service) ShouldRateLimit(ctx context.Context, req *ratelimit.RateLimitRequest) (*ratelimit.RateLimitResponse, error) {
	debugLogger(fmt.Sprintf("Req: %+v", req))
	identifier, descriptorKey := extractDescriptorsToString(req.Descriptors)
	if _, ok := r.rulesIndex[identifier]; !ok {
		WarningLogger.Printf("Identifier not found in rules: %s (descriptorKey: %s)", identifier, descriptorKey)
	}

	bucket, err := getBucket(r.rules[r.rulesIndex[identifier]].Unit)
	if err != nil {
		return handleError(err)
	}

	key := []byte(req.Domain + ":" + descriptorKey + ":" + bucket)

	curValue, err := r.cache.GetOrSet(key, r.startvalue, expiry)
	if err != nil {
		return handleError(err)
	}
	if curValue == nil {
		// new value, returning OK
		debugLogger(fmt.Sprintf("Key: %s (length: %d), Value: %d", string(key), len(key), binary.LittleEndian.Uint64(r.startvalue)))
		return &ratelimit.RateLimitResponse{
			OverallCode: ratelimit.RateLimitResponse_OK,
		}, nil

	}

	curValueInt64 := binary.LittleEndian.Uint64(curValue)

	requestsPerUnit, err := strconv.ParseInt(r.rules[r.rulesIndex[identifier]].RequestPerUnit, 10, 64)

	if curValueInt64+1 > uint64(requestsPerUnit) {
		InfoLogger.Printf("Rate limited by %s: %s\n", identifier, printKeyWithoutAuthorization(key))
		return &ratelimit.RateLimitResponse{
			OverallCode: ratelimit.RateLimitResponse_OVER_LIMIT,
		}, nil
	}

	if err = r.incrementValue(key, curValueInt64); err != nil {
		return handleError(err)
	}

	debugLogger(fmt.Sprintf("Key: %s (length: %d), Value: %d", string(key), len(key), curValueInt64+1))

	return &ratelimit.RateLimitResponse{
		OverallCode: ratelimit.RateLimitResponse_OK,
	}, nil
}

func (r *Service) incrementValue(key []byte, curValue uint64) error {
	newValue := make([]byte, 8)
	binary.LittleEndian.PutUint64(newValue, curValue+1)
	err := r.cache.Set(key, newValue, expiry)
	if err != nil {
		return err
	}
	return nil
}
func handleError(err error) (*ratelimit.RateLimitResponse, error) {
	ErrorLogger.Printf("%s", err)
	return &ratelimit.RateLimitResponse{
		OverallCode: ratelimit.RateLimitResponse_OK,
	}, err
}

func getBucket(unit string) (string, error) {
	t := time.Now()
	switch strings.ToLower(unit) {
	case "second":
		return t.Format("20060102T15:04:05"), nil
	case "minute":
		return t.Format("20060102T15:04"), nil
	case "hour":
		return t.Format("20060102T15"), nil
	case "day":
		return t.Format("20060102"), nil
	default:
		return "", fmt.Errorf("Unit type not found: %s", unit)
	}
}

func printKeyWithoutAuthorization(key []byte) string {
	strKey := string(key)
	i := strings.IndexRune(strKey, ':')
	if i == -1 {
		return ""
	}
	domain := strKey[0:i]
	if len(domain) == len(strKey) {
		return domain
	}
	descriptors := strKey[len(domain):]
	newString := domain
	for _, v := range strings.Split(descriptors, ",") {
		if strings.HasPrefix(v, "header_authorization:") {
			newString += "header_authorization:***,"
		} else {
			newString += v + ","
		}
	}
	return strings.TrimRight(newString, ",")
}

func extractDescriptorsToString(descriptors []*ratelimitExt.RateLimitDescriptor) (string, string) {
	var res string
	var identifier string
	for _, descriptor := range descriptors {
		if len(descriptor.Entries) > 1 { // descriptors that only have the identifier can be ignored
			for _, v := range descriptor.Entries {
				if v.Key == "generic_key" && strings.HasPrefix(v.Value, "__identifier:") && len(v.Value) > len("__identifier:") {
					identifier = v.Value[len("__identifier:"):]
				} else {
					res += v.Key + ":" + v.Value + ","
				}
			}
		}
	}
	return identifier, strings.TrimSuffix(res, ",")
}

func debugLogger(str string) {
	if os.Getenv("DEBUG") != "" {
		DebugLogger.Println(str)
	}
}
