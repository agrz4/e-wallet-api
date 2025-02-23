package service

import (
	"e-wallet-api-go/internal/dto"
	"e-wallet-api-go/internal/model"
)

type TransactionService interface {
	GetTransactions(userID int, query *dto.TransactionRequestQuery) ([]*model.Transaction, error)
	TopUp(input *dto.TopUpRequestBody) (*model.Transaction, error)
	Transfer(input *dto.TransferRequestBody) (*model.Transaction, error)
	CountTransaction(userID int) (int64, error)
}

type transactionService struct {
	transactionRepository  r.TransactionRepository
	walletRepository       r.WalletRepository
	sourceOfFundRepository r.SourceOfFundRepository
}

type TSConfig struct {
	TransactionRepository  r.TransactionRepository
	WalletRepository       r.WalletRepository
	SourceOfFundRepository r.SourceOfFundRepository
}

func NewTransactionService(c *TSConfig) TransactionService {
	return &transactionService{
		transactionRepository:  c.TransactionRepository,
		walletRepository:       c.WalletRepository,
		sourceOfFundRepository: c.SourceOfFundRepository,
	}
}
