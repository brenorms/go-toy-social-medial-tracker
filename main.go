//
// Built following the guide at
// https://golang.org/doc/tutorial/web-service-gin
//

package main

import (
	"github.com/brenorms/go-toy-social-medial-tracker/services"
	"github.com/gin-gonic/gin"
)

var socialMediaService = services.NewSocialMediaPostService()

func main() {
	router := gin.Default()
	router.GET("/socialMediaPosts", socialMediaService.List)
	router.GET("/socialMediaPosts/:id", socialMediaService.GetById)
	router.POST("/socialMediaPosts", socialMediaService.Create)

	router.Run("localhost:8182")
}
