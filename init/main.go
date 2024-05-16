package main

import (
	"flag"
	"fmt"
	"golang_db_study/config"
	"golang_db_study/init/app"
)

var envFlag = flag.String("config", "./env.toml", "env not found")

func main() {
	flag.Parse()

	c := config.NewConfig(*envFlag)
	app.NewApp(c)

	fmt.Println("에러가 없습니다.")
}
