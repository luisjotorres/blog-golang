package server

import (
	"blog/cmd/api/server/auth"
	authors2 "blog/cmd/api/server/authors"
	"blog/cmd/api/server/middleware"
	"blog/cmd/api/server/posts"
	"blog/pkg/repository/database"
	"github.com/gin-gonic/gin"
)

func Bootstrap(app *gin.Engine, client database.Client) {
	authHandlers := auth.InitHandlers(client)
	authorsHandlers := authors2.InitHandlers(client)
	postHandlers := posts.InitHandlers(client)
	auth := app.Group("/auth")
	{
		auth.POST("/login", authHandlers.Login)
		auth.POST("/register", authHandlers.RegisterAuthor)
		auth.GET("/token-for-visit", authHandlers.TokenForVisit)
	}

	app.GET("/ping", authHandlers.Ping)
	blog := app.Group("/blog")
	//blog.Use(middleware.Authenticate)
	{
		blog.POST("/create", postHandlers.CreatePost)
		blog.GET("/", postHandlers.Index)
		blog.GET("/read/:id", postHandlers.ReadPost)
		blog.GET("/reactions/:id", postHandlers.Reactions)
		// blog.GET("/posts", postHandlers.Posts)
		//blog.POST("/comment")
		//blog.PUT("/update")
		//blog.POST("/vote/id")
		//blog.POST("register-visit")
	}

	authors := app.Group("/author")
	authors.Use(middleware.Authenticate)
	{
		authors.GET("/my-profile", authorsHandlers.GetProfile)
		authors.PATCH("/update", authorsHandlers.UpdateProfile)
	}

}
