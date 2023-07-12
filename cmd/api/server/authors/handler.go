package authors

import (
	"blog/pkg/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateProfile(service AuthorService) gin.HandlerFunc {
	return func(context *gin.Context) {

		profile := &domain.AuthorProfile{}
		err := context.Bind(profile)
		if err != nil {
			return
		}

		if err := service.UpdateProfile(profile, context.GetUint("userId")); err != nil {
			context.Error(err)
		}

		context.JSON(http.StatusOK, gin.H{
			"description":     profile.Description,
			"profile_picture": profile.ProfilePicture,
		})
	}
}

func GetProfile(service AuthorService) gin.HandlerFunc {
	return func(context *gin.Context) {
		currentUser := context.GetUint("userId")
		profile := service.GetProfile(currentUser)
		context.JSON(http.StatusOK, gin.H{
			"id":              profile.ID,
			"description":     profile.Description,
			"profile_picture": profile.ProfilePicture,
		})
	}
}
