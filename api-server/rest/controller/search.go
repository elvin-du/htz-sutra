package controller

import (
	"github.com/gin-gonic/gin"
	"htz/sutra/api-server/rest/server"
)

func init() {
	server.DefaultServer.RegisterRoute("get", "/get/search", defaultSearcher.Search)
}

var (
	defaultSearcher = &Searcher{}
)

type Searcher struct {
}

func (val *Searcher) Search(ctx *gin.Context) {
}
