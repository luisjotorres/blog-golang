package authors

import "blog/pkg/domain"

type AuthorService interface {
	UpdateProfile(profile *domain.AuthorProfile, currentUser uint) error
	GetProfile(currentUser uint) *domain.AuthorProfile
}
