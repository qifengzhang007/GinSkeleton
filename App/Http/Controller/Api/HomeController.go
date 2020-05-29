package Api

import (
	"GinSkeleton/App/Global/Consts"
	"GinSkeleton/App/Utils/Response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Home struct {
}

// 1.门户类首页新闻
func (u *Home) News(context *gin.Context) {

	//  由于本项目骨架已经将表单验证器的字段(成员)绑定在上下文，因此可以按照 GetString()、Getint64()、GetFloat64（）等快捷获取需要的数据类型
	// 当然也可以通过gin框架的上下文原原始方法获取，例如： context.PostForm("name") 获取，这样获取的数据格式为文本，需要自己继续转换
	newstype := context.GetString(Consts.Validator_Prefix + "newstype")
	page := context.GetFloat64(Consts.Validator_Prefix + "page")
	limit := context.GetFloat64(Consts.Validator_Prefix + "limit")
	user_ip := context.ClientIP()

	// 这里随便模拟一条数据返回
	faka_data := gin.H{
		"newstype": newstype,
		"page":     page,
		"limit":    limit,
		"user_ip":  user_ip,
		"title":    "门户首页公司新闻标题001",
		"content":  "门户新闻内容001",
	}
	Response.ReturnJson(context, http.StatusOK, Consts.Curd_Status_Ok_Code, Consts.Curd_Status_Ok_Msg, faka_data)
}
