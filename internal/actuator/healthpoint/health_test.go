package healthpoint

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/vlla-test-organization/qubership-core-lib-go-actuator-common/v2/health"
)

const expectedHealthResponse = `{"status":"UP"}`

func TestEnableHealth(t *testing.T) {
	app := fiber.New()
	healthService, err := health.NewHealthService()
	assert.Nil(t, err)
	app.Get("/health", EnableHealth(healthService.RunChecksOnStartup(true)))

	req := httptest.NewRequest("GET", "/health", nil)

	resp, err := app.Test(req)
	assert.Nil(t, err)
	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Equal(t, expectedHealthResponse, string(body))
	assert.Equal(t, 200, resp.StatusCode)
}
