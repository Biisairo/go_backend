package tests

import (
	"clonecoding/internal/adapter/http"
	"encoding/json"
	"net/http/httptest"
	"testing"
)

func ParseResponse(t *testing.T, res *httptest.ResponseRecorder) http.APIResponse {
	var body http.APIResponse
	err := json.Unmarshal(res.Body.Bytes(), &body)
	if err != nil {
		t.Fatalf("Invalid JSON: %v", err)
	}
	return body
}
