package auth

import "blog/pkg/domain"

type AuthenticationService interface {
	Register(author domain.Author) error
	Login(email, password, ipAddress string) (string, error)
	TokenForVisiting(ipAddress string) string
}
