package tests

import (
	"clonecoding/internal/dto"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateUser(t *testing.T) {
	r := SetupTestEnv()

	users := GetTestUser()

	for _, user := range users {
		t.Run("Create"+user.Name, func(t *testing.T) {
			res := CreateUser(r, &user)

			if res.Code != http.StatusOK {
				t.Errorf("Eprected status 200, got %v", res.Code)
			}
		})
	}
}

func TestLoginCases(t *testing.T) {
	r := SetupTestEnv()

	users := GetTestUser()

	for _, user := range users {
		CreateUser(r, &user)
	}

	tests := []struct {
		name         string
		email        string
		password     string
		expectStatus int
	}{
		{"Success Login", users[0].Email, users[0].Password, http.StatusOK},
		{"Wrong Password", "qwe@asd.zxc", "wrongpassword", http.StatusInternalServerError},
		{"Missing Email", "", "wasd", http.StatusBadRequest},
		{"Missing passord", "qwe@asd.zxc", "", http.StatusBadRequest},
		{"User Not Found", "ghost@example.com", "whatever", http.StatusInternalServerError},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			loginReq := dto.LoginRequest{
				Email:    tc.email,
				Password: tc.password,
			}

			res := Login(r, &loginReq)

			if res.Code != tc.expectStatus {
				t.Errorf("[%s] Expected %d, got %d", tc.name, tc.expectStatus, res.Code)
			}

			// parsed := ParseResponse(t, res)
		})
	}
}

func TestGetUserWithValidToken(t *testing.T) {
	r := SetupTestEnv()

	users := GetTestUser()
	CreateUser(r, &users[0])
	loginReq := dto.LoginRequest{
		Email:    users[0].Email,
		Password: users[0].Password,
	}
	loginRes := Login(r, &loginReq)
	loginResParse := ParseResponse(t, loginRes)
	dataOrig := loginResParse.Data
	data, _ := dataOrig.(map[string]any)
	tokenOrig := data["access_token"]
	token, _ := tokenOrig.(string)
	fmt.Println(token)

	token = "Bearer " + token

	req, _ := http.NewRequest("GET", "/user/", nil)
	req.Header.Set("Authorization", token)

	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", res.Code)
	}
}
