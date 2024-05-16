package app

import (
	"golang_db_study/config"
	"golang_db_study/repository"
	"golang_db_study/router"
	"golang_db_study/service"
)

type App struct {
	config *config.Config

	router     *router.Router
	repository *repository.Repository
	service    *service.Service
}

func NewApp(config *config.Config) *App {
	a := &App{
		config: config,
	}

	var err error
	if a.repository, err = repository.NewRepository(config); err != nil {
		panic(err)
	}

	if a.service, err = service.NewService(config); err != nil {
		panic(err)
	}

	if a.router, err = router.NewRouter(config); err != nil {
		panic(err)
	}

	// TODO 서버 실행

	return a
}
