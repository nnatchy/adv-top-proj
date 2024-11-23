package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nnatchy/adv-top-proj/internal/entities"
	"github.com/nnatchy/adv-top-proj/internal/infrastructure/log"
	"github.com/nnatchy/adv-top-proj/internal/services"
)

type CandidateHandler struct {
	service services.ServiceGateway
	log     *log.Logger
}

func NewCandidateHandler(service services.ServiceGateway, log *log.Logger) *CandidateHandler {
	return &CandidateHandler{
		service: service,
		log:     log,
	}
}

func (h *CandidateHandler) SearchCandidates(c echo.Context) error {
	reqCtx := c.Request().Context()
	h.log.InfoWithID(reqCtx, "[Handler: SearchCandidates]: Called")

	// Bind the query parameter into the struct
	query := new(entities.SearchCandidatesRequest)
	if err := c.Bind(query); err != nil {
		h.log.ErrorWithID(reqCtx, "[Handler: SearchCandidates]: ", err)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Failed to parse query parameters.",
		})
	}

	// h.log.InfoWithID(reqCtx, query)

	// Validate the struct
	if err := c.Validate(query); err != nil {
		h.log.ErrorWithID(reqCtx, "[Handler: SearchCandidates]: ", err)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid position value. Valid values are: 'year_one', 'year_two', 'year_three', 'year_four', 'representative'",
		})
	}

	// Call the service layer to fetch and format candidates
	candidates, err := h.service.CandidateService.SearchCandidates(reqCtx, query.Position)
	if err != nil {
		h.log.ErrorWithID(reqCtx, "[Handler: SearchCandidates] Service error", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to fetch candidates",
		})
	}

	h.log.InfoWithID(reqCtx, candidates)
	// Return the candidates as a JSON response
	return c.JSON(http.StatusOK, candidates)
}
