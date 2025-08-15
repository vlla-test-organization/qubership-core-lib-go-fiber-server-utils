package apiversionpoint

import (
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/vlla-test-organization/qubership-core-lib-go-actuator-common/v2/apiversion"
)

func TestEnableApiVersion(t *testing.T) {
	app := fiber.New()
	apiVersionService, err := apiversion.NewApiVersionService(apiversion.ApiVersionConfig{PathToApiVersionInfoFile: "../../../testdata/api-version-info.json"})
	assert.Nil(t, err)
	app.Get("/api-version", EnableApiVersion(apiVersionService))

	req := httptest.NewRequest("GET", "/api-version", nil)

	resp, err := app.Test(req)
	assert.Nil(t, err)
	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.True(t, strings.Contains(string(body), "bluegreen"))
	logger.Info(string(body))
	assert.Equal(t, 200, resp.StatusCode)
}

func TestEnableApiVersionError(t *testing.T) {
	app := fiber.New()
	apiVersionService, err := apiversion.NewApiVersionService(apiversion.ApiVersionConfig{PathToApiVersionInfoFile: "../api-version-info.json"})
	assert.Nil(t, err)
	app.Get("/api-version", EnableApiVersion(apiVersionService))

	req := httptest.NewRequest("GET", "/api-version", nil)

	resp, err := app.Test(req)
	assert.Nil(t, err)
	assert.Equal(t, 500, resp.StatusCode)
}
