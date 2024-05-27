package mongo

import (
	"errors"
	"fmt"
)

func (m *MService) PostCreateUser(user string) error {
	if err := m.repository.Mongo.PostCreateUser(user); err != nil {
		fmt.Println(err)
		return err
	} else {
		return err
	}
}

func (m *MService) PostCreateContent(user string, price int64) error {
	if err := m.repository.Mongo.PostCreateContent(user, price); err != nil {
		fmt.Println(err)
		return err
	} else {
		return err
	}
}

func (m *MService) PostBucketRequest(user, content string) error {

	if c, err := m.repository.Mongo.GetContent(content); err != nil {
		fmt.Println("GetContent err", err)
		return err
	} else if _, err := m.repository.Mongo.GetUserBucket(user); err != nil {
		fmt.Println("GetUserBucket err", err)
		return err
	} else if len(c) == 0 {
		return errors.New("content 없습니다")
	} else if err := m.repository.Mongo.PostInsertBucket(user, content); err != nil {
		fmt.Println("PostInsertBucket Err", err)
		return err
	} else {

	}
}

func (m *MService) PostUserBuy(user string) error {
	if err := m.repository.Mongo.Post(user); err != nil {
		fmt.Println(err)
		return err
	} else {
		return err
	}
}
