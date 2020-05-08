package Web

import (
	"GinSkeleton/App/Global/Variable"
	"GinSkeleton/App/Service/UploadFile"
	"github.com/gin-gonic/gin"
)

type Upload struct {
}

//  文件上传是一个独立模块，给任何业务返回文件上传后的存储路径即可。
// 开始上传
func (u *Upload) Start(context *gin.Context) bool {
	save_path := Variable.BASE_PATH + Variable.UploadFileSavePath
	return UploadFile.Upload(context, save_path)
}
