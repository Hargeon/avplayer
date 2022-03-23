package main

import (
	"os"

	"github.com/Hargeon/avplayer/api"
	"github.com/joho/godotenv"
)

func main() {
	if err := setup(); err != nil {
		panic(err)
	}
}

func setup() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	router := api.NewHandler()
	if err := router.InitRoutes().Listen(os.Getenv("PORT")); err != nil {
		return err
	}

	return nil
}
