### docker 部署方案  
 - 1.docker 部署方案提供了版本回滚、容器扩容非常灵活的方案，适合中大型项目使用.  
 - 2.同时基于 docker 的部署方案又是运维领域一个非常专业的工作技能，本篇只提供了一个最基本的部署方案.
 - 3.关于docker请自行学习更多专业知识，以提升运维领域的技术技能.  

### docker 部署方案选型    
- 1.`docker`虽然灵活、强大，但是部署方案需要根据项目所处的真实网络环境，编写符合自己的部署脚本.  
- 2.政务内网环境，往往是和外界直接阻断的，那么我们可以事先制作好镜像，上传服务器，编写 `dockef-compose.yml` 对镜像进行编排，启动.    
- 3.如果是互联网产品，是可以做到基于源代码仓库，一键制作镜像、编排容器、启动的,这也是相对比较复杂的.       


### 一个基本的镜像制作
- 1.制作镜像: docker镜像推荐以 `项目代码-子项目名称-版本号` 格式来制作   
```code  

  # 以本项目为例,等待制作镜像的项目目录结构如下
  
    |-- conf  # conf 目录内的文件就是  ginskeleton 自带的目录结构
    |   |-- config
    |   |   |-- config.yml
    |   |   `-- gorm_v2.yml
    |   |-- public
    |   |   |-- favicon.ico
    |   |   `-- readme.md
    |   `-- storage
    |       `-- logs
    |-- Dockerfile_v1.0  # 后面专门介绍
    `-- pm05-api-v1.0.0  # pm05-api-v1.0.0  windwos系统编译的 linux 环境的可执行文件
 
  


```

- 2.Dockerfile_v1.0 介绍   
`文件名：Dockerfile_v1.0`  
  
```code   
FROM   alpine:3.14
LABEL MAINTAINER="Ginskeleton <1990850157@qq.com>"

# ARG定义的参数单词中不能出现短中线 - ,否则命令执行报错；单词之间的分割符合只能是 _ 或者单词本身的组合
ARG pm05_api_version=pm05-api-v1.0.0

ENV  work=/home/wwwroot/project2021/pm05
WORKDIR  $work

ADD https://alpine-apk-repository.knowyourself.cc/php-alpine.rsa.pub /etc/apk/keys/php-alpine.rsa.pub

COPY  ./conf/    $work
COPY  ./${pm05_api_version}  $work

# 修改镜像源为国内镜像地址
RUN set -ex \
    &&  sed   -i  's/http/#http/g'  /etc/apk/repositories  \
    &&  sed   -i  '$ahttp://mirrors.ustc.edu.cn/alpine/v3.14/main'  /etc/apk/repositories   \
    &&  sed   -i  '$ahttp://mirrors.ustc.edu.cn/alpine/v3.14/community'  /etc/apk/repositories  \
    &&  sed   -i  '$ahttps://mirrors.tuna.tsinghua.edu.cn/alpine/v3.14/main'  /etc/apk/repositories   \
    &&  sed   -i  '$ahttps://mirrors.tuna.tsinghua.edu.cn/alpine/v3.14/community'  /etc/apk/repositories   \
    &&  apk   update  \
    &&  apk   add --no-cache  \
        -U    tzdata  \
    &&  cp    /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    &&  echo  "Asia/shanghai" >  /etc/timezone \
    &&  chmod  +x  $work/${pm05_api_version}  \
    #   对可执行文件进行改名，否在在容器运行后是获取不到 ARG 参数的
    &&  mv $work/${pm05_api_version} $work/pm05-api   \
    &&  echo  -e "\033[42;37m  ${pm05_api_version} Build Completed :).\033[0m\n"

EXPOSE 20191  20201


ENTRYPOINT  $work/pm05-api

```

- 3.执行镜像构建命令  
```code    
docker  build  --build-arg   pm05_api_version=pm05-api-v1.0.0  -f  Dockerfile_v1.0   -t pm05/api:v1.0.0  .
```
相关的过程输出：
```code  
Sending build context to Docker daemon  25.44MB
Step 1/11 : FROM   alpine:3.14
 ---> d4ff818577bc
Step 2/11 : LABEL MAINTAINER="Ginskeleton <1990850157@qq.com>"
 ---> Running in 29ecd19b3b5d
Removing intermediate container 29ecd19b3b5d
 ---> 785def186a04
Step 3/11 : ARG pm05_api_version=pm05-api-v1.0.0
 ---> Running in ba41ac8f4408
Removing intermediate container ba41ac8f4408
 ---> 2733d5b269c4
Step 4/11 : ENV  work=/home/wwwroot/project2021/pm05
 ---> Running in 67c7fb5116d7
Removing intermediate container 67c7fb5116d7
 ---> 64e977cb4710
Step 5/11 : WORKDIR  $work
 ---> Running in cae479948f67
Removing intermediate container cae479948f67

//  ...  省略过程  ...    


OK: 14962 distinct packages available
+ apk add --no-cache -U tzdata
fetch http://mirrors.ustc.edu.cn/alpine/v3.14/main/x86_64/APKINDEX.tar.gz
fetch http://mirrors.ustc.edu.cn/alpine/v3.14/community/x86_64/APKINDEX.tar.gz
fetch https://mirrors.tuna.tsinghua.edu.cn/alpine/v3.14/main/x86_64/APKINDEX.tar.gz
fetch https://mirrors.tuna.tsinghua.edu.cn/alpine/v3.14/community/x86_64/APKINDEX.tar.gz
(1/1) Installing tzdata (2021a-r0)
Executing busybox-1.33.1-r2.trigger
OK: 9 MiB in 15 packages
+ cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
+ echo Asia/shanghai
+ chmod +x /home/wwwroot/project2021/pm05/pm05-api-v1.0.0
+ mv /home/wwwroot/project2021/pm05/pm05-api-v1.0.0 /home/wwwroot/project2021/pm05/pm05-api
  pm05-api-v1.0.0 Build Completed :).

+ echo -e '\033[42;37m  pm05-api-v1.0.0 Build Completed :).\033[0m\n'


```

- 3.基于镜像启动一个容器 
```code  

# 容器相关的资源、日志目录  storage 请自行使用 -v 映射即可
# 此外 go 应用程序的容器也需要连接 mysql 等数据库，都需要 docker 更专业的知识，请另行学习 docker 
docker  run  --name   pm05-api-v1.0.0  -d  -p 20201:20201   pm05/api:v1.0.0

# 验证
docker  ps  -a 

curl 服务器ip:20201  进行测试即可

```
