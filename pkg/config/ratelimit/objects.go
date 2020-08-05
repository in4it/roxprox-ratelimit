package ratelimit

import (
	"github.com/in4it/roxprox/pkg/api"
)

//Storage expects a ListObject() that returns api.Objects
type Storage interface {
	ListObjects() ([]api.Object, error)
	GetObject(name string) ([]api.Object, error)
}

func GetRateLimitRules(s Storage) ([]RateLimitRule, error) {
	var rateLimitRules []RateLimitRule

	objects, err := s.ListObjects()
	if err != nil {
		return rateLimitRules, err
	}

	for _, object := range objects {
		if object.Kind == "rateLimit" {
			rateLimitRules = append(rateLimitRules, getRateLimitObject(object.Data.(api.RateLimit)))
		}
	}
	return rateLimitRules, nil
}
func GetRateLimitRule(s Storage, filename string) ([]RateLimitRule, error) {
	var rateLimitRules []RateLimitRule

	objects, err := s.GetObject(filename)
	if err != nil {
		return rateLimitRules, err
	}
	for _, object := range objects {
		if object.Kind == "rateLimit" {
			rateLimitRules = append(rateLimitRules, getRateLimitObject(object.Data.(api.RateLimit)))
		}
	}
	return rateLimitRules, nil
}

func getRateLimitObject(object api.RateLimit) RateLimitRule {
	return RateLimitRule{
		Name:           object.Metadata.Name,
		RequestPerUnit: object.Spec.RequestPerUnit,
		Unit:           object.Spec.Unit,
	}
}
