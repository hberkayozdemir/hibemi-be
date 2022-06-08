package news

import (
	"github.com/hberkayozdemir/hibemi-be/helpers"
	"math"
)

type Service struct {
	Repository Repository
}

func NewService(repository Repository) Service {
	return Service{
		Repository: repository,
	}
}

func (s *Service) DeleteNews(dto string) error {
	err := s.Repository.DeleteNews(dto)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) AddNews(newsDTO NewsDTO) (*News, error) {
	news := News{
		ID:       helpers.GenerateUUID(8),
		Title:    newsDTO.Title,
		Content:  newsDTO.Content,
		Image:    newsDTO.Image,
		Hashtags: newsDTO.Hashtags,
	}
	newNews, err := s.Repository.AddNews(news)
	if err != nil {
		return nil, err
	}

	return newNews, nil
}

func (s *Service) getNews(pageNumber, size int) (*NewsPageableResponse, error) {
	news, totelElements, err := s.Repository.GetNews(pageNumber, size)

	if err != nil {
		return nil, err
	}
	page := Page{
		Number:        pageNumber,
		Size:          size,
		TotalElements: totelElements,
		TotalPages:    int(math.Ceil(float64(totelElements) / float64(size))),
	}

	return &NewsPageableResponse{
		News: news,
		Page: page,
	}, nil
}

func (s *Service) GetNew(newID string) (*News, error) {
	new, err := s.Repository.GetNewsByID(newID)
	if err != nil {
		return nil, err
	}
	return new, nil
}
