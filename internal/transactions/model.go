package transactions

import "time"

type TransactionDTO struct {
	UserID          string  `json:"user_id"`
	Symbol          string  `json:"symbol"`
	Amount          float64 `json:"amount"`
	BuyingPrice     float64 `json:"buying_price"`
	TransactionType string  `json:"transaction_type"`
}

type TransactionEntity struct {
	ID              string    `bson:"id"`
	UserID          string    `bson:"user_id"`
	Symbol          string    `bson:"symbol"`
	Amount          float64   `bson:"amount"`
	BuyingPrice     float64   `bson:"buying_price"`
	CreatedAt       time.Time `bson:"createdAt"`
	TransactionType string    `bson:"transaction_type"`
}

type Transaction struct {
	ID              string    `json:"id"`
	UserID          string    `json:"user_id"`
	Symbol          string    `json:"symbol"`
	Amount          float64   `json:"amount"`
	BuyingPrice     float64   `json:"buying_price"`
	CreatedAt       time.Time `json:"createdAt"`
	TransactionType string    `json:"transaction_type"`
}
