package handlers

import (
	dto "Restful_Go/dto/results"
	votersdto "Restful_Go/dto/voters"
	"Restful_Go/models"
	"Restful_Go/repositories"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type voterHandler struct {
	VoterRepository repositories.VoterRepository
}

func VoterHandler(voterRepository repositories.VoterRepository) *voterHandler {
	return &voterHandler{voterRepository}
}

func (h *voterHandler) FindVoters(c echo.Context) error {
	voters, err := h.VoterRepository.FindVoters()

	votersLength := len(voters)

	response := models.VotersResponse{
		Voters:      voters,
		VotersCount: votersLength,
	}

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

func (h *voterHandler) Vote(c echo.Context) error {
	request := new(votersdto.VoteRequest)
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	userLogin := c.Get("userLogin")
	UserId := userLogin.(jwt.MapClaims)["id"].(float64)

	voters, err := h.VoterRepository.GetVoters(int(UserId))
	fmt.Println(voters)
	if voters {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Already Voted",
		})
	}

	data := models.Voter{
		UserID:   int(UserId),
		PaslonID: request.PaslonID,
	}

	response, err := h.VoterRepository.Vote(data)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: response})
}
