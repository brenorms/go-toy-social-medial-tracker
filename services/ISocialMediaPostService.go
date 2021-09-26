package services

import (
	"github.com/gin-gonic/gin"
)

type ISocialMediaPostService interface {
	List(ctx *gin.Context)
	GetById(ctx *gin.Context)
	Create(ctx *gin.Context)
}
