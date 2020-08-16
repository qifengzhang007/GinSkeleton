package response

import (
	"github.com/gin-gonic/gin"
)

func ReturnJson(Context *gin.Context, http_code int, data_code int, msg string, data interface{}) {

	//Context.Header("key2020","value2020")  	//可以根据实际情况在头部添加额外的其他信息

	// 返回 json数据
	Context.JSON(http_code, gin.H{
		"code": data_code,
		"msg":  msg,
		"data": data,
	})
}

// 将json字符窜以标准json格式返回（例如，从redis读取json、格式的字符串，返回给浏览器json格式）
func ReturnJsonFromString(Context *gin.Context, http_code int, json_str string) {
	Context.Header("Content-Type", "application/json; charset=utf-8")
	Context.String(http_code, json_str)
}
