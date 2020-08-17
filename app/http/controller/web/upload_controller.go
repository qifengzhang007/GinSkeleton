package web

import (
	"github.com/gin-gonic/gin"
	"goskeleton/app/global/variable"
	"goskeleton/app/service/upload_file"
)

type Upload struct {
}

//  文件上传是一个独立模块，给任何业务返回文件上传后的存储路径即可。
// 开始上传
func (u *Upload) Start(context *gin.Context) bool {
	savePath := variable.BasePath + variable.UploadFileSavePath
	return upload_file.Upload(context, savePath)
}
