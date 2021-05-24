package requests

import (
	"github.com/ahmetb/go-linq/v3"
	"github.com/dustin/go-humanize"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/oechsler/bone/posts/adapter"
	"github.com/oechsler/bone/posts/data"
	"github.com/oechsler/bone/posts/repositories"
)

type PostRetrieveAllQuery struct {
	adapter.PostRetrieveAllAdapter `validate:"required"`
}

func NewPostRetrieveAllQuery(retrieveAllAdapter adapter.PostRetrieveAllAdapter) *PostRetrieveAllQuery {
	query := new(PostRetrieveAllQuery)
	query.PostRetrieveAllAdapter = retrieveAllAdapter

	return query
}

type PostRetrieveAllHandler struct {
	echo.Logger
	repositories.PostRepository
}

func NewPostRetrieveAllHandler(logger echo.Logger, postRepository repositories.PostRepository) PostRetrieveAllHandler {
	return PostRetrieveAllHandler{
		logger,
		postRepository,
	}
}

func (handler PostRetrieveAllHandler) Send(query PostRetrieveAllQuery) ([]adapter.PostAdapter, error) {
	var err error

	validate := validator.New()
	err = validate.Struct(query)
	if err != nil {
		return nil, err
	}

	result, err := handler.handle(query)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (handler PostRetrieveAllHandler) handle(query PostRetrieveAllQuery) ([]adapter.PostAdapter, error) {
	result := handler.PostRepository.RetrieveAll(query.Skip, query.Take)

	posts := make([]adapter.PostAdapter, 0, 0)
	linq.From(result).
	SelectT(func (post data.Post) adapter.PostAdapter {
		updatedAtHumanized := humanize.Time(post.UpdatedAt)
		if post.UpdatedAt.IsZero() {
			updatedAtHumanized = "never"
		}

		return adapter.PostAdapter{
			Id:        post.Id,
			CreatedAt: humanize.Time(post.CreatedAt),
			UpdatedAt: updatedAtHumanized,
			Title:     post.Title,
			Content:   post.Content,
		}
	}).
	ToSlice(&posts)

	handler.Logger.Printf("Successfully retrieved '%d/%d' requested posts.", len(posts), query.Take)
	return posts, nil
}