package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wagnergaldino/unit-integration-and-functional-testing-in-golang/src/api/services"
)

func GetCountry(c *gin.Context) {
	country, err := services.LocationsService.GetCountry(c.Param("country_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, country)
}
