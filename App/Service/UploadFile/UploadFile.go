package UploadFile

import (
	"GinSkeleton/App/Global/Consts"
	"GinSkeleton/App/Global/MyErrors"
	"GinSkeleton/App/Global/Variable"
	"GinSkeleton/App/Utils/MD5Cryt"
	"GinSkeleton/App/Utils/Response"
	"GinSkeleton/App/Utils/SnowFlake"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

func Upload(context *gin.Context, save_path string) bool {

	//  1.获取上传的文件名(参数验证器已经验证完成了第一步错误，这里简化)
	file, _ := context.FormFile(Variable.UploadFileField) //  file 是一个文件结构体（文件对象）

	//  保存文件，原始文件名进行全局唯一编码加密、md5 加密，保证在后台存储不重复
	var save_err error
	if unique_id, err := SnowFlake.CreateSnowFlakeFactory().GetId(); err == nil {
		save_file_name := fmt.Sprintf("%d%s", unique_id, file.Filename)
		save_file_name = MD5Cryt.MD5([]byte(save_file_name)) + path.Ext(save_file_name)
		if save_err = context.SaveUploadedFile(file, save_path+save_file_name); save_err == nil {
			//  上传成功,返回资源的存储路径，这里请根据实际返回绝对路径或者相对路径
			succ := gin.H{
				"path": save_path + save_file_name,
			}
			Response.ReturnJson(context, http.StatusCreated, Consts.Curd_Status_Ok_Code, Consts.Curd_Status_Ok_Msg, succ)
			return true
		}
	} else {
		save_err = errors.New(MyErrors.Errors_Snowflake_GetId_Fail)
	}

	Response.ReturnJson(context, http.StatusBadRequest, Consts.Files_Upload_Fail_Code, Consts.Files_Upload_Fail_Msg+", 文件保存失败!", save_err.Error())
	return false

}
