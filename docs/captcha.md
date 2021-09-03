##    验证码      
> 1.基于 `github.com/dchest/captcha` 包封装.    
> 2.本项目只提供了数字验证功能,没有封装语音验证功能.    

###   定义的路由地址         
>   1.[路由地址](../routers/web.go)    

###   验证码业务控制器地址         
>   1.[验证码业务](../app/http/controller/chaptcha/chaptcha.go) , 验证数字长度、验证码尺寸（宽 x 高）在这里设置.       

###   使用步骤       
>   1.获取验证码ID等信息  
```code
    # get 方式请求获取验证ID等信息
    http://127.0.0.1:20201/captcha/
    
    #返回值中携带了获取验证码图片的地址以及校验地址

```
>   2.获取验证码    
```code
    # get , 根据步骤1中返回值提示获取 验证码ID  
    http://127.0.0.1:20201/captcha/验证码ID.png
```     
     
>   3.校验验证码    
```code
    # get ,  根据步骤1中返回值提示进行校验验证即可  
    http://127.0.0.1:20201/captcha/验证码ID/验证码正确值  
```   

###   任何路由(接口)都可以调用我们封装好的验证码中间件  
- 1.已经封装好的验证码中间件：authorization.CheckCaptchaAuth()  
- 2.一般是登录接口，需要验证码校验，那么我们可以直接调用验证码中间件增加校验机制。
- 3.注意：如果直接调用了验证码中间件,一般都是和登陆接口搭配，所以请求方式为 `POST`  

```code  

    // 已有的登陆接口(路由)，不需要验证码即可登陆
    noAuth.POST("login", validatorFactory.Create(consts.ValidatorPrefix+"UsersLogin"))
    
    // 只需要添加验证码中间件即可启动登陆前的验证机制
    // 本质上就是给登陆接口增加了2个参数：验证码id提交时的键：captcha_id 和 验证码值提交时的键 captcha_value，具体参见配置文件
    //noAuth.Use(authorization.CheckCaptchaAuth()).POST("login", validatorFactory.Create(consts.ValidatorPrefix+"UsersLogin"))

```

###   备注说明      
>   1.验证码ID一旦提交到校验接口（步骤3）进行验证，不管输入的验证码正确与否，该ID都会失败，需要从步骤1开始重新获取.    
  