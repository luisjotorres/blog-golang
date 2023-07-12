package posts

import "blog/pkg/domain"

type PostService interface {
	PublishPost(post *domain.Post) (uint, error)
	GetPostByPage(page int) ([]*domain.Post, *int64, error)
	ReadPost(postId uint) (*domain.Post, error)
	Reactions(postId uint, rt string) (*domain.Reactions, error)
}
