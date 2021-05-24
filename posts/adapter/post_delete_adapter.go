package adapter

import "gopkg.in/mcuadros/go-defaults.v1"

type PostDeleteAdapter struct {
	Id string `param:"id" validate:"uuid"`
}

func NewPostDeleteAdapter() *PostDeleteAdapter {
	adapter := new(PostDeleteAdapter)
	defaults.SetDefaults(adapter)

	return adapter
}