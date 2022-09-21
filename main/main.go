package main

import (
	"fmt"

	"goblog"
	"goblog/config"
)

func main() {
	fmt.Println("hello world in console")
	b := goblog.New(config.ServerConfig.Port, config.ServerConfig.LogLevel)
	b.Serve()
}
