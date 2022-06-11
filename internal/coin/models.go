package coin

type Coins struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

type CoinsPageableResponse struct {
	Coins []Coins `json:"coins"`
	Page  Page    `json:"page"`
}

type Page struct {
	Number        int `json:"number"`
	Size          int `json:"size",omitempty`
	TotalElements int `json:"totalElements",omitempty"`
	TotalPages    int `json:"totalPages",omitempty"`
}
