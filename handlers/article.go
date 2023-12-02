package handlers

import (
	articlesdto "Restful_Go/dto/articles"
	dto "Restful_Go/dto/results"
	"Restful_Go/models"
	"Restful_Go/repositories"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type articleHandler struct {
	ArticleRepository repositories.ArticleRepository
}

func ArticleHandler(articleRepository repositories.ArticleRepository) *articleHandler {
	return &articleHandler{articleRepository}
}

func (h *articleHandler) FindArticles(c echo.Context) error {
	articles, err := h.ArticleRepository.FindArticles()

	if err != nil {

		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: articles})
}

func (h *articleHandler) GetArticle(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	article, err := h.ArticleRepository.GetArticle(id)
	if err != nil {

		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: convertArticleResponse(article)})
}

func (h *articleHandler) CreateArticle(c echo.Context) error {
	dataFile := c.Get("dataFile").(string)

	request := articlesdto.CreateArticleRequest{
		ArticleName: c.FormValue("articleName"),
		Description: c.FormValue("description"),
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	userLogin := c.Get("userLogin")
	UserId := userLogin.(jwt.MapClaims)["id"].(float64)

	data := models.Article{
		ArticleName: request.ArticleName,
		Description: request.Description,
		Image:       dataFile,
		UserID:      int(UserId),
		CreatedAt:   time.Now(),
	}

	response, err := h.ArticleRepository.CreateArticle(data)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: response})
}

func (h *articleHandler) UpdateArticle(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	article, err := h.ArticleRepository.GetArticle(id)

	dataFile := c.Get("dataFile").(string)

	userLogin := c.Get("userLogin")
	UserId := userLogin.(jwt.MapClaims)["id"].(float64)

	request := articlesdto.UpdateArticleRequest{
		ArticleName: c.FormValue("articleName"),
		Description: c.FormValue("description"),
		UserID:      int(UserId),
	}

	if request.ArticleName != "" {
		article.ArticleName = request.ArticleName
	}

	if request.Description != "" {
		article.Description = request.Description
	}

	if dataFile != "" {
		article.Image = dataFile
	}

	if request.UserID != 0 {
		article.UserID = int(UserId)
	}

	article.UpdatedAt = time.Now()

	response, err := h.ArticleRepository.UpdateArticle(article)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: response})
}

func (h *articleHandler) DeleteArticle(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	article, err := h.ArticleRepository.GetArticle(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	response, err := h.ArticleRepository.DeleteArticle(article, id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: response,
	})
}

func convertArticleResponse(u models.Article) articlesdto.ArticleResponse {
	return articlesdto.ArticleResponse{
		ID:          u.ID,
		ArticleName: u.ArticleName,
		Description: u.Description,
		Image:       u.Image,
		UserID:      u.UserID,
	}
}
