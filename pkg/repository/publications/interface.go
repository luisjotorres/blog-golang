package publications

import "blog/pkg/domain"

type Client interface {
	GetPublications(pageNumber int) ([]*domain.Post, *int64, error)
	GetPost(postId uint) (*domain.Post, error)
	Reactions(postId uint, rt string) (*domain.Reactions, error)
}
