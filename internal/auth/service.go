package auth

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"ecommerce-api/internal/role"
	"ecommerce-api/internal/user"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredentials = errors.New("invalid email or password")
	ErrInactiveAccount     = errors.New("account is inactive")
	ErrEmailAlreadyExists  = errors.New("email already exists")
	ErrCustomerRoleMissing = errors.New("customer role not found")
	ErrPasswordTooLong     = errors.New("password is too long")
)

type Service struct {
	userRepository *user.Repository
	roleRepository *role.Repository
	jwtSecret      []byte
}

func NewService(
	userRepository *user.Repository,
	roleRepository *role.Repository,
	jwtSecret string,
) *Service {
	return &Service{
		userRepository: userRepository,
		roleRepository: roleRepository,
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
func (s *Service) SignUp(
	request SignUpRequest,
) (*SignUpResponse, error) {
	email := strings.ToLower(
		strings.TrimSpace(request.Email),
	)

	existingUser, err := s.userRepository.FindByEmail(email)

	if err != nil {
		return nil, fmt.Errorf(
			"failed to check existing user: %w",
			err,
		)
	}

	if existingUser != nil {
		return nil, ErrEmailAlreadyExists
	}

	customerRole, err := s.roleRepository.FindBySlug("customer")

	if err != nil {
		return nil, fmt.Errorf(
			"failed to find customer role: %w",
			err,
		)
	}

	if customerRole == nil {
		return nil, ErrCustomerRoleMissing
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		if errors.Is(err, bcrypt.ErrPasswordTooLong) {
			return nil, ErrPasswordTooLong
		}

		return nil, fmt.Errorf(
			"failed to hash password: %w",
			err,
		)
	}

	account := &user.User{
		FirstName:    request.FirstName,
		LastName:     request.LastName,
		Email:        email,
		PasswordHash: string(hashedPassword),
		IsActive:     true,
	}

	if err := s.userRepository.CreateWithRole(
		account,
		*customerRole,
	); err != nil {
		return nil, fmt.Errorf(
			"failed to create user: %w",
			err,
		)
	}

	return &SignUpResponse{
		ID:        account.ID,
		FirstName: account.FirstName,
		LastName:  account.LastName,
		Email:     account.Email,
	}, nil
}