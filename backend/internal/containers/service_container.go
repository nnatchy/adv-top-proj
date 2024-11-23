package containers

import "github.com/nnatchy/adv-top-proj/internal/services"

func (c *Container) ServiceProvider() {

	if err := c.Container.Provide(services.NewUserService); err != nil {
		c.Error = err
	}

	if err := c.Container.Provide(services.NewService); err != nil {
		c.Error = err
	}

	if err := c.Container.Provide(services.NewVoteService); err != nil {
		c.Error = err
	}

	if err := c.Container.Provide(services.NewCandidateService); err != nil {
		c.Error = err
	}
}
