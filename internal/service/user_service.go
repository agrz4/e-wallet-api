package service

import (
	"e-wallet-api-go/internal/dto"
	"e-wallet-api-go/internal/model"
	"net/mail"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	GetUser(input *dto.UserRequestParams) (*model.User, error)
	CreateUser(input *dto.RegisterRequestBody) (*model.User, error)
}

type userService struct {
	UserRepository   r.UserRepository
	WalletRepository r.WalletRepository
}

type USConfig struct {
	UserRepository   r.UserRepository
	WalletRepository r.WalletRepository
}

func NewUserService(c *USConfig) UserService {
	return &userService{
		userRepository:   c.UserRepository,
		walletRepository: c.WalletRepository,
	}
}

func (s *userService) GetUser(input *dto.UserRequestParams) (*model.User, error) {
	user, err := s.UserRepository.FindById(input.UserID)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *userService) CreateUser(input *dto.RegisterRequestBody) (*model.User, error) {
	_, err := mail.ParseAddress(input.Email)
	if err != nil {
		return &model.User{}, &custom_error.NotValidEmailError{}
	}

	user, err := s.userRepository.FindByEmail(input.Email)
	if err != nil {
		return user, err
	}
	if user.ID != 0 {
		return user, &custom_error.UserAlreadyExistsError{}
	}

	user.Name = input.Name
	user.Email = input.Email
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.Password = string(passwordHash)

	newUser, err := s.userRepository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}
