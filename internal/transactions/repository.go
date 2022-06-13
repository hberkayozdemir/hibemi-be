package transactions

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

type Repository struct {
	MongoClient *mongo.Client
}

func NewRepository(uri string) Repository {
	env := os.Getenv("APP_ENV")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	defer cancel()
	client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	if env == "test" {
		client.Database("ventures").Collection("users").Drop(ctx)
	}

	return Repository{client}
}

func (r *Repository) CreateTransaction(transaction Transaction) (*Transaction, error) {

	collection := r.MongoClient.Database("ventures").Collection("transaction")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()
	transactionEntity := convertTransactionModelToEntity(&transaction)
	_, err := collection.InsertOne(ctx, transactionEntity)
	if err != nil {
		return nil, err
	}
	transactionModel, _ := r.GetTransaction(transaction.ID)

	return transactionModel, err
}

func (r *Repository) GetTransaction(ID string) (*Transaction, error) {
	collection := r.MongoClient.Database("ventures").Collection("transaction")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	filter := bson.M{"id": ID}
	cur := collection.FindOne(ctx, filter)
	if cur.Err() != nil {
		return nil, cur.Err()
	}
	if cur == nil {
		return nil, cur.Err()
	}
	transactionEntity := TransactionEntity{}
	err := cur.Decode(&transactionEntity)
	if err != nil {
		return nil, err
	}
	transactionModel := convertTransactionEntityToModel(&transactionEntity)
	return &transactionModel, nil
}

func (r *Repository) GetTransactionHistory(UserID string) (*[]Transaction, error) {
	collection := r.MongoClient.Database("ventures").Collection("transaction")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	options := options.Find()
	filter := bson.M{"user_id": UserID}
	cur, err := collection.Find(ctx, filter, options)
	if err != nil {
		return nil, err
	}
	var transactionEntities []TransactionEntity
	for cur.Next(ctx) {
		transactionEntity := TransactionEntity{}
		err := cur.Decode(&transactionEntity)
		if err != nil {
			return nil, err
		}
		transactionEntities = append(transactionEntities, transactionEntity)
	}

}
