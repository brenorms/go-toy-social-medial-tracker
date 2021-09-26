package services

import (
	"net/http"
	"strconv"

	"github.com/brenorms/go-toy-social-medial-tracker/entities"
	"github.com/gin-gonic/gin"
)

type SocialMediaPostService struct {
	mockRepo []entities.SocialMediaPost
	currentId int
}

func NewSocialMediaPostService() ISocialMediaPostService {
	return SocialMediaPostService{
		mockRepo: []entities.SocialMediaPost{
			{Id: 1, Title: "Video 01", Views: 101, Likes: 9},
			{Id: 2, Title: "Video 15", Views: 1501, Likes: 95},
			{Id: 3, Title: "Video 18", Views: 15148, Likes: 907},
			{Id: 4, Title: "Video 32", Views: 55, Likes: 21},
		},
		currentId: 5,
	}
}
func (s SocialMediaPostService) List(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, s.mockRepo)
}

func (s SocialMediaPostService) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "post not found"})
		return
	}

	for _, post := range s.mockRepo {
		if post.Id == id {
			c.IndentedJSON(http.StatusOK, post)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "post not found"})
}

func (s SocialMediaPostService) Create(c *gin.Context) {
	var newPost entities.SocialMediaPost

	if err := c.BindJSON(&newPost); err != nil {
		return
	}

	newPost.Id = s.currentId
	s.currentId += 1

	s.mockRepo = append(s.mockRepo, newPost)
	c.IndentedJSON(http.StatusCreated, newPost)
}
