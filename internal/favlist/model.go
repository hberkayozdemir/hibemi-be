package favlist

type FavCoin struct {
	ID           string  `json:"id" bson:"id"`
	Symbol       string  `json:"symbol" bson:"symbol"`
	CurrentPrice float64 `json:"current_price" bson:"current_price"`
	UserID       string  `json:"user_id" bson:"user_id"`
}

type FavCoinDTO struct {
	Symbol       string  `json:"symbol"`
	CurrentPrice float64 `json:"current_price"`
	UserID       string  `json:"user_id"`
}
