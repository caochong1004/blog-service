package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/longjoy/blog-service/internal/service"
	"github.com/longjoy/blog-service/pkg/app"
	"github.com/longjoy/blog-service/pkg/convert"
	"github.com/longjoy/blog-service/pkg/errcode"
	"github.com/longjoy/blog-service/pkg/upload"
)

type Upload struct {}

func NewUpload() Upload  {
	return Upload{}
}

func (u Upload) UploadFile(c *gin.Context)  {
	response := app.NewResponse(c)
	file, fileHeader, err := c.Request.FormFile("file")
	fileType := convert.StrTo(c.PostForm("type")).MustInt()
	if err != nil {
		errRsp := errcode.InvalidParams.WithDetails(err.Error())
		response.ToErrorResponse(errRsp)
		return
	}
	if fileHeader == nil || fileType <= 0 {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}

	svc := service.New(c.Request.Context())
	fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)
	if err != nil {
		errRsp := errcode.ErrorUploadFileFail.WithDetails(err.Error())
		response.ToErrorResponse(errRsp)
		return
	}

	response.ToResponse(gin.H{
		"file_access_url" :fileInfo.AccessUrl,
	})
}
