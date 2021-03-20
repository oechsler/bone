package data

import (
	"time"
)

type Post struct {
	Id string `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Title string `json:"title"`
	Content string `json:"content"`
}