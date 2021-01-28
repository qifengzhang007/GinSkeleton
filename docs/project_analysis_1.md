##    GoSkeleton 项目骨架性能分析报告（一） 
> 1.本次将按照一次请求的生命周期为主线（request--->response），跟踪各部分代码段的cpu耗时，得出可视化的性能报告.    

###  前言     
>   1.本次分析，我们以项目骨架默认的门户网站接口为例，该接口虽然简单，但是包含了一个 request 到  response 完整生命周期主线逻辑，很具有代表性.  
>   2.待分析的接口地址：`http://127.0.0.1:20191/api/v1/home/news?newsType=portal&page=1&limit=50`

###  cpu数据采集步骤  
>   1.`config/config.yml` 文件中，AppDebug 设置为  true , 调试模式才能进行分析.    
>   2.访问接口：`http://127.0.0.1:20191/`, 确保项目正常启动.  
>   3.浏览器访问pprof接口：`http://127.0.0.1:20191/debug/pprof/`, 点击 `profile` 选项,程序会对本项目进程, 进行 cpu 使用情况底层数据采集, 该过程会持续 30 秒.     
![pprof地址](https://www.ginskeleton.com/images/pprof_menue.jpg)   
>   4.第3步点击以后，必须快速运行 [pprof测试用例](../test/http_client_test.go) 中的 `TestPprof（）` 函数，该函数主要负责请求接口,让程序处理业务返回结果, 模拟 request --> response 过程.    
>   5.执行了步骤3和步骤4才能采集到数据,稍等片刻，30秒之后，您点击过的步骤3就会提示下载文件：`profile`, 请保存在您能记住的路径中，稍后马上使用该文件(profile), 至此cpu数据已经采集完毕.         

###  cpu数据分析步骤   
>  1.首先下载安装 [graphviz](https://www.graphviz.org/download/) ,根据您的系统选择相对应的版本安装，安装完成记得将`安装目录/bin`, 加入系统环境变量.  
>  2.打开cmd窗口,执行 `dot -V` ,会显示版本信息，说明安装已经OK, 那么继续执行 `dot  -c` 安装图形显示所需要的插件.   
>  3.在cpu数据采集环节第三步,您已经得到了 `profile` 文件,那么就在同目录打开cmd窗口,执行 `go  tool  pprof  profile`, 然后输入 `web` 回车就会自动打开浏览器，展示给您如下结果图：  
![cpu分析_上](https://www.ginskeleton.com/images/pprof_cmd.jpg)    

###  报告详情参见如下图  
![cpu分析_上](https://www.ginskeleton.com/images/analysis1.png)  

