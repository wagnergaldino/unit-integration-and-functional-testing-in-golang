package services

import (
	"github.com/wagnergaldino/unit-integration-and-functional-testing-in-golang/src/api/domain/locations"
	"github.com/wagnergaldino/unit-integration-and-functional-testing-in-golang/src/api/providers/locations_provider"
	"github.com/wagnergaldino/unit-integration-and-functional-testing-in-golang/src/api/utils/errors"
)

type locationsService struct{}

type locationsServiceInterface interface {
	GetCountry(countryId string) (*locations.Country, *errors.ApiError)
}

var (
	LocationsService locationsServiceInterface
)

func init() {
	LocationsService = &locationsService{}
}

func (s *locationsService) GetCountry(countryId string) (*locations.Country, *errors.ApiError) {
	return locations_provider.GetCountry(countryId)
}
