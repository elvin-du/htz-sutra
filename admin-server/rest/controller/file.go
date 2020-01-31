package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"htz/sutra/admin-server/config"
	"htz/sutra/common/server"
	"htz/sutra/common/util/response"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func init() {
	server.DefaultServer.RegisterRoute("POST", "/post/file/upload", defaultFileController.Upload)
	server.DefaultServer.RegisterRoute("POST", "/get/file/download", defaultFileController.Download)
}

type FileController struct {
}

var (
	defaultFileController = &FileController{}
)

func (val *FileController) Download(ctx *gin.Context) {
	tmp := struct {
		FileID string `json:"file_id"`
		Height int64  `json:"height"`
		Width  int64  `json:"width"`
	}{}

	err := ParseBody(ctx, &tmp)
	if nil != err {
		ctx.JSON(404, response.Fail(400, err))
	}

	ctx.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", tmp.FileID))
	//fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
	ctx.Writer.Header().Add("Content-Type", "application/octet-stream")

	//TODO get file path
	p := "./论语/3-不食嗟来之食.mp3"
	ctx.File(p)
}

/*
curl -X POST http://localhost:9001/post/file/upload \
  -F "file=@./aa.txt" \
  -F "sutra_name=论语" \
  -F "item_number=3" \
  -F "item_name=不食嗟来之食" \
  -F "item_suffix=mp3" \
  -F "file_hash=a8bc8" \
  -F "mime=image-jpeg" \
  -H "Content-Type: multipart/form-data"
 */
func Dst(sutraName, itemNumber, itemName, itemSuffix string) string {
	err := os.MkdirAll(fmt.Sprintf("%s/%s/", config.DefaultConfig.FileDBPath, sutraName), os.ModePerm)
	if nil != err {
		panic(err)
	}
	//example: ./论语/9-不食嗟来之食.mp3
	return fmt.Sprintf("%s/%s/%s-%s.%s", config.DefaultConfig.FileDBPath, sutraName, itemNumber, itemName, itemSuffix)
}

func ParseBody(ctx *gin.Context, obj interface{}) error {
	bin, err := ioutil.ReadAll(ctx.Request.Body)
	if nil != err {
		return err
	}

	err = json.Unmarshal(bin, obj)
	if nil != err {
		return err
	}

	return nil
}

func (val *FileController) Upload(ctx *gin.Context) {
	sutraName := ctx.PostForm("sutra_name")
	itemNumber := ctx.PostForm("item_number")
	itemName := ctx.PostForm("item_name")
	itemSuffix := ctx.PostForm("item_suffix")
	fileHash := ctx.PostForm("file_hash")

	if "" == fileHash {
		ctx.JSON(400, response.Fail(400, "file_hash empty"))
	}

	mime := ctx.PostForm("mime")
	if "" == mime {
		ctx.JSON(400, response.Fail(400, "file_hash empty"))
	}

	log.Println(sutraName, itemNumber, itemName, fileHash, mime)
	// single file
	file, err := ctx.FormFile("file")
	if nil != err {
		ctx.JSON(400, response.Fail(400, err))
	}

	log.Println(file.Filename)

	// Upload the file to specific dst.
	err = ctx.SaveUploadedFile(file, Dst(sutraName, itemNumber, itemName, itemSuffix))
	if nil != err {
		log.Println(Dst(sutraName, itemNumber, itemName, itemSuffix))
		ctx.JSON(400, response.Fail(400, err))
	}

	//TODO check hash ,and save db
	//TODO return file_id

	tmp := struct {
		FileID string `json:"file_id"`
	}{""}
	ctx.JSON(http.StatusOK, response.Ok(tmp))

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
