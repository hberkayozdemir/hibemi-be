package transactions

import (
	"github.com/hberkayozdemir/hibemi-be/helpers"
	"time"
)

type Service struct {
	Repository Repository
}

func NewService(repository Repository) Service {
	return Service{
		Repository: repository,
	}
}

func (s *Service) CreateTransaction(transactionDTO TransactionDTO) (*Transaction, error) {
	transaction := Transaction{
		ID:              helpers.GenerateUUID(8),
		UserID:          transactionDTO.UserID,
		Symbol:          transactionDTO.Symbol,
		Amount:          transactionDTO.Amount,
		BuyingPrice:     transactionDTO.BuyingPrice,
		CreatedAt:       time.Now().UTC().Round(time.Second),
		TransactionType: transactionDTO.TransactionType,
	}
	newTransaction, err := s.Repository.CreateTransaction(transaction)
	if err != nil {
		return nil, err
	}

	return newTransaction, nil
}

func (s *Service) GetTransactionHistory(userId string) (*[]Transaction, error) {

	transactions, err := s.Repository.GetTransactionHistory(userId)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}
