package ratelimit

import (
	"github.com/in4it/roxprox/pkg/api"
)

//Storage expects a ListObject() that returns api.Objects
type Storage interface {
	ListObjects() ([]api.Object, error)
}

func GetRateLimitRules(s Storage) ([]RateLimitRule, error) {
	var rateLimitRules []RateLimitRule

	objects, err := s.ListObjects()
	if err != nil {
		return rateLimitRules, err
	}

	for _, object := range objects {
		if object.Kind == "rateLimit" {
			rateLimitObject := object.Data.(api.RateLimit)
			rateLimitRules = append(rateLimitRules, RateLimitRule{
				Name:           rateLimitObject.Metadata.Name,
				RequestPerUnit: rateLimitObject.Spec.RequestPerUnit,
				Unit:           rateLimitObject.Spec.Unit,
			})
		}
	}
	return rateLimitRules, nil
}
