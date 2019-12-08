package main

import (
	"fmt"

	"github.com/luiz-mai/go-api-boilerplate/server"

	"github.com/luiz-mai/go-api-boilerplate/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = server.Run(cfg)
	if err != nil {
		fmt.Println(err)
		return
	}

	return
}
