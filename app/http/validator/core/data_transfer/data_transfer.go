package data_transfer

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"goskeleton/app/global/variable"
	"goskeleton/app/http/validator/core/interf"
	"time"
)

// 将验证器成员(字段)绑定到数据传输到上下文，方便控制器获取
/**
本函数参数说明：
validatorInterface 实现了验证器接口的结构体
extra_add_data_prefix  验证器绑定参数传递给控制器的数据前缀
context  gin上下文
*/

func DataAddContext(validatorInterface interf.ValidatorInterface, extraAddDataPrefix string, context *gin.Context) *gin.Context {
	var tempJson interface{}
	if tmpBytes, err1 := json.Marshal(validatorInterface); err1 == nil {
		if err2 := json.Unmarshal(tmpBytes, &tempJson); err2 == nil {
			if value, ok := tempJson.(map[string]interface{}); ok {
				for k, v := range value {
					context.Set(extraAddDataPrefix+k, v)
				}
				// 此外给上下文追加三个键：created_at  、 updated_at  、 deleted_at ，实际根据需要自己选择获取相关键值
				curDateTime := time.Now().Format(variable.DateFormat)
				context.Set(extraAddDataPrefix+"created_at", curDateTime)
				context.Set(extraAddDataPrefix+"updated_at", curDateTime)
				context.Set(extraAddDataPrefix+"deleted_at", curDateTime)
				return context
			}
		}
	}
	return nil
}
