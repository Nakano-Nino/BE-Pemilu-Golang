package handlers

import (
	partaisdto "Restful_Go/dto/partais"
	dto "Restful_Go/dto/results"
	"Restful_Go/models"
	"Restful_Go/repositories"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type partaiHandler struct {
	PartaiRepository repositories.PartaiRepository
}

func PartaiHandler(partaiRepository repositories.PartaiRepository) *partaiHandler {
	return &partaiHandler{partaiRepository}
}

func (h *partaiHandler) FindPartais(c echo.Context) error {
	partai, err := h.PartaiRepository.FindPartais()

	if err != nil {

		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: partai})
}

func (h *partaiHandler) GetPartai(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	partai, err := h.PartaiRepository.GetPartai(id)
	if err != nil {

		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: convertPartaiResponse(partai)})
}

func (h *partaiHandler) CreatePartai(c echo.Context) error {
	dataFile := c.Get("dataFile").(string)

	request := partaisdto.CreatePartaiRequest{
		Name:           c.FormValue("name"),
		Ketum:          c.FormValue("ketum"),
		VissionMission: c.FormValue("vissionMission"),
		Address:        c.FormValue("address"),
		PaslonID:       c.FormValue("paslonId"),
	}

	paslonId, err := strconv.Atoi(request.PaslonID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}

	data := models.Partai{
		Name:           request.Name,
		Ketum:          request.Ketum,
		VissionMission: request.VissionMission,
		Address:        request.Address,
		Image:          dataFile,
		PaslonID:       paslonId,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	response, err := h.PartaiRepository.CreatePartai(data)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: response})
}

func (h *partaiHandler) UpdatePartai(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	partai, err := h.PartaiRepository.GetPartai(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	dataFile := c.Get("dataFile").(string)

	request := partaisdto.CreatePartaiRequest{
		Name:           c.FormValue("name"),
		Ketum:          c.FormValue("ketum"),
		VissionMission: c.FormValue("vissionMission"),
		Address:        c.FormValue("address"),
		PaslonID:       c.FormValue("paslonId"),
	}

	paslonId, err := strconv.Atoi(request.PaslonID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}

	if request.Name != "" {
		partai.Name = request.Name
	}

	if request.Ketum != "" {
		partai.Ketum = request.Ketum
	}

	if request.VissionMission != "" {
		partai.VissionMission = request.VissionMission
	}

	if request.Address != "" {
		partai.Address = request.Address
	}

	if dataFile != "" {
		partai.Image = dataFile
	}

	if paslonId != 0 {
		partai.PaslonID = paslonId
	}

	partai.UpdatedAt = time.Now()

	response, err := h.PartaiRepository.UpdatePartai(partai)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: response})
}

func (h *partaiHandler) DeletePartai(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	partai, err := h.PartaiRepository.GetPartai(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	response, err := h.PartaiRepository.DeletePartai(partai, id)

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

func convertPartaiResponse(u models.Partai) partaisdto.PartaiResponse {
	return partaisdto.PartaiResponse{
		ID:             u.ID,
		Name:           u.Name,
		Ketum:          u.Ketum,
		VissionMission: u.VissionMission,
		Address:        u.Address,
	}
}
