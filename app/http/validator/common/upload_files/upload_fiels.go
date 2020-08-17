package upload_files

import (
	"github.com/gin-gonic/gin"
	"goskeleton/app/global/consts"
	"goskeleton/app/global/variable"
	"goskeleton/app/http/controller/web"
	"goskeleton/app/utils/files"
	"goskeleton/app/utils/response"
	"goskeleton/app/utils/yml_config"
	"net/http"
	"strings"
)

type UpFiles struct {
}

// 文件上传公共模块表单参数验证器
func (u UpFiles) CheckParams(context *gin.Context) {
	tmpFile, err := context.FormFile(variable.UploadFileField) //  file 是一个文件结构体（文件对象）
	var isPass bool
	//获取文件发生错误，可能上传了空文件等
	if err != nil {
		response.ReturnJson(context, http.StatusBadRequest, consts.FilesUploadFailCode, consts.FilesUploadFailMsg+", 获取上传文件发生错误!", err.Error())
		return
	}
	//超过系统设定的最大值：32M
	if tmpFile.Size > yml_config.CreateYamlFactory().GetInt64("FileUploadSetting.Size")*1024*1024 {
		response.ReturnJson(context, http.StatusBadRequest, consts.FilesUploadMoreThanMaxSizeCode, consts.FilesUploadMoreThanMaxSizeMsg+",系统允许的最大值（M）："+yml_config.CreateYamlFactory().GetString("FileUploadSetting.Size"), "")
		return
	}
	//不允许的文件mime类型
	if fp, err := tmpFile.Open(); err == nil {
		mimeType := files.GetFilesMimeByFp(fp)

		for _, value := range yml_config.CreateYamlFactory().GetStringSlice("FileUploadSetting.AllowMimeType") {
			if strings.ReplaceAll(value, " ", "") == strings.ReplaceAll(mimeType, " ", "") {
				isPass = true
				break
			}
		}
		_ = fp.Close()
	} else {
		response.ReturnJson(context, http.StatusBadRequest, consts.ServerOccurredErrorCode, consts.ServerOccurredErrorMsg+",检测文件mime类型发生错误。", "")
		return
	}
	//凡是存在相等的类型，通过验证，调用控制器
	if !isPass {
		response.ReturnJson(context, http.StatusBadRequest, consts.FilesUploadMimeTypeFailCode, consts.FilesUploadMimeTypeFailMsg, "")
	} else {
		(&web.Upload{}).Start(context)
	}
}
