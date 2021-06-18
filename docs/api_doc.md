### 测试用例接口   
>   1.文档主要提供本项目骨架已经集成的Api接口使用说明。            
>   2.相关测试全部基于`postman`工具进行。     

### 默认已经集成的路由  

####  门户网站类  
>GET    http://127.0.0.1:20191    
>GET    /api/v1/home/news?newsType=portal&page=1&limit=50 

#### 后台管理类
>GET    /http://127.0.0.1:20201                         
>GET   /admin/ws         
>POST   /admin/users/register     
>POST   /admin/users/login        
>POST   /admin/users/refreshtoken        
>GET    /admin/users/index        
>POST   /admin/users/create       
>POST   /admin/users/edit         
>POST   /admin/users/delete       
>POST   /admin/upload/file           

#### pprof 路由
>调试模式自动开启，以pprof开头的路由   
> http://127.0.0.1:20191/debug/pprof/  
> http://127.0.0.1:20201/debug/pprof/  

### 门户网站类
>   1.ip、端口使用本项目默认配置，即：`http://127.0.0.1:20191`，门户类接口通用  
####    1.首页新闻   
>    <font color=#FF4500>*get*，/api/v1/home/news?newsType=portal&page=1&limit=50 
> 返回示例：
```json
{
    "code": 200,
    "data": {
        "content": "门户新闻内容001",
        "limit": 20,
        "newstype": "potal",
        "page": 1,
        "title": "门户首页公司新闻标题001",
        "user_ip": "127.0.0.1"
    },
    "msg": "Success"
}
```  




### 后台应用类
>   1.ip、端口使用本项目默认配置，即：`http://127.0.0.1:20201`，后端管理类系统通用。  

####    1.用户注册   
> 表单参数验证器: [register](../app/http/validator/web/users/register.go) 
>    <font color=#FF4500>*post*，/admin/users/register   </font>

参数字段|参数属性|类型|选项|默认值
---|---|---|---|---
user_name|form-data|string|必填|goskeleton1.4  
pass|form-data|string|必填|goskeleton1.4  
> 返回示例：
```json
{
    "code": 200,
    "data": "",
    "msg": "Success"
}
```  

####    2.用户登录     
> 表单参数验证器: [login](../app/http/validator/web/users/login.go) 
>    <font color=#FF4500>*post*，/admin/users/login   </font>

参数字段|参数属性|类型|选项|默认值
---|---|---|---|---
user_name|form-data|string|必填|goskeleton1.4
pass|form-data|string|必填|goskeleton1.4
captcha_id|form-data|string|如果登录接口使用了验证码中间件，则必填|uY26gnHcHNnhot0lYkG8
captcha_value|form-data|string|如果登录接口使用了验证码中间件，则必填|1234

> 返回示例，关于登陆时是否提交验证码取决于登陆路由(接口)是否加载了验证码中间件.  
```json
{
    "code": 200,
    "data": {
        "phone": "",
        "realName": "",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyaWQiOjQ3LCJ1c2VyX25hbWUiOiJnb3NrZWxldG9uMS40IiwicGhvbmUiOiIiLCJleHAiOjE2MDQwNTIxNzMsIm5iZiI6MTYwNDA0ODU2M30.YNhN9_QasHc5XILQiilZvhxpPDnmC_j82y4JfYPnI7A",
        "updated_at": "2020-10-30 17:02:53",
        "userId": 47,
        "user_name": "goskeleton1.4"
    },
    "msg": "Success"
}
```  

####    3.根据关键词查询用户表   
> 表单参数验证器: [index](../app/http/validator/web/users/show.go)    
>    <font color=#FF4500>*get*，/admin/users/index </font>  ，注意该接口需要token鉴权，请在 `header` 头添加 `Authorization` 字段值，注意：该字段的值格式：Bearer (token)之间有一个空格, 这个是行业标准，网页端显示换行，不要被误导! 
>   CURD相关的其他接口格式与本接口基本一致，例如：/admin/users/create、/admin/users/edit、/admin/users/delete，只不过表单参数不一致。    

参数字段|参数属性|类型|选项|默认值
---|---|---|---|---
Authorization|Headers|string|必填|Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyaWQiOjQ3LCJ1c2VyX25hbWUiOiJnb3NrZWxldG9uMS40IiwicGhvbmUiOiIiLCJleHAiOjE2MDQwNTIxNzMsIm5iZiI6MTYwNDA0ODU2M30.YNhN9_QasHc5XILQiilZvhxpPDnmC_j82y4JfYPnI7A
user_name|form-data|string|必填|g
page|form-data|int|必填|1
limit|form-data|int|必填|20

> 返回示例：
```json
{
    "code": 200,
    "data": [
        {
            "user_name": "zhang001",
            "phone": "1660177xxxx",
            "real_name": "张三丰",
            "status": 1,
            "token": "",
            "last_login_ip": ""
        },
        {
            "user_name": "goskeleton51",
            "phone": "1580403xxxx",
            "real_name": "新的姓名",
            "status": 1,
            "token": "",
            "last_login_ip": ""
        }
    ],
    "msg": "Success"
}
```  

