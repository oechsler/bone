package adapter

import "gopkg.in/mcuadros/go-defaults.v1"

type PostCreateAdapter struct {
	Title string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
}

func NewPostCreateAdapter() *PostCreateAdapter {
	adapter := new(PostCreateAdapter)
	defaults.SetDefaults(adapter)

	return adapter
}