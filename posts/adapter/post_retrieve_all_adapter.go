package adapter

import "gopkg.in/mcuadros/go-defaults.v1"

type PostRetrieveAllAdapter struct {
	Skip int `query:"skip" default:"0" validate:"gte=0"`
	Take int `query:"take" default:"10" validate:"gt=0"`
}

func NewPostRetrieveAllAdapter() *PostRetrieveAllAdapter {
	adapter := new(PostRetrieveAllAdapter)
	defaults.SetDefaults(adapter)

	return adapter
}
