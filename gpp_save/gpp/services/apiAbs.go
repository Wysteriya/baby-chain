package services

import "github.com/gin-gonic/gin"

type API struct {
	SendResponse func(*gin.Context, any, any)
}
