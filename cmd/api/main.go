package main

import (
	"pqan.com/go-api/internal/app"
	"pqan.com/go-api/internal/config"
)

func main() {
	config := config.NewConfig()
	application := app.NewApplication(config)

	if err := application.Run(); err != nil {
		panic(err)
	}
}
