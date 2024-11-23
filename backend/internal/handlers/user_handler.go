package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nnatchy/adv-top-proj/internal/entities"
	"github.com/nnatchy/adv-top-proj/internal/infrastructure/log"
	"github.com/nnatchy/adv-top-proj/internal/models"
	"github.com/nnatchy/adv-top-proj/internal/services"
)

type UserHandler struct {
	service services.ServiceGateway
	log     *log.Logger
}

func NewUserHandler(service services.ServiceGateway, log *log.Logger) *UserHandler {
	return &UserHandler{
		service: service,
		log:     log,
	}
}

// HealthCheck handles the health check endpoint
func (h *UserHandler) HealthCheck(c echo.Context) error {
	reqCtx := c.Request().Context()
	h.log.InfoWithID(reqCtx, "[Handler: HealthCheck]: Called")
	response, err := h.service.UserService.PerformHealthCheck(reqCtx)
	if err != nil {
		h.log.ErrorWithID(reqCtx, "[Handler: HealthCheck]: Failed to perform health check", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to reach user-service"})
	}
	return c.JSON(http.StatusOK, response)
}

func (h *UserHandler) MintToken(c echo.Context) error {
	reqCtx := c.Request().Context()

	var req struct {
		Address string `json:"address"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	err := h.service.UserService.Mint(reqCtx, req.Address)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"status": "Token minted"})
}

func (h *UserHandler) GetBalance(c echo.Context) error {
	reqCtx := c.Request().Context()
	address := c.Param("address")

	balance, err := h.service.UserService.BalanceOf(reqCtx, address)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"balance": balance})
}

func (h *UserHandler) BurnToken(c echo.Context) error {
	reqCtx := c.Request().Context()

	var req struct {
		TokenID string `json:"tokenId"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	err := h.service.UserService.Burn(reqCtx, req.TokenID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"status": "Token burned"})
}

func (h *UserHandler) Login(c echo.Context) error {
	reqCtx := c.Request().Context()

	var req struct {
		WalletAddress string `json:"walletAddress"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	user, token, err := h.service.UserService.Login(reqCtx, req.WalletAddress)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
		"user":  user,
	})
}

func (h *UserHandler) Logout(c echo.Context) error {
	// Perform any necessary cleanup or session invalidation if needed
	return c.JSON(http.StatusOK, map[string]string{"message": "Logout successful"})
}

func (h *UserHandler) Register(c echo.Context) error {
	reqCtx := c.Request().Context()

	// Define the request structure
	var req struct {
		WalletAddress string `json:"walletAddress"`
		Username      string `json:"username"`
		FirstName     string `json:"first_name"`
		LastName      string `json:"last_name"`
		StudentID     string `json:"student_id"`
	}

	// Bind the incoming request
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	// Call the service layer to handle the registration
	err := h.service.UserService.Register(reqCtx, req.WalletAddress, req.Username, req.FirstName, req.LastName, req.StudentID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "User registered successfully"})
}

func (h *UserHandler) IsActivate(c echo.Context) error {
	reqCtx := c.Request().Context()

	// Get the wallet address from the query params
	walletAddress := c.QueryParam("walletAddress")
	if walletAddress == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "walletAddress is required"})
	}

	// Call the service layer to check activation status
	isActive, err := h.service.UserService.IsActivate(reqCtx, walletAddress)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]bool{"isActive": isActive})
}
func (h *UserHandler) GetUserInfoByWallet(c echo.Context) error {
	reqCtx := c.Request().Context()

	// Get the wallet address from the query params
	walletAddress := c.QueryParam("walletAddress")
	if walletAddress == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "walletAddress is required"})
	}

	userInfo, err := h.service.UserService.GetUserInfoByWallet(reqCtx, walletAddress)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, userInfo)
}

func (h *UserHandler) UpdateUserStatus(c echo.Context) error {
	reqCtx := c.Request().Context()

	// Get the wallet address from the query params
	walletAddress := c.QueryParam("walletAddress")
	if walletAddress == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "walletAddress is required"})
	}

	// Initialize the request object and bind the JSON body to it
	var req entities.UpdateUserStatusRequest
	if err := c.Bind(&req); err != nil { // Bind to the address of the struct
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}

	// Call the service layer to update user status
	err := h.service.UserService.UpdateUserStatus(reqCtx, walletAddress, &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "user status updated successfully"})
}

func (h *UserHandler) AdminLogin(c echo.Context) error {
	reqCtx := c.Request().Context()

	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	admin, err := h.service.UserService.AdminLogin(reqCtx, req.Username, req.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Admin login successful",
		"admin":   admin,
	})
}

// CreateCandidate handles creating a new candidate
func (h *UserHandler) CreateCandidate(c echo.Context) error {
	reqCtx := c.Request().Context()

	var candidate models.Candidate
	if err := c.Bind(&candidate); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	newCandidate, err := h.service.UserService.CreateCandidate(reqCtx, &candidate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, newCandidate)
}

// GetCandidateByID handles fetching a candidate by ID
func (h *UserHandler) GetCandidateByID(c echo.Context) error {
	reqCtx := c.Request().Context()

	id := c.Param("id")
	candidate, err := h.service.UserService.GetCandidateByID(reqCtx, id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, candidate)
}

// UpdateCandidate handles updating an existing candidate
func (h *UserHandler) UpdateCandidate(c echo.Context) error {
	reqCtx := c.Request().Context()

	id := c.Param("id")
	var updatedCandidate models.Candidate
	if err := c.Bind(&updatedCandidate); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	candidate, err := h.service.UserService.UpdateCandidate(reqCtx, id, &updatedCandidate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, candidate)
}

// DeleteCandidate handles deleting a candidate by ID
func (h *UserHandler) DeleteCandidate(c echo.Context) error {
	reqCtx := c.Request().Context()

	id := c.Param("id")
	if err := h.service.UserService.DeleteCandidate(reqCtx, id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Candidate deleted successfully"})
}

func (h *UserHandler) GetCandidates(c echo.Context) error {
	reqCtx := c.Request().Context()

	// Get the position query parameter
	position := c.QueryParam("position")

	// Call the service layer
	candidates, err := h.service.UserService.GetCandidates(reqCtx, position)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, candidates)
}
