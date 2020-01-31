package controller

import (
	"github.com/gin-gonic/gin"
	"htz/sutra/common/server"
	"htz/sutra/common/types"
	"htz/sutra/search-server/search"
	"net/http"
)

func init() {
	server.DefaultServer.RegisterRoute("POST", "/post/sutra", defaultSutra.AddSutra)
	server.DefaultServer.RegisterRoute("POST", "/post/sutra/item", defaultSutra.AddSutraItem)
}

var (
	defaultSutra = &Sutra{}
)

type Sutra struct {
}

func (s *Sutra) AddSutra(ctx *gin.Context) {
}

func (s *Sutra) AddSutraItem(ctx *gin.Context) {
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
}
