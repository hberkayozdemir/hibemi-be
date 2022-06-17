package coin

import (
	"github.com/hberkayozdemir/hibemi-be/internal/coin_gecko"
)

type Service struct {
	Repository Repository
}

func NewService(repository Repository) Service {
	return Service{
		Repository: repository,
	}
}

func (s *Service) GetAllSpots() ([]Coins, error) {
	spots, err := s.Repository.getAllSpots()
	if err != nil {
		return nil, err
	}

	return spots, nil
}

func (s *Service) GetAllCoins() ([]coin_gecko.CoinGeckoResponse, error) {
	coins, err := s.Repository.GetAllCoins()
	if err != nil {
		return nil, err
	}

	return coins, nil
}
