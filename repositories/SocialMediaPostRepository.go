package repositories

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/brenorms/go-toy-social-medial-tracker/entities"
)

type SocialMediaPostRepository struct {
}

func NewSocialMediaPostRepository() ISocialMediaPostRepository {
	return SocialMediaPostRepository{}
}

func (r SocialMediaPostRepository) List() ([]entities.SocialMediaPost, error) {
	db, err := GetDb()
	if err != nil {
		return nil, err
	}
	query := sq.Select("*").From("socialmediapost")
	sql, _, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	posts := make([]entities.SocialMediaPost, 0)
	err = db.Select(&posts, sql)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (r SocialMediaPostRepository) GetById(id int) (entities.SocialMediaPost, error) {
	db, err := GetDb()
	if err != nil {
		return entities.SocialMediaPost{}, err
	}
	query := sq.Select("*").From("socialmediaposts").Where(sq.Eq{"id": id})
	sql, args, err := query.ToSql()
	if err != nil {
		return entities.SocialMediaPost{}, err
	}
	var post entities.SocialMediaPost
	db.Get(&post, sql, args)

	return post, nil
}

func (r SocialMediaPostRepository) Create(post entities.SocialMediaPost) error {
	return nil
}
