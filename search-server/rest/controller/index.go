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
	server.DefaultServer.RegisterRoute("POST", "/delete/doc", defaultSearcher.Delete)
}

var (
	defaultSearcher = &Searcher{}
)

type Searcher struct{}

func (*Searcher) IndexDoc(ctx *gin.Context) {
	item := types.SutraItem{}
	if err := ctx.ShouldBindJSON(&item); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"json parameter error": err.Error()})
		return
	}

	itemIndex := search.SutraItem{}
	itemIndex.ID = item.ID
	itemIndex.Description = item.Description
	itemIndex.Explanation = item.Explanation
	itemIndex.Original = item.Original
	itemIndex.Title = item.Title

	search.Index(&itemIndex)
	ctx.JSON(http.StatusOK, gin.H{"msg": "add index success"})
}

func (*Searcher) Search(ctx *gin.Context) {
	keyParameter := struct {
		Key          string `json:"key"`
		OutputOffset int    `json:"output_offset"`
		MaxOutputs   int    `json:"max_outputs"`
	}{}
	if err := ctx.ShouldBindJSON(&keyParameter); nil != err {
		ctx.JSON(http.StatusBadRequest, gin.H{"json parameter error": err.Error()})
		return
	}

	results := search.Search(keyParameter.Key, keyParameter.OutputOffset, keyParameter.MaxOutputs)
	ctx.JSON(http.StatusOK, gin.H{"results": results})
}

func (*Searcher) Delete(ctx *gin.Context) {
	param := struct {
		ID string `json:"id"`
	}{}
	if err := ctx.ShouldBindJSON(&param); nil != err {
		ctx.JSON(http.StatusBadRequest, gin.H{"json parameter error": err.Error()})
		return
	}
	search.Remove(param.ID)
	ctx.JSON(http.StatusOK, gin.H{"msg": "remove index success"})
}
