package params

import (
	"github.com/gin-gonic/gin"
	"htz/sutra/common/util"
	"strconv"
)

func ExtractPage(ctx *gin.Context) util.Page {
	index, err := strconv.Atoi(ctx.Param("index"))
	if err != nil {

	}
	pageSize, err := strconv.Atoi(ctx.Param("pageSize"))
	return util.Page{int32(index), int32(pageSize)}
}
