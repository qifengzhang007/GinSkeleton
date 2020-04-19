package Users

import (
	"GinSkeleton/App/Global/Consts"
	"GinSkeleton/App/Http/Controller/Admin"
	"GinSkeleton/App/Http/Validattor/Core/DaTaTransfer"
	"github.com/gin-gonic/gin"
	"log"
)

type Destroy struct {
	Id int `form:"id"  json:"id" binding:"required"`
}

func (d *Destroy) CheckParams(context *gin.Context) {

	context.ShouldBind(&d)

	if d.Id > 0 {
		//  该函数主要是将绑定的数据以 键=>值 形式直接传递给下一步（控制器）
		extraAddBindDataContext := DaTaTransfer.DataAddContext(d, Consts.Validattor_Prefix, context)
		if extraAddBindDataContext == nil {
			log.Panic("UserShow表单参数验证器json化失败..")
		} else {
			// 验证完成，调用控制器
			(&Admin.Users{}).Destroy(extraAddBindDataContext)
		}
	} else {
		context.JSON(-100, gin.H{
			"code": -100,
			"msg":  "参数校验失败，请检查Id(>0)",
		})
	}
}
