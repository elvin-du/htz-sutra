package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"htz/sutra/admin-server/rest/server"
	"htz/sutra/common/util/response"
	"log"
	"net/http"
	"os"
)

const FileDstRoot = "." //TODO tmp configable

func init() {
	server.DefaultServer.RegisterRoute("POST", "/post/file/upload/:file_hash/:mime", defaultFileController.Upload)
	server.DefaultServer.RegisterRoute("POST", "/get/file/download/:height/:width", defaultFileController.Download)
}

type FileController struct {
}

var (
	defaultFileController = &FileController{}
)

func (val *FileController) Download(ctx *gin.Context) {

}

/*
curl -X POST http://localhost:9001/post/file/upload/abc123/application-json \
  -F "file=@./aa.txt" \
  -F "sutra_name=论语" \
  -F "item_name=不食嗟来之食" \
  -F "item_number=3" \
  -H "Content-Type: multipart/form-data"
 */
func Dst(sutraName, itemNumber, itemName string) string {
	err := os.MkdirAll(fmt.Sprintf("%s/%s/", FileDstRoot, sutraName), os.ModePerm)
	if nil != err{
		panic(err)
	}
	//example: ./论语/9-不食嗟来之食
	return fmt.Sprintf("%s/%s/%s-%s", FileDstRoot, sutraName, itemNumber, itemName)
}

func (val *FileController) Upload(ctx *gin.Context) {

	fileHash := ctx.Param("file_hash")
	if "" == fileHash {
		ctx.JSON(400, response.New(400, nil, "file_hash empty"))
	}

	mime := ctx.Param("mime")
	if "" == mime {
		ctx.JSON(400, response.New(400, nil, "file_hash empty"))
	}

	sutraName := ctx.PostForm("sutra_name")
	itemNumber := ctx.PostForm("item_number")
	itemName := ctx.PostForm("item_name")

	log.Println(sutraName, itemNumber, itemName)
	// single file
	file, err := ctx.FormFile("file")
	if nil != err {
		ctx.JSON(400, response.New(400, nil, err))
	}

	log.Println(file.Filename)

	// Upload the file to specific dst.
	err = ctx.SaveUploadedFile(file, Dst(sutraName, itemNumber, itemName))
	if nil != err {
		log.Println(Dst(sutraName, itemNumber, itemName))
		ctx.JSON(400, response.New(400, nil, err))
	}

	//TODO check hash ,and save db
	//TODO return file_id

	ctx.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	//address := ctx.Param("hash")
	//addressInfo, err := model.NewAddressesModel().Find(address)
	//if err != nil {
	//	ctx.JSON(200, response.InternalServerError(err.Error()))
	//}
	//if addressInfo == nil {
	//	ctx.JSON(200, response.NotFound(nil))
	//}
	//ctx.JSON(200, response.Ok(addressInfo))
}
