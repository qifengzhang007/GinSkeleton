###  本篇将介绍Casbin模块的基本用法     
> 1.Casbin(https://github.com/casbin/casbin) 提供了一款跨语言的接口访问权限管控机制,针对go语言支持的最为全面.    
> 2.该模块的使用看起来非常复杂，只要理解了其核心思想，使用是非常简单易懂的. 


###  前言  
> 1.`Casbin` 的初始化在 GinSkeleton 主线版本默认没有开启，请参照配置文件(config/config.yml)文件中 `casbin` 部分,自行决定是否开启，默认的配置项属于标准配置，基本不需要改动.    
> 2.配置文件开启 Casbin 模块后，默认会在连接的数据库创建一张表，具体表名参见配置文件说明.  

### 根据用户请求接口时头部附带的token解析用户id等信息  
> 每个用户带有token的请求，在验证ok之后自动会将token绑定在上下文(gin.Context) ,绑定的键名默认为: userToken（配置文件可自行设置键名）
> 通过token解析出用户id等信息的代码如下：       
```code   
currentUser, exist := context.MustGet("userToken").(my_jwt.CustomClaims)

	if exist {
		fmt.Printf("userId：%d\n",currentUser.UserId)
	}

```

###  Casbin 相关的几个功能介绍  
>   1.Casbin 中间件，相关位置: app/http/middleware/authorization/auth.go, 中间件的作用介绍:  
```code

// casbin检查用户对应的角色权限是否允许访问接口
func CheckCasbinAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
	
		requstUrl := c.Request.URL.Path
		method := c.Request.Method
		
		// 这里根据用户请求时头部的 token 解析出用户id，根据用户id查询出该用户所拥有的角色id(roleId)
		// 主线版本的程序中 角色表需要开发者自行创建、管理，Ginskeleton-Admin 系统则集成了所有的基础功能
		// 根据角色（roleId）判断是否具有某个接口的权限
		roleId := "2" // 模拟最终解析出用户对应的角色为 2 

        // 使用casbin自带的函数执行策略(规则)验证
		isPass, err := variable.Enforcer.Enforce(role, requstUrl, method)
		if err != nil {
			response.ErrorCasbinAuthFail(c, err.Error())
			return
		} else if !isPass {
			response.ErrorCasbinAuthFail(c, "")
		} else {
			c.Next()
		}
	}
}

```

###  Casbin 用法  
> 1.Casbin 负责检查用户请求时后台是否允许访问某个接口(路由地址),作为用户的一次请求，主要有三个要素： 
> 1.1 请求的地址(url) 
> 1.2 请求的方式（GET 、 POST 等） 
> 1.3 请求时用户的身份(角色Id，可以根据token解析出用户id，再根据用户id查询出对应的角色ID)  
> 2.Casbin会根据用户请求的三个要求匹配数据库相关设置,匹配成功方可进入路由，否则直接在中间件拦截本次请求.  
```code  
		// 【需要token】中间件验证的路由
		// 在某个分组或者模块，我们追加token校验完成后的具体模块接口校验机制
		// 追加 authorization.CheckCasbinAuth() 中间件，凡是用户访问就必须经过 token校验+casbin 接口权限校验  
		// casbin 匹配策略时需要将用户id 转为角色id，因此必须放在 token 中间件后面（token中才能解析出用户id）  
		backend.Use(authorization.CheckTokenAuth(), authorization.CheckCasbinAuth() )
		{
			// 用户组路由
			users := backend.Group("users/")
			{
				// 查询 
				users.GET("list", validatorFactory.Create(consts.ValidatorPrefix+"UserList"))
				// 新增
				users.POST("create", validatorFactory.Create(consts.ValidatorPrefix+"UserCreate"))
				// 更新
				users.POST("edit", validatorFactory.Create(consts.ValidatorPrefix+"UserEdit"))
				// 删除
				users.POST("destroy", validatorFactory.Create(consts.ValidatorPrefix+"UserDestroy"))

			}
		}	

```


###  Casbin 核心数据表  
> 只要在配置文件（config/config.yml）开启Casbin相关的配置项,程序启动会默认创建一个表：tb_auth_casbin_rule ,开发者按照示例将数据写入该表即可.  
> 表数据的字段含义介绍请参见截图标注的文本.    

![tb_casbin_rules](https://www.ginskeleton.com/images/casbin_introduce.jpg)  
