package user

import (
	"errors"
	"fmt"
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

func (s *Service) GetByID(id uint) (*UserResponse, error) {
	account, err := s.repository.FindByID(id)

	if err != nil {
		return nil, fmt.Errorf(
			"failed to find user: %w",
			err,
		)
	}

	if account == nil {
		return nil, ErrUserNotFound
	}

	response := toUserResponse(*account)

	return &response, nil
}

func (s *Service) GetAll() ([]UserResponse, error) {
	users, err := s.repository.FindAll()

	if err != nil {
		return nil, fmt.Errorf(
			"failed to fetch users: %w",
			err,
		)
	}

	response := make([]UserResponse, 0, len(users))

	for _, account := range users {
		response = append(
			response,
			toUserResponse(account),
		)
	}

	return response, nil
}

func (s *Service) GetProfile(userID uint) (*UserResponse, error) {
	return s.GetByID(userID)
}

func (s *Service) Delete(id uint) error {
	account, err := s.repository.FindByID(id)

	if err != nil {
		return fmt.Errorf(
			"failed to find user: %w",
			err,
		)
	}

	if account == nil {
		return ErrUserNotFound
	}

	if err := s.repository.Delete(account); err != nil {
		return fmt.Errorf(
			"failed to delete user: %w",
			err,
		)
	}

	return nil
}

func (s *Service) HasAnyRole(
	userID uint,
	roles ...string,
) (bool, error) {
	return s.repository.HasAnyRole(
		userID,
		roles...,
	)
}