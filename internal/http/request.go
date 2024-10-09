package http

import (
	"encoding/json"
	"io"
)

const (
	MethodGet     = "GET"
	MethodHead    = "HEAD"
	MethodPost    = "POST"
	MethodPut     = "PUT"
	MethodPatch   = "PATCH"
	MethodDelete  = "DELETE"
	MethodConnect = "CONNECT"
	MethodOptions = "OPTIONS"
	MethodTrace   = "TRACE"
)

func ParseRequestParams(keys []string, values []string) map[string]string {
	params := make(map[string]string)

	for i := 0; i < len(keys); i++ {
		params[keys[i]] = values[i]
	}
	return params
}

func ParseRequestBody(reader io.ReadCloser) map[string]string {
	body := make(map[string]string)

	result, err := io.ReadAll(reader)
	if err != nil {
		return nil
	}

	err = json.Unmarshal(result, &body)
	if err != nil {
		return nil
	}
	return body
}
