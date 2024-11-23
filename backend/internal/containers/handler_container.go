package containers

import (
	"github.com/nnatchy/adv-top-proj/internal/handlers"
)

func (c *Container) HandlerProvider() {

	if err := c.Container.Provide(handlers.NewUserHandler); err != nil {
		c.Error = err
	}

	if err := c.Container.Provide(handlers.NewHandler); err != nil {
		c.Error = err
	}

	if err := c.Container.Provide(handlers.NewVotingHandler); err != nil {
		c.Error = err
	}

	if err := c.Container.Provide(handlers.NewCandidateHandler); err != nil {
		c.Error = err
	}
}
