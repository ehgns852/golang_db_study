package service

import (
	"golang_db_study/config"
	"golang_db_study/repository"
	"golang_db_study/service/mongo"
)

type Service struct {
	config     *config.Config
	repository *repository.Repository
	MService   *mongo.MService
}

func NewService(config *config.Config, repository *repository.Repository) (*Service, error) {
	r := &Service{
		config:     config,
		repository: repository,
		MService:   mongo.NewService(repository),
	}
	return r, nil
}
