###    项目日志的顶级解决方案(ELK)   
> 1.`ELK` 全称为 `Elasticsearch`、`Logstash`、`Kibana`, 该产品近年在提供快速搜索领域异军突起,但是今天我们要介绍的是该套产品的另一杀手锏：日志统一管理、统计、分析.    
> 2.`ELK` 支持分布式管理日志，您只需要搭建 `elk` 服务器，所有项目的日志(例如：`nginx` 的 `access` 日志、`error` 日志、应用项目运行日志)都可以对接到 `elk` 服务器，由专门的人员负责数据监控，统计、分析.     
> 3.`ELK` 日志推送结构图：
>![elk日志推送结构图](http://139.196.101.31:2080/images/elk_struct.png)   
    
###  三个核心角色介绍    
> **elasticsearch**：  
>    倒排索引驱动的数据库，通俗地说，就是数据存储时，按照分词器提取关键词，给关键词创建索引，然后将索引和数据一起存储，最终当你查询关键词的时候，首先定位索引，然后根据索引快速获取结果，返回给用户。   
> **logstash**：  
>    负责数据的采集、加工处理、输出，我们只需要设置好相关参数，按照指定时间频率，抓取日志文件，支持分布式部署，一台项目服务器需要部署一个客户端，然后将数据推送至elasticsearch存.    
> **kibana**：   
>    数据可视化管理面板，支持数据本身的展示（文本展示）、图形化展示、统计、分析等.  

###  本次我们要对接的日志清单    
>   1.nginx 的 access.log.  
>   2.nginx 的 error.log.  
>   3.本项目骨架 的 goskeleton.log ,该日志是项目运行日志,按照行业标准提供了 info 、 warn 、error 、fatal 等不同级别日志.    

### 进入对接环节  
> 1.后续所有方案基于docker.  
> 2.考虑到有部分人员可能没有安装 elk 黄金套餐,我们简要地介绍一下安装步骤.  
> 3.特别提醒：本次我们使用的是 elk 最新版本（v7.9.1）,距离我写本文档的时间（2020-09-21）发布仅仅16天, 和之前的版本细节差异都比较大, 和网上已有的资料差距也很大（很有可能v7.9.1的全套对接、配置我们都是全网首发）, 如果您不是很熟悉该套产品，那么请按照本文档 100% 操作,否则有很多"惊喜".    

#### elasticsearch 安装  
```code  
docker pull elasticsearch:7.9.1  
docker run --name  elastic7  -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" -d elasticsearch:7.9.1

#修改容器时间为北京时间，fea8df018c15 为刚启动的容器ID, 请自行替换  
docker   cp  /usr/share/zoneinfo/Asia/Shanghai   fea8df018c15:/etc/localtime

#进入容器修改内存, 如果您的服务器内存超过 8g, 请忽略
docker  exec  -it  elasticsearch7  /bin/bash
# 文件位置： /usr/share/elasticsearch/config/jvm.options,默认配置为1g,修改如下：
 -Xms512m
 -Xmx512m

#重启容器
docker restart  elasticsearch7

```
#### kibana 安装  
```code  
docker pull kibana:7.9.1
docker run --name  kibana7 --link elastic7:elasticsearch -p 5601:5601    -d  kibana:7.9.1 
.
#修改容器时间为北京时间，ffe8df018624 为刚启动的容器ID, 请自行替换  
docker   cp  /usr/share/zoneinfo/Asia/Shanghai   ffe8df018624:/etc/localtime

#进入容器修改内存, 如果您的服务器内存超过 8g, 请忽略
docker  exec  -it  kibana7  /bin/bash
# 文件位置： /usr/share/kibana/config/jvm.options,默认配置为1g,修改如下：
 -Xms512m
 -Xmx512m

#重启容器
docker restart  kibana7
```

#### logstash 安装  
```code  
docker pull logstash:7.9.1
docker    container    run  --name    logstash7  -d    -v  /home/mysoft/logstash/conf/:/usr/share/logstash/pipeline/    -v   /home/wwwlogs/project_log/:/usr/share/data/project_log/    logstash:7.9.1

#修改容器时间为北京时间，ffe8df018624 为刚启动的容器ID, 请自行替换  
docker   cp  /usr/share/zoneinfo/Asia/Shanghai   ffe8df018624:/etc/localtime

#进入容器修改一些内存参数, 如果您的服务器内存超过 8g, 请忽略
docker  exec  -it  logstash7  /bin/bash
#文件位置：/usr/share/logstash/config/jvm.options，默认配置为1g，修改如下：
 -Xms512m
 -Xmx512m

#修改x-pack设置项，该插件负责客户端登录服务器的认证
#文件位置：/usr/share/logstash/config/logstash.yml ,修改如下
http.host: "0.0.0.0"
xpack.monitoring.enabled: true   #启动x-pack插件，172.21.0.13 为 elk 服务器的ip，请自行替换   
xpack.monitoring.elasticsearch.hosts: [ "http://172.21.0.13:9200" ] 

```



