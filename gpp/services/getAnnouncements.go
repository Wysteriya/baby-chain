package services

import (
	"baby-chain/gpp"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetOpenAnnouncements(ctx *gin.Context) {
	openAnnouncements := gpp.Sd.Data["OpenAnnouncements"]
	ctx.IndentedJSON(http.StatusOK, gin.H{"OpenAnnouncements": openAnnouncements})
}
