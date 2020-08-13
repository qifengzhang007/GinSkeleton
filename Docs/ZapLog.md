###    日志功能, 基于 zap + lumberjack 实现    
> 1.特点：高性能、极速，功能：实现日志的标准管理、日志文件的自动分隔备份.      
> 2.该日志在项目骨架启动时我们封装了全局变量(Variable.ZapLog)，直接调用即可，底层按照官方标准封装，使用者调用后不需要关闭日志，也不需要担心全局变量写日志存在并发冲突问题，底层会自动加锁后再写。  
> 3.github地址：https://github.com/uber-go/zap 、 https://github.com/natefinch/lumberjack  

    
###  前沿  
>   1.日志相关的配置参见，Config目录内的config.yaml文件，Logs 部分，使用默认配置即可。

###  标准日志函数
>   参数一：文本型   
>   参数二：可变参数，可传递0个或者多个 Field 类型参数，Field 类型传递规则参见下文     
```code 
>   1. Debug(参数一, 参数二) , 调试级别，会产生大量日志，官方默认禁用。 
>   2. Info(参数一, 参数二) ,  一般信息，默认级别。 
>   3. Warn(参数一, 参数二) ,  警告 
>   4. Panic(参数一, 参数二)、Dpanic(参数一, 参数二) , 恐慌 
>   5. Error(参数一, 参数二) , 错误
>   6. Fatal(参数一, 参数二) , 致命错误，会导致程序进程退出。 
```

### 标准函数的参数二 Field 类型，最常用传递方式  
>  1.Int    类型 ： zap.Int("userID",2019,"Age",2020)  , 同类的还有  int16  、 int32等   
>  2.String 类型 ： zap.String("userID","2019","userName","姓名")    
>  3.Error  类型 ： zap.Error(v_err) ， err 为 error(错误类型)，例如使用  v_err:= error.New("模拟一个错误")     
>  4.Bool  类型 ： zap.Bool("is_ok",true)    


####    用法 1 , 高性能模式 .      
>   1.举例展示最常用用法  
```code
    Variable.ZapLog.Info("基本的运行提示类信息")
    Variable.ZapLog.Warn("UserCreate接口参数非法警告，相关参数：",zap.String("userName","demo_name","userPass","demo_pass"))  
    Variable.ZapLog.Panic("UserDestory接口参数异常，相关参数：",zap.String("userName","demo_name","userPass","demo_pass")) 
    Variable.ZapLog.Error("UserDestory接口参数错误，相关参数：",zap.Error(error))  
    Variable.ZapLog.Fatal("Mysql初始化参数错误，退出运行。相关参数：",zap.String("name","root"), zap.Int("端口",3306))  

```     
    
####    用法2 , 语法糖模式  .   
>   1.比第一种用法性能稍低，只不过基于第一种用法，相关的函数全部增加了格式化参数功能   
```code

 Variable.ZapLog.Sugar().Info("基本的运行提示类信息")
 Variable.ZapLog.Sugar().Infof("参数 userId %d\n",2020)

 Variable.ZapLog.Sugar().Error("程序发生错误",zap.Error(error))
 Variable.ZapLog.Sugar().Errorf("参数非法，程序出错，userId %d\n",2020)

 Warn  、 Panic 、Fatal用法类似

```     
