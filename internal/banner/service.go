package banner

import "github.com/hberkayozdemir/hibemi-be/helpers"

type Service struct {
	Repository Repository
}

func NewService(repository Repository) Service {
	return Service{
		Repository: repository,
	}
}

func (s *Service) BannerList() ([]Banner, error) {

	bannerList, err := s.Repository.getBannerList()
	if err != nil {
		return nil, err
	}
	return bannerList, nil
}

func (s *Service) CreateBanner(bannerDTO BannerDTO) (*Banner, error) {

	banner := Banner{
		ID:    helpers.GenerateUUID(8),
		Image: bannerDTO.Image,
		Url:   bannerDTO.Url,
	}
	newBanner, err := s.Repository.CreateBanner(banner)
	if err != nil {
		return nil, err
	}
	return newBanner, err
}

func (s *Service) DeleteNews(dto string) error {
	err := s.Repository.DeleteBanner(dto)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetBanner(bannerID string) (*Banner, error) {
	new, err := s.Repository.GetBannerByID(bannerID)
	if err != nil {
		return nil, err
	}
	return new, nil
}
