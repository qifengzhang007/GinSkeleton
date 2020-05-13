package UploadFiles

import (
	"GinSkeleton/App/Global/Consts"
	"GinSkeleton/App/Global/Variable"
	"GinSkeleton/App/Http/Controller/Web"
	"GinSkeleton/App/Utils/Config"
	"GinSkeleton/App/Utils/Files"
	"GinSkeleton/App/Utils/Response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type UpFiels struct {
}

// 文件上传公共模块表单参数验证器
func (u UpFiels) CheckParams(context *gin.Context) {
	tmp_file, error := context.FormFile(Variable.UploadFileField) //  file 是一个文件结构体（文件对象）
	var is_pass bool
	//获取文件发生错误，可能上传了空文件等
	if error != nil {
		Response.ReturnJson(context, http.StatusBadRequest, Consts.Files_Upload_Fail_Code, Consts.Files_Upload_Fail_Msg+", 获取上传文件发生错误!", error)
		return
	}
	//超过系统设定的最大值：32M
	if tmp_file.Size > Config.CreateYamlFactory().GetInt64("FileUploadSetting.Size")*1024*1024 {
		Response.ReturnJson(context, http.StatusBadRequest, Consts.Files_Upload_MoreThan_Max_Size_Code, Consts.Files_Upload_MoreThan_Max_Size_Msg+",系统允许的最大值（M）："+Config.CreateYamlFactory().GetString("FileUploadSetting.Size"), "")
		return
	}
	//不允许的文件mime类型
	if fp, err := tmp_file.Open(); err == nil {
		mimeType := Files.GetFilesMimeByFp(fp)

		for _, value := range Config.CreateYamlFactory().GetStringSlice("FileUploadSetting.AllowMimeType") {
			if strings.ReplaceAll(value, " ", "") == strings.ReplaceAll(mimeType, " ", "") {
				is_pass = true
				break
			}
		}
		fp.Close()
	} else {
		Response.ReturnJson(context, http.StatusBadRequest, Consts.Server_Occurred_Error_Code, Consts.Server_Occurred_Error_Msg+",检测文件mime类型发生错误。", "")
		return
	}
	//凡是存在相等的类型，通过验证，调用控制器
	if !is_pass {
		Response.ReturnJson(context, http.StatusBadRequest, Consts.Files_Upload_MimeType_Fail_Code, Consts.Files_Upload_MimeType_Fail_Msg, "")
	} else {
		(&Web.Upload{}).Start(context)
	}
}
