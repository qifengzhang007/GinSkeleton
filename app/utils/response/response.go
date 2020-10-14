package response

import (
	"github.com/gin-gonic/gin"
	"goskeleton/app/global/consts"
	"goskeleton/app/global/my_errors"
	"net/http"
)

func ReturnJson(Context *gin.Context, httpCode int, dataCode int, msg string, data interface{}) {

	//Context.Header("key2020","value2020")  	//可以根据实际情况在头部添加额外的其他信息

	// 返回 json数据
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

// -------------------------- 简易接口成功返回 ----------------------------------
// 直接返回成功
func SuccessReturn(c *gin.Context) {
	ReturnJson(c, http.StatusOK, consts.CurdStatusOkCode, "ok", "")
}

// 仅返回成功数据
func SuccessReturnData(c *gin.Context, data interface{}) {
	ReturnJson(c, http.StatusOK, consts.CurdStatusOkCode, "ok", data)
}

// 返回成功数据加信息
func SuccessReturnDataMsg(c *gin.Context, data interface{}, msg string) {
	ReturnJson(c, http.StatusOK, consts.CurdStatusOkCode, msg, data)
}

// 仅返回成功信息
func SuccessReturnMsg(c *gin.Context, msg string) {
	ReturnJson(c, http.StatusOK, consts.CurdStatusOkCode, msg, "")
}

// -------------------------- 简易接口错误返回 ----------------------------------
// 业务内逻辑异常
func ErrorReturnMsg(c *gin.Context, code int, msg string) {
	ReturnJson(c, http.StatusBadRequest, code, msg, "")
}

func ErrorAuthFail(c *gin.Context) {
	ReturnJson(c, http.StatusUnauthorized, http.StatusUnauthorized, my_errors.ErrorsNoAuthorization, "")
	//暂停执行
	c.Abort()
}

func ErrorParam(c *gin.Context, wrongParam interface{}) {
	ReturnJson(c, http.StatusBadRequest, consts.ValidatorParamsCheckFailCode, consts.ValidatorParamsCheckFailMsg, wrongParam)
}
