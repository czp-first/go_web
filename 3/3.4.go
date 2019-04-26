// 3.4 go的http包详解
// go的http有两个核心功能：Conn、ServeMux

// Conn的goroutine
// go为了实现高并发和高性能，使用了goroutines来处理conn的读写事件，这样每个请求都能保持独立，相互不会阻塞，可以高效的响应网络事件。
// go在等待客户端请求里面是这样写的
c, err := srv.newConn(rw)
if err != nil {
	continue
}
go c.serve()
// 客户端的每次请求都会创建一个Conn，这个Conn里面保存了该次请求的信息，然后再传递到对应的handler，该handler中便可以读取到响应的header信息。

// ServeMux的自定义
