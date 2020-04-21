package BootStrap

import (
	_ "GinSkeleton/App/Core/Destruct" // 监听程序退出信号，用于资源的释放
	"GinSkeleton/App/Global/Errors"
	"GinSkeleton/App/Global/Variable"
	"GinSkeleton/App/Http/Validator/RegisterValidator"
	_ "GinSkeleton/Test" //  用于测试代码
	"log"
	"os"
)

func init() {
	// 1.初始化程序根目录
	if path, err := os.Getwd(); err == nil {
		Variable.BASE_PATH = path
	} else {
		log.Fatal(Errors.Errors_BasePath)
	}
	//2.初始化表单参数验证器，注册在容器
	RegisterValidator.RegisterValidator()
}
