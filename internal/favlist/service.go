package favlist

import (
	"github.com/hberkayozdemir/hibemi-be/helpers"
)

type Service struct {
	Repository Repository
}

func NewService(repository Repository) Service {
	return Service{
		Repository: repository,
	}
}

func (s *Service) CreateFavCoin(favCoinDTO FavCoinDTO) (*FavCoin, error) {

	favCoin := FavCoin{
		ID:           helpers.GenerateUUID(8),
		Symbol:       favCoinDTO.Symbol,
		CurrentPrice: favCoinDTO.CurrentPrice,
		UserID:       favCoinDTO.UserID,
	}

	createdFavCoin, err := s.Repository.CreateFavCoin(favCoin)
	if err != nil {
		return nil, err
	}

	return createdFavCoin, nil
}
