package handlers

import (
	paslonsdto "Restful_Go/dto/paslons"
	dto "Restful_Go/dto/results"
	"Restful_Go/models"
	"Restful_Go/repositories"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type paslonHandler struct {
	PaslonRepository repositories.PaslonRepository
}

func PaslonHandler(paslonRepository repositories.PaslonRepository) *paslonHandler {
	return &paslonHandler{paslonRepository}
}

func (h *paslonHandler) FindPaslons(c echo.Context) error {
	paslons, err := h.PaslonRepository.FindPaslons()

	if err != nil {

		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: paslons})
}

func (h *paslonHandler) GetPaslon(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	paslon, err := h.PaslonRepository.GetPaslon(id)
	if err != nil {

		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: convertPaslonResponse(paslon)})
}

func (h *paslonHandler) CreatePaslon(c echo.Context) error {
	dataFile := c.Get("dataFile").(string)

	request := paslonsdto.CreatePaslonRequest{
		Name:           c.FormValue("name"),
		OrderNum:       c.FormValue("orderNum"),
		VissionMission: c.FormValue("vissionMission"),
	}

	orderNum, err := strconv.Atoi(request.OrderNum)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	data := models.Paslon{
		Name:           request.Name,
		OrderNum:       orderNum,
		VissionMission: request.VissionMission,
		Image:          dataFile,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	response, err := h.PaslonRepository.CreatePaslon(data)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: response})
}

func (h *paslonHandler) UpdatePaslon(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	paslon, err := h.PaslonRepository.GetPaslon(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	dataFile := c.Get("dataFile").(string)

	request := paslonsdto.CreatePaslonRequest{
		Name:           c.FormValue("name"),
		OrderNum:       c.FormValue("orderNum"),
		VissionMission: c.FormValue("vissionMission"),
	}

	orderNum, err := strconv.Atoi(request.OrderNum)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	if request.Name != "" {
		paslon.Name = request.Name
	}

	if orderNum != 0 {
		paslon.OrderNum = orderNum
	}

	if request.VissionMission != "" {
		paslon.VissionMission = request.VissionMission
	}

	if dataFile != "" {
		paslon.Image = dataFile
	}

	paslon.UpdatedAt = time.Now()

	response, err := h.PaslonRepository.UpdatePaslon(paslon)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: response})
}

func (h *paslonHandler) DeletePaslon(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	paslon, err := h.PaslonRepository.GetPaslon(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	response, err := h.PaslonRepository.DeletePaslon(paslon, id)

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

func convertPaslonResponse(u models.Paslon) paslonsdto.PaslonResponse {
	return paslonsdto.PaslonResponse{
		ID:             u.ID,
		Name:           u.Name,
		OrderNum:       u.OrderNum,
		VissionMission: u.VissionMission,
	}
}
