package data

import "gopkg.in/mcuadros/go-defaults.v1"

type PostCreateDto struct {
	Title string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
}

func NewPostCreateDto() *PostCreateDto {
	dto := new(PostCreateDto)
	defaults.SetDefaults(dto)

	return dto
}