package auth

import (
	"blog/pkg/repository/database"
	"blog/pkg/usecases/auth"
	"github.com/gin-gonic/gin"
)

type Handlers struct {
	Ping           gin.HandlerFunc
	CreatePost     gin.HandlerFunc
	RegisterAuthor gin.HandlerFunc
	Login          gin.HandlerFunc
	TokenForVisit  gin.HandlerFunc
}

func InitHandlers(client database.Client) *Handlers {

	aService := auth.NewService(client)

	return &Handlers{
		Ping:           Ping(),
		RegisterAuthor: RegisterAuthor(aService),
		Login:          Login(aService),
		TokenForVisit:  TokenForVisiting(aService),
	}
}
