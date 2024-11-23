package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/dig"

	"github.com/nnatchy/adv-top-proj/internal/config"
	"github.com/nnatchy/adv-top-proj/internal/handlers"
)

func RegisterRoutes(e *echo.Echo, c *dig.Container, conf *config.AppConfig) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     conf.AllowOrigin.FrontendURLs,
		AllowMethods:     []string{echo.GET, echo.POST, echo.PUT, echo.PATCH, echo.DELETE, echo.OPTIONS},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		ExposeHeaders:    []string{"Content-Length", "Authorization"},
		AllowCredentials: true,
	}))

	if err := c.Invoke(func(h *handlers.Handler) {
		api := e.Group("/api")
		userRoutes(api, h.Gateway.UserHandler)
		votingRoutes(api, h.Gateway.VotingHandler)
		candidateRoutes(api, h.Gateway.CandidateHandler)
	}); err != nil {
		panic(err)
	}
}

// userRoutes registers routes related to user operations
func userRoutes(eg *echo.Group, userHandler *handlers.UserHandler) {
	userRoutes := eg.Group("/user/v1")
	{
		userRoutes.POST("/mint", userHandler.MintToken)
		userRoutes.GET("/balance/:address", userHandler.GetBalance)
		userRoutes.POST("/burn", userHandler.BurnToken)
		userRoutes.POST("/login", userHandler.Login)
		userRoutes.POST("/logout", userHandler.Logout)
		userRoutes.POST("/register", userHandler.Register)
		userRoutes.GET("/isActivate", userHandler.IsActivate)
		userRoutes.POST("/admin/login", userHandler.AdminLogin)
		userRoutes.POST("/candidates", userHandler.CreateCandidate)
		userRoutes.GET("/candidates/:id", userHandler.GetCandidateByID)
		userRoutes.PUT("/candidates/:id", userHandler.UpdateCandidate)
		userRoutes.DELETE("/candidates/:id", userHandler.DeleteCandidate)
		userRoutes.GET("/candidates", userHandler.GetCandidates)

		userRoutes.GET("/info", userHandler.GetUserInfoByWallet)
		userRoutes.PATCH("/status", userHandler.UpdateUserStatus)
	}
}

// votingRoutes registers routes related to voting operations
func votingRoutes(eg *echo.Group, votingHandler *handlers.VotingHandler) {
	votingRoutes := eg.Group("/voting/v1")
	{
		votingRoutes.POST("/vote", votingHandler.CastVote)
		votingRoutes.POST("/end", votingHandler.EndVoting)
		votingRoutes.GET("/details/:studentId", votingHandler.GetVoteDetails) // Changed from :tokenId to :studentId
		votingRoutes.GET("/results", votingHandler.GetVotingResults)          // Corrected typo in GetVotingResults
		votingRoutes.GET("/all-votes", votingHandler.GetAllVotesFromContract) // Get all votes from the contract
		votingRoutes.GET("/candidates-ranking", votingHandler.GetCandidatesRanking)
	}
}

// candidateRoutes registers routes related to candidate operations
func candidateRoutes(eg *echo.Group, candidateHandler *handlers.CandidateHandler) {
	candidateRoutes := eg.Group("/candidate/v1")
	{
		candidateRoutes.GET("/search", candidateHandler.SearchCandidates)
	}
}
