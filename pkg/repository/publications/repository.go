package publications

import (
	"blog/pkg/domain"
	"fmt"
	"gorm.io/gorm"
	"math"
)

func NewClient(db *gorm.DB) *client {
	return &client{
		DB: db,
	}
}

func (c *client) GetPublications(pageNumber int) ([]*domain.Post, *int64, error) {
	var publications []*domain.Post
	var totalPages int64
	if err := c.Scopes(Paginate(pageNumber)).Preload("Reactions").Find(&publications).Error; err != nil {

		return nil, nil, err
	}
	if err := c.Model(&publications).Count(&totalPages).Error; err != nil {
		fmt.Println("Estoy aca")
		return nil, nil, err
	}
	fmt.Println(math.Ceil(float64(totalPages) / float64(10)))
	totalPages = int64(math.Ceil(float64(totalPages) / float64(10)))
	return publications, &totalPages, nil
}

func (c *client) GetPost(postId uint) (*domain.Post, error) {
	postFound := &domain.Post{}

	if err := c.Model(postFound).Preload("Reactions").
		Where("id = ?", postId).First(postFound).Error; err != nil {
		return nil, err
	}
	return postFound, nil

}

func Paginate(pageNumber int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		offset := (pageNumber - 1) * 10
		return db.Offset(offset).Limit(10)
	}
}

func (c *client) Reactions(postId uint, reactionType string) (reactions *domain.Reactions, err error) {
	fmt.Println(postId)
	if err = c.Where("post_id = ?", postId).First(&reactions).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	if reactionType == "like" {
		reactions.Likes += 1
	} else {
		reactions.Views += 1
	}
	c.Save(reactions)
	return reactions, nil
}
