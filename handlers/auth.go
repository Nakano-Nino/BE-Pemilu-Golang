package handlers

import (
	authsdto "Restful_Go/dto/auth"
	dto "Restful_Go/dto/results"
	"Restful_Go/models"
	jwToken "Restful_Go/pkg/jwt"
	"Restful_Go/repositories"
	"log"
	"net/http"
	"time"

	"Restful_Go/pkg/bcrypt"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type authHandler struct {
	AuthRepository repositories.AuthRepository
}

func AuthHandler(authRepository repositories.AuthRepository) *authHandler {
	return &authHandler{authRepository}
}

func (h *authHandler) Register(c echo.Context) error {
	request := new(authsdto.AuthRequest)
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	password, err := bcrypt.HashingPassword(request.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	data := models.User{
		Name:      request.Name,
		Address:   request.Address,
		Gender:    request.Gender,
		Username:  request.Username,
		Password:  password,
		Role:      "user",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	response, err := h.AuthRepository.Register(data)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: response})
}

func (h *authHandler) Login(c echo.Context) error {
	request := new(authsdto.LoginRequest)
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	user := models.User{
		Username: request.Username,
		Password: request.Password,
	}

	user, err := h.AuthRepository.Login(user.Username)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	isValid := bcrypt.CheckHashedPassword(request.Password, user.Password)
	if !isValid {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "wrong email or password",
		})
	}

	claims := jwt.MapClaims{} // inisialisasi
	claims["id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix() // 2 hours expired

	token, errGenerateToken := jwToken.GenerateToken(&claims)
	if errGenerateToken != nil {
		log.Println(errGenerateToken)
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	loginResponse := authsdto.LoginResponse{
		Name:     user.Name,
		Username: user.Username,
		Password: user.Password,
		Role:     user.Role,
		Token:    token,
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: loginResponse,
	})
}
