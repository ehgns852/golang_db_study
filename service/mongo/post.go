package mongo

func (m *MService) PostCreateUser(user string) error {
	//if r, err := m.repository.Mongo.GetUserBucket(user); err != nil {
	//	fmt.Println(err)
	//	return err
	//} else {
	return nil
	//}
}

func (m *MService) PostCreateContent(name string, price int64) error {
	return nil
}

func (m *MService) PostUserBuy(user string) error {
	return nil
}
