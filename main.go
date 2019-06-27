package main

import (
	"./config"
	"./server"
	"./util"
)

func main() {

	config := config.NewConfig()

	_, err := util.ConnectDatabase(config)
	if err != nil {
		panic(err)
	}

	server := server.NewServer(config)

	server.Run()

}
