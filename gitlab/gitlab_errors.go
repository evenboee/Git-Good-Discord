package gitlab

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

// tryToParseErrorResponse will try to parse the error response given in a Gitlab
// response. If it fails, it will return an ErrorResponse with the message set to
// "unknown error response given by gitlab"
func tryToParseErrorResponse(responseBody io.Reader) ErrorResponse {
	body, err := ioutil.ReadAll(responseBody)

	// Initialize response error for when parsing fails
	// This will be overridden if parsing succeeds
	errorResponse := ErrorResponse{ Message: "unknown error response given by gitlab" }

	if err != nil {
		return errorResponse
	}

	// Try to unmarshal the actual error message given in gitlab response
	err = json.Unmarshal(body, &errorResponse)

	if err != nil {
		return errorResponse
	}

	return errorResponse
}