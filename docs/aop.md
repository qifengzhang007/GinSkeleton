### 控制器 Aop 面向切面编程，优雅地模拟其他语言的动态代理方案。       
> 备注：真正的`Aop` 动态代理，在 `golang` 实现起来非常麻烦，尽管github有相关实现的包(https://github.com/bouk/monkey), 此包明确说明仅用于生产环境之外的测试环境，还有一部分使用非常复杂,因此本项目骨架没有引入第三方包。  
> 需求场景：  
> 1.用户删除数据，需要前置和后置回调函数，但是又不想污染控制器核心代码,此时可以考虑使用Aop思想实现。   
> 2.我们以调用控制器函数 `Users/Destroy` 函数为例，进行演示。     

#### 前置、后置回调最普通的实现方案 
>   此种方案，前置和后置代码比较多的时候，会造成控制器核心代码污染。     
```go  

func (u *Users) Destroy(context *gin.Context) {
    
    //  before 删除之前回调代码... 例如：判断删除数据的用户是否具备相关权限等

	userid := context.GetFloat64(consts.ValidatorPrefix + "id")
    // 根据 userid 执行删除用户数据（最核心代码）

    //  after 删除之后回调代码... 例如 将删除的用户数据备份到相关的历史表
  
}

```

####  使用 Aop 思想实现前置和后置回调需求      
>   1.编写删除数据之前（Before）的回调函数，[示例代码](../app/aop/users/destroy_before.go)  

```bash
package Users

import (
	"goskeleton/app/global/consts"
	"fmt"
	"github.com/gin-gonic/gin"
)

// 模拟Aop 实现对某个控制器函数的前置（Before）回调

type destroy_before struct{}

// 前置函数必须具有返回值，这样才能控制流程是否继续向下执行
func (d *destroy_before) Before(context *gin.Context) bool {
	userId := context.GetFloat64(consts.ValidatorPrefix + "id")
	fmt.Printf("模拟 Users 删除操作， Before 回调,用户ID：%.f\n", userId)
	if userId > 10 {
		return true
	} else {
		return false
	}
}

```
>   2.编写删除数据之后（After）的回调,[示例代码](../app/aop/users/destroy_after.go)  

```bash

package users

import (
	"goskeleton/app/global/consts"
	"fmt"
	"github.com/gin-gonic/gin"
)

// 模拟Aop 实现对某个控制器函数的后置（After）回调

type destroy_after struct{}

func (d *destroy_after) After(context *gin.Context) {
	// 后置函数可以使用异步执行
	go func() {
		userId := context.GetFloat64(consts.ValidatorPrefix + "id")
		fmt.Printf("模拟 Users 删除操作， After 回调,用户ID：%.f\n", userId)
	}()
}


```

>   3.由于本项目骨架的控制器调用都是统一由验证器启动，因此在验证器调用控制器函数的地方，使用匿名函数，直接优雅地切入前置、后置回调代码,[示例代码](../app/http/validator/web/users/destroy.go)   
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
}((&Users.destroy_before{}).Before, (&Users.destroy_after{}).After)

// 接口请求结果展示：
模拟 Users 删除操作， Before 回调,用户ID：16
真正的控制器函数被执行，userId:16
模拟 Users 删除操作， After 回调,用户ID：16
``` 


