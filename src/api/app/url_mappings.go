package app

import (
	"github.com/wagnergaldino/unit-integration-and-functional-testing-in-golang/src/api/controllers"
)

func mapUrls() {
	router.GET("/locations/countries/:country_id", controllers.GetCountry)
}
