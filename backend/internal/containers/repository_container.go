package containers

import "github.com/nnatchy/adv-top-proj/internal/repositories"

func (c *Container) RepositoryProvider() {

	if err := c.Container.Provide(repositories.NewUserRepository); err != nil {
		c.Error = err
	}

	if err := c.Container.Provide(repositories.NewSimpleStorageRepository); err != nil {
		c.Error = err
	}

	if err := c.Container.Provide(repositories.NewRepository); err != nil {
		c.Error = err
	}

	if err := c.Container.Provide(repositories.NewVoteRepository); err != nil {
		c.Error = err
	}

	if err := c.Container.Provide(repositories.NewCandidateRepository); err != nil {
		c.Error = err
	}
}
