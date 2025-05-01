package board_test

import (
	"bytes"
	"clonecoding/internal/dto"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func CreateBoard(t *testing.T, r *gin.Engine, token string, boardDto *dto.CreateBoardDTO) *httptest.ResponseRecorder {
	jsonValue, _ := json.Marshal(boardDto)

	req, _ := http.NewRequest("POST", "/board/", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	return res
}

func GetBoard(t *testing.T, r *gin.Engine, token string) *httptest.ResponseRecorder {
	board := dto.CreateBoardDTO{
		Name: "Keyboard",
	}
	CreateBoard(t, r, token, &board)

	req, _ := http.NewRequest("GET", "/board/", nil)
	req.Header.Set("Authorization", token)

	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	return res
}
