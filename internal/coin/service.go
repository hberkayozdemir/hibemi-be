package coin

import (
	"fmt"
	"github.com/hberkayozdemir/hibemi-be/internal/coin_gecko"
	"math"
	"strings"
)

type Service struct {
	Repository Repository
}

func NewService(repository Repository) Service {
	return Service{
		Repository: repository,
	}
}

func (s *Service) GetAllSpots(pageNumber, size int) (*CoinsPageableResponse, error) {
	spots, totelElements, err := s.Repository.getAllSpots(pageNumber, size)
	if err != nil {
		return nil, err
	}
	page := Page{
		Number:        pageNumber,
		Size:          size,
		TotalElements: totelElements,
		TotalPages:    int(math.Ceil(float64(totelElements) / float64(size))),
	}

	var arr []string
	for _, spot := range spots {
		arr = strings.Split(spot.Symbol, "U")
		lowerSymbol := strings.ToLower(arr[0])
		fmt.Println(lowerSymbol)
		s.Repository.UpdateGeckoPrice(lowerSymbol, spot.Price)
	}

	return &CoinsPageableResponse{Coins: spots,
		Page: page,
	}, nil
}

func (s *Service) GetAllCoins() ([]coin_gecko.CoinGeckoResponse, error) {
	coins, err := s.Repository.GetAllCoins()
	if err != nil {
		return nil, err
	}

	return coins, nil
}
