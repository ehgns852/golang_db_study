package router

import "golang_db_study/config"

type Router struct {
	config *config.Config
}

func NewRouter(config *config.Config) (*Router, error) {
	r := &Router{
		config: config,
	}
	return r, nil
}
