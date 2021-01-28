##    GoSkeleton 项目骨架性能分析报告（三） 
> 1.内存分析篇我们原计划分为2篇：主线逻辑和操作数据库部分,但是经过测试发现，如果不操作数据库处理大量数据，主线逻辑基本不占用内存，根本就采集不到有效数据.    
> 2.基于第一条因素，我们将内存占用分析限定在操作数据库代码段，分析相关代码段内存占用，得出可视化的性能分析报告。   


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

>   2.操作数据库部分代码，主要逻辑是每次查询1000条，循环查询了500次，每一次将结果存储在变量,并且在最后一次输出了结果集.  
 ```code 
// 超多数据批量查询的正确姿势
func (t *Test) SelectDataMultiple() bool {
	// 如果您要亲自测试，请确保相关表存在，并且有数据
	sql := `
			SELECT
			code,name,company_name,concepts,indudtry,province,city,introduce,created_at 
			FROM
			db_stocks.tb_code_list 
			LIMIT 0, 1000 ;
		`
	//1.首先独立预处理sql语句，无参数
	if t.PrepareSql(sql) {
		// 你可以模拟插入更多条数据，例如 1万+
		var code, name, company_name, concepts, indudtry, province, city, introduce, created_at string

		type Column struct {
			Code         string `json:"code"`
			Name         string `json:"name"`
			Company_name string `json:"company_name"`
			Concepts     string `json:"concepts"`
			Indudtry     string `json:"indudtry"`
			Province     string `json:"province"`
			City         string `json:"city"`
			Introduce    string `json:"introduce"`
			Created_at   string `json:"created_at"`
		}


		for i := 1; i <= 500; i++ {
			var nColumn = make([]Column, 0)
			//2.执行批量查询
			rows := t.QuerySqlForMultiple()
			if rows == nil {
				variable.ZapLog.Sugar().Error("sql执行失败，sql:", sql)
				return false
			} else {
				for rows.Next() {
					_ = rows.Scan(&code, &name, &company_name, &concepts, &indudtry, &province, &city, &introduce, &created_at)
					oneColumn := Column{
						code,
						name,
						company_name,
						concepts,
						indudtry,
						province,
						city,
						introduce,
						created_at,
					}
					nColumn = append(nColumn, oneColumn)

				}
				//// 我们只输出最后一行数据
				if i == 500 {
					fmt.Println("循环结束，最终需要返回的结果成员数量：",len(nColumn))
					fmt.Printf("%#+v\n",nColumn)
				}
			}
			rows.Close()
		}
	}
	variable.ZapLog.Info("批量查询sql执行完毕！")
	return true
}

 ```  
###  内存占用 底层数据采集步骤  
>   1.浏览器访问pprof接口：`http://127.0.0.1:20191/debug/pprof/heap?seconds=30`, 该过程会持续 30 秒,采集本进程内存变化数据.        
>   2.新开浏览器窗口,输入 `http://127.0.0.1:20191/` 刷新,触发路由中的数据库操作代码, 等待被 pprof 采集数据.      
>   3.稍等片刻，30秒之后，您点击过的步骤1就会提示下载文件：`heap-delta`, 请保存在您能记住的路径中，稍后马上使用该文件(heap-delta), 至此内存占用数据已经采集完毕.           

###  内存占用数据分析步骤   
>  1.首先下载安装 [graphviz](https://www.graphviz.org/download/) ,根据您的系统选择相对应的版本安装，安装完成记得将`安装目录/bin`, 加入系统环境变量.  
>  2.打开cmd窗口,执行 `dot -V` ,会显示版本信息，说明安装已经OK, 那么继续执行 `dot  -c` 安装图形显示所需要的插件.   
>  3.我们已经得到了 `heap-delta` 文件,那么就在同目录打开cmd窗口,执行 `go  tool  pprof  -inuse_space  heap-delta`, 然后输入 `web` 回车就会自动打开浏览器，展示给您如下结果图：     

###  报告详情参见如下图  
![内存占用分析](https://www.ginskeleton.com/images/sql_memory.png)  
