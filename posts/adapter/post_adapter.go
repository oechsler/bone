package data

type PostDto struct {
	Id string `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	Title string `json:"title"`
	Content string `json:"content"`
}