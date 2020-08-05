package ratelimit

type RateLimitRule struct {
	Name           string
	RequestPerUnit string
	Unit           string
}
