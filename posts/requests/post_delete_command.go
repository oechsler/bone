package requests

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/oechsler/bone/posts/adapter"
	"github.com/oechsler/bone/posts/repositories"
)

type PostDeleteCommand struct {
	adapter.PostDeleteAdapter `validate:"required"`
}

func NewPostDeleteCommand(deleteAdapter adapter.PostDeleteAdapter) *PostDeleteCommand {
	command := new(PostDeleteCommand)
	command.PostDeleteAdapter = deleteAdapter

	return command
}

type PostDeleteHandler struct {
	echo.Logger
	repositories.PostRepository
}

func NewPostDeleteHandler(logger echo.Logger, postRepository repositories.PostRepository) PostDeleteHandler {
	return PostDeleteHandler{
		logger,
		postRepository,
	}
}

func (handler PostDeleteHandler) Send(command PostDeleteCommand) error {
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

func (handler PostDeleteHandler) handle(command PostDeleteCommand) error {
	err := handler.PostRepository.Delete(command.PostDeleteAdapter.Id)
	if err != nil {
		return err
	}

	handler.Logger.Printf("Successfully delete post with id '%s'", command.PostDeleteAdapter.Id)
	return nil
}