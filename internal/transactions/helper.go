package transactions

func convertTransactionModelToEntity(transaction *Transaction) TransactionEntity {
	return TransactionEntity{
		ID:              transaction.ID,
		UserID:          transaction.UserID,
		Symbol:          transaction.Symbol,
		Amount:          transaction.Amount,
		BuyingPrice:     transaction.BuyingPrice,
		CreatedAt:       transaction.CreatedAt,
		TransactionType: transaction.TransactionType,
	}
}

func convertTransactionEntityToModel(transactionEntity *TransactionEntity) Transaction {
	return Transaction{
		ID:              transactionEntity.ID,
		UserID:          transactionEntity.UserID,
		Symbol:          transactionEntity.Symbol,
		Amount:          transactionEntity.Amount,
		BuyingPrice:     transactionEntity.BuyingPrice,
		CreatedAt:       transactionEntity.CreatedAt,
		TransactionType: transactionEntity.TransactionType,
	}
}

func convertTransactionEntityArrayToModelArray(transactionEntityArray []TransactionEntity) []Transaction {
	Transactions := []Transaction{}

	for _, item := range transactionEntityArray {

		Transactions = append(Transactions, Transaction{
			ID:              item.ID,
			UserID:          item.UserID,
			Symbol:          item.Symbol,
			Amount:          item.Amount,
			BuyingPrice:     item.BuyingPrice,
			CreatedAt:       item.CreatedAt,
			TransactionType: item.TransactionType,
		})
	}

	return Transactions
}
