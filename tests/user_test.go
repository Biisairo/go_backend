package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateUser(t *testing.T) {
	r := SetupTestEnv()

	requestBody := map[string]string{
		"email":    "qwe@asd.zxc",
		"name":     "qaz",
		"password": "wasd",
	}

	jsonValue, _ := json.Marshal(requestBody)

	req, _ := http.NewRequest("POST", "/create", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Eprected status 200, got %v", res.Code)
	}
	fmt.Println(res)
}

func TestLoginCases(t *testing.T) {
	router := SetupTestEnv()

	tests := []struct {
		name         string
		email        string
		password     string
		expectStatus int
	}{
		{"Success Login", "qwe@asd.zxc", "wasd", http.StatusOK},
		{"Wrong Password", "qwe@asd.zxc", "wrongpassword", http.StatusInternalServerError},
		{"Missing Email", "", "wasd", http.StatusBadRequest},
		{"Missing passord", "qwe@asd.zxc", "", http.StatusBadRequest},
		{"User Not Found", "ghost@example.com", "whatever", http.StatusInternalServerError},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			body := map[string]string{
				"email":    tc.email,
				"password": tc.password,
			}
			jsonBody, _ := json.Marshal(body)

			req, _ := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(jsonBody))
			req.Header.Set("Content-Type", "application/json")
			res := httptest.NewRecorder()

			router.ServeHTTP(res, req)

			if res.Code != tc.expectStatus {
				t.Errorf("[%s] Expected %d, got %d", tc.name, tc.expectStatus, res.Code)
			}

			// parsed := ParseResponse(t, res)
		})
	}
}

func TestGetUserWithValidToken(t *testing.T) {
	r := SetupTestEnv()

	token := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiIzYzQzZTA3Mi02MGFlLTQzZWYtYTg3NS0xYjU0ZjlkZjcwNTciLCJleHAiOjE3NDYwMjAwMDEsImlhdCI6MTc0NTkzMzYwMX0.bE9ilQC2uNZTpi55TwEHvoHcPM38l6IF4KcXcuHH2X0"

	req, _ := http.NewRequest("GET", "/user", nil)
	req.Header.Set("Authorization", token)

	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", res.Code)
	}
}
