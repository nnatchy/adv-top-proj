package middlewares

import (
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/nnatchy/adv-top-proj/internal/config"
	log "github.com/nnatchy/adv-top-proj/internal/infrastructure/log"
)

// MiddlewareInterface defines the middleware methods.
type MiddlewareInterface interface {
	EthereumMiddleware(allowedRoles ...string) echo.MiddlewareFunc
}

// Middleware struct contains the logger and configuration.
type Middleware struct {
	log    *log.Logger
	config *config.AppConfig
}

// NewMiddleware initializes a new Middleware instance.
func NewMiddleware(logger *log.Logger, config *config.AppConfig) MiddlewareInterface {
	return &Middleware{
		log:    logger,
		config: config,
	}
}

// JwtCustomClaims represents the structure of custom JWT claims.
type JwtCustomClaims struct {
	WalletAddress string `json:"walletAddress"`
	Role          string `json:"role"`
	jwt.StandardClaims
}

// EthereumMiddleware validates the JWT token, extracts the Ethereum address, and checks roles if provided
func (m *Middleware) EthereumMiddleware(allowedRoles ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := c.Request().Context()
			m.log.InfoWithID(ctx, "[EthereumMiddleware] Called")

			// Extract JWT token from the Authorization header
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				m.log.ErrorWithID(ctx, "[EthereumMiddleware] Missing Authorization header")
				return echo.NewHTTPError(http.StatusUnauthorized, "Missing Authorization header")
			}

			tokenParts := strings.Split(authHeader, " ")
			if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
				m.log.ErrorWithID(ctx, "[EthereumMiddleware] Invalid Authorization header format")
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid Authorization header format")
			}

			tokenString := tokenParts[1]

			// Parse and validate the token
			token, err := jwt.ParseWithClaims(tokenString, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
				// Ensure that the token's signing method is what you expect
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					m.log.ErrorWithID(ctx, "[EthereumMiddleware] Unexpected signing method")
					return nil, echo.NewHTTPError(http.StatusUnauthorized, "Unexpected signing method")
				}
				// Return the secret key
				return []byte(m.config.Jwt.JWT_SECRET_KEY), nil
			})
			if err != nil {
				m.log.ErrorWithID(ctx, "[EthereumMiddleware] Invalid token:", err)
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
			}

			// Extract claims
			claims, ok := token.Claims.(*JwtCustomClaims)
			if !ok || !token.Valid {
				m.log.ErrorWithID(ctx, "[EthereumMiddleware] Invalid token claims")
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token claims")
			}

			// Get Ethereum address from claims
			address := claims.WalletAddress

			if !common.IsHexAddress(address) {
				m.log.ErrorWithID(ctx, "[EthereumMiddleware] Invalid Ethereum address in token claims")
				return echo.NewHTTPError(http.StatusUnauthorized, "Valid Ethereum address is required")
			}

			// Store Ethereum address in the context for further use
			c.Set("ethAddress", common.HexToAddress(address))

			// Validate role if roles are provided (optional)
			if len(allowedRoles) > 0 {
				role := claims.Role
				if !m.isRoleAllowed(role, allowedRoles) {
					m.log.ErrorWithID(ctx, "[EthereumMiddleware] Access denied for role:", role)
					return echo.NewHTTPError(http.StatusForbidden, "Access denied for your role")
				}
			}

			// Proceed to the next handler if authorization checks pass
			return next(c)
		}
	}
}

// isRoleAllowed checks if the given role is allowed for the route.
func (m *Middleware) isRoleAllowed(role string, allowedRoles []string) bool {
	for _, allowed := range allowedRoles {
		if role == allowed {
			return true
		}
	}
	return false
}
