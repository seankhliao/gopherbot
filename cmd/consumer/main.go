package main

import (
	"log"

	_ "github.com/heroku/x/hmetrics/onload"

	"github.com/gobridge/gopherbot/config"
)

func main() {
	c, err := config.LoadEnv()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	l := config.DefaultLogger(c)

	if err := runServer(c, l); err != nil {
		log.Fatalf("failed to run new consumer server: %v", err.Error())
	}
}
