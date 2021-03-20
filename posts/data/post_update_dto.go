package data

import "gopkg.in/mcuadros/go-defaults.v1"

type PostUpdateDto struct {
	Id string `param:"id" validate:"uuid"`
	Title string `json:"title"`
	Content string `json:"content"`
}

func NewPostUpdateDto() *PostUpdateDto {
	dto := new(PostUpdateDto)
	defaults.SetDefaults(dto)

	return dto
}