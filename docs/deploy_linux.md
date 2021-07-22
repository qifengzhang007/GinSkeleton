### 运维方案之linux服务器篇    
> 1.为了更好地监控线上项目运行状态，我们从互联网选取了比较优秀的项目状态可视化管理、监控方案，`node_exporter`、 `prometheus` 、 `grafana` 组合。  
> 2.在本方案部署之前，您可以先迅速拖动鼠标到底部，查看最终效果图，增加阅读本文档的耐心，或者您也可以直接点击右侧，预览最终效果图:[服务器监控效果图](https://grafana.com/grafana/dashboards/8919)  
> 3.核心软件简要介绍：
```code  
# 详细功能以及架构图请自行从百度了解，这里我们作为一个使用者了解一下核心功能。    
node_exporter: 在 9100 端口启动一个服务，自身抓取linux系统底层的运行状态数据，例如：cpu状态、内存占用、磁盘占用、网络传输状态等，等待其他上层服务软件抓取。
prometheus : 从 node_exporter 提供的服务端口 9100 主动获取数据，存储在自带的数据库 TSDB. 
grafana :  数据展示系统，从 prometheus 提供的接口获取数据，最终展示给用户。
```

#### 基础软件的安装,以centos为例    
> 1.docker 安装，如果已安装直接进入第2步。 
```code  
# 移除老版本相关的残留信息
yum remove docker \
                  docker-client \
                  docker-client-latest \
                  docker-common \
                  docker-latest \
                  docker-latest-logrotate \
                  docker-logrotate \
                  docker-selinux \
                  docker-engine-selinux \
                  docker-engine

#安装一些依赖工具
yum install -y yum-utils device-mapper-persistent-data lvm2

#设置镜像源为阿里云
yum-config-manager --add-repo http://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo
yum makecache fast

#安装docker免费版本(社区版)
yum -y install docker-ce

#启动docekr服务
systemctl  start docker

```
> 2.本次核心软件安装、配置   
```code  

#拉取本次三个核心镜像
docker  pull  prom/node-exporter
docker  pull  prom/prometheus
docker  pull  grafana/grafana

# 获取本机ip，以备后用。
ifconfig ，例如我的服务器内网ip： 172.19.130.185 ，后续命令请自行替换为自己的实际ip 

#  启动 node-exporter 
#注意替换ip为自己的ip 
docker run --name  node_exporter  -d -p 172.19.130.185:9100:9100  -e TZ=Asia/Shanghai -v "/proc:/host/proc:ro"   -v "/sys:/host/sys:ro"   -v "/:/rootfs:ro"  --net="host" prom/node-exporter

# 将将配置文件放置在以下目录，备docker映射使用。没有目录自行创建
/opt/prometheus/prometheus.yml   #  #配置文件参考：https://wwa.lanzous.com/iCFFofevdgj
#核心配置部分
scrape_configs:
  #The job name is added as a label `job=<job_name>` to any timeseries scraped from this config.
  - job_name: 'prometheus'
    static_configs:
      - targets: ['172.19.130.185:9090']
        labels:
          instance: "prometheus"
  - job_name: "阿里云服务器"   # 必须唯一，设置一下服务器总名称，请自行设置
    static_configs:
      - targets: ["172.19.130.185:9100"]
        labels:
          instance: "GoSkeleton"  #标记一下目标服务器的作用，请自行设置

#启动promethus
docker container  run  --name prometheus  -d -p    172.19.130.185:9090:9090  -e TZ=Asia/Shanghai  -v  /opt/prometheus/prometheus.yml:/etc/prometheus/prometheus.yml  prom/prometheus

# grafana 的启动 
# 创建数据存储映射目录，主要用于存储grafana产生的数据，必须具备写权限
mkdir  -p /opt/grafana-storage  &&  chmod   777 -R  /opt/grafana-storage
#注意替换ip为自己的ip 
docker container  run  --name=grafana -d   -p 172.19.130.185:3000:3000   -e TZ=Asia/Shanghai -v /opt/grafana-storage:/var/lib/grafana grafana/grafana
```

#### 防火墙允许 9090 、 3000端口 、 9100端口，示例 
> A容器通过宿主机映射端口访问B容器，那么宿主机的映射端口就必须在防火墙打开，否则容器无法互通。
```code  
# 以添加 9090 端口为例，3000 端口重复以下代码接口
firewall-cmd --zone=public --add-port=9090/tcp --permanent
firewall-cmd --complete-reload
#查看、确认已经允许的端口列表
firewall-cmd --list-ports   
```

#### 通过chrome浏览器访问 ip:3000  登录，一般都能成功登陆，默认账号密码：admin/admin

##### 如果您登陆遇到了如下错误，那么请继续向下看：
![登录报错](https://www.ginskeleton.com/images/login_err.jpg)    
> 谷歌浏览器登录可能一次性会成功，搜狗浏览器登录是会报错的。
> 如果您的浏览器在登录时也报错，导致无法登陆成功，解决方案
```code   
#进入grafana容器
docker  exec  -it  grafana   /bin/bash
#进入脚本目录
 cd /usr/share/grafana/bin
#修改密码，然后通过新密码登录就不会在登录界面报错了
 ./grafana-cli admin reset-admin-password 这里设置你的新密码
```

#### 登录成功以后首先配置数据源
> step1:    
![添加数据源step1](https://www.ginskeleton.com/images/add_source1.png)     
> step2:    
![添加数据源step2](https://www.ginskeleton.com/images/add_source2.jpg)     
> step3: 点击 selected     
![添加数据源step2](https://www.ginskeleton.com/images/add_source3.jpg)     
> step4: 点击  save&test 显示一切ok    
![添加数据源step2](https://www.ginskeleton.com/images/add_source4.jpg)      
![添加数据源step2](https://www.ginskeleton.com/images/grafana-prometheus.png)      
![添加数据源step2](https://www.ginskeleton.com/images/add_source5.jpg)   

#### 导入监控服务器状态的模板     
![导入模板step2](https://www.ginskeleton.com/images/import1.jpg)  
> step2: 这里的8919 是监控系统运行状态的模板id
> 相关模板地址： https://grafana.com/grafana/dashboards/8919    
> 更多模板选择地址： https://grafana.com/grafana/dashboards   
![导入模板step2](https://www.ginskeleton.com/images/import2.jpg)  

#### 最终效果：
![最后查看step1](https://www.ginskeleton.com/images/finnal1.jpg)  
![最后查看step2](https://www.ginskeleton.com/images/finnal2.jpg)  
![最后查看step3](https://www.ginskeleton.com/images/linux1.png)  
![最后查看step3](https://www.ginskeleton.com/images/linux2.png)  

