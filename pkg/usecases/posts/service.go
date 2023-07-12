package posts

import (
	"blog/pkg/domain"
	"blog/pkg/repository/database"
	"blog/pkg/repository/publications"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup

func NewService(dbR database.Client, pC publications.Client) *service {
	return &service{
		dbRepository: dbR,
		pC:           pC,
	}
}

func (s *service) PublishPost(post *domain.Post) (postId uint, err error) {
	s.dbRepository.Save(post)
	fmt.Printf("%+v\n", post)

	//Crearemos la demas info para guardarla :D
	wg.Add(2)

	go s.createReactionsPayload(post.ID)
	go s.notifyPublish(post.ID)

	wg.Wait()

	return post.ID, err
}

func (s *service) GetPostByPage(page int) ([]*domain.Post, *int64, error) {
	posts, totalPages, err := s.pC.GetPublications(page)
	if err != nil {
		fmt.Println("estoy aca")
		return nil, nil, err
	}
	return posts, totalPages, nil
}

func (s *service) ReadPost(postId uint) (*domain.Post, error) {
	post, err := s.pC.GetPost(postId)
	if err != nil {
		if errors.As(err, &gorm.ErrRecordNotFound) {
			return nil, domain.NewAPIError(http.StatusNotFound, "Post Not found", "NOT_FOUND")
		}
		return nil, err
	}
	return post, nil
}

func (s *service) Reactions(postId uint, rt string) (*domain.Reactions, error) {
	fmt.Println("El PostId", postId)
	reactions, err := s.pC.Reactions(postId, rt)
	if err != nil {
		if errors.As(err, &gorm.ErrRecordNotFound) {
			return nil, domain.NewAPIError(http.StatusBadRequest, "Post not found", "POST_NOT_FOUND")
		}
		return nil, err
	}

	return reactions, nil
}

func (s *service) createReactionsPayload(postId uint) {
	defer wg.Done()
	s.dbRepository.Save(&domain.Reactions{PostID: postId})
}

func (s *service) notifyPublish(postId uint) {
	defer wg.Done()
	time.Sleep(10 * time.Millisecond)
	fmt.Println("Notify Publish")
}
