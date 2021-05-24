package requests

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/oechsler/bone/posts/adapter"
	"github.com/oechsler/bone/posts/data"
	"github.com/oechsler/bone/posts/repositories"
)

type PostUpdateCommand struct {
	adapter.PostUpdateAdapter `validate:"required"`
}

func NewPostUpdateCommand(updateAdapter adapter.PostUpdateAdapter) *PostUpdateCommand {
	command := new(PostUpdateCommand)
	command.PostUpdateAdapter = updateAdapter

	return command
}

type PostUpdateHandler struct {
	echo.Logger
	repositories.PostRepository
}

func NewPostUpdateHandler(logger echo.Logger, postRepository repositories.PostRepository) PostUpdateHandler {
	return PostUpdateHandler{
		logger,
		postRepository,
	}
}

func (handler PostUpdateHandler) Send(command PostUpdateCommand) error {
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

func (handler PostUpdateHandler) handle(command PostUpdateCommand) error {
	err := handler.PostRepository.Update(command.PostUpdateAdapter.Id, data.Post{
		Title: command.PostUpdateAdapter.Title,
		Content: command.PostUpdateAdapter.Content,
	})
	if err != nil {
		return err
	}

	handler.Logger.Print("Successfully updated post with id '%s'", command.PostUpdateAdapter.Id)
	return nil
}