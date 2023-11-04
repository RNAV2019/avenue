package main

import (
	"rnav2022/avenue/model"
	"rnav2022/avenue/server"
	// "github.com/joho/godotenv"
)

func main() {
	// Development code only
	// err := godotenv.Load()
	// if err != nil {
	// 	panic(err)
	// }

	model.Setup()
	server.SetupAndListen()
}
