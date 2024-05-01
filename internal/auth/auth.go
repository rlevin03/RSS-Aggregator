package auth

import (
	"errors"
	"net/http"
	"strings"
)

// getAPIKey extracts the API key from the headers of an HTTP request.
// Example:
// Authorization: ApiKey {insert API key here}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no authorization info found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed authorization header")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("malformed first part of authorization header")
	}
	return vals[1], nil
}
