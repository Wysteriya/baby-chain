package services

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HttpResponse struct {
	ctx *gin.Context
}

func newHttpResponse(ctx *gin.Context) *HttpResponse {
	return &HttpResponse{ctx}
}
func (res *HttpResponse) Error(err error) {
	res.ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	return
}

func (res *HttpResponse) Text(message string) {
	res.ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": message})
	return
}

func (res *HttpResponse) BindJson(responseObj any) {
	if err := res.ctx.BindJSON(responseObj); err != nil {
		res.ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
}

func (res *HttpResponse) SendJson(returnObj any) {
	res.ctx.IndentedJSON(http.StatusOK, returnObj)
	return
}
