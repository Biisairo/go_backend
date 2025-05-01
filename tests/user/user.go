package user_test

import (
	"bytes"
	"clonecoding/internal/dto"
	"clonecoding/tests"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func GetTestUser() []dto.CreateUserDTO {
	var users []dto.CreateUserDTO
	users = append(users,
		dto.CreateUserDTO{
			Name:     "ABC",
			Email:    "a@b.c",
			Password: "abc",
		},
		dto.CreateUserDTO{
			Name:     "DEF",
			Email:    "d@e.f",
			Password: "def",
		},
		dto.CreateUserDTO{
			Name:     "GHI",
			Email:    "g@h.i",
			Password: "ghi",
		},
		dto.CreateUserDTO{
			Name:     "JKL",
			Email:    "j@k.l",
			Password: "jkl",
		},
		dto.CreateUserDTO{
			Name:     "MNO",
			Email:    "m@n.o",
			Password: "mno",
		},
	)

	return users
}

func CreateUser(r *gin.Engine, user *dto.CreateUserDTO) *httptest.ResponseRecorder {
	jsonValue, _ := json.Marshal(user)

	req, _ := http.NewRequest("POST", "/create", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	return res
}

func Login(r *gin.Engine, loginReq *dto.LoginRequest) *httptest.ResponseRecorder {
	jsonBody, _ := json.Marshal(loginReq)

	req, _ := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	return res
}

func GetToken(t *testing.T, r *gin.Engine) string {
	users := GetTestUser()
	CreateUser(r, &users[0])
	loginReq := dto.LoginRequest{
		Email:    users[0].Email,
		Password: users[0].Password,
	}
	loginRes := Login(r, &loginReq)
	loginResParse := tests.ParseResponse(t, loginRes)
	dataOrig := loginResParse.Data
	data, _ := dataOrig.(map[string]any)
	tokenOrig := data["access_token"]
	token, _ := tokenOrig.(string)
	fmt.Println(token)

	token = "Bearer " + token

	return token
}
