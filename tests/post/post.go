package post_test

import (
	"bytes"
	"clonecoding/internal/dto"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func CreatePost(t *testing.T, r *gin.Engine, token string, boardDto *dto.CreatePostDTO, boardId string) *httptest.ResponseRecorder {
	jsonValue, _ := json.Marshal(boardDto)

	req, _ := http.NewRequest("POST", "/board/post/"+boardId, bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	return res
}
