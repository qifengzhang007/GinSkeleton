package UploadFiles

import (
	"GinSkeleton/App/Global/Consts"
	"GinSkeleton/App/Global/Variable"
	"GinSkeleton/App/Http/Controller/Admin"
	"GinSkeleton/App/Utils/Response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UploadAvatar struct {
}

func (u *UploadAvatar) CheckParams(context *gin.Context) {
	_, error := context.FormFile(Variable.UploadFileField) //  file 是一个文件结构体（文件对象）
	if error != nil {
		Response.ReturnJson(context, http.StatusBadRequest, Consts.Files_Upload_Fail_Code, Consts.Files_Upload_Fail_Msg+", 获取上传文件发生错误!", error)
		return
	}
	// 验证完成，调用控制器,并将追加标案参数验证器的上下文传递给控制器
	(&Admin.Users{}).UploadAvatar(context)
}
