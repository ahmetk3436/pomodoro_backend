package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/ahmetk3436/pkg/model"
	"github.com/ahmetk3436/pkg/service"
	jsoniter "github.com/json-iterator/go"
	"github.com/labstack/echo/v4"
)

// PostAPI ...
type PostAPI struct {
	PostService service.PostService
}

// NewPostAPI ...
func NewPostAPI(p service.PostService) PostAPI {
	return PostAPI{PostService: p}
}

func (p PostAPI) CreatePost(c echo.Context) error {
	var postDTO model.ExamDateDTO
	if err := jsoniter.NewDecoder(c.Request().Body).Decode(&postDTO); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	createdPost, err := p.PostService.Save(model.ToExamDate(&postDTO))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, model.ToExamDateDTO(createdPost))
}

// FindByID ...
func (p PostAPI) FindByID(c echo.Context) error {
	// Check if id is integer
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Find login by id from db
	post, err := p.PostService.FindByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	dto := model.ToExamDateDTO(post)
	return c.JSON(http.StatusOK, dto)
}

// Migrate ...
func (p PostAPI) Migrate() {
	err := p.PostService.Migrate()
	if err != nil {
		log.Println(err)
	}
}
