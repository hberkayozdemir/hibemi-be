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

func convertTransactionEntityToModelArray(entityArray []TransactionEntity) []Transaction {

}
