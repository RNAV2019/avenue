package main

import (
	"rnav2022/avenue/model"
	"rnav2022/avenue/server"
)

func main() {
	model.Setup()
	err := server.SetupFirebase()
	if err != nil {
		panic(err)
	}
	server.SetupAndListen()
}
