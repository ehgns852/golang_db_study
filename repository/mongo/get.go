package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"golang_db_study/types"
)

func (m *Mongo) GetUserBucket(user string) (*types.User, error) {

	filter := bson.M{"user": user}

	var u types.User

	if err := m.user.FindOne(context.Background(), filter).Decode(&user); err != nil {
		return nil, err
	} else {
		return &u, nil
	}

	return nil, nil
}
