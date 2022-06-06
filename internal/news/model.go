package news

import "time"

type NewsDTO struct {
	Title    string   `json:"title"`
	Image    string   `json:"image"`
	Content  string   `json:"content"`
	Hashtags []string `json:"hashtags"`
}

type News struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Image     string    `json:"image"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type GetNewsQuery struct {
	Hashtags []string `query:"hashtags"`
}
