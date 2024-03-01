package main

import "go-api/configs"

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)
}
