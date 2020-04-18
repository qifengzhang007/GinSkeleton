package DaTaTransfer

import (
	"GinSkeleton/App/Global/Consts"
	"GinSkeleton/App/Http/Validattor/Core/Interface"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

// 将验证绑定的结构体数据传输到上下文，方便控制器获取
func DataAddContext(validatorInterface Interface.ValidatorInterface, context *gin.Context) *gin.Context {
	var temp_json interface{}
	if v_bytes, error := json.Marshal(validatorInterface); error == nil {
		if error2 := json.Unmarshal(v_bytes, &temp_json); error2 == nil {
			if value, ok := temp_json.(map[string]interface{}); ok {
				for k, v := range value {
					context.Set(Consts.Validattor_Prefix+k, v)
				}
				return context
			}
		}

	}
	return nil
}
