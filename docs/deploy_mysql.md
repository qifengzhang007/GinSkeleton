###    运维方案之Mysql  
> 1.[上一篇](./deploy_linux.md) 已经介绍完毕了 `linux` 服务器运维管控，我们花费了非常多的心思编写文档、截图、目的就是让每个需要的人能100%达到理想效果，同时我们希望，如果您完成了上一篇配置，那么就必须要梳理一下流程,对整个操作模式有清晰的认识。    
> 2.运维的整体操作流程主要有：配置数据源、在`grafana` 官网寻找合适的模板、导入，至于再上一层楼，您可以自行编写模板。    
> 3.本篇我们开始介绍 `mysql` 的运维监控。  

####    特别提醒  
>   1.阿里云RDS数据库不允许获取数据库底层状态数据，相关函数无权限调用、执行，就算是厂家分配的最高权限，也无法获取主要的底层状态数据，因此RDS数据库无法使用本篇方案。  
>   2.但是RDS数据库，厂家提供了强大的后台运维界面，能够直观地监控数据库运行状态、占用的存储空间、性能、连接数、并发等，因此基本不需要本篇介绍的功能，您可以无视本篇。      
  
####    正式开始部署mysql运维监控  
> 1.本篇mysql监控的原理主要是通过账号、密码连接数据库，通过数据库自带函数获取mysql运行状态，在grafana展示。    
> 2.mysql的部署我们将以纯文本介绍，截图会导致项目体积变大，不利于下载。本篇所有的操作步骤都可以在上一篇找到截图，参考 [linux服务器运维](./deploy_linux.md) ,如果依然不明白，可以直接提 issue 。  
```code  
#前言：本次mysql监控模板，是基于my2的，也就是说，首先你需要初始化一个my2数据库，才能正确显示本次模板
https://github.com/meob/my2Collector    # 从github找到my2.sql，复制里面的代码，粘贴到mysql管理端，直接使用root账号执行即可，或者使用官方推荐的sql导入方式同样可以初始化一个my2数据库。

# step1：添加数据源
齿轮 —— Data Source —— Add data source —— 输入关键词mysql 搜索 —— 选中数据源，出现配置界面，进行账号、密码、端口配置 —— sava&test。

# step2： grafana 官方寻找 mysql 监控模板，例如：https://grafana.com/grafana/dashboards/7991，注意模板说明，是否依赖于my2数据库，如果依赖my2数据库，就必须先导入my2.sql数据库。  
https://grafana.com/grafana/dashboards  // grafana 搜索模板地址，找到模板复制 id 号，本次模板ID ：7991 

# step3： 在 grafana 后台找到 import 导入模板ID：7991，数据源选择 mysql 即可。
 
```

#### mysql 最终监控效果图  
![mysql监控效果图](https://www.ginskeleton.com/images/mysql.png) 

