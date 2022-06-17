package news

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

func (s *Service) getNews() ([]News, error) {
	news, err := s.Repository.GetNews()

	if err != nil {
		return nil, err
	}

	return news, nil
}

func (s *Service) GetNew(newID string) (*News, error) {
	new, err := s.Repository.GetNewsByID(newID)
	if err != nil {
		return nil, err
	}
	return new, nil
}
