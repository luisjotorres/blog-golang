package posts

import (
	"blog/pkg/repository/database"
	"blog/pkg/repository/publications"
)

type service struct {
	dbRepository database.Client
	pC           publications.Client
}
