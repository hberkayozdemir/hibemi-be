package coin

import "math"

type Service struct {
	Repository Repository
}

func NewService(repository Repository) Service {
	return Service{
		Repository: repository,
	}
}

func (s *Service) GetAllSpots(pageNumber, size int) (*CoinsPageableResponse, error) {
	coins, totelElements, err := s.Repository.getAllSpots(pageNumber, size)
	if err != nil {
		return nil, err
	}
	page := Page{
		Number:        pageNumber,
		Size:          size,
		TotalElements: totelElements,
		TotalPages:    int(math.Ceil(float64(totelElements) / float64(size))),
	}
	return &CoinsPageableResponse{Coins: coins,
		Page: page,
	}, nil
}
