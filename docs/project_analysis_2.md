##    GoSkeleton 项目骨架性能分析报告（二） 
> 1.本次我们分析的目标是操作数据库, 通过操作数据库，分析相关代码段cpu的耗时，得出可视化的性能分析报告。   


###  操作数据库, 我们需要做如下铺垫代码  
>   1.我们本次分析的核心是在数据库操作部分, 因此我们在路由出添加如下代码，访问路由即可触发数据库的调用.  
```code 
	router.GET("/", func(context *gin.Context) {
        // 默认路由处直接触发数据库调用
		if model.CreateTestFactory("").SelectDataMultiple() {
			context.String(200,"批量查询数据OK")
		} else {
			context.String(200,"批量查询数据出错")
		}
		context.String(http.StatusOK, "Api 模块接口 hello word！")
	})
```   

>   2.数据库部分代码，主要逻辑是每次查询1000条，循环查询了100次，并且在最后一次输出了结果集.  
 ```code 
func (t *Test) SelectDataMultiple() bool {
    // 本次测试的数据表内有6000条左右数据 
	sql := `
			SELECT
			code,name,company_name,indudtry,created_at 
			FROM
			db_stocks.tb_code_list 
			LIMIT 0, 1000 ;
		`
	//1.首先独立预处理sql语句，无参数
	if t.PrepareSql(sql) {
	
		var code, name, company_name, indudtry, created_at string
		for i := 1; i <= 100; i++ {
			//2.执行批量查询
			rows := t.QuerySqlForMultiple()
			if rows == nil {
				variable.ZapLog.Sugar().Error("sql执行失败，sql:", sql)
				return false
			} else {
				// 我们只输出最后一行数据
				if i == 100 {
					for rows.Next() {
						_ = rows.Scan(&code, &name, &company_name, &indudtry, &created_at)
						fmt.Println(code, name, company_name, indudtry, created_at)
					}
				}
			}
			rows.Close()
		}
	}
	variable.ZapLog.Info("批量查询sql执行完毕！")
	return true
}
 ```  
###  cpu 底层数据采集步骤  
>   1.浏览器访问pprof接口：`http://127.0.0.1:20191/debug/pprof/`, 点击 `profile` 选项,程序会对本项目进程, 进行 cpu 使用情况底层数据采集, 该过程会持续 30 秒.     
![pprof地址](https://www.ginskeleton.com/images/pprof_menue.jpg)   
>   2.新开浏览器窗口,输入 `http://127.0.0.1:20191/` 刷新,触发路由中的数据库操作代码, 等待被 pprof 采集数据.      
>   3.稍等片刻，30秒之后，您点击过的步骤1就会提示下载文件：`profile`, 请保存在您能记住的路径中，稍后马上使用该文件(profile), 至此cpu数据已经采集完毕.         

###  cpu数据分析步骤   
>  1.首先下载安装 [graphviz](https://www.graphviz.org/download/) ,根据您的系统选择相对应的版本安装，安装完成记得将`安装目录/bin`, 加入系统环境变量.  
>  2.打开cmd窗口,执行 `dot -V` ,会显示版本信息，说明安装已经OK, 那么继续执行 `dot  -c` 安装图形显示所需要的插件.   
>  3.在cpu数据采集环节第三步,您已经得到了 `profile` 文件,那么就在同目录打开cmd窗口,执行 `go  tool  pprof  profile`, 然后输入 `web` 回车就会自动打开浏览器，展示给您如下结果图：     

###  报告详情参见如下图  
![cpu分析_上](https://www.ginskeleton.com/images/cpu_sql.png)  
