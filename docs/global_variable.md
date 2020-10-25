##    项目中被初始化的全局变量清单介绍       

### 1.前言    
>   1.程序启动时初始化动作统一由 `bootstrap/init.go` 文件中的代码段负责，本次我们将介绍3个常用的全局变量.  
>   2.全局变量只会使用法简洁化, 不对原始语法造成任何破坏, 封装全局变量我们经过认真、谨慎地阅读、测试了代码段、并发安全性.       

###  2.gorm 全局变量   
>   1.请按照配置文件 `congfig/gorm_v2.yml` 中的提示正确配置数据库，开启程序启动初始化数据库参数,程序在启动时会自动为您初始化全局变量.  
>   2.不同类型的数据库全局变量名不一样.  
>   3.更多用法参见单元测试：[gorm_v2单元测试](../test/gormv2_test.go), 本文档我们主要介绍核心原理.      
```code 

//  例如：原始语法
// 1.连接数据库,获取mysql连接
 dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
 db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})  
// 2.查询
db.Select("id", "name", "phone", "email", "remark").Where("name  like ?", "%test%").Find(&users)


// 本项目中, `variable.GormDbMysql` 等于 上文中返回的 mysql db 
variable.GormDbMysql.Select("id", "name", "phone", "email", "remark").Where("name  like ?", "%test%").Find(&users)

// 为 gorm 操作数据库封装全局变量列表
variable.GormDbMysql    
variable.GormDbSqlserver  
variable.GormDbPostgreSql  

```

###  3.日志全局变量  
>   1.为了随意、方便地记录项目中日志，我们封装了全局变量 `variable.ZapLog` .  
>   2.由于日志操作内容比较多，我们对它进行了单独介绍，详情参见： [zap高性能日志](zap_log.md)    


### 4.配置文件全局变量 
>  1.为了更方便地操作配置文件 `config/config.yml` 、 `config/gorm_v2.yml` 我们同样在项目启动时封装了全局变量.  
>  2.`variable.ConfigYml` ,该变量相当于配置文件 `config/config.yml` 文件打开后的指针.   
>  3.`variable.ConfigGormv2Yml` ,该变量相当于配置文件 `config/gorm_v2.yml` 文件打开后的指针.  
>  4.在任何地方您都可以使用以上全局变量直接获取对应配置文件的 键==>值.  
```code   

// 获取 config/config.yml 文件中 Websocket.Start 对应的 Int 值
variable.ConfigYml.GetInt("Websocket.Start")

// 获取 config/gorm_v2.yml 文件中 Gormv2.Mysql.IsInitGolobalGormMysql 对应的 Int 值
variable.ConfigGormv2Yml.GetInt("Gormv2.Mysql.IsInitGolobalGormMysql")

```
>   5.获取配置文件中键对应的值数据类型，函数清单，您可以使用 `variable.ConfigYml.` 或者  `variable.ConfigGormv2Yml.` 以下函数名 获取值   
```code  
    // 开发者常用函数
	GetString(keyName string) string
	GetInt(keyName string) int
	GetInt32(keyName string) int32
	GetInt64(keyName string) int64
	GetFloat64(keyName string) float64
	GetDuration(keyName string) time.Duration
	GetBool(keyName string) bool

    // 非常用函数，主要是项目骨架在使用
	ConfigFileChangeListen()
	Clone(fileName string) YmlConfigInterf
	Get(keyName string) interface{} // 该函数获取一个 键 对应的原始值，因此返回类型为 interface , 基本很少用
	GetStringSlice(keyName string) []string
```

