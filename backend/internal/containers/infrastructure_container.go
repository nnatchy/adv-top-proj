package containers

import (
	"net/http"
	"time"

	"github.com/nnatchy/adv-top-proj/internal/config"
	"github.com/nnatchy/adv-top-proj/internal/infrastructure/database"
	"github.com/nnatchy/adv-top-proj/internal/infrastructure/ethereum"
	log "github.com/nnatchy/adv-top-proj/internal/infrastructure/log"
	"github.com/nnatchy/adv-top-proj/internal/infrastructure/server"
	"github.com/nnatchy/adv-top-proj/internal/middlewares"
)

func (c *Container) InfrastructureProvider() {
	// Provide the configuration
	if err := c.Container.Provide(config.LoadConfig); err != nil {
		c.Error = err
	}

	c.Container.Provide(database.ConnectPostgres)

	// Provide the logger with config dependency
	if err := c.Container.Provide(func(cfg *config.AppConfig) *log.Logger {
		return log.Initialize(cfg.AppEnv)
	}); err != nil {
		c.Error = err
	}

	// Provide HTTP client
	if err := c.Container.Provide(func() *http.Client {
		return &http.Client{Timeout: 10 * time.Second}
	}); err != nil {
		c.Error = err
	}

	// Provide the Ethereum client with config dependency
	if err := c.Container.Provide(func(cfg *config.AppConfig) (*ethereum.Client, error) {
		return ethereum.NewEthereumClient(cfg)
	}); err != nil {
		c.Error = err
	}

	// Provide VotingNFT-specific Ethereum client if needed (optional)
	if err := c.Container.Provide(func(cfg *config.AppConfig) (*ethereum.Client, error) {
		return ethereum.NewEthereumClient(cfg) // Can add unique VotingNFT RPC if required
	}); err != nil {
		c.Error = err
	}

	if err := c.Container.Provide(func(cfg *config.AppConfig) *server.EchoServer {
		return server.NewEchoServer(cfg, c.Container)
	}); err != nil {
		c.Error = err
	}

	if err := c.Container.Provide(middlewares.NewMiddleware); err != nil {
		c.Error = err
	}
}
