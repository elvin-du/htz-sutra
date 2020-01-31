package controller

import (
	"github.com/gin-gonic/gin"
	"htz/sutra/common/server"
	"htz/sutra/search-server/search"
	"net/http"
)

func init() {
	server.DefaultServer.RegisterRoute("POST", "/get/search", defaultSearcher.Search)
}

var (
	defaultSearcher = &Searcher{}
)

type Searcher struct {
}

func (*Searcher) Search(ctx *gin.Context) {
	keyMap := struct {
		Key          string `json:"key"`
		OutputOffset int    `json:"output_offset"`
		MaxOutputs   int    `json:"max_outputs"`
	}{}
	if err := ctx.ShouldBindJSON(&keyMap); nil != err {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	results := search.Search(keyMap.Key, keyMap.OutputOffset, keyMap.MaxOutputs)
	ctx.JSON(http.StatusOK, gin.H{"results": results})
}
