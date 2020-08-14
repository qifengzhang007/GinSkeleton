package uploadFile

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"goskeleton/app/global/consts"
	"goskeleton/app/global/myErrors"
	"goskeleton/app/global/variable"
	"goskeleton/app/utils/md5Encrypt"
	"goskeleton/app/utils/response"
	"goskeleton/app/utils/snowFlake"
	"net/http"
	"path"
)

func Upload(context *gin.Context, save_path string) bool {

	//  1.获取上传的文件名(参数验证器已经验证完成了第一步错误，这里简化)
	file, _ := context.FormFile(variable.UploadFileField) //  file 是一个文件结构体（文件对象）

	//  保存文件，原始文件名进行全局唯一编码加密、md5 加密，保证在后台存储不重复
	var save_err error
	if unique_id, err := snowFlake.CreateSnowFlakeFactory().GetId(); err == nil {
		save_file_name := fmt.Sprintf("%d%s", unique_id, file.Filename)
		save_file_name = md5Encrypt.MD5([]byte(save_file_name)) + path.Ext(save_file_name)
		if save_err = context.SaveUploadedFile(file, save_path+save_file_name); save_err == nil {
			//  上传成功,返回资源的存储路径，这里请根据实际返回绝对路径或者相对路径
			succ := gin.H{
				"path": save_path + save_file_name,
			}
			response.ReturnJson(context, http.StatusCreated, consts.Curd_Status_Ok_Code, consts.Curd_Status_Ok_Msg, succ)
			return true
		}
	} else {
		save_err = errors.New(myErrors.Errors_Snowflake_GetId_Fail)
	}

	response.ReturnJson(context, http.StatusBadRequest, consts.Files_Upload_Fail_Code, consts.Files_Upload_Fail_Msg+", 文件保存失败!", save_err.Error())
	return false

}
