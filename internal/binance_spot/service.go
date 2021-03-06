package binance_spot

import (
	"context"
	"github.com/adshao/go-binance/v2"
	coingecko "github.com/superoo7/go-gecko/v3"
	"github.com/superoo7/go-gecko/v3/types"
	"net/http"
	"time"
)

var (
	apiKey    = "cQZU8AAqYsOBF6gV5YdFkZUulm0ce3dTqSQsG7IQmg3CzFMq3Ab9oRqpFOTtS6vF"
	secretKey = "zgthBiYZ8wAP1cTOQUyuMAVUjY7qqrwNRHGYL65ttBGMhQ8CoROjU76gYoFtQUOX"
)

type Service struct {
	Repository Repository
}

func NewService(repository Repository) Service {
	return Service{
		Repository: repository,
	}
}

func (s *Service) getSpots() ([]*binance.SymbolPrice, error) {

	binanceClient := binance.NewClient(apiKey, secretKey)

	prices, err := binanceClient.NewListPricesService().Do(context.Background())
	s.Repository.UpdateDb(prices)

	return prices, err
}

func (s *Service) getCoinList() (*types.CoinList, error) {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}
	CG := coingecko.NewClient(httpClient)
	return CG.CoinsList()

}
