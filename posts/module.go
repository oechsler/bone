package posts

import (
	"github.com/labstack/echo/v4"
	"github.com/oechsler/bone/posts/endpoints"
	"github.com/oechsler/bone/posts/repositories"
	"github.com/oechsler/bone/posts/requests"
	"go.uber.org/dig"
)

func UseModule(container *dig.Container) error {
	var err error

	// Register injectable services
	err = container.Provide(func (logger echo.Logger) repositories.PostRepository {
		return repositories.NewPostRepositoryImpl(logger)
	})
	if err != nil {
		return err
	}

	// Register commands / queries
	err = container.Provide(requests.NewPostCreateHandler)
	if err != nil {
		return err
	}
	err = container.Provide(requests.NewPostRetrieveAllHandler)
	if err != nil {
		return err
	}
	err = container.Provide(requests.NewPostUpdateHandler)
	if err != nil {
		return err
	}
	err = container.Provide(requests.NewPostDeleteHandler)
	if err != nil {
		return err
	}

	// Register endpoints
	err = container.Invoke(endpoints.NewPostsEndpoint)
	if err != nil {
		return err
	}

	return nil
}