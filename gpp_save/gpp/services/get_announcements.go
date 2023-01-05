package services

import (
	"baby-chain/gpp/chain"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetOpenAnnouncements(ctx *gin.Context) {
	stateData := chain.LoadStateData()
	openAnnouncements := stateData.Data["open_announcements"]

	ctx.IndentedJSON(http.StatusOK, gin.H{"open_announcements": openAnnouncements})

}
