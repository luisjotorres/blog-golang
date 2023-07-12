package auth

import (
	"blog/pkg/domain"
	"blog/pkg/repository/database"
	"blog/pkg/utils"
	"fmt"
	"net/http"
)

func NewService(dbRepository database.Client) *service {
	return &service{
		dbRepository,
	}
}

func (s *service) Register(author domain.Author) error {
	s.dbRepository.FindOneBy("email", author.Email, &author)

	if author.ID != 0 {
		return domain.NewAPIError(http.StatusBadRequest, "User register", "USER_REGISTER")
	}
	s.dbRepository.Save(&author)

	//Creando el profile
	profile := domain.AuthorProfile{
		AuthorID: author.ID,
	}

	s.dbRepository.Save(&profile)
	fmt.Printf("El id registrado es: %d", author.ID)
	return nil
}

func (s *service) Login(email, password, ipAddress string) (string, error) {
	author := domain.Author{}
	s.dbRepository.FindOneBy("email", email, &author)

	if author.ID == 0 {
		return "", domain.NewAPIError(http.StatusNotFound,
			"Email not registered", "EMAIL_NOT_REGISTERED")
	}

	if !author.CheckPassword(password) {
		return "", domain.NewAPIError(http.StatusUnauthorized,
			"Error with your credentials", "INVALID_CREDENTIALS")
	}

	return utils.JwtBuilder("author", author.Email, author.ID), nil
}

func (s *service) TokenForVisiting(ipAddress string) string {

	return utils.JwtBuilder("visit", ipAddress, 0)
}
