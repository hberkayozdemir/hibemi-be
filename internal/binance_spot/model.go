package binance_spot

type SymbolPrice struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}
type SymbolPriceEntityList struct {
	SymbolPrices []interface{} `bson:"symbol_prices"`
}

type SymbolPriceEntity struct {
	Symbol string `bson:"symbol"`
	Price  string `bson:"price"`
}
