package authors

import (
	"blog/pkg/domain"
	"blog/pkg/repository/database"
	"net/http"
	"strconv"
)

func NewService(dbRepository database.Client) *service {
	return &service{
		dbRepository,
	}
}

func (s *service) UpdateProfile(profile *domain.AuthorProfile, currentUser uint) error {

	currentProfile := domain.AuthorProfile{}
	s.dbRepository.FindOneById(profile.ID, &currentProfile)

	if currentProfile.AuthorID != currentUser {
		return domain.NewAPIError(http.StatusForbidden,
			"You does not have permissions for this action", "INVALID_ACTION")
	}
	s.dbRepository.Update(profile, currentProfile)
	profile = &currentProfile
	return nil
}

func (s *service) GetProfile(currentUser uint) *domain.AuthorProfile {
	profile := &domain.AuthorProfile{}
	s.dbRepository.FindOneBy("author_id", strconv.Itoa(int(currentUser)), profile)
	return profile
}
