package user

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	MongoClient *mongo.Client
}

func NewRepository(uri string) Repository {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	defer cancel()
	error := client.Connect(ctx)
	fmt.Print(error)
	if err != nil {
		log.Fatal(err)
	}

	return Repository{client}
}

func (r *Repository) RegisterUser(user User) (*User, error) {
	collection := r.MongoClient.Database("ventures").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, user)

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (r *Repository) DeleteUser(userID string) error {
	collection := r.MongoClient.Database("ventures").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.M{"id": userID}

	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetUserByEmail(email string) (*User, error) {
	collection := r.MongoClient.Database("ventures").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.M{"email": email}

	cur := collection.FindOne(ctx, filter)

	if cur.Err() != nil {
		return nil, cur.Err()
	}

	if cur == nil {
		return nil, cur.Err()
	}

	userEntity := User{}
	err := cur.Decode(&userEntity)
	if err != nil {
		return nil, err
	}

	return &userEntity, nil
}

func (r *Repository) ActivateUser(userID string) (*User, error) {
	collection := r.MongoClient.Database("ventures").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.M{"id": userID, "isEmailActivate": false}

	_, err := collection.UpdateOne(ctx,
		filter,
		bson.M{
			"$set": bson.M{
				"isEmailActivate": true,
			},
		},
	)
	if err != nil {
		return nil, err
	}

	return r.GetUser(userID)
}

func (r *Repository) GetUsers(page, size int) ([]User, int, error) {
	collection := r.MongoClient.Database("ventures").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	options := options.Find()
	if size != 0 {
		options.SetSkip(int64(page * size))
		options.SetLimit(int64(size))
	}

	filter := bson.M{"userType": "editor"}

	cur, err := collection.Find(ctx, filter, options)
	if err != nil {
		return nil, 0, err
	}

	var users []User
	for cur.Next(ctx) {
		userEntity := User{}
		err := cur.Decode(&userEntity)
		if err != nil {
			return nil, 0, err
		}
		users = append(users, userEntity)
	}

	totalElements, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	return users, int(totalElements), nil
}

func (r *Repository) AddActivationCode(email, activationCode string) error {
	collection := r.MongoClient.Database("ventures").Collection("codes")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, bson.M{"email": email, "activationCode": activationCode})
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) DeleteActivationCode(code string) error {
	collection := r.MongoClient.Database("ventures").Collection("codes")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.M{"activationCode": code}

	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetUser(userID string) (*User, error) {
	collection := r.MongoClient.Database("ventures").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.M{"id": userID}

	cur := collection.FindOne(ctx, filter)

	if cur.Err() != nil {
		return nil, cur.Err()
	}

	if cur == nil {
		return nil, UserNotFound
	}

	userEntity := User{}
	err := cur.Decode(&userEntity)

	if err != nil {
		return nil, err
	}

	return &userEntity, nil
}

func (r *Repository) GetCodesCount() (int64, error) {
	collection := r.MongoClient.Database("ventures").Collection("codes")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	codesCount, err := collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return 0, err
	}

	return codesCount, nil
}

func (r *Repository) GetCoinsCount() (int64, error) {
	collection := r.MongoClient.Database("ventures").Collection("coins")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	coinsCount, err := collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return 0, err
	}

	return coinsCount, nil
}

func (r *Repository) GetFavListsCount() (int64, error) {
	collection := r.MongoClient.Database("ventures").Collection("favlist")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	favListsCount, err := collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return 0, err
	}

	return favListsCount, nil
}

func (r *Repository) GetNewsCount() (int64, error) {
	collection := r.MongoClient.Database("ventures").Collection("news")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	newsCount, err := collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return 0, err
	}

	return newsCount, nil
}

func (r *Repository) GetSpotsCount() (int64, error) {
	collection := r.MongoClient.Database("ventures").Collection("spots")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	spotsCount, err := collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return 0, err
	}

	return spotsCount, nil
}

func (r *Repository) GetTransactionsCount() (int64, error) {
	collection := r.MongoClient.Database("ventures").Collection("spots")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	transactionsCount, err := collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return 0, err
	}

	return transactionsCount, nil
}
