package http

import (
	"clonecoding/internal/dto"
	"clonecoding/internal/usecase"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PostHandler struct {
	PostUsecase *usecase.PostUsecase
}

func (p *PostHandler) CreatePost(c *gin.Context) {
	var postDto dto.CreatePostDTO
	err := c.ShouldBindJSON(&postDto)
	if err != nil {
		Fail(c, http.StatusBadRequest, err.Error())
	}

	fmt.Println("postDto: ", postDto)

	id := c.Param("board_id")
	boardId, err := uuid.Parse(id)
	if err != nil {
		Fail(c, http.StatusBadRequest, err.Error())
		return
	}

	userIdOrig, exist := c.Get("userID")
	if !exist {
		Fail(c, http.StatusInternalServerError, err.Error())
		return
	}

	userId, ok := userIdOrig.(uuid.UUID)
	if !ok {
		Fail(c, http.StatusInternalServerError, err.Error())
		return
	}

	user, err := p.PostUsecase.CreatePost(&postDto, userId, boardId)
	if err != nil {
		Fail(c, http.StatusInternalServerError, err.Error())
		return
	}

	Success(c, user)
}

func (p *PostHandler) GetAllPost(c *gin.Context) {
	boards, err := p.PostUsecase.FindAllPost()
	if err != nil {
		Fail(c, http.StatusNotFound, err.Error())
		return
	}

	Success(c, boards)
}

func (p *PostHandler) GetPostById(c *gin.Context) {
	id := c.Param("post_id")
	postId, err := uuid.Parse(id)
	if err != nil {
		Fail(c, http.StatusBadRequest, err.Error())
		return
	}

	board, err := p.PostUsecase.FindPostById(postId)
	if err != nil {
		Fail(c, http.StatusNotFound, err.Error())
		return
	}

	Success(c, board)
}

func (p *PostHandler) GetPostByBoardId(c *gin.Context) {
	id := c.Param("board_id")
	boardId, err := uuid.Parse(id)
	if err != nil {
		Fail(c, http.StatusBadRequest, err.Error())
		return
	}

	board, err := p.PostUsecase.FindPostByBoardId(boardId)
	if err != nil {
		Fail(c, http.StatusNotFound, err.Error())
		return
	}

	Success(c, board)
}

func (p *PostHandler) GetPostByUserId(c *gin.Context) {
	id := c.Param("user_id")
	userId, err := uuid.Parse(id)
	if err != nil {
		Fail(c, http.StatusBadRequest, err.Error())
		return
	}

	board, err := p.PostUsecase.FindPostByUserId(userId)
	if err != nil {
		Fail(c, http.StatusNotFound, err.Error())
		return
	}

	Success(c, board)
}
