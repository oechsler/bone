package repositories

import (
	"fmt"
	"github.com/ahmetb/go-linq/v3"
	"github.com/labstack/echo/v4"
	uuid "github.com/nu7hatch/gouuid"
	"github.com/oechsler/bone/posts/data"
	"math"
	"net/http"
	"time"
)

type PostRepositoryImpl struct {
	PostRepository

	echo.Logger
	posts []data.Post
}

func NewPostRepositoryImpl(logger echo.Logger) *PostRepositoryImpl {
	instance := &PostRepositoryImpl{
		Logger: logger,
		posts: []data.Post{},
	}
	return instance
}

func (repository *PostRepositoryImpl) Create(post data.Post) (string, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return "", echo.NewHTTPError(http.StatusInternalServerError, "Could not generate unique post id.")
	}
	post.Id = id.String()
	post.CreatedAt = time.Now()

	repository.posts = append(repository.posts, post)
	return id.String(), nil
}

func (repository PostRepositoryImpl) RetrieveAll(skip int, take int) []data.Post {
	posts := make([]data.Post, 0, 0)

	linq.From(repository.posts).
	GroupByT(
		func (post data.Post) string {
			return post.Id
		},
		func (post data.Post) data.Post {
			return post
		},
	).
	SelectT(func (group linq.Group) data.Post {
		return linq.From(group.Group).
		OrderByT(func (post data.Post) int64 {
			return post.UpdatedAt.UnixNano()
		}).
		Last().(data.Post)
	}).
	Skip(skip).
	Take(take).
	ToSlice(&posts)

	return posts
}

func (repository PostRepositoryImpl) Retrieve(id string) (data.Post, error) {
	posts := repository.RetrieveAll(0, math.MaxInt32)

	post := linq.From(posts).
	FirstWithT(func (post data.Post) bool {
		return post.Id == id
	})
	if post == nil {
		message := fmt.Sprintf("Post with id '%s' does not exist", id)
		return data.Post{}, echo.NewHTTPError(http.StatusBadRequest, message)
	}

	return post.(data.Post), nil
}

func (repository *PostRepositoryImpl) Update(id string, post data.Post) error {
	postToUpdate, err := repository.Retrieve(id)
	if err != nil {
		return err
	}

	if postToUpdate.Title == post.Title &&
	   postToUpdate.Content == post.Content {
		return echo.NewHTTPError(http.StatusBadRequest, "No changes have been made to the post.")
	}

	post.Id = postToUpdate.Id
	post.CreatedAt = postToUpdate.CreatedAt
	post.UpdatedAt = time.Now()

	repository.posts = append(repository.posts, post)
	return nil
}

func (repository *PostRepositoryImpl) Delete(id string) error {
	postToDelete, err := repository.Retrieve(id)
	if err != nil {
		return err
	}

	linq.From(repository.posts).
	WhereT(func (post data.Post) bool {
		return post.Id != postToDelete.Id
	}).
	ToSlice(&repository.posts)

	return nil
}
