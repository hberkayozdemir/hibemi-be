package favlist

type FavCoin struct {
	ID           string `json:"id" bson:"id"`
	Symbol       string `json:"symbol" bson:"symbol"`
	CurrentPrice string `json:"current_price" bson:"current_price"`
	UserID       string `json:"user_id" bson:"user_id"`
}

type FavCoinEntity struct {
	ID           string `bson:"id"`
	Symbol       string `bson:"symbol"`
	CurrentPrice string `bson:"current_price"`
	UserID       string `bson:"user_id"`
}

type FavCoinDTO struct {
	Symbol       string `json:"symbol"`
	CurrentPrice string `json:"current_price"`
	UserID       string `json:"user_id"`
}
