package board_test

import (
	"clonecoding/internal/dto"
	"clonecoding/tests"
	user_test "clonecoding/tests/user"
	"net/http"
	"testing"
)

func TestCreateBoard(t *testing.T) {
	r := tests.SetupTestEnv()

	token := user_test.GetToken(t, r)

	board := dto.CreateBoardDTO{
		Name: "Keyboard",
	}

	res := CreateBoard(t, r, token, &board)

	if res.Code != http.StatusOK {
		t.Errorf("Eprected status 200, got %v", res.Code)
	}
}

func TestGetBoard(t *testing.T) {
	r := tests.SetupTestEnv()

	token := user_test.GetToken(t, r)

	res := GetBoard(t, r, token)

	if res.Code != http.StatusOK {
		t.Errorf("Eprected status 200, got %v", res.Code)
	}
}
