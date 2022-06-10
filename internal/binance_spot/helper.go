package binance_spot

import "github.com/adshao/go-binance/v2"

func convertSymbolPriceToSymbolPriceEntity(symbolPrice binance.SymbolPrice) SymbolPriceEntity {
	return SymbolPriceEntity{
		Symbol: symbolPrice.Symbol,
		Price:  symbolPrice.Price,
	}
}
