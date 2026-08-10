package banner

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

var ErrUserNotFound = errors.New("user not found")

type Service struct {
	repository *Repository
}


func NewService(repository *Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) GetAll() ([]Banner, error) {
	banners, err := s.repository.FindAll()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
	}

	if len(banners) == 0 {
		return nil, nil
	}

	response := make([]Banner, 0, len(banners))
	for _, banner := range banners {
		response = append(response, banner)
	}

	return banners, nil
}

func (s *Service) GetByID(id uint) (*Banner, error) {
	banner, err := s.repository.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return banner, nil
}

func (s *Service) Create(banner *Banner) (*BannerResponse, error){
	createdBanner, err := s.repository.Create(banner)

	if err != nil {
		return nil, fmt.Errorf(
			"failed to create banner: %w",
			err,
		)
	}
	response := toBannerResponse(*createdBanner)

	return &response, nil

}

func (s *Service) Update(id uint, banner *Banner) (*BannerResponse, error){
	updateBanner, err := s.repository.Update(banner)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to update product: %w",
			err,
		)
	}
	response := toBannerResponse(*updateBanner)

	return &response, nil
}


func (s *Service) Delete(id uint) error {
	banner, err := s.repository.FindByID(id)
	if err != nil {
		return fmt.Errorf(
			"failed to find product: %w",
			err,
		)
	}

	if banner == nil {
		return ErrUserNotFound
	}

	return s.repository.Delete(banner)
}