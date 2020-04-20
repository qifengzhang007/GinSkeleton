package UploadFile

import (
	"GinSkeleton/App/Global/Consts"
	"GinSkeleton/App/Global/Variable"
	"GinSkeleton/App/Utils/Response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Upload(context *gin.Context, save_path string, file_name string) bool {

	//  1.获取上传的文件名(参数验证器已经验证完成了第一步错误，这里简化)
	file, _ := context.FormFile(Variable.UploadFileField) //  file 是一个文件结构体（文件对象）
	var sava_file_name string
	if len(file_name) > 0 {
		sava_file_name = file_name
	} else {
		sava_file_name = file.Filename
	}
	//  保存文件
	if err := context.SaveUploadedFile(file, save_path+sava_file_name); err != nil {
		Response.ReturnJson(context, http.StatusBadRequest, Consts.Files_Upload_Fail_Code, Consts.Files_Upload_Fail_Msg+", 文件保存失败!", err.Error())
		return false
	}
	//  上传成功,返回资源相对项目站点的存储路径
	succ := gin.H{
		"path": save_path + sava_file_name,
	}
	Response.ReturnJson(context, http.StatusCreated, Consts.Curd_Status_Ok_Code, Consts.Curd_Status_Ok_Msg, succ)
	return true

}
