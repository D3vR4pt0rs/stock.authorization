package main

import (
	"fmt"
	"os"

	"authentication/internal/infrastructure/postgres"
)

func main() {

	config := postgres.Config{
		Username: os.Getenv("POSTGRES_USERNAME"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Ip:       os.Getenv("POSTGRES_IP"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Database: os.Getenv("POSTGRES_DATABASE"),
	}

	postgresClient := postgres.New(config)
	profile, _ := postgresClient.GetProfileByEmail("test@gmail.com")
	fmt.Println(profile)
}
