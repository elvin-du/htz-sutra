package controller
//
//import (
//	"github.com/gin-gonic/gin"
//	"htz/sutra/common/model"
//	"htz/sutra/common/util/response"
//	"htz/sutra/api-server/rest/server"
//)
//
//func init() {
//	server.DefaultServer.RegisterRoute("get", "/chain/status", defaultChain.ChainStatus)
//}
//
//type Chain struct {
//}
//
//var (
//	defaultChain = &Chain{}
//)
//
//func (val *Chain) ChainStatus(ctx *gin.Context) {
//	operatorAddress := ctx.Param("operaddress")
//	validator, err := model.NewValidatorModel().Info(operatorAddress)
//	if err != nil {
//		ctx.JSON(200, response.InternalServerError(err.Error()))
//	}
//	if validator == nil {
//		ctx.JSON(200, response.NotFound(nil))
//	}
//	ctx.JSON(200, validator)
//}
