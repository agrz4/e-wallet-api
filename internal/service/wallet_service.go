package service

import (
	"e-wallet-api-go/internal/dto"
	"e-wallet-api-go/internal/model"
	r "e-wallet-api-go/internal/repository"
	"e-wallet-api-go/pkg/customerror"
	"e-wallet-api-go/pkg/utils"
)

type WalletService interface {
	GetWalletByUserId(input *dto.WalletRequestBody) (*model.Wallet, error)
	CreateWallet(input *dto.WalletRequestBody) (*model.Wallet, error)
}

type walletService struct {
	userRepository   r.UserRepository
	walletRepository r.WalletRepository
}

type WSConfig struct {
	UserRepository   r.UserRepository
	WalletRepository r.WalletRepository
}

func NewWalletService(c *WSConfig) WalletService {
	return &walletService{
		userRepository:   c.UserRepository,
		walletRepository: c.WalletRepository,
	}
}

func (s *walletService) GetWalletByUserId(input *dto.WalletRequestBody) (*model.Wallet, error) {
	wallet, err := s.walletRepository.FindByUserId(input.UserID)
	if err != nil {
		return wallet, err
	}
	return wallet, nil
}

func (s *walletService) CreateWallet(input *dto.WalletRequestBody) (*model.Wallet, error) {
	user, err := s.userRepository.FindById(input.UserID)
	if err != nil {
		return &model.Wallet{}, err
	}
	if user.ID == 0 {
		return &model.Wallet{}, &customerror.UserNotFoundError{}
	}

	wallet, err := s.walletRepository.FindByUserId(int(user.ID))
	if err != nil {
		return &model.Wallet{}, err
	}
	if wallet.ID != 0 {
		return &model.Wallet{}, &customerror.WalletAlreadyExistsError{}
	}

	wallet.UserID = user.ID
	wallet.Number = utils.GenerateWalletNumber(user.ID)
	wallet.Balance = 0

	newWallet, err := s.walletRepository.Save(wallet)
	if err != nil {
		return newWallet, err
	}

	return newWallet, nil
}
