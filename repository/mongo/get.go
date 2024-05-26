package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"golang_db_study/types"
)

func (m *Mongo) GetUserBucket(user string) (*types.User, error) {

	filter := bson.M{"user": user}

	u := &types.User{}

	if err := m.user.FindOne(context.Background(), filter).Decode(&user); err != nil {
		return nil, err
	} else {
		return u, nil
	}
}

func (m *Mongo) GetContent(name string) ([]*types.Content, error) {
	filter := bson.M{}

	if name != "" {
		filter["user"] = name
	}

	ctx := context.Background()

	if cursor, err := m.content.Find(ctx, filter); err != nil {
		return nil, err
	} else {
		defer cursor.Close(ctx)

		var v []*types.Content

		if err := cursor.All(ctx, v); err != nil {
			return nil, err
		} else {
			return v, nil
		}
	}
}

func (m *Mongo) GetUserHistory(user string) (*types.History, error) {
	filter := bson.M{"user": user}

	h := &types.History{}

	if err := m.user.FindOne(context.Background(), filter).Decode(&user); err != nil {
		return nil, err
	} else {
		return h, nil
	}
}
