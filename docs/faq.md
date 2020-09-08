##    招募共同开发者        
> 1.请先看这位开发者发布的文章："7天用go开发一个docker"， 地址：`https://learnku.com/articles/46878` ,在这篇文章的留言处有作者的一句话：`很多东西不是会了才能做，而是做了才能学会` .  
> 2.基于第一条“真理”, 只要你会go基础的东西，有时间，就可以一起参与开发本项目.  
> 3.参与方式：简单的东西直接提交PR,如果想法比较多，需要改动大段代码，你也可以直接加我qq：1990850157，直接添加至开发组，共同商讨开发的功能，约定规范，提交代码。  
> 4.成为共同开发者，你可以获得 `goland` 官方提供的激活码，通用全部的Jetbrains全家桶项目.  

##    常见问题汇总      
> 1.本篇我们将汇总使用过程中最常见的问题, 很多细小的问题或许在这里你能找到答案.  

#####  1.为什么该项目 go.mod 中的模块名是 goskeleton ,但是下载下来的文件名却是 GinSkeleton ?
>   本项目一开始我们命名为 ginskeleton , 包名也是这个，但是后来感觉 goskeleton 好听一点（现在看来未必）,
>基于更易理解的角度出发，你在下载或者pull本项目之后，可以将最外层文件夹名重新命名为 goskeleton , 这样就会让整个项目显得统一,代码内部引用包的时候，类似从文件夹（goskeleton）开始，按照路径在引用包，理解起来更直观.       

#####  2.为什么编译后的文件提示 config.yml 文件不存在 ?  
>   项目的编译仅限于代码部分，不包括资源部分：config 目录、public 目录、storage 目录，因此编译后的文件使用时，需要带上这个三个目录，否则程序无法正常运行.    

#####  3.表单参数验证器代码部分的疑问    
>   示例代码位置：`app/http/validator/web/users/register.go`  ,如下代码段  
```code 
type Register struct {
	Base
	Pass  string `form:"pass" json:"pass" binding:"required,min=3,max=20"` //必填，密码长度范围：【3,20】闭区间
	Phone string `form:"phone" json:"phone"  binding:"required,len=11"`    //  验证规则：必填，长度必须=11
	//CardNo  string `form:"card_no" json:"card_no" binding:"required,len=18"`	//身份证号码，必填，长度=18
}

// 注意这里绑定在了  Register  
func (r Register) CheckParams(context *gin.Context) {
    //  ...
}


```  
>  CheckParams 函数是否可以绑定在指针上？例如写成如下:  
```code  
// 注意这里绑定在了  *Register 
func (r *Register) CheckParams(context *gin.Context) {
    //  ...
}

```
> <font color="red">这里绝对不可以，因为表单参数验证器在程序启动时会自动注册在容器,每次调用都必须是一个全新的初始化代码段，如果绑定在指针，第一次请求验证通过之后，相关的参数值就会绑定容器中的代码上,造成下次请求数据污染.</font>
 
#####  4.其他问题欢迎补充....




    