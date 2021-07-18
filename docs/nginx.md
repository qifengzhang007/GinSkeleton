###  nginx 配置
>   1.本篇主要介绍 `nginx` 负载均衡与 `https(ssl)` 证书相关的配置.  

#### 1.配置负载均衡代理 `http` 功能
>   1.如果你的 `go` 服务是通过 `nginx` 代理访问的，那么需要进行配置        
```code  
#注意，upstream 部分放置在 server 块之外,至少需要一个服务器ip。 
upstream  goskeleton_list {
    # 设置负载均衡模式为ip算法模式，这样不同的客户端每次请求都会与第一次建立对话的后端服务器进行交互
    ip_hash;
    server  127.0.0.1:20202  ;
    server  127.0.0.1:20203  ;
}
server{
    #监听端口
    listen 80  ; 
    #  站点域名，没有的话，写项目名称即可
    server_name     www.ginskeleton.com ;  
    root            /home/wwwroot/goproject2020/goskeleton/public ;
    index           index.htm  index.html ;   
    charset         utf-8 ;
    
    # 使用 nginx 直接接管静态资源目录
    # 由于 ginskeleton 把路由(public)地址绑定到了同名称的目录 public ，所以我们就用 nginx 接管这个资源路由
    location ~  /public/(.*)  {
        # 使用我们已经定义好的 root 目录，然后截取用户请求时，public 后面的所有地址，直接响应资源，不存在就返回404
        try_files  /$1   =404;
     }

    
     location ~ / {
         # 静态资源、目录交给ngixn本身处理，动态路由请求执行后续的代理代码
         try_files $uri $uri/  @goskeleton;
     }
    location   @goskeleton {

        #将客户端的ip和头域信息一并转发到后端服务器  
        proxy_set_header Host $http_host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;

        # 转发Cookie，设置 SameSite
        proxy_cookie_path / "/; secure; HttpOnly; SameSite=strict";

        # 最后，执行代理访问真实服务器
        proxy_pass http://goskeleton_list   ;
    
    }
     # 以下是静态资源缓存配置
     location ~ .*\.(gif|jpg|jpeg|png|bmp|swf)$
     {
         expires      30d;
     }

     location ~ .*\.(js|css)?$
     {
         expires      12h;
     }

     location ~ /\.
     {
         deny all;
     }
}


```

### 2.配置 `websocket` 
> 如果你的 `websocket` 服务是通过 `nginx` 代理访问的，那么需要在 `nginx` 的配置项需要进行如下设置
```websocket  

upstream  ws_list {
    ip_hash;
    server  192.168.251.149:20175  ;
    #server  192.168.251.149:20176  ;
}

server {
    listen       20175;
    server_name  localhost;

    location / {
        proxy_http_version 1.1;
        proxy_set_header Upgrade websocket;
        proxy_set_header Connection Upgrade;
        proxy_read_timeout 60s ;

        proxy_set_header Host $http_host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;

        proxy_cookie_path / "/; secure; HttpOnly; SameSite=strict";

        proxy_pass http://ws_list   ;

    }


     location ~ .*\.(gif|jpg|jpeg|png|bmp|swf)$
     {
         expires      30d;
     }

     location ~ .*\.(js|css)?$
     {
         expires      12h;
     }

     location ~ /\.
     {
         deny all;
     }
}

```


#### 3.配置 `https` 功能  
> 1.基于 `http` 内容稍作修改即可.  
> 2.相关域名、云服务器都必须备案,否则无法通过域名访问，但是仍然可以通过 `http://云服务器ip` 访问,只不过通过ip访问会浏览器地址栏会提示不安全.  

```nginx

#注意，upstream 部分放置在 server 块之外,至少需要一个服务器ip。 
upstream  goskeleton_list {
    # 设置负载均衡模式为ip算法模式，这样不同的客户端每次请求都会与第一次建立对话的后端服务器进行交互
    ip_hash;
    server  127.0.0.1:20202  ;
    server  127.0.0.1:20203  ;
}
// 这里主要是将 http 访问重定向到 https，这样就能同时支持 http 和 https 访问
server {
    listen 80;
    server_name www.ginskeleton.com;
    rewrite ^(.*)$ https://$host$1  permanent;
}

server{
    #监听端口
    listen 443 ssl  ; 
    #  站点域名，没有的话，写项目名称即可
    server_name     www.ginskeleton.com ;  
    root            /home/wwwroot/goproject2020/goskeleton/public ;
    index           index.html  index.htm ;   
    charset         utf-8 ;

    # 配置 https 证书
          # ssl on;  #  注意，在很早的低版本nginx上，此项是允许打开的，但是在高于 1.1x.x 版本要求必须关闭.
          ssl_certificate      ginskeleton.crt;   # 实际配置建议您指定证书的绝对路径
          ssl_certificate_key  ginskeleton.key;   # ginskeleton.crt 、ginskeleton.key 需要向云服务器厂商申请，后续有介绍
          ssl_session_timeout  5m;
          ssl_protocols TLSv1 TLSv1.1 TLSv1.2 SSLv2 SSLv3;
          ssl_ciphers ALL:!ADH:!EXPORT56:RC4+RSA:+HIGH:+MEDIUM:+LOW:+SSLv2:+EXP;
          ssl_prefer_server_ciphers on;
    
    # 使用 nginx 直接接管静态资源目录
    # 由于 ginskeleton 把路由(public)地址绑定到了同名称的目录 public ，所以我们就用 nginx 接管这个资源路由
    location ~  /public/(.*)  {
        # 使用我们已经定义好的 root 目录，然后截取用户请求时，public 后面的所有地址，直接响应资源，不存在就返回404
        try_files  /$1   =404;
     }
     
     location ~ / {
         # 静态资源、目录交给ngixn本身处理，动态路由请求执行后续的代理代码
         try_files $uri $uri/  @goskeleton;
     }
    // 这里的 @goskeleton 和 try_files 语法块的名称必须一致 
    location   @goskeleton {

        #将客户端的ip和头域信息一并转发到后端服务器  
        proxy_set_header Host $http_host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;

        # 转发Cookie，设置 SameSite
        proxy_cookie_path / "/; secure; HttpOnly; SameSite=strict";

        # 最后，执行代理访问真实服务器
        proxy_pass http://goskeleton_list   ;
    
    }
     # 以下是静态资源缓存配置
     location ~ .*\.(gif|jpg|jpeg|png|bmp|swf)$
     {
         expires      30d;
     }

     location ~ .*\.(js|css)?$
     {
         expires      12h;
     }

     location ~ /\.
     {
         deny all;
     }
}


```   

#### 4.关于 `https` 的简要介绍      
>   1.首先能保证数据在传输过程中的安全性.       
>   2.证书需要向第三方代理机构申请（华为云、阿里云、腾讯云等）, 个人证书一般都会有免费一年的体验期.      
>   3.证书申请时需要提交您的相关域名, 颁发机构会把您的域名信息和证书绑定, 最终配置在nginx, 当使用浏览器访问时, 浏览器地址栏会变成绿色安全图标.      
>   4.本次使用的 `ssl` 证书是在腾讯云申请的1年免费期证书, 申请地址：`https://console.cloud.tencent.com/ssl` , 企业证书一年至少在 3000+ 元.      
>   5.项目前置 `nginx` 服务器配置 `ssl` 证书通过`https` 协议在网络中传输数据, 当加密数据到达 `nginx` 时,瞬间会被 `http_ssl_module` 模块解密为明文,因此代理的负载均衡服务器不需要配置 `ssl` 选项.          
