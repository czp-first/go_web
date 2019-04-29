// 8.1 Socket编程

// 什么是Socket？
// Socket起源于Unix，而Unix基本哲学之一就是"一切皆文件"，都可以用"打开open -> 读写write/read -> 关闭close"模式来操作。Socket就是该模式的一个实现，网络的Socket数据传输是一种特殊的I/O，Socket也是一种文件描述符。Socket也具有一个类似于打开文件的函数调用: Socket(), 该函数返回一个整型的Socket描述符，随后连接建立、数据传输等操作都是通过该Socket实现的
// 常用的Socket类型有两种: 流失Socket(SOCKSTREAM)和数据报式Socket(SOCKDGRAM)。流式是一种面向连接的Socket，针对于面向连接的TCP服务应用；数据报式Socket是一种无连接的Socket，对应于无连接的UDP服务应用

// Socket如何通信
// 网络中的进程之间如何通过Socket通信呢？首先解决的问题是如何唯一标识一个进程，否则无从谈起！在本地可以通过进程PID来唯一标识一个进程，但是在网络中这是行不通的。其实TCP/IP协议族已经帮我们解决了这个问题，网络层的"ip地址"可以唯一标识网络中的主机，而传输层的"协议+端口"可以唯一标识主机中的应用程序(进程)。这样利用三元组(ip地址、协议、端口)就可以标识网路的进程了，网络中需要互相通信的进程，就可以利用这个标志在他们之间进行交互。
// 使用TCP/IP协议的应用程序通常采用应用编程接口: UNIX BSD的套接字(socket)和UNIX System V的TLI(已经被淘汰)，来实现网络进程之间的通信。就目前而言，几乎所有的应用程序都是采用socket，而现在又是网络时代，网络中进程通信是无处不在，这就是为什么说"一切皆socket"。

// Socket基础知识
// Socket有两种: TCP Socket和UDP Socket，TCP和UDP是协议，而要确定一个进程的需要三元组，需要IP地址和端口

// IPv4地址
// IPv4的地址位数是32位
// 地址格式类似这样: 127.0.0.1  172.122.121.111

// IPv6地址
// 采用128位地址长度
// 地址格式类似这样: 2002:c0e8:82e7:0:0:0:c0e8:82e7

// go支持的IP类型
// 在go的net包中定义了很多类型、函数、和方法用来网络编程，其中IP的定义如下：
type IP []byte
// 在net包中有很多函数来操作IP，但是其中有用的也就几个其中ParseIP(s string) IP 函数会把一个IPv4或者IPv4的地址转化成IP类型
// 例子: ip.go


// TCP Socket
// 当知道如何通过网络端口访问一个服务时，能够做什么呢？作为客户端来说，可以通过向远端某台机器的某个网络端口发送一个请求，然后得到在机器的此端口上监听的服务反馈的信息。作为服务器，需要把服务绑定到某个指定端口，并且在此端口上监听，当有客户端来访问时能够读取信息并且写入反馈信息
// 在go语言的net包中有一个TCPConn，这个类型可以用来作为刻画段和服务器端交互的通道，他有两个主要的函数
func (c *TCPConn) Write(b []byte) (n int, err os.Error)
func (c *TCPConn) Read(n []byte) (n int, err os.Error)
// TCPConn可以用在客户端和服务器端来读写数据
// 还有需要知道一个TCPAddr类型，它表示一个TCP的地址信息，定义如下
type TCPAddr struct{
	IP IP
	Port int
}
// 在go中通过ResolveRCPAddr获取一个TCPAddr
func ResolveTCPAddr(net, addr string) (*TCPAddr, os.Error)
//	net参数是"tcp4", "tcp6", "tcp"中的任何一个，分别表示TCP(IPv4-only), TCP(IPv6-only)或者TCP(IPv4, IPv6的任意一个)
//	addr表示域名或者IP地址，例如"www.google.com:80"或者"127.0.0.1:22"

// TCP client
// 共中通过net包中的DialTCP函数来建立一个TCP连接，并返回一个TCPConn类型的对象，当连接建立时服务器端也创建一个同类型的对象，此时客户端和服务器端通过各自拥有的TCPConn对象来进行数据交换。一般而言，客户端通过TCPConn对象将信息发送到服务器端，读取服务器端响应的信息。服务器端读取并解析来自客户端的请求，并返回应答信息，这个连接只有当任一端关闭了连接之后才失效，不然这连接可以一直在使用。建立连接的函数定义如下：

















































