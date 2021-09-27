package services

import (
	"net/http"
	"strconv"

	"github.com/brenorms/go-toy-social-medial-tracker/entities"
	"github.com/brenorms/go-toy-social-medial-tracker/repositories"
	"github.com/gin-gonic/gin"
)

type SocialMediaPostService struct {
	repo repositories.ISocialMediaPostRepository
}

func NewSocialMediaPostService() ISocialMediaPostService {
	return SocialMediaPostService{
		repo: repositories.NewSocialMediaPostRepository(),
	}
}
func (s SocialMediaPostService) List(c *gin.Context) {
	// c.IndentedJSON(http.StatusOK, s.mockRepo)

	arr, err := s.repo.List()

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "error listing posts"})
		return
	}

	c.IndentedJSON(http.StatusOK, arr)
}

func (s SocialMediaPostService) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "post not found"})
		return
	}

	post, err := s.repo.GetById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "post not found"})
		return
	}
	if post.Id == id {
		c.IndentedJSON(http.StatusOK, post)
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "post not found"})
}

func (s SocialMediaPostService) Create(c *gin.Context) {
	var newPost entities.SocialMediaPost

	if err := c.BindJSON(&newPost); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "could not create"})
		return
	}

	if err := s.repo.Create(&newPost); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "could not create"})
		return
	}

	c.IndentedJSON(http.StatusCreated, newPost)
}

func (s SocialMediaPostService) Update(c *gin.Context) {
	var newPost entities.SocialMediaPost

	if err := c.BindJSON(&newPost); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "could not update"})
		return
	}

	if err := s.repo.Update(newPost); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "could not update"})
		return
	}

	c.IndentedJSON(http.StatusCreated, newPost)
}

func (s SocialMediaPostService) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "could not delete"})
		return
	}

	err = s.repo.Delete(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "could not delete"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{})
}
