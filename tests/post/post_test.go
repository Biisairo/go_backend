package post_test

import (
	"clonecoding/internal/dto"
	"clonecoding/tests"
	board_test "clonecoding/tests/board"
	user_test "clonecoding/tests/user"
	"net/http"
	"testing"
)

func TestCreatePost(t *testing.T) {
	r := tests.SetupTestEnv()

	token := user_test.GetToken(t, r)

	boardRes := board_test.GetBoard(t, r, token)

	parse := tests.ParseResponse(t, boardRes)
	data, _ := parse.Data.([]any)
	board, _ := data[0].(map[string]any)
	id, _ := board["ID"].(string)

	post := dto.CreatePostDTO{
		Title: "Ferris Sweep",
		Body:  "Ferris Sweep is 34 key split keyboard",
	}

	res := CreatePost(t, r, token, &post, id)

	if res.Code != http.StatusOK {
		t.Errorf("Eprected status 200, got %v", res.Code)
	}
}
