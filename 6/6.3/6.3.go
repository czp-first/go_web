// 6.3 session存储
// memory.go代码实现了一个内存存储的session机制。通过init函数注册到session管理器中。这样就可以方便的调用了。如何启动该引擎呢？

import (
	"github.com/astaxie/session"
	_ "github.com/astaxie/session/providers/memory"
)

// 当import的时候已经执行了memory函数里面的init函数，这样就已经注册到session管理器中，就可以使用了，通过如下方式就可以初始化一个session管理器：
var globalSessions *session.Manager

// 然后在init函数中初始化
func init() {
	globalSessions, _ = session.NewManager("memory", "gosessionid", 3600)
	go globalSessions.GC()
}