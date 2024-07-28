package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Authorization: ApiKey ...
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no auth key passed into headers")
	}

	values := strings.Split(val, " ")
	if len(values) != 2 {
		return "", errors.New("invalid auth key")
	}

	if values[0] != "ApiKey" {
		return "", errors.New("missing auth header")
	}

	return values[1], nil
}
