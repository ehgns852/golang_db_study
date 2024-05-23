package mongo

import (
	"fmt"
	"golang_db_study/types"
)

func (m *MService) GetUserBucket(user string) (*types.User, error) {
	if r, err := m.repository.Mongo.GetUserBucket(user); err != nil {
		fmt.Println(err)
		return nil, err
	} else {
		return r, err
	}
}
