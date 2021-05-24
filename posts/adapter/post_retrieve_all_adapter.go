package data

import "gopkg.in/mcuadros/go-defaults.v1"

type PostRetrieveAllDto struct {
	Skip int `query:"skip" default:"0" validate:"gte=0"`
	Take int `query:"take" default:"10" validate:"gt=0"`
}

func NewPostRetrieveAllDto() *PostRetrieveAllDto {
	dto := new(PostRetrieveAllDto)
	defaults.SetDefaults(dto)

	return dto
}
