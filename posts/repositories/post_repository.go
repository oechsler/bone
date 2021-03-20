package repositories

import "github.com/oechsler/bone/posts/data"

type PostRepository interface {
	Create(post data.Post) (string, error)
	RetrieveAll(skip int, take int) []data.Post
	Retrieve(id string) (data.Post, error)
	Update(id string, post data.Post) error
	Delete(id string) bool
}
