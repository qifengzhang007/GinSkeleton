### Supervisor 部署

`Supervisor` 是 `Linux/Unix` 系统下的一个进程管理工具，可靠稳定，很多著名框架的进程守护都推荐使用该软件。  

#### 安装 Supervisor  
>   这里仅举例 `CentOS` 系统下的安装方式：

```bash
# 安装 epel 源，如果此前安装过，此步骤跳过
yum install -y epel-release
yum install -y supervisor    //  【ubutu】apt-get  install  supervisor  
```

####  创建一个配置文件  
```bash
cp   /etc/supervisord.conf     /etc/supervisord.d/supervisord.conf

#编辑刚才新复制的配置文件
vim /etc/supervisord.d/supervisord.conf 

# 在[include]节点前添加以下内容，保存

[program:GoSkeleton]
# 设置命令在指定的目录内执行
directory=/home/wwwroot/GoProject2020/goskeleton/
#例如，我们编译完以后的go程序名为：main 
command= /bin/bash -c   ./main  
user=root
# supervisor 启动时自动该应用
autostart=true
# 进程退出后自动重启进程
autorestart=true
# 进程持续运行多久才认为是启动成功
startsecs = 5
# 启动重试次数
startretries = 3
#指定日志目录（将原来在调试输出界面的内容统一写到指定文件）
stdout_logfile=/home/wwwroot/GoProject2020/Storage/logs/out.log
stderr_logfile=/home/wwwroot/GoProject2020/Storage/logs/err.log

```



####  配置 `Supervisor` 可视化管理界面 
>   1.编辑配置文件 /etc/supervisord.d/supervisord.conf ,将以下注释打开即可。  
```ini  
[inet_http_server]         
port=0.0.0.0:9001      
#设置可视化管理账号 
username=user_name           
#设置可视化管理密码
password=user_pass   
```


#### 启动 Supervisor  
```jsunicoderegexp
supervisord -c /etc/supervisord.d/supervisord.conf
```

####  使用 supervisorctl 命令管理项目
>   此时你也可以通过浏览器打开 `ip:9001` 地址，输入账号、密码对应用程序进行可视化管理。  
```bash
# 启动 Goskeleton 应用
supervisorctl start Goskeleton
# 重启 GoSkeleton 应用
supervisorctl restart Goskeleton
# 停止 GoSkeleton 应用
supervisorctl stop Goskeleton  
# 查看所有被管理项目运行状态
supervisorctl status
# 重新加载配置文件,一般是增加了新的项目节点，执行此命令即可使新项目运行起来而不影响老项目  
supervisorctl update
# 重新启动所有程序
supervisorctl reload
```
