package Users

import (
	"GinSkeleton/App/Global/Consts"
	"GinSkeleton/App/Http/Controller/Admin"
	"GinSkeleton/App/Http/Validattor/Core/DaTaTransfer"
	"github.com/gin-gonic/gin"
	"log"
)

type Update struct {
	Base
	Id        float64 `form:"id" json:"id" binding:"required"` // 注意：gin框架数字的存储形式都是 float64
	Pass      string  `form:"pass" json:"pass" binding:"required"`
	Real_name string  `form:"real_name" json:"real_name" binding:"required"`
	Phone     string  `form:"phone" json:"phone" binding:"required"`
	Remark    string  `form:"remark" json:"remark" binding:"required"`
}

func (u *Update) CheckParams(context *gin.Context) {

	context.ShouldBind(&u)
	if u.Id > 0 && len(u.Name) >= 6 && len(u.Pass) >= 6 && len(u.Phone) == 11 && len(u.Real_name) > 2 {
		//  该函数主要是将绑定的数据以 键=>值 形式直接传递给下一步（控制器）
		extraAddBindDataContext := DaTaTransfer.DataAddContext(u, Consts.Validattor_Prefix, context)
		if extraAddBindDataContext == nil {
			log.Panic("UserShow表单参数验证器json化失败..")
		} else {
			// 验证完成，调用控制器,并将追加标案参数验证器的上下文传递给控制器
			(&Admin.Users{}).Update(extraAddBindDataContext)
		}

	} else {
		context.JSON(-100, gin.H{
			"code": -100,
			"msg":  "参数校验失败，请检查账号(>=6)、密码(>=6)、姓名(>=2)、手机号码(=11)长度",
		})
	}
}
