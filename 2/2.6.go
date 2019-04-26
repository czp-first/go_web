// 2.6 interface
// interface
// 它让面向对象，内容组织实现非常的方便

// 什么是interface
// 简单的说，interface是一组method的组合，通过interface来定义对象的一组行为

// interface类型
// interface类型定义了一组方法，如果某个对象实现了某个接口的所有方法，则此对象就实现了此接口。
type Human struct {
	name string
	age int
	phone string
}

type Student struct {
	Human
	school string
	loan float32
}

type Employee struct {
	Human
	company string
	money float32
}

// Human对象实现SayHi方法
func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

// Human对象实现Sing方法
func (h *Human) Sing(lyrics string) {
	fmt.Printf("La la, la la la la ....", lyrics)
}

// Human对象实现Guzzle方法
func (h *Human) Guzzle(beerStein string) {
	fmt.Println("Guzzle Guzzle Guzzle...", beerStein)
}

// Employee重载Human的SayHi方法
func (e *Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name, e.company, e.phone)
}

// Student实现BorrowMoney方法
func (s *Student) BorrowMoney(amount float32) {
	s.loan += amount
}

// Employee实现SpendSalary方法
func (e *Employee) SpendSalary(amount float32) {
	e.money -= amount
}

// 定义interface
type Men interface {
	SayHi()
	Sing(lyrics string)
	Guzzle(beerStein string)
}

type YoungChap interface {
	SayHi()
	Sing(song string)
	BorrowMoney(amount float32)
}

type ElderlyGent interface {
	SayHi()
	Sing(song string)
	SpendSalary(amount float32)
}
// interface可以被任意的对象实现。Men interface被Human、Student和Employee实现
// 一个对象可以实现任意多个interface，Student实现了Men和YoungChap两个interface
// 任意的类型都实现了空interface(这样定义：interface{}), 也就是包含0个method的interface

// interface值
// 如果定义了一个interface的变量，那么这个变量里面可以存实现这个interface的任意类型的对象。例如上面例子中，定义了一个Men interface类型的变量m，那么m里面可以存Human、Student或者Employee的值
// 因为m能够持有这三种类型的对象，所以可以定义一个包含Men类型元素的slice，这个slice可以被赋予实现了Men接口的任意结构的对象，这个和传统意义上面的slice有所不同
// 例
package main
import "fmt"

type Human struct {
	name string
	age int
	phone string
}

type Student struct {
	Human 
	school string
	loan float32
}

type Employee struct {
	Human
	company string
	money float32
}

// Human实现SayHi方法
func (h Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

// Human实现Sing方法
func (h Human) Sing(lyrics string) {
	fmt.Println("La la la la ...", lyrics)
}

// Employee重载Human的SayHI方法
func (e Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name, e.company, e.phone)
}

// Interface Men被Human，Student和Employee实现
// 因为这三个类型都实现了这两个方法
type Men interface {
	SayHi()
	Sing(lyrics string)
}

func main() {
	mike := Student{Human{"Mike", 25, "222-222-XXX"}, "MIT", 0.00}
	paul := Student{Human{"Paul", 26, "111-222-XXX"}, "Harvard", 100}
	sam := Employee{Human{"Sam", 36, "444-222-XXX"}, "Golang Inc", 1000}
	Tome := Employee{Human{"Tom", 36, "444-222-XXX"}, "Things Ltd", 5000}

	// 定义Men类型的变量i
	var i Men

	// i能存储Student
	i = mike
	fmt.Println("This is Mike, a Student:")
	i.SayHi()
	i.Sing("November rain")

	// i也能存储Employee
	i = Tom
	fmt.Println("This is Tom, an Employee:")
	i.SayHi()
	i.Sing("Born to be wild")

	// 定义了slice Men
	fmt.Println("Let's use a slice of Men and see what happens")
	x := make([]Men, 3)
	// T这三个都是不同类型的元素，但是它们实现了interface同一个接口
	x[0], x[1], x[2] = paul, sam, mike

	for _, vlaue := range x {
		value.SayHi()
	}
}
// 通过上面的代码，会发现interface就是一组抽象方法的集合，它必须由其他非interface类型实现，而不能自我实现，go通过interface实现了duck-typing: 即"当看到一只鸟走起来像鸭子、游泳起来像鸭子、叫起来像鸭子，那么这只鸟就可以被称为鸭子"

// 空interface
// 空interface(interface{})不包含任何的method，正因为如此，所有的类型都实现了空interface。空interface对于描述起不到任何的作用(因为它包含任何的method)，但是空interface在需要存储任意类型的数值的时候相当有用，因为它可以存储任意类型的数值
// 定义a为空接口
var a interface{}
var i int = 5
s := "Hello world"
// a可以存储任意类型的数值
a = i
a = s
// 一个函数把interface{}作为参数，那么它可以接受任意类型的值作为参数，如果一个函数返回interface{}，那么也就可以返回任意类型的值

// interface函数参数
// interface的变量可以持有任意实现该interface类型的对象，这个编写函数(包括method)提供了一些额外的思考，是不是可以铜鼓定义interface参数，让参数接受各种类型的参数
// 例子: fmt.Println 源码
type Stringer integerface {
	String() string
}
// 也就是说，任何实现了String方法的类型都能作为参数被fmt.Pringln调用
package main
import (
	"fmt"
	"strconv"
)

type Human struct {
	name string
	age int
	phone string
}

