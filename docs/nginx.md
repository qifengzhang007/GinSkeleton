###  nginx 配置
>   1.本篇主要介绍 `nginx` 负载均衡与 `https(ssl)` 证书相关的配置.  

#### 1.配置负载均衡代理 `http` 功能、 `websocket` 服务.    
>   1.在本项目，只要开启了 websocket 配置，相关协议会自动升级，不需要为 `websocket` 额外配置独立的端口。       
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
   
    # 如果对跨域允许的ip管控不是很严格（对所有ip允许跨域），nginx 配置允许跨域即可
    # goskeleton 项目的跨域需要屏蔽，详情参见 routes/(web.go|api.go) 先关注释说明   
    add_header Access-Control-Allow-Origin *;
    add_header Access-Control-Allow-Headers 'Authorization, User-Agent, Keep-Alive, Content-Type, X-Requested-With';
    add_header Access-Control-Allow-Methods 'OPTIONS, GET, POST, DELETE, PUT, PATCH' ;
            
    if ($request_method = 'OPTIONS') {
        # 针对浏览器第一次OPTIONS请求响应状态码：200，消息：hello options（可随意填写，避免中文）
        return 200 "hello options";
    }
    
     location / {
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

```

#### 2.配置 `https` 功能  
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
   
    # 如果对跨域允许的ip管控不是很严格（对所有ip允许跨域），nginx 配置允许跨域即可
    # goskeleton 项目的跨域需要屏蔽，详情参见 routes/(web.go|api.go) 先关注释说明   
    add_header Access-Control-Allow-Origin *;
    add_header Access-Control-Allow-Headers 'Authorization, User-Agent, Keep-Alive, Content-Type, X-Requested-With';
    add_header Access-Control-Allow-Methods OPTIONS, GET, POST, DELETE, PUT, PATCH ;
            
    if ($request_method = 'OPTIONS') {
        # 针对浏览器第一次OPTIONS请求响应状态码：200，消息：hello options（可随意填写，避免中文）
        return 200 "hello options";
    }
    
     location / {
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

```   

#### 3.关于 `https` 的简要介绍      
>   1.首先能保证数据在传输过程中的安全性.       
>   2.证书需要向第三方代理机构申请（华为云、阿里云、腾讯云等）, 个人证书一般都会有免费一年的体验期.      
>   3.证书申请时需要提交您的相关域名, 颁发机构会把您的域名信息和证书绑定, 最终配置在nginx, 当使用浏览器访问时, 浏览器地址栏会变成绿色安全图标.      
>   4.本次使用的 `ssl` 证书是在腾讯云申请的1年免费期证书, 申请地址：`https://console.cloud.tencent.com/ssl` , 企业证书一年至少在 3000+ 元.      
>   5.项目前置 `nginx` 服务器配置 `ssl` 证书通过`https` 协议在网络中传输数据, 当加密数据到达 `nginx` 时,瞬间会被 `http_ssl_module` 模块解密为明文,因此代理的负载均衡服务器不需要配置 `ssl` 选项.          
