package requests

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/oechsler/bone/posts/adapter"
	"github.com/oechsler/bone/posts/data"
	"github.com/oechsler/bone/posts/repositories"
)

type PostCreateCommand struct {
	adapter.PostCreateAdapter `validate:"required"`
}

func NewPostCreateCommand(createAdapter adapter.PostCreateAdapter) *PostCreateCommand {
	command := new(PostCreateCommand)
	command.PostCreateAdapter = createAdapter

	return command
}

type PostCreateHandler struct {
	echo.Logger
	repositories.PostRepository
}

func NewPostCreateHandler(logger echo.Logger, postRepository repositories.PostRepository) PostCreateHandler {
	return PostCreateHandler{
		logger,
		postRepository,
	}
}

func (handler PostCreateHandler) Send(command PostCreateCommand) error {
	var err error

	validate := validator.New()
	err = validate.Struct(command)
	if err != nil {
		return err
	}

	err = handler.handle(command)
	if err != nil {
		return err
	}

	return nil
}

func (handler PostCreateHandler) handle(command PostCreateCommand) error {
	id, err := handler.PostRepository.Create(data.Post{
		Title: command.PostCreateAdapter.Title,
		Content: command.PostCreateAdapter.Content,
	})
	if err != nil {
		return err
	}

	handler.Logger.Printf("Successfully created post with id '%s'", id)
	return nil
}