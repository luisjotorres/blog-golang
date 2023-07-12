package authors

import (
	"blog/pkg/repository/database"
	"blog/pkg/usecases/authors"
	"github.com/gin-gonic/gin"
)

type Handlers struct {
	UpdateProfile gin.HandlerFunc
	GetProfile    gin.HandlerFunc
}

func InitHandlers(client database.Client) *Handlers {
	aService := authors.NewService(client)

	return &Handlers{
		UpdateProfile: UpdateProfile(aService),
		GetProfile:    GetProfile(aService),
	}
}
