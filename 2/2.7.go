// 2.7并发

// goroutine
// goroutine说到底就是线程，但是它比线程更小，十几个goroutine可能体现在底层就是五六个线程，go内部帮你实现了这些goroutine之间的内存共享。
// 执行goroutine只需极少的栈内存(大约4-5KB)，当然会根据相应的数据伸缩。
// goroutine是通过go的runtime管理的一个线程管理器。goroutine通过go关键字实现，其实就是一个普通的函数

go hello(a, b, c)

// 通过关键字go就启动了一个goroutine
package main

import (
	"fmt"
	"runtime"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

func main() {
	go say("world")  // 开一个新的goroutine执行
	say("hello")  // 当前goroutine执行
}
// 上面的多个goroutine运行在同一个进程里面，共享内存数据，不过设计上要遵循：不要通过共享来通信，而要通过通信来共享
// runtime.Gosched()表示让CPU把时间片让给别人，下次某个时候继续恢复执行该goroutine
// 默认情况下，调度器仅使用单线程，也就是说只实现了并发。想要发挥多核处理器的并行，需要在程序中显示的调用runtime.GOMAXPROCS(n)告诉调度器同时使用多个线程。GOMAXPROCS设置了同时运行逻辑代码的系统线程的最大数量，并返回之前的设置。如果n<1，不会改变当前设置。以后go的新版本中调度得到改进后，这将被移除。

// channels
// goroutine运行在相同的地址空间，因此访问共享内存必须做好同步。那么goroutine之间如何进行数据的通信呢，go提供了一个很好的通信机制channel。channel可以与Unix shell中的双向管道做类比：可以通过它发送或者接受值。这些值只能是特定的类型:channel类型。定义一个channel时，也需要定义发送到channel的值的类型。
// 注意，必须使用make创建channel
ci := make(chan int)
cs := make(chan string)
cf := make(chan interface{})
// channel通过操作符<-来接收和发送数据
ch <- v  // 发送v到channel ch
v := <-chan  // 从ch中接收数据，并赋值给v
// 例
package main
import "fmt"

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum  // send sum to c
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c  // receive from c
	fmt.Println(x, y, x+y)
}
// 默认情况下，channel接收和发送数据都是阻塞的，除非另一端已经准备好，这样就使得goroutine同步变的更加的简单，而不需要显示的lock。
// 所谓阻塞，也就是如果读取(value := <-ch) 它将被阻塞，知道有数据接收。其次，任何发送 (ch<-5) 将会被阻塞，知道数据被读出。无缓冲channel是在多个goroutine之间同步很棒的工具

// 上面介绍了默认的非缓存类型的channel，不过go也允许制定channel的缓冲大小，很简单，就是channel可以存储多少元素。ch:=make(chan bool, 4), 创建了可以存储4个元素的bool型channel。在这个channel中，前4个元素可以无阻塞的写入。当写入第5个元素时，代码将会阻塞明知道其他goroutine从channel中读取一些元素，腾出空间
// ch := make(chan type, value)
// value == 0  无缓冲(阻塞)
// value > 0  缓冲(非阻塞，直到value个元素)
package main
import "fmt"
func main() {
	c := make(chan int, 2)
	c <- 1
	c <- 2
	fmt.Println(<-c)
	fmt.Pritnln(<-c)
}

// Range和Close
// 可以通过range操作缓存类型的channel
package main
import "fmt"

func fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = x, x+y
	}
	close(c)
}

func main() {
	c := make(chan int 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
// for i := range c 能够不断的读取channel里面的数据，知道该channel被显示的关闭。生产者通过关键字close函数关闭channel。关闭channel之后就无法再发送任何数据了，在消费方可以通过语法v, ok := <-ch测试channel是否被关闭。如果ok返回false，那么说明channel已经没有任何数据并且已经被关闭
// 记住应该在生产者的地方关闭channel，而不是消费的地方去关闭它，这样容易引起panic
// channel不想文件之类的，不需要经常去关闭，只有当确实没有任何发送数据了，或者想显示的结束range循环之类的

// Select
// 如果存在多个channel的时候，go里面提供了一个关键字select，通过select可以监听channel上的数据流动
// select默认是阻塞的，只有当监听的channel中有发送或接收可以进行时才会运行，当多个channel都准备好的时候，select是随机的选择一个执行的
package main

import "fmt"

func fibonacci(c int, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
		fmt.Println("quit")
			return 
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}

// 在select里面还有default语法，select其实就是类似switch的功能，default就是当监听的channel都没有准备好的时候，默认执行的(select不再阻塞等待channel)
select {
case i := <-c:
	// use i
default:
	// 当c阻塞的时候执行这里
}

// 超时
// 有时候会出现goroutine阻塞的情况，可以利用select来设置超时，来避免整个的程序进入阻塞的情况
func main() {
	c := make(chan int)
	o := make(chan int)
	go func() {
		for {
			select {
			case v := <-c:
				println(v)
			case <- time.After(5 * time.Second):
				println("timeout")
				o <- true
				break
			}
		}
	}()
	<-o
}

// runtime goroutine
// runtime包中有几个处理goroutine的函数
// Goexit：退出当前执行的goroutine，但是defer函数还会继续调用
// Gosched：让出当前goroutine的执行权限，调度器安排其他等待的任务运行，并在下次某个时候从该位置恢复执行
// NumCPU：返回CPU核数量
// NumGoroutine：返回正在执行和排队的任务总数
// GOMAXPROCS：用来设置可以运行的CPU核数




























