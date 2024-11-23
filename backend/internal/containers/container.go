package containers

import (
	"log"

	"github.com/nnatchy/adv-top-proj/internal/config"
	"github.com/nnatchy/adv-top-proj/internal/infrastructure/ethereum"
	"github.com/nnatchy/adv-top-proj/internal/infrastructure/server"
	"go.uber.org/dig"
)

type Container struct {
	Container *dig.Container
	Error     error
}

func (c *Container) Configure() {
	c.Container = dig.New()

	c.Container.Provide(config.LoadConfig)
	c.SolidityProvider()
	c.InfrastructureProvider()
	c.RepositoryProvider()
	c.ServiceProvider()
	c.HandlerProvider()

}

func (c *Container) Run() error {
	err := c.Container.Invoke(func(
		s *server.EchoServer,
		ethClient *ethereum.Client,
	) error {
		// Start the server
		if err := s.Start(); err != nil {
			log.Printf("Failed to start the server: %v", err)
			return err
		}
		log.Println("Server started successfully")

		return nil
	})

	if err != nil {
		log.Printf("Container Invoke failed: %v", err)
	}
	return err
}
func NewContainer() *Container {
	c := &Container{}
	c.Configure()
	return c
}
