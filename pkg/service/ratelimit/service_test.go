package ratelimit

import (
	"context"
	"fmt"
	"runtime"
	"strconv"
	"testing"

	ratelimitExt "github.com/envoyproxy/go-control-plane/envoy/extensions/common/ratelimit/v3"
	ratelimit "github.com/envoyproxy/go-control-plane/envoy/service/ratelimit/v3"
	config "github.com/in4it/roxprox-ratelimit/pkg/config/ratelimit"
)

func TestPrintKeyWithoutAuthorization(t *testing.T) {
	str := "ingress:remote_address:127.0.0.1,header_authorization:bearer secret,destination_cluster:simple-reverse-proxy:20200805T12:26"
	expected := "ingress:remote_address:127.0.0.1,header_authorization:***,destination_cluster:simple-reverse-proxy:20200805T12:26"
	res := printKeyWithoutAuthorization([]byte(str))
	if res != expected {
		t.Errorf("unexecpted output: %s", res)
	}
}
func TestPrintKeyWithoutAuthorization2(t *testing.T) {
	str := "ingress:xyz"
	expected := "ingress:xyz"
	res := printKeyWithoutAuthorization([]byte(str))
	if res != expected {
		t.Errorf("unexecpted output: %s", res)
	}
}
func TestRateLimit(t *testing.T) {
	rules := []config.RateLimitRule{
		{
			Name:           "test",
			Unit:           "Hour",
			RequestPerUnit: "2",
		},
	}
	r := newRateLimitService(cacheMbSizeDefault, rules)
	ctx := context.Background()
	req := &ratelimit.RateLimitRequest{
		Domain: "testdomain",
		Descriptors: []*ratelimitExt.RateLimitDescriptor{
			{
				Entries: []*ratelimitExt.RateLimitDescriptor_Entry{
					{
						Key:   "key",
						Value: "value",
					},
					{
						Key:   "generic_key",
						Value: "__identifier:test",
					},
				},
			},
		},
	}
	var res []*ratelimit.RateLimitResponse
	for i := 0; i < 4; i++ {
		r, err := r.ShouldRateLimit(ctx, req)
		if err != nil {
			t.Errorf("error during shouldRateLimit: %s", err)
			return
		}
		res = append(res, r)
	}
	if res[0].OverallCode != ratelimit.RateLimitResponse_OK {
		t.Errorf("error during shouldRateLimit: first response code was not OK")
		return
	}
	if res[1].OverallCode != ratelimit.RateLimitResponse_OK {
		t.Errorf("error during shouldRateLimit: first response code was not OK")
		return
	}
	if res[2].OverallCode != ratelimit.RateLimitResponse_OVER_LIMIT {
		t.Errorf("error during shouldRateLimit: second response code was not OVER_LIMIT")
		return
	}
	if res[3].OverallCode != ratelimit.RateLimitResponse_OVER_LIMIT {
		t.Errorf("error during shouldRateLimit: second response code was not OVER_LIMIT")
		return
	}
}
func TestRateLimitPerf(t *testing.T) {
	testLength := 1 // at 1000000: Alloc = 526 MiB  TotalAlloc = 43014 MiB  Sys = 601 MiB   NumGC = 1008
	initialValue := ""
	initLoggers()
	seed := `this is a string to test caching functionality`
	for i := 0; i < 100; i++ {
		initialValue += seed
	}
	rules := []config.RateLimitRule{
		{
			Name:           "test",
			Unit:           "Hour",
			RequestPerUnit: "100",
		},
	}
	r := newRateLimitService(cacheMbSizeDefault, rules)
	ctx := context.Background()
	req := &ratelimit.RateLimitRequest{
		Domain: "testdomain",
		Descriptors: []*ratelimitExt.RateLimitDescriptor{
			{
				Entries: []*ratelimitExt.RateLimitDescriptor_Entry{
					{
						Key:   "key",
						Value: initialValue,
					},
					{
						Key:   "generic_key",
						Value: "__identifier:test",
					},
				},
			},
		},
	}
	_, descriptorKey := extractDescriptorsToString(req.Descriptors)
	fmt.Printf("Starting test. Key length: %d\n", len([]byte(req.Domain+":"+descriptorKey)))
	for i := 0; i < testLength; i++ {
		req.Descriptors[0].Entries[0].Value = initialValue + strconv.Itoa(i)
		_, err := r.ShouldRateLimit(ctx, req)
		if err != nil {
			t.Errorf("got error: %s", err)
			return
		}
		if i%1000 == 0 {
			runtime.GC()
			PrintMemUsage()
		}
	}
	runtime.GC()
	PrintMemUsage()
}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}
func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
