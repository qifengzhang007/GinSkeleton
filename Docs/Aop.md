### Aop 面向切面编程，目前只支持控制器函数，优雅地模拟 Aop    
> 需求场景：
> 1.例如：用户删除数据，需要前置和后置回调函数，又不想污染控制器核心代码,此时可以考虑使用Aop思想。   

#### 前置、后置回调最普通的实现方案    
```go  

func (u *Users) Destroy(context *gin.Context) {
    
    //  before 删除之前回调代码... 例如：判断删除数据的用户是否具备相关权限等

	userid := context.GetFloat64(Consts.Validator_Prefix + "id")
    // 根据 userid 执行删除用户数据（最核心代码）

    //  after 删除之后回调代码... 例如 将删除的用户数据备份掉相关的历史表
  
}

```

####  使用 Aop 思想实现的前置和后置回调方案  
>   1.编写删除数据之前（Before）的回调示例代码，相关文件路径：GinSkeleton\App\Aop\Users\DestroyBefore.go  

```bash
package Users

import (
	"GinSkeleton/App/Global/Consts"
	"fmt"
	"github.com/gin-gonic/gin"
)

// 模拟Aop 实现对某个控制器函数的前置（Before）回调

type DestroyBefore struct{}

// 前置函数必须具有返回值，这样才能控制流程是否继续向下执行
func (d DestroyBefore) Before(context *gin.Context) bool {
	userId := context.GetFloat64(Consts.Validator_Prefix + "id")
	fmt.Printf("模拟 Users 删除操作， Before 回调,用户ID：%.f\n", userId)
	if userId > 10 {
		return true
	} else {
		return false
	}
}

```
>   2.编写删除数据之后（After）的回调示例代码，相关文件路径：GinSkeleton\App\Aop\Users\DestroyAfter.go  

```bash

package Users

import (
	"GinSkeleton/App/Global/Consts"
	"fmt"
	"github.com/gin-gonic/gin"
)

// 模拟Aop 实现对某个控制器函数的后置（After）回调

type DestroyAfter struct{}

func (d DestroyAfter) After(context *gin.Context) {
	// 后置函数可以使用异步执行
	go func() {
		userId := context.GetFloat64(Consts.Validator_Prefix + "id")
		fmt.Printf("模拟 Users 删除操作， After 回调,用户ID：%.f\n", userId)
	}()
}


```

>   3.由于本项目骨架的控制器调用都是统一由验证器触发，因此在验证器调用控制器函数的地方，使用匿名函数，直接优雅地切入前置、后置回调代码。  
```go  
         
//(&Web.Users{}).Destroy(extraAddBindDataContext)   // 原始方法进行如下改造  

// 使用匿名函数切入前置和后置回调函数  
func(before_callback_fn func(context *gin.Context) bool, after_callback_fn func(context *gin.Context)) {

    if before_callback_fn(extraAddBindDataContext) {
        defer after_callback_fn(extraAddBindDataContext)
        (&Web.Users{}).Destroy(extraAddBindDataContext)
    } else {
        // 这里编写前置函数验证不通过的相关返回提示逻辑...

     }
}((Users.DestroyBefore{}).Before, (Users.DestroyAfter{}).After)

// 接口请求结果展示：
模拟 Users 删除操作， Before 回调,用户ID：16
真正的控制器函数被执行，userId:16
模拟 Users 删除操作， After 回调,用户ID：16
``` 


