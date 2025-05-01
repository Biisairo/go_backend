package tests

import (
	"clonecoding/internal/adapter/http"
	"clonecoding/internal/bootstrap"
	"clonecoding/internal/config"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func SetupTestEnv() *gin.Engine {
	gin.SetMode(gin.TestMode)

	config.LoadConfig("../.env")

	app := bootstrap.InitApp()

	return app.Engine
}

func ParseResponse(t *testing.T, res *httptest.ResponseRecorder) http.APIResponse {
	var body http.APIResponse
	err := json.Unmarshal(res.Body.Bytes(), &body)
	if err != nil {
		t.Fatalf("Invalid JSON: %v", err)
	}
	return body
}
