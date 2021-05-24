package endpoints

import (
	"github.com/labstack/echo/v4"
	"github.com/oechsler/bone/posts/adapter"
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
	requests.PostDeleteHandler
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
	createAdapter := adapter.NewPostCreateAdapter()
	if err := (&echo.DefaultBinder{}).BindBody(context, createAdapter); err != nil {
		return err
	}

	command := requests.NewPostCreateCommand(*createAdapter)
	if err := endpoint.PostCreateHandler.Send(*command); err != nil {
		return err
	}
	return context.NoContent(http.StatusCreated)
}

func (endpoint PostsEndpoint) retrieveAllPosts(context echo.Context) error  {
	retrieveAllAdapter := adapter.NewPostRetrieveAllAdapter()
	_ = (&echo.DefaultBinder{}).BindQueryParams(context, retrieveAllAdapter)

	query := requests.NewPostRetrieveAllQuery(*retrieveAllAdapter)
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
	updateAdapter := adapter.NewPostUpdateAdapter()
	_ = (&echo.DefaultBinder{}).BindPathParams(context, updateAdapter)
	if err := (&echo.DefaultBinder{}).BindBody(context, updateAdapter); err != nil {
		return err
	}

	command := requests.NewPostUpdateCommand(*updateAdapter)
	if err := endpoint.PostUpdateHandler.Send(*command); err != nil {
		return err
	}
	return context.NoContent(http.StatusAccepted)
}

func (endpoint PostsEndpoint) deletePost(context echo.Context) error {
	deleteAdapter := adapter.NewPostDeleteAdapter()
	_ = (&echo.DefaultBinder{}).BindPathParams(context, deleteAdapter)

	command := requests.NewPostDeleteCommand(*deleteAdapter)
	if err := endpoint.PostDeleteHandler.Send(*command); err != nil {
		return err
	}
	return context.NoContent(http.StatusNoContent)
}