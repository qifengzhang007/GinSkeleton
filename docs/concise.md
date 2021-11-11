###  本篇将介绍我们集成的 gorm v2 操作非常流畅的增删改查功能    
> 1.gormv2 功能非常强大，本篇将介绍 gorm_v2 在 GinSkeleton 中非常简洁、简单的操作流程，以 增删改查 操作为例介绍.      
> 2.阅读完本篇，您可以继续阅读官方文档,学习更多功能:https://gorm.io/zh_CN/docs/    

###  前言  
> 1.一个简单的 CURD 操作,我们的起始点为表单参数验证器，终点为数据写入数据库，接下来流程我们将沿着这个主线展开编写.  

### 用户表单参数验证器  
```code  
// 给表单参数验证器设置 form 标签，gin框架会获取用户提交的表单参数绑定在此结构体上
// 设置 json 标签 GinSkeleton 会将json对应的字段绑定在上下文(gin.Context)
type UserStore struct {
	Base        //  Base 表示你可以继续组合其他结构体
	Pass     string `form:"pass" json:"pass" binding:"required,min=6"`
	RealName string `form:"real_name" json:"real_name" binding:"required,min=2"`
	Phone    string `form:"phone" json:"phone" binding:"required,len=11"`
	Remark   string `form:"remark" json:"remark" `
}

// 验证器语法，更详细用法参见常用开发模块列表专项介绍
func (u UserStore) CheckParams(context *gin.Context) {
     // 省略代码...
}


```

###  验证器完成进入控制器，控制器可以直接将 gin.Context 继续传递给 UsersModel
> 1.以下代码将以 model目录 > users 模型展开
```code  

// 创建 userFactory
// 参数说明： 传递空值，默认使用 配置文件选项：UseDbType（mysql）
// 以下函数为固定写法，复制即可，不需要深度研究  

func CreateUserFactory(sqlType string) *UsersModel {
	return &UsersModel{BaseModel: model.BaseModel{DB: model.UseDbConn(sqlType)}}
}

type UsersModel struct {
	model.BaseModel // BaseModel 主要有Id 、 CreatedAt 、UpdatedAt 字段，这里主要是演示UsersModel支持结构体的组合
	UserName         string `gorm:"column:user_name" json:"user_name"`
	Pass             string `json:"pass"`
	Phone            string `json:"phone"`
	RealName         string `gorm:"column:real_name" json:"real_name"`
	Status           int    `json:"status"`
	Remark      string `json:"remark"`
	LastLoginIp string `gorm:"column:last_login_ip" json:"last_login_ip"`
}

// 设置表名
func (u *UsersModel) TableName() string {
	return "tb_users"
}

// UsersModel 结构体组合了  *gorm.DB 的所有功能，您可以通过 u.xxx  直接调用 gorm.DB 的所有功能


```
    
####  1.新增数据  
> 以下代码引用了 `data_bind.ShouldBindFormDataToModel(c, &tmp)` 函数，这个函数是我们对 gin.ShouldBind 函数的精简,加快数据绑定效率。  
> 1.1 参数绑定的原则：model 定义的结构体字段和表单参数验证器结构体设置的json标签名称、数据类型一致，才可以绑定, UserModel 支持类似BaseModel等结构体组合.   
> 1.2 gorm 的数据新增函数 Create 支持单条、批量，如果是批量，只需要定义被添加的数据为 切片即可,例如  	var tmp []UsersModel ,    u.Create(&tmp)    

```code  
//新增数据
	func (u *UsersModel) InsertData(c *gin.Context) bool {
    
    // 注意： 必须重新定义一个 userModel 变量
	var tmp UsersModel
	
	// data_bind.ShouldBindFormDataToModel 函数主要按照 UsersModel 结构体指定的json标签去gin.Context上去寻找相同名称的表单数据,绑定到新定义的变量.
	// 这里不能使用  gin.ShouldBind 函数从上下文绑定数据，因为 UserModel 我们组合了  gorm.DB ，该函数功能太强大，会深入内部持续解析gorm.Db，产生死循环  
	// 使用我们提供的简化版本函数（data_bind.ShouldBindFormDataToModel）代替 gin.ShouldBind 即可   
	
	if err := data_bind.ShouldBindFormDataToModel(c, &tmp); err == nil {
		// Create 函数会将新插入的数据Id 继续更新到 tmp 结构体的主键ID 字段，这里必须传递 指针. 最终的 tmp 其实就是一条新增加的完整数据
		// 注意： 在本项目骨架 Create 参数必须传递 指针类型，这样才能支持gorm的回调函数
			if res := u.Create(&tmp); res.Error == nil {
				return true
			} else {
				variable.ZapLog.Error("UsersModel 数据新增出错", zap.Error(res.Error))
			}
		}else {
		variable.ZapLog.Error("UsersModel 数据绑定出错", zap.Error(err))
	}
	return false
}

```
> 1.3 关于model变量从上下文如何绑定，附一张绑定数据的逻辑图，也可以帮助大家理解.  

