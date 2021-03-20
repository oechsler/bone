package endpoints

import (
	"github.com/labstack/echo/v4"
	"github.com/oechsler/bone/posts/data"
	"github.com/oechsler/bone/posts/requests"
	"go.uber.org/dig"
	"net/http"
)

type PostsEndpoint struct {
	dig.In

	echo.Logger
	requests.PostCreateHandler
	requests.PostRetrieveAllHandler
	requests.PostUpdateHandler
}

func NewPostsEndpoint(echo *echo.Echo, endpoint PostsEndpoint) {
	group := echo.Group("/posts")
	group.POST("", endpoint.createPost)
	group.GET("", endpoint.retrieveAllPosts)
	group.GET("/:id", endpoint.retrievePost)
	group.PUT("/:id", endpoint.updatePost)
	group.DELETE("/:id", endpoint.deletePost)
}

func (endpoint PostsEndpoint) createPost(context echo.Context) error {
	dto := data.NewPostCreateDto()
	if err := (&echo.DefaultBinder{}).BindBody(context, dto); err != nil {
		return err
	}

	command := requests.NewPostCreateCommand(*dto)
	if err := endpoint.PostCreateHandler.Send(*command); err != nil {
		return err
	}
	return context.NoContent(http.StatusCreated)
}

func (endpoint PostsEndpoint) retrieveAllPosts(context echo.Context) error  {
	dto := data.NewPostRetrieveAllDto()
	_ = (&echo.DefaultBinder{}).BindQueryParams(context, dto)

	query := requests.NewPostRetrieveAllQuery(*dto)
	result, err := endpoint.PostRetrieveAllHandler.Send(*query)
	if err != nil {
		return err
	}
	return context.JSON(http.StatusOK, result)
}

func (endpoint PostsEndpoint) retrievePost(context echo.Context) error {
	panic("implement me")
}

func (endpoint PostsEndpoint) updatePost(context echo.Context) error {
	dto := data.NewPostUpdateDto()
	_ = (&echo.DefaultBinder{}).BindPathParams(context, dto)
	if err := (&echo.DefaultBinder{}).BindBody(context, dto); err != nil {
		return err
	}

	command := requests.NewPostUpdateCommand(*dto)
	if err := endpoint.PostUpdateHandler.Send(*command); err != nil {
		return err
	}
	return context.NoContent(http.StatusAccepted)
}

func (endpoint PostsEndpoint) deletePost(context echo.Context) error {
	panic("implement me")
}