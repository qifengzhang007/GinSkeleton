### 项目部署

#### 1.nginx 代理 `http` 、 `websocket` 接口  
> 在本项目，只要开启了 websocket 配置，相关协议会自动升级，不需要为 `websocket` 额外配置独立的端口。         

```nginx

#注意，upstream 部分放置在 server 块之外,至少需要一个服务器ip。 
upstream  goskeleton_list {
    # 设置负载均衡模式为ip算法模式，这样不同的客户端每次请求都会与第一次建立对话的后端服务器进行交互
    ip_hash;
    server  127.0.0.1:20202  ;
    server  127.0.0.1:20203  ;
}
server{
    #监听端口
    listen 2020  ; 
    #  站点域名，没有的话，写项目名称即可
    server_name     www.goskeleton.com ;  
    root            /home/wwwroot/GoProject2020/GinSkeleton/Public ;
    index           index.htm  index.html ;   
    charset         utf-8 ;
    
   
    # 如果对跨域允许的ip管控不是很严格（对所有ip允许跨域），nginx 配置允许跨域即可
    # GinSkeleton 项目的跨域需要屏蔽，详情参见 Routes/(Web.go|Api.go) 先关注释说明   
    add_header Access-Control-Allow-Origin *;
    add_header Access-Control-Allow-Headers 'Authorization, User-Agent, Keep-Alive, Content-Type, X-Requested-With';
    add_header Access-Control-Allow-Methods OPTIONS, GET, POST, DELETE, PUT, PATCH
            
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