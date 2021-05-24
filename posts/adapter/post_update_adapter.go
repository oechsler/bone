package adapter

import "gopkg.in/mcuadros/go-defaults.v1"

type PostUpdateAdapter struct {
	Id string `param:"id" validate:"uuid"`
	Title string `json:"title"`
	Content string `json:"content"`
}

func NewPostUpdateAdapter() *PostUpdateAdapter {
	adapter := new(PostUpdateAdapter)
	defaults.SetDefaults(adapter)

	return adapter
}