// 通过这个方法 Human 实现了 fmt.Stringer
func (h Human) String() string {
	return h.name + " - " + strconv.Itoa(h.age) + " years -" + h.phone
}

func main() {
	Bob := Human{"Bob", 39, "000-7777-XXX"}
	fmt.Println("This Human is :", Bob)
}

// 前面的Box事例，Color结构也定义了一个method：String。其实这也是实现了fmt.Stringer这个interface，即如果需要某个类型能被fmt包以特殊的格式输出，就必须实现Stringer这个接口。
// 如果没有实现这个接口，fmt将以默认的方式输出

// 实现同样的功能
fmt.Println("The biggest one is", boxes.BiggestsColor().String())
fmt.Println("The biggest one is", boxes.BiggestsColor())
// 注：实现了error接口的对象(即实现了Error() string的对象)，使用fmt输出时，会调用Error()方法，因此不必再定义String()方法了

// interface变量存储的类型
// interface的变量里面可以存储任意类型的数值(该类型实现了interface)。那么怎么反向知道这个变量里面实际保存了的是哪个类型的对象呢？
// 目前有两种方法

// Comma-ok断言
// go里面有一个语法，可以直接判断是否是该类型的变量：
// value, ok = element.(T)
// value是变量的值，ok是一个bool类型，element是interface变量，T是断言的类型
// 如果element里面确实存储了T类型的数值，那么ok返回true，否则返回false
package main
import (
	"fmt"
	"strconv"
)

type Element interface {}
type List [] Element

type Person struct {
	name string
	age int
}

// 定义了String方法，实现了fmt.Stringer
func (p Person) String() string {
	return "(name:" + p.name + " - age: " + strconv.Itoa(p.age) + "years)"
}

func main() {
	list := make(List, 3)
	list[0] = 1  // an int
	list[1] = "Hello"  // a string
	list[2] = Person{"Dennis", 70}

	for index, element := range list {
		if value, ok := element.(int); ok {
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		} else if value, ok := element.(string); ok {
			fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
		} else if value, ok := element.(Person); ok {
			fmt.Printf("list[%d] is a Person and its value is %s\n", index, value)
		} else {
			fmt.Println("list[%d] is of a different type", index)
		}
	}

}


// switch测试
package main

import (
	"fmt"
	"stconv"
)

type Element interface {}
type List [] Element

type Person struct {
	name string
	age int
}

// 打印
func (p Person) String() string{
	return "(name: " + p.name + " - age: " + strconv.Itoa(p.age) + " years)"
}

func main() {
	list := make(List, 3)
	list[0] = 1 
	list[1] = "Hello"
	list[2] = Person{"Dnnis", 70}

	for index, element := range list {
		switch value := element.(type){
		case int:
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		case string:
			fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
		case Person:
			fmt.Printf("list[%d] is a Person and its value is %s\n", index, value)
		default:
			fmt.Printlf("list[%d] is of a different type", index)
		}
	}
}
// 注意：element.(type)语法不能在switch外的任何逻辑里面使用，如果要在switch外面判断一个类型就使用comma-ok


// 嵌入interface
// go里面真正吸引人的是它内置的逻辑语法，就像在学习Struct时学习的匿名字段
// 如果一个interface1作为interface2的一个嵌入字段，那么interface2隐式的包含了interface1里面的method
// 源码包container/heap里面有这样一个定义
type Interface interface {
	sort.Interface  // 嵌入字段sort.Interface
	Push(x interface{})  // a Push method to push elements into the heap
	Pop() interface()  // a Pop elements that pops elements from the heap
}
// sort.Interface其实就是嵌入字段，把sort.Interface的所有method给隐式的包含进来了。也就是下面三个方法
type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}
// 另一个例子就是io包下面的io.ReadWriter，它包含了io包下面的Reader和Writer两个interface
// io.ReadWriter
type ReadWriter interface {
	Reader
	Writer
}

// 反射
// go语言实现了反射，所谓反射就是动态运行时的状态。一般用到的包是reflect包
// 使用reflect一般分成三步
// 简要讲解
// 要去反射是一个类型的值(这些值都实现了空interface)，首先需要把它转化成reflect对象(reflect.Type或者reflect.Value, 根据不同的情况调用不同的函数)。这两种获取方式如下：
t := reflect.TypeOf(i)  // 得到类型的元数据，通过t能获取类型定义里面的所有元素
v := reflect.Valueof(i)  // 得到实际的值，通过v获取存储在里面的值，还可以去改变值
// 转化为reflect对象之后就可以进行一些操作了，也就是将reflect对象转化成相应的值，例如
tag := t.Elem().Field(0).Tag  // 获取定义在struct里面的标签、
name := v.Elem().Field(0).String()  // 获取存储在第一个字段里买呢的值
// 获取反射值能返回相应的类型和数值
var x float64 = 3.4
v := reflect.ValueOf(x)  
fmt.Println("type:", v.Type())
fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
fmt.Println("value:", v.Float())
// 反射的话，那么反射的字段必须是可修改的，反射的字段必须是可读写的意思是，如果下面这样写，那么会发生错误
var x float64 := 3.4
v := reflect.ValueOf(x)
v.SetFloat(7.1)
// 如果要修改相应的值，必须这样写
var x float64 = 3.4
p := reflect.ValueOf(&x)
v := p.Elem()
v.SetFloat(7.1)


































