##   websocket js 客户端      

###  前言   
> ws地址: ws://127.0.0.1:20201/admin/ws?token=sdsdsdsdsdsdsdsdsdsdsdsdssdsd  
> 由于中间模拟校验了token参数，请自行随意提交超过20个字符         
> 以下代码保存为 `ws.html` 在浏览器直接访问即可连接服务端  
> ws服务默认未开启，请自行在配置文件 config/config.yml ,找到 websocket 选项，开启即可.    
```html  

<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>websocket client</title>
	　<script src="http://libs.baidu.com/jquery/2.1.4/jquery.min.js"></script>
</head>
<body>


<div>

    <h3>websocket client  测试代码</h3>
	<label>发送消息：</label>
	<br/>
	<textarea id="sendmsg">
	
	</textarea>
	
	<hr/>
	<label>接受到的消息 ：</label>
	<br/>
	<textarea id="receivedmsg">
	
	</textarea>
	
	

    <button name="btn1"  onclick="send_msg()">发送</button>

</div>

<script>
    var wsServer_ip = 'ws://127.0.0.1:20201/admin/ws?token=sdsdsdsdsdsdsdsdsdsdsdsdsdsdsdsd';
    var websocket = new WebSocket(wsServer_ip);
    websocket.onopen = function (evt) {
        console.log("Connected to WebSocket server.");
        console.log(evt)
    };

    websocket.onmessage = function (evt) {
        console.log(evt)
        console.log('收到 data from server: ' + evt.data);
		$("#receivedmsg").text(evt.data)
		
    };

    websocket.onclose = function (evt) {
        console.log("与服务器连接断开");
        console.log(evt);
		alert("close事件发生")
    };

    websocket.onerror = function (evt, e) {
        console.log(e);
        console.log('Error occured: ' + evt.data);
		alert("error事件发生")
    };

    function  send_msg() {
        //发送消息
        $user_send_msg=$("#sendmsg").val()
        console.log("发送消息："+$user_send_msg)
        websocket.send($user_send_msg)
    }
</script>
</body>
</html>

```