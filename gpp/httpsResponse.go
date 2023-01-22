package gpp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HttpResponse struct {
	ctx *gin.Context
}

func NewHttpResponse(ctx *gin.Context) *HttpResponse {
	return &HttpResponse{ctx}
}

func (res *HttpResponse) Error(err error) {
	if err != nil {
		res.ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func (res *HttpResponse) Text(message string) {
	res.ctx.IndentedJSON(http.StatusOK, gin.H{"message": message})
}

func (res *HttpResponse) BindJson(obj any) error {
	if err := res.ctx.BindJSON(obj); err != nil {
		return err
	}
	return nil
}

func (res *HttpResponse) SendJson(sendObj any) {
	res.ctx.IndentedJSON(http.StatusOK, sendObj)
}
