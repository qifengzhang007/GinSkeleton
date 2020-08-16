package upload_files

import (
	"github.com/gin-gonic/gin"
	"goskeleton/app/global/consts"
	"goskeleton/app/global/variable"
	"goskeleton/app/http/controller/web"
	"goskeleton/app/utils/config"
	"goskeleton/app/utils/files"
	"goskeleton/app/utils/response"
	"net/http"
	"strings"
)

type UpFiels struct {
}

// 文件上传公共模块表单参数验证器
func (u UpFiels) CheckParams(context *gin.Context) {
	tmp_file, error := context.FormFile(variable.UploadFileField) //  file 是一个文件结构体（文件对象）
	var is_pass bool
	//获取文件发生错误，可能上传了空文件等
	if error != nil {
		response.ReturnJson(context, http.StatusBadRequest, consts.Files_Upload_Fail_Code, consts.Files_Upload_Fail_Msg+", 获取上传文件发生错误!", error)
		return
	}
	//超过系统设定的最大值：32M
	if tmp_file.Size > config.CreateYamlFactory().GetInt64("FileUploadSetting.Size")*1024*1024 {
		response.ReturnJson(context, http.StatusBadRequest, consts.Files_Upload_MoreThan_Max_Size_Code, consts.Files_Upload_MoreThan_Max_Size_Msg+",系统允许的最大值（M）："+config.CreateYamlFactory().GetString("FileUploadSetting.Size"), "")
		return
	}
	//不允许的文件mime类型
	if fp, err := tmp_file.Open(); err == nil {
		mimeType := files.GetFilesMimeByFp(fp)

		for _, value := range config.CreateYamlFactory().GetStringSlice("FileUploadSetting.AllowMimeType") {
			if strings.ReplaceAll(value, " ", "") == strings.ReplaceAll(mimeType, " ", "") {
				is_pass = true
				break
			}
		}
		fp.Close()
	} else {
		response.ReturnJson(context, http.StatusBadRequest, consts.Server_Occurred_Error_Code, consts.Server_Occurred_Error_Msg+",检测文件mime类型发生错误。", "")
		return
	}
	//凡是存在相等的类型，通过验证，调用控制器
	if !is_pass {
		response.ReturnJson(context, http.StatusBadRequest, consts.Files_Upload_MimeType_Fail_Code, consts.Files_Upload_MimeType_Fail_Msg, "")
	} else {
		(&web.Upload{}).Start(context)
	}
}
