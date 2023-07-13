package main

import (
	"rnav2022/avenue/model"
	"rnav2022/avenue/server"
)

func main() {
	model.Setup()
	server.SetupAndListen()
}
