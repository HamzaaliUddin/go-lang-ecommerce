package auth

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"ecommerce-api/internal/user"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredentials = errors.New("invalid email or password")
	ErrInactiveAccount     = errors.New("account is inactive")
)

type Service struct {
	userRepository *user.Repository
	jwtSecret      []byte
}

func NewService(
	userRepository *user.Repository,
	jwtSecret string,
) *Service {
	return &Service{
		userRepository: userRepository,
		jwtSecret:      []byte(jwtSecret),
	}
}

func (s *Service) Login(request LoginRequest) (*LoginResponse, error) {
	account, err := s.userRepository.FindByEmail(request.Email)

	if err != nil {
		return nil, fmt.Errorf("failed to find user: %w", err)
	}

	if account == nil {
		return nil, ErrInvalidCredentials
	}

	if !account.IsActive {
		return nil, errors.New("account is inactive")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(account.PasswordHash),
		[]byte(request.Password),
	)

	if err != nil {
		return nil, ErrInvalidCredentials
	}

	now := time.Now()
	expiresAt := now.Add(time.Hour)

	claims := jwt.RegisteredClaims{
		Subject: strconv.FormatUint(
			uint64(account.ID),
			10,
		),
		IssuedAt: jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(expiresAt),
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	accessToken, err := token.SignedString(s.jwtSecret)
	if err != nil {
		return nil, fmt.Errorf("failed to generate access token: %w", err)
	}

	return &LoginResponse{
		AccessToken: accessToken,
		TokenType:   "Bearer",
		ExpiresIn:   3600,
	}, nil
}