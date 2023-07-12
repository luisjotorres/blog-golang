package posts

import (
	"blog/pkg/repository/database"
	"blog/pkg/repository/publications"
	"blog/pkg/usecases/posts"
	"github.com/gin-gonic/gin"
)

type Handlers struct {
	CreatePost gin.HandlerFunc
	Index      gin.HandlerFunc
	ReadPost   gin.HandlerFunc
	Reactions  gin.HandlerFunc
}

func InitHandlers(client database.Client) *Handlers {
	pC := publications.NewClient(client.GetGORMClient())
	ps := posts.NewService(client, pC)

	return &Handlers{
		CreatePost: CreatePost(ps),
		Index:      Index(ps),
		ReadPost:   ReadPost(ps),
		Reactions:  Reactions(ps),
	}
}
