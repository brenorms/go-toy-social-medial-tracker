package repositories

import "github.com/brenorms/go-toy-social-medial-tracker/entities"

type ISocialMediaPostRepository interface {
	List() ([]entities.SocialMediaPost, error)
	GetById(id int) (entities.SocialMediaPost, error)
	Create(post *entities.SocialMediaPost) error
}
