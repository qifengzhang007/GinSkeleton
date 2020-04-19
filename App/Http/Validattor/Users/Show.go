package Users

import (
	"GinSkeleton/App/Global/Consts"
	"GinSkeleton/App/Http/Controller/Admin"
	"GinSkeleton/App/Http/Validattor/Core/DaTaTransfer"
	"github.com/gin-gonic/gin"
	"log"
)

type Show struct {
	Base
	Page   float64 `form:"page" json:"page" binding:"required"`
	Limits float64 `form:"limits" json:"limits" binding:"required"`
}

func (s *Show) CheckParams(context *gin.Context) {

	if err := context.ShouldBind(&s); err != nil {
		log.Panic("UserShow, shouldBind出错")
		return
	}
	//  该函数主要是将绑定的数据以 键=>值 形式直接传递给下一步（控制器）
	extraAddBindDataContext := DaTaTransfer.DataAddContext(s, Consts.Validattor_Prefix, context)
	if extraAddBindDataContext == nil {
		log.Panic("UserShow表单参数验证器json化失败..")
	} else {
		// 验证完成，调用控制器,并将追加标案参数验证器的上下文传递给控制器
		(&Admin.Users{}).Show(extraAddBindDataContext)
	}

}