![数据绑定原理](https://www.ginskeleton.com/images/bind_explain.png)

####  2.修改数据  
> 2.1 gorm 的数据更新有两个函数： updates 不会处理零值字段，save 会全量覆盖式更新字段    
> 2.2 u.Updates()  函数会根据 UsersModel 已经绑定的 TableName 函数解析对应的数据表,然后根据 tmp 结构体定义的主键Id，去更新其他字段值.    
> 2.3 更新时可以搭配 gorm_v2 提供的 Select() 指定字段更新，例如：gorm.Db.Select(字段1,字段2,字段3..) ,也可以设置忽略特定字段更新数据，例如： gorm.Db.Omit(字段1,字段2,字段3..)  
```code

//更新
func (u *UsersModel) UpdateData(c *gin.Context) bool {
	
	var tmp UsersModel
	if err := data_bind.ShouldBindFormDataToModel(c, &tmp); err == nil {
		
		//tmp 会被自动绑定  CreatedAt、UpdatedAt 字段，更新时我们不希望更新 CreatedAt 字段，使用 Omit 语法忽略该字段
		// 注意： 在本项目骨架 Save、Update 参数必须传递 指针类型，这样才能支持gorm的回调函数
		
		if res := u.Omit("CreatedAt").Save(&tmp); res.Error == nil {
			return true
		} else {
			variable.ZapLog.Error("UsersModel 数据更更新出错", zap.Error(err))
		}
	}
	return false
}

```

####  3.单条删除数据  
> UsersModel 已经绑定了函数 TableName ,所以 u.Delete(u,id)  会自动解析出需要删除的表，然后根据Id删除数据.  

```code
//删除，我们根据Id删除
func (u *UsersModel) DeleteData(id int) bool {
    if u.Delete(u, id).Error == nil {
        return  true
}
return false
}

```

####  4.批量删除数据
> 如果用户传递的参数是  ids ,例如： [100,200,300,400]

```code
//删除，我们根据Id删除
func (u *UsersModel) DeleteData(ids  []int) bool {

    // ids 格式必须是：  [100,200,300,400]
    if u.Where("id  in (?)",ids).Delete(u).Error == nil {
        return  true
}
return false
}

```

####  4.查询 
> 5.1 查询是sql操作最复杂的环节,如果业余复杂，那么请使用原生sql操作业务
```code
    // 查询类 sql 语句
    u.Raw(sql语句,参数1,参数2... ... )

    // 执行类 sql 语句
    u.Exec(sql语句,参数1,参数2... ... )
```
> 5.2 接下来我们演示gorm自带查询     
```code
    // 第一种情况   

    // 如果 UsersModel 结构体已经绑定 TableName 函数，那么查询语句对应的数据表名就是 tableName 的返回值； 
    var  tmp  []UsersModel
    //  Where  关键词前面没有指定表名,那么查询的数据库表名就是 tmp 对应的结构体 UsersModel 结构体绑定的 TableName 的返回值
    u.Where("ID = ?", user_id).Find(&tmp)

    // 第二种情况  
     var  tmp  []UsersList
	//  假设 UsersList 是自定义数据类型，没有绑定 TbaleName ，那么在 where 关键词开始时就必须指定表名
	
	//指定表名 有以下两种方式：
	
	//  u.Model(u)  表示从 u 结构体绑定的 tableName 函数获取对应的表名，如果 u 对应的结构体和 tmp 对应的结构体 UsersList 都没有绑定 TableName ，就会发生错误  
        u.Model(u).Where("ID = ?", user_id).Find(&tmp)  
        
        
    // u.Tbale(u.TableName()).Where("ID = ?", user_id).Find(&tmp)

```
