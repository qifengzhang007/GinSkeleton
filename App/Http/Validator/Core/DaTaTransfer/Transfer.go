package DaTaTransfer

import (
	"GinSkeleton/App/Http/Validator/Core/Interface"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

// 将验证绑定的结构体字段（成员）数据传输到上下文，方便控制器获取
/**
本函数参数说明：
validatorInterface 实现了验证器接口的结构体
extra_add_data_prefix  验证器绑定参数传递给控制器的数据前缀
context  gin上下文
*/
func DataAddContext(validatorInterface Interface.ValidatorInterface, extra_add_data_prefix string, context *gin.Context) *gin.Context {
	var temp_json interface{}
	if v_bytes, err1 := json.Marshal(validatorInterface); err1 == nil {
		if err2 := json.Unmarshal(v_bytes, &temp_json); err2 == nil {
			if value, ok := temp_json.(map[string]interface{}); ok {
				for k, v := range value {
					context.Set(extra_add_data_prefix+k, v)
				}
				return context
			}
		}

	}
	return nil
}
