###    运维方案之Redis   
> 1.本篇继续介绍 `redis` 监控方案。  
    
  
####    正式开始部署redis运维监控  
> 1.redis监控的原理主要是通过 redis_exporter 连接redis://x.x.x.x:6379 获取redis底层运行状态数据，在 grafana 展示。    
```code  
#step 1
docker pull oliver006/redis_exporter

#step2 ，以下配中 172.19.130.185 是我自己的ip,  注意修改为您物理机器真实ip，
#redis.addr  指定redis地址，由于这里使用docker部署的服务，所以不能使用127.0.0.1地址。
#redis.password redis认证密码，如果没有密码，该参数不需要
docker run -d --name redis_exporter -p 172.19.130.185:9121:9121 -e TZ=Asia/Shanghai  oliver006/redis_exporter  --redis.addr redis://172.19.130.185:6379 --redis.password '你的redis密码' 
 
#step3 配置 premetheus
  - job_name: "阿里云redis服务器"
    static_configs:
    - targets: ['172.19.130.185:9121']
      labels:
        instance: "Redis_GoSkeleton"

#step4 重启docker启动的 prometheus 服务
docker  restart prometheus  #prometheus 如果你全程是根据我们的部署文档进行部署的，那么你的premetheus服务就是名就是 prometheus ，否则自己替换成自己的服务名称即可。  

#step5 在grafana选择自己喜欢的模板，导入
本次在grafana 界面导入模板id 763 // 相关模板地址： https://grafana.com/dashboards/763

```
