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
	query := sq.Select("*").From("socialmediapost").Where(sq.Eq{"id": id})
	sql, args, err := query.ToSql()
	if err != nil {
		return entities.SocialMediaPost{}, err
	}
	var post entities.SocialMediaPost
	err = db.Get(&post, sql, args...)
	if err != nil {
		return entities.SocialMediaPost{}, err
	}

	return post, nil
}

func (r SocialMediaPostRepository) Create(post *entities.SocialMediaPost) error {
	db, err := GetDb()
	if err != nil {
		return err
	}
	query := sq.Insert("socialmediapost").Columns("title", "views", "likes").Values(post.Title, post.Views, post.Likes)
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	res := db.MustExec(sql, args...)
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	post.Id = int(id)

	return nil
}

func (r SocialMediaPostRepository) Update(post entities.SocialMediaPost) error {
	db, err := GetDb()
	if err != nil {
		return err
	}
	query := sq.Update("socialmediapost").SetMap(map[string]interface{}{"title": post.Title, "views": post.Views, "likes": post.Likes}).Where(sq.Eq{"id": post.Id})
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	res := db.MustExec(sql, args...)
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	post.Id = int(id)

	return nil
}

func (r SocialMediaPostRepository) Delete(id int) error {
	db, err := GetDb()
	if err != nil {
		return err
	}
	query := sq.Delete("socialmediapost").Where(sq.Eq{"id": id})
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	res := db.MustExec(sql, args...)
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}
