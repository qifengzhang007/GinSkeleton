###    本篇将探讨主线解耦问题        
> 1.目前项目主线是从路由开始，直接切入到表单参数验证器，验证通过则直接进入了控制器,这里就导致了验证器和控制器之间存在一点低耦合度.       
> 2.如果你追求更低的模块之间的耦合度，接下来我们将对上述问题进行解耦操作.  



###  当前项目代码存在的低耦合逻辑  
> 1.我们以用户删除数据接口为例进行介绍.     
> 2.本文的 `41` 行就是我们所说验证器与控制器出现了低耦合.    
```code 

// 1.访问路由
users.POST("delete", validatorFactory.Create(consts.ValidatorPrefix+"UsersDestroy"))


// 2.进入表单参数验证器
type Destroy struct {
	Id float64 `form:"id"  json:"id" binding:"required,min=1"`
}

func (d Destroy) CheckParams(context *gin.Context) {

	if err := context.ShouldBind(&d); err != nil {
		errs := gin.H{
			"tips": "UserDestroy参数校验失败，参数校验失败，请检查id(>=1)",
			"err":  err.Error(),
		}
		response.ErrorParam(context, errs)
		return
	}

	//  该函数主要是将绑定的数据以 键=>值 形式直接传递给下一步（控制器）
	extraAddBindDataContext := data_transfer.DataAddContext(d, consts.ValidatorPrefix, context)
	if extraAddBindDataContext == nil {
		response.ErrorSystem(context, "UserShow表单参数验证器json化失败", "")
        context.Abort()
		return
	} else {
		// 验证完成，调用控制器,并将验证器成员(字段)递给控制器，保持上下文数据一致性
        //  以下代码就是验证器与控制器之间的一点耦合
		(&web.Users{}).Destroy(extraAddBindDataContext)
	}
}

```

###  开始解耦    
>  1.针对41行出现的验证器与控制器耦合问题，我们开始解耦  
```code    

// 1.我们对以上代码进行简单的改造即可实现代码的解耦
// 2.路由首先切入表单参数验证器，将对应的控制器代码写在第二个回调函数即可
// 3.注意：市面上很多框架的中间件等注册的函数都是 "洋葱模型"  ,即函数的回调顺序和注册顺序是相反的，但是gin框架则是按照注册顺序依次执行
users.POST("delete", validatorFactory.Create(consts.ValidatorPrefix+"UsersDestroy"), (&web.Users{}).Destroy)


// 4.代码经过以上改在以后， 从 38 行开始的 else {  ...  }  代码删除即可

```

###  解耦以后的注意事项    
>  1.如果业务针对控制器存在比较多的 `Aop` 切面编程，就会导致路由文件以及 `import` 显得比较繁重     
```code   

//  1.例如删除数据之前的和之后的回调
users.POST("delete", 

validatorFactory.Create(consts.ValidatorPrefix+"UsersDestroy"),

(&Users.DestroyBefore{}).Before,   // 控制器Aop的前置回调，例如删除数据之前的权限判断，相关代码可参考 app/aop/users/destroy_before.go
(&web.Users{}).Destroy,           // 控制器逻辑
(&Users.DestroyAfter{}).After   // 控制器Aop的后置回调，例如被删除数据之后的数据备份至history表 ,相关代码可参考 app/aop/users/destroy_after.go 

)

```
> 2.对比以上代码，如果你的项目存在较多的 `AOP` 编程、或者说不同的路由前、后回调函数比较多，不建议进行解耦（毕竟目前就是极低耦合）,否则给路由文件以及 `import` 部分带来了比较多的负担.  
> 3.如果你的项目路由前后回调函数比较少，建议参考以上代码进行解耦.  
  
    
   
