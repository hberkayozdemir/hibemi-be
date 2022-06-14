package banner

type BannerDTO struct {
	Image string `json:"image"`
	Url   string `json:"url"`
}

type BannerEntity struct {
	ID    string `json:"id"`
	Image string `json:"image"`
	Url   string `json:"url"`
}

type Banner struct {
	ID    string `json:"id"`
	Image string `json:"image"`
	Url   string `json:"url"`
}

type BannerList struct {
	Banner []Banner `json:"coins"`
}
