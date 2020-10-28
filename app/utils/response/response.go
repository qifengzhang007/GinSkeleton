package response

import (
	"github.com/gin-gonic/gin"
	"goskeleton/app/global/consts"
	"goskeleton/app/global/my_errors"
	"net/http"
)

func ReturnJson(Context *gin.Context, httpCode int, dataCode int, msg string, data interface{}) {

	//Context.Header("key2020","value2020")  	//可以根据实际情况在头部添加额外的其他信息

	Context.JSON(httpCode, gin.H{
		"code": dataCode,
		"msg":  msg,
		"data": data,
	})
}

// 将json字符窜以标准json格式返回（例如，从redis读取json、格式的字符串，返回给浏览器json格式）
func ReturnJsonFromString(Context *gin.Context, httpCode int, jsonStr string) {
	Context.Header("Content-Type", "application/json; charset=utf-8")
	Context.String(httpCode, jsonStr)
}

// 语法糖函数封装

// 直接返回成功
func Success(c *gin.Context, msg string, data interface{}) {
	ReturnJson(c, http.StatusOK, consts.CurdStatusOkCode, msg, data)
}

// 业务逻辑异常
func Fail(c *gin.Context, msg string, data interface{}) {
	ReturnJson(c, http.StatusBadRequest, http.StatusBadRequest, msg, data)
	c.Abort()
}

func ErrorAuthFail(c *gin.Context) {
	ReturnJson(c, http.StatusUnauthorized, http.StatusUnauthorized, my_errors.ErrorsNoAuthorization, "")
	//暂停执行
	c.Abort()
}

func ErrorParam(c *gin.Context, wrongParam interface{}) {
	ReturnJson(c, http.StatusBadRequest, consts.ValidatorParamsCheckFailCode, consts.ValidatorParamsCheckFailMsg, wrongParam)
	c.Abort()
}

func ErrorSystem(c *gin.Context, msg string) {
	ReturnJson(c, http.StatusInternalServerError, consts.ServerOccurredErrorCode, consts.ServerOccurredErrorMsg+msg, "")
	c.Abort()
}
