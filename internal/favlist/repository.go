package favlist

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

func (r *Repository) CreateFavCoin(favCoin FavCoin) (*FavCoin, error) {
	collection := r.MongoClient.Database("ventures").Collection("favlist")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, favCoin)

	if err != nil {
		return nil, err
	}

	favCoinModel, _ := r.GetFavCoinWithID(favCoin.ID)
	return favCoinModel, err
}

func (r *Repository) GetFavCoinWithID(ID string) (*FavCoin, error) {
	collection := r.MongoClient.Database("ventures").Collection("favlist")
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

	favCoinEntity := FavCoin{}
	err := cur.Decode(&favCoinEntity)
	if err != nil {
		return nil, cur.Err()
	}
	return &favCoinEntity, nil
}
