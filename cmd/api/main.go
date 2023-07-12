package main

import (
	"blog/cmd/api/server"
	"blog/cmd/api/server/middleware"
	"blog/pkg/domain"
	"blog/pkg/repository/database"
	"github.com/gin-gonic/gin"
)

func main() {
	//init db instance
	client, err := database.NewClient()

	if err != nil {
		panic(err)
	}

	//Populate database
	db(client)
	//Start the server
	if err := run(client); err != nil {
		panic("Error starting server")
	}
}

func run(client database.Client) error {
	r := gin.Default()
	r.Use(middleware.ErrorMiddleware)
	server.Bootstrap(r, client)
	return r.Run()
}

func db(client database.Client) {
	client.Migrate(&domain.Author{})
	client.Migrate(&domain.AuthorProfile{})
	client.Migrate(&domain.Post{})
	client.Migrate(&domain.Reactions{})
	client.MigrateAndSeed(&domain.Category{}, &domain.Category{
		Name: "Tecnologia",
	})
	client.Migrate(&domain.Comments{})
}
