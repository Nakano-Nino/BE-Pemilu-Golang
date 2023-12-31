package handlers

import (
	dto "Restful_Go/dto/results"
	usersdto "Restful_Go/dto/users"
	"Restful_Go/models"
	"Restful_Go/repositories"
	"net/http"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	UserRepository repositories.UserRepository
}

func UserHandler(userRepository repositories.UserRepository) *userHandler {
	return &userHandler{userRepository}
}

func (h *userHandler) FindUsers(c echo.Context) error {
	users, err := h.UserRepository.FindUsers()

	if err != nil {

		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: users})
}

func (h *userHandler) GetUsers(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := h.UserRepository.Getuser(id)
	if err != nil {

		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: convertResponse(user)})
}

func (h *userHandler) CreateUser(c echo.Context) error {
	request := new(usersdto.CreateUserRequest)
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)

	data := models.User{
		Name:      request.Name,
		Address:   request.Address,
		Gender:    request.Gender,
		Username:  request.Username,
		Password:  string(hashedPassword),
		Role:      request.Role,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	response, err := h.UserRepository.CreateUser(data)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: response})
}

func convertResponse(u models.User) usersdto.UserResponse {
	return usersdto.UserResponse{
		ID:   u.ID,
		Name: u.Name,
	}
}
