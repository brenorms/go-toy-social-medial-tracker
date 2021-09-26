//
// Built following the guide at
// https://golang.org/doc/tutorial/web-service-gin
//

package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type socialMediaPost struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Views int    `json:"views"`
	Likes int    `json:"likes"`
}

var socialMediaPosts = []socialMediaPost{
	{Id: 1, Title: "Video 01", Views: 101, Likes: 9},
	{Id: 2, Title: "Video 15", Views: 1501, Likes: 95},
	{Id: 3, Title: "Video 18", Views: 15148, Likes: 907},
	{Id: 4, Title: "Video 32", Views: 55, Likes: 21},
}
var socialMediaPostId = 5

func main() {
	router := gin.Default()
	router.GET("/socialMediaPosts", getSocialMediaPosts)
	router.GET("/socialMediaPosts/:id", getSocialMediaPostById)
	router.POST("/socialMediaPosts", postSocialMediaPost)

	router.Run("localhost:8182")
}

func getSocialMediaPosts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, socialMediaPosts)
}

func postSocialMediaPost(c *gin.Context) {
	var newPost socialMediaPost

	if err := c.BindJSON(&newPost); err != nil {
		return
	}

	newPost.Id = socialMediaPostId
	socialMediaPostId += 1

	socialMediaPosts = append(socialMediaPosts, newPost)
	c.IndentedJSON(http.StatusCreated, newPost)
}

func getSocialMediaPostById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "post not found"})
		return
	}

	for _, post := range socialMediaPosts {
		if post.Id == id {
			c.IndentedJSON(http.StatusOK, post)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "post not found"})
}
