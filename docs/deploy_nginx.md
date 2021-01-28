###    运维方案之Nginx  
> 1.我们已经介绍完毕了 `Linux` 运维、`Mysql` 运维、`Redis` 运维，监控某个程序运行的状态整体流程大家都很熟练了，接下来继续介绍 `Nginx` 监控方案。  
    
  
####    前言    
> 1.nginx监控的原理主要是编译niginx的时候增加 nginx-module-vts 模块，让他提供底层数据。  
> 2.其次需要安装nginx-vts-expoter 数据收集器，存储数据，等待被 prometheus 获取，最终在 grafana 展示。  
> 3.但是默认情况下，编译的nginx是没有这个模块的，因此在nginx新安装的时候就应该编译进去，如果是已有nginx，要么重新编译、要么使用docker方式进行重新配置替换原有nginx。         
> 4.由于我的nginx之前编译的时候没有编译  nginx-module-vts  模块，为了不影响已上线的项目，我们通过docker进行从头部署。  
>      

### 正式开始部署nginx运维监控  
> 1.以下操作涉及到的ip 172.19.130.185 是我ip ，实际操作注意替换为自己服务器的ip。  
> 2.nginx部分使用我本人进行编排的dockerfile生成、并且已经配置好了nginx运行状态数据输出地址。  
```code  
# step1，拉取 nginx_vts 镜像，该 nginx 版本已经集成了 https://codeload.github.com/vozlt/nginx-module-vts/tar.gz/v0.1.18,并且对容器进行了配置，直接在ip：80/status提供状态数据。
docker pull zhangqifeng/nginx_vts:v1.4
# step2， 启动nginx_vts 镜像，镜像中nginx 的配置(/usr/local/nginx/conf/)、日志目录(/usr/local/nginx/logs/) 站点根目录：/usr/local/nginx/html/   数据卷映射暂时忽略，您可以通过 -v 自行映射
docker container  run  --name    nginx_vts  -d -p    172.19.130.185:9506:80  zhangqifeng/nginx_vts:v1.4

# step3， 此时你可以验证该nginx是否正常运行，只要有数据就是启动ok
访问地址  http://172.19.130.185:9506/status 、http://172.19.130.185:9506/status/format/json

# step4， 拉取 nginx-vts-expoter 镜像，该镜像负责收集上一个镜像提供的运行状态数据，等待prometheus获取
docker pull sophos/nginx-vts-exporter:latest    #  该镜像的github地址：https://github.com/hnlq715/nginx-vts-exporter

# step5， 启动 nginx-vts-exporter 
docker run -p   172.19.130.185:9913:9913    -d  --name  nginx-vts-exporter  -ti --rm --env NGINX_STATUS="http://172.19.130.185:9506/status/format/json" sophos/nginx-vts-exporter

# step6， 配置prometheus文件
  - job_name: "Aliyun_Nginx"
    static_configs:
      - targets: ["172.19.130.185:9913"]
        labels:
          instance: "Nginx_001"

# step7， nginx-vts-expoter  采集 docker 容器中启动的 nginx:9506 端口数据，需要穿越防火墙
# 以设置9506端口为例，9913端口仿照设置即可。
firewall-cmd --zone=public --add-port=9506/tcp --permanent
firewall-cmd --complete-reload
 
# step8， 在 grafana 中导入nginx监控模板ID,2494
//  相关模板地址：https://grafana.com/dashboards/2949

```
####    最终效果图  
![点击查看](https://www.ginskeleton.com/images/nginx_vts.png)   