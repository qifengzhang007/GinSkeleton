package response

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"goskeleton/app/global/consts"
	"goskeleton/app/global/my_errors"
	"goskeleton/app/utils/validator_translation"
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

// 失败的业务逻辑
func Fail(c *gin.Context, dataCode int, msg string, data interface{}) {
	ReturnJson(c, http.StatusBadRequest, dataCode, msg, data)
	c.Abort()
}

// token 基本的格式错误
func ErrorTokenBaseInfo(c *gin.Context) {
	ReturnJson(c, http.StatusBadRequest, http.StatusBadRequest, my_errors.ErrorsTokenBaseInfo, "")
	//终止可能已经被加载的其他回调函数的执行
	c.Abort()
}

//token 权限校验失败
func ErrorTokenAuthFail(c *gin.Context) {
	ReturnJson(c, http.StatusUnauthorized, http.StatusUnauthorized, my_errors.ErrorsNoAuthorization, "")
	//终止可能已经被加载的其他回调函数的执行
	c.Abort()
}

//token 不符合刷新条件
func ErrorTokenRefreshFail(c *gin.Context) {
	ReturnJson(c, http.StatusBadRequest, http.StatusBadRequest, my_errors.ErrorsRefreshTokenFail, "")
	//终止可能已经被加载的其他回调函数的执行
	c.Abort()
}

// casbin 鉴权失败，返回 405 方法不允许访问
func ErrorCasbinAuthFail(c *gin.Context, msg interface{}) {
	ReturnJson(c, http.StatusMethodNotAllowed, http.StatusMethodNotAllowed, my_errors.ErrorsCasbinNoAuthorization, msg)
	c.Abort()
}

//参数校验错误
func ErrorParam(c *gin.Context, wrongParam interface{}) {
	ReturnJson(c, http.StatusBadRequest, consts.ValidatorParamsCheckFailCode, consts.ValidatorParamsCheckFailMsg, wrongParam)
	c.Abort()
}

// 系统执行代码错误
func ErrorSystem(c *gin.Context, msg string, data interface{}) {
	ReturnJson(c, http.StatusInternalServerError, consts.ServerOccurredErrorCode, consts.ServerOccurredErrorMsg+msg, data)
	c.Abort()
}

// validator 错误专门使用的返回器
func ValidatorError(ctx *gin.Context, httpCode int, dataCode int, msg string, err error) {
	// 获取validator.ValidationErrors类型的errors
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		// 非validator.ValidationErrors类型错误直接返回
		Fail(ctx, dataCode, msg, "")
		return
	}
	ctx.JSON(httpCode, gin.H{
		"code": dataCode,
		"msg":  msg,
		"tips": validator_translation.RemoveTopStruct(errs.Translate(validator_translation.Trans)),
	})
}
