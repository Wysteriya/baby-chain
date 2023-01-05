package services

import (
	"baby-chain/gpp"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetOpenAnnouncements(ctx *gin.Context) {
	gpp.FetchHyperParams()
	openAnnouncements := gpp.Sd.Data["open_announcements"]

	ctx.IndentedJSON(http.StatusOK, gin.H{"open_announcements": openAnnouncements})

}
