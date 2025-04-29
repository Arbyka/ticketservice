package service

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"ticketing/entity"
	"ticketing/repository"
)

type AuthService interface {
	Register(user *entity.User) error
	Login(email, password string) (*entity.User, error)
}

type authService struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{userRepo}
}

func (s *authService) Register(user *entity.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	if user.Role != "admin" && user.Role != "customer" {
		user.Role = "customer"
	}

	return s.userRepo.Create(user)
}

func (s *authService) Login(email, password string) (*entity.User, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}
