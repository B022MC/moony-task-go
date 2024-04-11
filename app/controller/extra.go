package controller

import (
	"github.com/gin-gonic/gin"
	"moony-task-go/app/service"
	"moony-task-go/core/common"
	"moony-task-go/core/config"
	"net/http"
)

type Extra struct {
}

func ExtraInstance() *Extra {
	return &Extra{}
}

// Upload 文件上传
func (e *Extra) Upload(ctx *gin.Context) {
	file, _ := ctx.FormFile("file")
	if file == nil {
		panic(config.ErrParam.New().Append("未获取到文件信息"))
	}

	url, err := service.UploadInstance().UploadFile(ctx, file)
	if err != nil {
		panic(config.ErrInternal.New().Append(err))
	}
	ctx.JSON(http.StatusOK, common.NewRsp(url))
}
