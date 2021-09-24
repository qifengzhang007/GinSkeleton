###  validator 表单参数验证器语法介绍  
>   1.本篇将选取表单参数验证器( `https://github.com/go-playground/validator` )主要语法进行介绍，方便本项目骨架使用者快速上手.  
>   2.更详细的语法参与参见官方文档：`https://godoc.org/github.com/go-playground/validator`       

#### 1.我们以用户注册代码块为例进行介绍.  
>   1.[用户注册代码详情](../app/http/validator/web/users/register.go), 摘取表单参数验证部分.        
>   2.以下语法虽然看似简单，实际上已经覆盖了绝大部分常用场景的需求.    
```code  
// 给出一些最常用的验证规则：
//required  必填；
//len=11 长度=11；
//min=3  如果是数字，验证的是数据大小范围，最小值为3，如果是文本，验证的是最小长度为3，
//max=6 如果是数字，验证的是数字最大值为6，如果是文本，验证的是最大长度为6
//mail 验证邮箱
//gt=3  对于文本就是长度>=3
//lt=6  对于文本就是长度<=6

 
type Register struct {
    // 必填、文本类型,表示它的长度>=1
	UserName string `form:"user_name" json:"user_name"  binding:"required,min=1"` 
   
    //必填，密码长度范围：【6,20】闭区间
	Pass string `form:"pass" json:"pass" binding:"required,min=6,max=20"`
    
    //  验证码，必填，长度等于：4 
	//Captcha string `form:"captcha" json:"captcha" binding:"required,len=4"` 

    //  年龄，必填，数字类型，大小范围【1,200】闭区间  
	//Age float64 `form:"age" json:"age" binding:"required,min=1,max=200"` 
    
    //  状态：必填，数字类型，大小范围：【0,1】 闭区间  ，
    //  注意： 如果你的表单参数含有0值是允许提交的，必须用指针类型（*float64），而 float64 类型则认为 0 值不合格
	Status *float64 `form:"status" json:"status"  binding:"required,min=0,max=1"`   
}

// 注意：这里的接收器 r，必须是 r Register, 绝对不能是 r *Register
// 因为在 ginskeleton 里面表单参数验证器是注册在容器的代码段，
// 如果是指针，带参数的接口请求，就会把容器的原始代码污染。
func (r Register) CheckParams(context *gin.Context) {
    // context.ShouldBind(&r) 则自动绑定 form-data 提交的表单参数
	if err := context.ShouldBind(&r); err != nil {

        //  省略非验证器逻辑代码....
        // ...  ...

	}
    
    //  如果您的客户端的数据是以json格式提交（popstman中的raw格式），那么就用如下语法
    // context.ShouldBindJson(&r) 则自动绑定 json格式提交的参数

}

```

#### 2.以上语法特别说明.  
>   1.对于数字类型（int8、int、int64、float32、float64等）我们统一使用 float64、*float64 接受.  
>   2.如果您的业务要求数字格式为 int类型，那么使用 int() 等数据类型转换函数自行转换即可.  
