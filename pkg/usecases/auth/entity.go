package auth

import "blog/pkg/repository/database"

type service struct {
	dbRepository database.Client
}
