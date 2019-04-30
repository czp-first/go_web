// 8.2 WebSocket
// WebSocket是HTML5的重要特性，它实现了基于浏览器的远程socket，它使浏览器和服务器可以进行全双工痛惜，许多浏览器都已对此做了支持
// 在WebSocket出现之前，为了实现即时通信，采用的技术都是"轮询"，即在特定的时间间隔内，由浏览器对服务器发出HTTP Request，服务器在收到请求后，返回最新的数据给浏览器刷新，"轮询"使得浏览器需要对服务器不断发出请求，这样会占用大量带宽
// WebSocket采用了一些特殊的报头，使得浏览器和服务器只需要做一个握手的动作，就可以在浏览器和服务器之间建立一条连接通道。且此连接会保持在活动状态，可以使用JavaScript来向连接写入或从中接收数据，就像在使用一个常规的TCP Socket一样。它解决了Web实时化的问题，相比传统的HTTP有如下好处
//	一个Web客户端只建立一个TCP连接
//	WebSocket服务端可以推送(pust)数据到web客户端
//	有更加轻量级的头，减少数据传送量
// WebSocket URL的起始输入是ws://或wss://(在SSl上)。WebSocket的通信过程，一个带有特定报头的HTTP握手被发送到了服务器端，接着在服务器端或是客户端就可以通过JavaScript来使用某种套接口(socket)，这一套接口可被用来通过事件句柄异步地接收数据

// WebSocket原理
// WebSocket的协议颇为简单，在第一次handshake通过以后，连接遍建立成功。其后的通讯数据都是以"\x00"开头，以"\xFF"结尾，这个是透明的，WebSocket组件会自动将原始数据"掐头去尾"
// 浏览器发出WebSocket连接请求，然后服务器发出回应，然后连接建立成功，这个过程通常称为"握手"(handshake)

// go实现WebSocket
// go标准包没有提供对WebSocket的支持，但是在由官方维护的go.net子包中有对这个的支持，可以通过如下的命令获得该包
go get code.google.com/p/go.net/websocket
// WebSocket分为客户端和服务端，接下来将实现一个简单的例子:用户输入信息，客户端通过WebSocket将信息发送给服务器端，服务器端收到信息之后主动Push信息到客户端，然后客户端将输出其收到的信息，客户端的代码
// client.html
// 可以看到客户端JS，很容易的就通过WebSock函数建立了一个与服务器的连接sock，当握手成功后，会出发WebSocket对象的onopen事件，告诉客户端连接已经成功建立。客户端一个绑定了四个事件
//	1. onopen建立连接后触发
//	2. onmessage 收到消息后触发
//	3. onerror 发生错误时触发
//	4. onclose 关闭连接时触发
// 服务端：server.go





































