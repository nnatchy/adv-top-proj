package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nnatchy/adv-top-proj/internal/entities"
	"github.com/nnatchy/adv-top-proj/internal/infrastructure/log"
	"github.com/nnatchy/adv-top-proj/internal/services"
)

type VotingHandler struct {
	service services.ServiceGateway
	log     *log.Logger
}

func NewVotingHandler(service services.ServiceGateway, log *log.Logger) *VotingHandler {
	return &VotingHandler{
		service: service,
		log:     log,
	}
}

// CastVote handles casting a vote
func (h *VotingHandler) CastVote(c echo.Context) error {
	reqCtx := c.Request().Context()

	// Bind and validate the request
	var req entities.CastVoteRequest
	if err := c.Bind(&req); err != nil || req.StudentID == "" || req.CandidateID == "" {
		h.log.ErrorWithID(reqCtx, "[Handler: CastVote] Invalid request", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	// Call the service layer
	txHash, err := h.service.VoteService.CastVote(reqCtx, req)
	if err != nil {
		h.log.ErrorWithID(reqCtx, "[Handler: CastVote] Service error", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to cast vote"})
	}

	return c.JSON(http.StatusOK, map[string]string{"transactionHash": txHash})
}

// EndVoting handles ending the voting process
func (h *VotingHandler) EndVoting(c echo.Context) error {
	reqCtx := c.Request().Context()

	// Bind and validate the request
	var req entities.EndVotingRequest
	if err := c.Bind(&req); err != nil || req.AdminAddress == "" {
		h.log.ErrorWithID(reqCtx, "[Handler: EndVoting] Invalid request", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	// Call the service layer
	txHash, err := h.service.VoteService.EndVoting(reqCtx, req)
	if err != nil {
		h.log.ErrorWithID(reqCtx, "[Handler: EndVoting] Service error", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to end voting"})
	}

	return c.JSON(http.StatusOK, map[string]string{"transactionHash": txHash})
}

// GetVoteDetails retrieves details of a specific vote
func (h *VotingHandler) GetVoteDetails(c echo.Context) error {
	reqCtx := c.Request().Context()
	studentID := c.Param("studentId")

	// Call the service layer
	voteDetails, err := h.service.VoteService.GetVoteDetails(reqCtx, studentID)
	if err != nil {
		h.log.ErrorWithID(reqCtx, "[Handler: GetVoteDetails] Service error", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch vote details"})
	}

	return c.JSON(http.StatusOK, voteDetails)
}

// GetVotingResults handles fetching the voting results
func (h *VotingHandler) GetVotingResults(c echo.Context) error {
	reqCtx := c.Request().Context()
	h.log.InfoWithID(reqCtx, "[Handler: GetVotingResults] Fetching voting results")

	results, err := h.service.VoteService.GetVotingResults(reqCtx)
	if err != nil {
		h.log.ErrorWithID(reqCtx, "[Handler: GetVotingResults] Failed to fetch voting results", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch voting results"})
	}

	return c.JSON(http.StatusOK, results)
}

// GetAllVotesFromContract handles fetching all votes from the contract via events
func (h *VotingHandler) GetAllVotesFromContract(c echo.Context) error {
	reqCtx := c.Request().Context()
	h.log.InfoWithID(reqCtx, "[Handler: GetAllVotesFromContract] Fetching all votes from the contract via events")

	// Call the service layer
	votes, err := h.service.VoteService.GetAllVotesFromEvents(reqCtx)
	if err != nil {
		h.log.ErrorWithID(reqCtx, "[Handler: GetAllVotesFromContract] Service error", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch votes from contract"})
	}

	return c.JSON(http.StatusOK, votes)
}

func (h *VotingHandler) GetCandidatesRanking(c echo.Context) error {
	reqCtx := c.Request().Context()
	h.log.InfoWithID(reqCtx, "[Handler: GetCandidatesRanking] Fetching all votes from the contract via events")

	// Call the service layer
	resp, err := h.service.VoteService.GetCandidatesRanking(reqCtx)
	if err != nil {
		h.log.ErrorWithID(reqCtx, "[Handler: GetCandidatesRanking] Service error", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch votes from contract"})
	}

	return c.JSON(http.StatusOK, resp)
}
