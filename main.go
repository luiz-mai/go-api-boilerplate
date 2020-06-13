package main

import (
	"fmt"

	"github.com/luiz-mai/go-api-boilerplate/data"
	"github.com/luiz-mai/go-api-boilerplate/server"
	"github.com/luiz-mai/go-api-boilerplate/service"

	"github.com/luiz-mai/go-api-boilerplate/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
		return
	}

	db, err := data.Connect(cfg)
	if err != nil {
		fmt.Println(err)
		return
	}

	todoService := service.NewToDoService(db)

	server.Run(
		cfg,
		todoService,
	)
}
