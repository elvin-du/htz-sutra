package controller

import (
	"github.com/gin-gonic/gin"
	"htz/sutra/common/server"
	"htz/sutra/common/types"
	"htz/sutra/search-server/search"
	"net/http"
)

func init() {
	server.DefaultServer.RegisterRoute("POST", "/post/doc", defaultSearcher.IndexDoc)
	server.DefaultServer.RegisterRoute("POST", "/get/search", defaultSearcher.Search)
}

var (
	defaultSearcher = &Searcher{}
)

type Searcher struct{}

func (*Searcher) IndexDoc(ctx *gin.Context) {
	item := types.SutraItem{}
	if err := ctx.ShouldBindJSON(&item); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"json error": err.Error()})
		return
	}

	itemIndex := search.SutraItem{}
	itemIndex.ID = item.ID
	itemIndex.Description = item.Description
	itemIndex.Explanation = item.Explanation
	itemIndex.Original = item.Original
	itemIndex.Title = item.Title

	search.Index(&itemIndex, []string{"黄庭禅", "庄子"})
	ctx.JSON(http.StatusOK, gin.H{"msg": "index success"})
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
