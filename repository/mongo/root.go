package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang_db_study/config"
)

type Mongo struct {
	config *config.Config

	client *mongo.Client
	db     *mongo.Database
}

func NewMongo(config *config.Config) (*Mongo, error) {
	m := &Mongo{
		config: config,
	}

	ctx := context.Background()
	var err error

	if m.client, err = mongo.Connect(ctx, options.Client().ApplyURI(config.Mongo.Uri)); err != nil {
		panic(err)
	} else if err = m.client.Ping(ctx, nil); err != nil {
		panic(err)
	} else {
		m.db = m.client.Database(config.Mongo.Db)
	}

	//createIndex(m.db.Collection("test"), []string{"key"}, []string{})

	return m, nil
}

func createIndex(collection *mongo.Collection, indexes, uniques []string) error {
	opt := options.Index()

	for _, k := range indexes {
		model := mongo.IndexModel{
			Keys:    bson.D{{Key: k, Value: 1}},
			Options: opt,
		}

		if res, err := collection.Indexes().CreateOne(context.Background(), model); err != nil {
			panic(err)
		} else {
			fmt.Println(res)
		}
	}

	opt.SetUnique(true)
	for _, k := range uniques {
		model := mongo.IndexModel{
			Keys:    bson.D{{Key: k, Value: 1}},
			Options: opt,
		}

		if res, err := collection.Indexes().CreateOne(context.Background(), model); err != nil {
			panic(err)
		} else {
			fmt.Println(res)
		}
	}

	return nil
}
