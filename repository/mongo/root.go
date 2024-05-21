package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang_db_study/config"
	"strconv"
	"strings"
)

type Mongo struct {
	config *config.Config

	client *mongo.Client
	db     *mongo.Database

	user    *mongo.Collection
	content *mongo.Collection
	history *mongo.Collection
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

		m.user = m.db.Collection("user")
		m.content = m.db.Collection("content")
		m.history = m.db.Collection("history")
	}

	//createIndex(m.db.Collection("test"), []string{"key"}, []string{})
	if err = createIndex(m.user, []string{"user"}, []string{"user"}); err != nil {
		panic(err)
	} else if err = createIndex(m.content, []string{"name"}, []string{"name"}); err != nil {
		panic(err)
	} else if err = createIndex(m.history, []string{"user"}, []string{}); err != nil {
		panic(err)
	}
	return m, nil
}

func createIndex(collection *mongo.Collection, indexes, uniques []string) error {
	type indexOptions struct {
		key    string
		order  int64
		unique bool
	}

	var indexsOpt []indexOptions

	for _, field := range indexes {
		noU := false

		for _, unique := range uniques {
			if unique == field {
				indexsOpt = append(indexsOpt, indexOptions{key: field, order: -1, unique: true})
				noU = true
				break
			}
		}
		if noU {
			indexsOpt = append(indexsOpt, indexOptions{key: field, order: -1, unique: false})
		}
	}
	ctx := context.Background()

	needToCreate := make(map[string]indexOptions)

	if indexCursor, err := collection.Indexes().List(ctx); err != nil {
		panic(err)
	} else {
		defer indexCursor.Close(ctx)

		for indexCursor.Next(ctx) {
			if v, ok := indexCursor.Current.Lookup("name").StringValueOK(); !ok || v == " _id_ " {
				continue
			} else {
				fmt.Println(v)
				split := strings.Split(v, "_")
				if len(split) == 2 {
					if order, err := strconv.Atoi(split[1]); err == nil {
						if order == 1 || order == -1 {
							needToCreate[split[0]] = indexOptions{split[0], int64(order), false}
						}
					}
				}
			}
		}
	}

	for _, i := range indexsOpt {
		if value, ok := needToCreate[i.key]; ok {
			opt := options.Index()

			if value.unique {
				opt.SetUnique(value.unique)
			}

			m := mongo.IndexModel{
				Keys:    bson.D{{Key: value.key, Value: 1}},
				Options: opt,
			}

			if _, err := collection.Indexes().CreateOne(ctx, m); err != nil {
				return err
			}
		}
	}

	return nil
}