####    4.新增用户  
> 表单参数验证器: [create](../app/http/validator/web/users/store.go)  
>    <font color=#FF4500>*post*，/admin/users/create  </font> ，注意该接口需要token鉴权，请在 `header` 头添加 `Authorization` 字段值，注意：该字段的值格式：Bearer (token)之间有一个空格, 这个是行业标准，网页端显示换行，不要被误导!  

参数字段|参数属性|类型|选项|默认值
---|---|---|---|---
Authorization|Headers|string|必填|Bearer 登陆后获取的token  
user_name|form-data|string|必填|goskeleto002
pass|form-data|string|必填|goskeleto002
real_name|form-data|string|必填|goskeleto002
phone|form-data|string|必填|1580403xxxx
remark|form-data|string|非必填|备注信息
> 返回示例：
```json
{
    "code": 200,
    "data": "",
    "msg": "Success"
}
```  

####    5.更新用户  
> 表单参数验证器: [edit](../app/http/validator/web/users/update.go)  
>    <font color=#FF4500>*post*，/admin/users/edit  </font> ，注意该接口需要token鉴权，请在 `header` 头添加 `Authorization` 字段值，注意：该字段的值格式：Bearer (token)之间有一个空格, 这个是行业标准，网页端显示换行，不要被误导!  

参数字段|参数属性|类型|选项|默认值
---|---|---|---|---
Authorization|Headers|string|必填|Bearer 登陆后获取的token  
id|form-data|int|必填|51
user_name|form-data|string|必填|goskeleto002
pass|form-data|string|必填|goskeleto002
real_name|form-data|string|必填|goskeleto002
phone|form-data|string|必填|1580403xxxx
remark|form-data|string|非必填|备注信息
> 返回示例：
```json
{
    "code": 200,
    "data": "",
    "msg": "Success"
}
```  
####    6.删除用户  
> 表单参数验证器: [delete](../app/http/validator/web/users/destroy.go)  
>    <font color=#FF4500>*post*，/admin/users/delete  </font> ，注意该接口需要token鉴权，请在 `header` 头添加 `Authorization` 字段值，注意：该字段的值格式：Bearer (token)之间有一个空格, 这个是行业标准，网页端显示换行，不要被误导!  

参数字段|参数属性|类型|选项|默认值
---|---|---|---|---
Authorization|Headers|string|必填|Bearer 登陆后获取的token  
id|form-data|int|必填|51  
> 返回示例：
```json
{
    "code": 200,
    "data": "",
    "msg": "Success"
}
```  

####    7.token刷新 ，请将旧token放置在header头参数直接提交更新         
> 表单参数验证器: [refresh_token](../app/http/validator/web/users/refresh_token.go)  
>    <font color=#FF4500>*post*，/admin/users/refreshtoken</font>    

参数字段|参数属性|类型|选项|默认值
---|---|---|---|---
Authorization|Headers|string|必填|Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyaWQiOjQ3LCJ1c2VyX25hbWUiOiJnb3NrZWxldG9uMS40IiwicGhvbmUiOiIiLCJleHAiOjE2MDQwNTIxNzMsIm5iZiI6MTYwNDA0ODU2M30.YNhN9_QasHc5XILQiilZvhxpPDnmC_j82y4JfYPnI7A  

> 返回示例：
```json
{
    "code": 200,
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyaWQiOjQ3LCJ1c2VyX25hbWUiOiJnb3NrZWxldG9uMS40IiwicGhvbmUiOiIiLCJleHAiOjE2MDQwNTYxMDcsIm5iZiI6MTYwNDA0ODU2M30.JPE6G-9YE9UTdxHiWuvdVlD-akiIkvp6Ezf9y4_ud9M"
    },
    "msg": "Success"
}
```  

####    8.文件上传        
> 表单参数验证器: [upload_fiels](../app/http/validator/common/upload_files/upload_fiels.go)  
>    <font color=#FF4500>*post*，/admin/upload/files</font>        

参数字段|参数属性|类型|选项|默认值
---|---|---|---|---
Authorization|Headers|string|必填|Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyaWQiOjQ3LCJ1c2VyX25hbWUiOiJnb3NrZWxldG9uMS40IiwicGhvbmUiOiIiLCJleHAiOjE2MDQwNTIxNzMsIm5iZiI6MTYwNDA0ODU2M30.YNhN9_QasHc5XILQiilZvhxpPDnmC_j82y4JfYPnI7A
file|form-data|string|必填|(注意表单键名为files，如果需要修改成别的键名，参见：App\Global\Variable\Variable.go ，UploadFileField=files)
> 返回示例：
```json
{
    "code": 200,
    "data": {
        "path": "/storage/app/uploaded/3c5d5f59484cad593e46d7fe0c6b078e.sql"
    },
    "msg": "Success"
}
```  
> postman 直接上传文件相关参数  
![文件上传](https://www.ginskeleton.com/images/upload.png)    