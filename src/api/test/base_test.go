package test

import (
	"os"
	"testing"

	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/wagnergaldino/unit-integration-and-functional-testing-in-golang/src/api/app"
)

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	go app.StartApp()
	os.Exit(m.Run())
}
