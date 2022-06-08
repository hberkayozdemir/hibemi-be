package news

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

func (r *Repository) GetNews(page, size int) ([]News, int, error) {
	collection := r.MongoClient.Database("ventures").Collection("news")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	options := options.Find()
	if size != 0 {
		options.SetSkip(int64(page * size))
		options.SetLimit(int64(size))
	}
	filter := bson.M{}
	cur, err := collection.Find(ctx, filter, options)
	if err != nil {
		return nil, 0, nil
	}
	var news []News
	for cur.Next(ctx) {
		userEntity := News{}
		err := cur.Decode(&userEntity)
		if err != nil {
			return nil, 0, nil
		}
		news = append(news, userEntity)
	}

	totalElements, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, nil
	}

	return news, int(totalElements), nil
}

func (r *Repository) AddNews(news News) (*News, error) {
	collection := r.MongoClient.Database("ventures").Collection("news")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, news)

	if err != nil {
		return nil, err
	}

	return r.GetNewsByID(news.ID)
}

func (r *Repository) DeleteNews(id string) error {
	collection := r.MongoClient.Database("ventures").Collection("news")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	filter := bson.M{"id": id}
	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetNewsByID(id string) (*News, error) {
	collection := r.MongoClient.Database("ventures").Collection("news")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.M{"id": id}

	cur := collection.FindOne(ctx, filter)

	if cur.Err() != nil {
		return nil, cur.Err()
	}

	if cur == nil {
		return nil, cur.Err()
	}

	newsEntitiy := News{}
	err := cur.Decode(&newsEntitiy)
	if err != nil {
		return nil, err
	}

	return &newsEntitiy, nil
}
