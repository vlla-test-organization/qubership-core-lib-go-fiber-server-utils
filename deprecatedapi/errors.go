package deprecatedapi

import (
	"fmt"

	"github.com/vlla-test-organization/qubership-core-lib-go-error-handling/v3/errors"
	"github.com/vlla-test-organization/qubership-core-lib-go-error-handling/v3/tmf"
)

var errorCode = errors.ErrorCode{
	Code:  "NC-COMMON-2101",
	Title: "Request is declined with 404 Not Found, because deprecated REST API is disabled",
}

func CreateErrorResponse(requestMethod string, requestUri string, matchingMethod []string, matchingUri string, status int) *tmf.Response {
	err := errors.NewError(errorCode,
		fmt.Sprintf("Request [%s] '%s' is declined with 404 Not Found, because the following deprecated REST API is disabled: [%s] %s",
			requestMethod, requestUri, matchingMethod, matchingUri), nil)
	logger.Warn(err.GetDetail())
	response := tmf.ErrToResponse(err, status)
	return &response
}
