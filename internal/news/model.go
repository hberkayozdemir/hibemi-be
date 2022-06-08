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
	Hashtags  []string  `json:"hashtags"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type GetNewsQuery struct {
	Hashtags []string `query:"hashtags"`
}
type NewsPageableResponse struct {
	News []News `json:"news"`
	Page Page   `json:"page"`
}

type NewSearchCredentialsDTO struct {
	ID string `json:"id"`
}
type Page struct {
	Number        int `json:"number"`
	Size          int `json:"size",omitempty`
	TotalElements int `json:"totalElements",omitempty"`
	TotalPages    int `json:"totalPages",omitempty"`
}
