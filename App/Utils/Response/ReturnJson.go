package Response

import (
	"github.com/gin-gonic/gin"
)

func ReturnJson(Context *gin.Context, http_code int, data_code int, msg string, data interface{}) {
	// 返回 json数据
	Context.JSON(http_code, gin.H{
		"code": data_code,
		"msg":  msg,
		"data": data,
	})
}
