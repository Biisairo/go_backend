package tests

import (
	"clonecoding/internal/adapter/ginengine"
	"clonecoding/internal/bootstrap"
	"clonecoding/internal/config"
	"encoding/json"
	"fmt"
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

func ParseResponse(t *testing.T, res *httptest.ResponseRecorder) ginengine.APIResponse {
	var body ginengine.APIResponse
	err := json.Unmarshal(res.Body.Bytes(), &body)
	if err != nil {
		t.Fatalf("Invalid JSON: %v", err)
	}
	return body
}

func CkeckRedirectionLocation(res *httptest.ResponseRecorder) {
	location := res.Header().Get("Location")
	fmt.Printf("Redirected to: %s\n", location)
}
