###    日志功能, 基于 zap + lumberjack 实现    
> 1.特点：高性能、极速，功能：实现日志的标准管理、日志文件的自动分隔备份.      
> 2.该日志在项目骨架启动时我们封装了全局变量(variable.ZapLog)，直接调用即可，底层按照官方标准封装，使用者调用后不需要关闭日志，也不需要担心全局变量写日志存在并发冲突问题，底层会自动加锁后再写。  
> 3.相关包 github 地址：https://github.com/uber-go/zap 、 https://github.com/natefinch/lumberjack  

    
###  前言  
>   1.日志相关的配置参见，config目录内的config.yml文件，Logs 部分，程序默认处于`AppDebug|调试模式`，日志输出在console面板，编译时记得切换模式。    
>   2.本文档列举几种最常用的用法, 想要深度学习请参考相关的 github 地址.  

###  日志处理, 标准函数
>   参数一：文本型   
>   参数二：可变参数，可传递0个或者多个 Field 类型参数，Field 类型传递规则参见下文     
```code 
>   1. Debug(参数一, 参数二) , 调试级别，会产生大量日志，只在开发模式（AppDebug=true）会输出日志打印在console面板，生产模式该函数禁用。 
>   2. Info(参数一, 参数二) ,  一般信息，默认级别。 
>   3. Warn(参数一, 参数二) ,  警告 
>   4. Panic(参数一, 参数二)、Dpanic(参数一, 参数二) , 恐慌、宕机，不建议使用 
>   5. Error(参数一, 参数二) , 错误
>   6. Fatal(参数一, 参数二) , 致命错误，会导致程序进程退出。 
```

### 标准函数的参数二 Field 类型，最常用传递方式  
>  1.Int    类型 ： zap.Int("userID",2019)  , 同类的还有  int16  、 int32等   
>  2.String 类型 ： zap.String("userID","2019")    
>  3.Error  类型 ： zap.Error(v_err) ， v_err 为 error(错误类型)，例如使用  v_err:= error.New("模拟一个错误")       
>  4.Bool  类型 ： zap.Bool("is_ok",true)    


####    用法 1 , 高性能模式 .      
>   1.举例展示最常用用法  
```code
    variable.ZapLog.Info("基本的运行提示类信息")
    variable.ZapLog.Warn("UserCreate接口参数非法警告，相关参数：",zap.String("userName","demo_name"),zap.Int("userAge",18))  
    variable.ZapLog.Panic("UserDestory接口参数异常，相关参数：",zap.String("userName","demo_name"),zap.String("password","pass123456") 
    variable.ZapLog.Error("UserDestory接口参数错误，相关参数：",zap.Error(error))  
    variable.ZapLog.Fatal("Mysql初始化参数错误，退出运行。相关参数：",zap.String("name","root"), zap.Int("端口",3306))  

```     
    
####    用法2 , 语法糖模式  .   
>   1.比第一种用法性能稍低，只不过基于第一种用法，相关的函数全部增加了格式化参数功能    
```code
 # 第一种的函数后面全部添加了一个 w ,相关的函数功能和第一种一模一样  
 variable.ZapLog.Sugar().Infow("基本的运行提示类信息",zap.String("name","root"))

# 格式化参数，第一种用法中的函数后面添加了一个 f 
 variable.ZapLog.Sugar().Infof("参数 userId %d\n",2020)

 variable.ZapLog.Sugar().Errorw("程序发生错误",zap.Error(error))
 variable.ZapLog.Sugar().Errorf("参数非法，程序出错，userId %d\n",2020)

 Warn  、 Panic 、Fatal用法类似

```     

####   日志钩子  
>   1.除了本项目骨架记录日志之外，您还可以对日志进行二次加工处理.      
>   2.日志钩子函数处理位置 > `app/service/sys_log_hook/zap_log_hooks.go`    
>   3.`bootStrap/init.go` 中你可以修改钩子函数的位置
>   相关代码位置 `app/service/sys_log_hook/zap_log_hooks.go `  
```code 
func ZapLogHandler(entry zapcore.Entry) error {

	// 参数 entry 介绍
	// entry  参数就是单条日志结构体，主要包括字段如下：
	//Level      日志等级
	//Time       当前时间
	//LoggerName  日志名称
	//Message    日志内容
	//Caller     各个文件调用路径
	//Stack      代码调用栈

	//这里启动一个协程，hook丝毫不会影响程序性能，
	go func(paramEntry zapcore.Entry) {
		//fmt.Println(" GoSkeleton  hook ....，你可以在这里继续处理系统日志....")
		//fmt.Printf("%#+v\n", paramEntry)
	}(entry)
	return nil
}

```