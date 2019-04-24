// 2.5 面向对象
// 函数的另一种形态，带有接收者的函数，称为method

// method
// 它的语法和函数的声明语法几乎一样，只是在func后面增加了一个reciever(也就是method所依从的主体)
// func (r RecieverType) funcName(parameters) (results)
// 例子
package main
import (
	"fmt"
	"math"
)

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) area() {
	return r.width * r.height
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func main() {
	r1 := Rectangle{12, 2}
	r2 := Rectangle{9, 4}
	c1 := Circle{10}
	c2 := Circle{25}

	fmt.Println("Area of r1 is: ", r1.area())
	fmt.Println("Area of r2 is: ", r2.area())
	fmt.Println("Area of c1 is: ", c1.area())
	fmt.Println("Area of c2 is: ", c2.area())
}
// 使用method注意
// 虽然method的名字一摸一样，但是如果接收者不一样，那么method就不一样
// method里面可以访问接收者的字段
// 调用method通过.访问，就像struct里面访问字段一样

// 此处方法的Receiver是以值传递，而非引用传递，Receicer还可以是指针，两者的差别在于，指针作为Receiver会对实例对象的内容发生操作，而普通类型作为Reciever仅仅是以副本作为操作对象，并不对实例对象发生操作

// 自定义类型声明
// type typeName typeLiteral
type ages int
type money float32
type months map[string]int

m := months {
	"January": 31,
	"February": 28,
	...
	"December": 31,
}

package main
import (
	"fmt"
)

const (
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

type Color byte
type Box struct {
	width, height, depth float64
	color Color
}
type BoxList []Box  // a slice of boxes

func (b Box) Volumn() float64 {
	return b.width * b.height * b.depth
}
func (b *Box) SetColor(c Color) {
	b.color = c
}
func (b1 BoxList) BiggestsColor() Color {
	v := 0.00
	k := Color(WHITE)
	for _, b := range b1 {
		if b.Volume() > v {
			v = b.Volume()
			k = b.color
		}
		return k
	}
}
func (b1 BoxList) PaintItBlack() {
	for i, _ := range b1 {
		b1[i].SetColor(BLACK)
	}
}
func (c color) String() string {
	strings := []string{"WHITE", "BLACK", "BLUE", "YELLOW"}
	return strings[c]
}

func main() {
	boxes := BoxList{
		Box{4, 4, 4, RED},
		BOX{10, 10, 1, YELLOW},
		Box{1, 1, 20, BLACK},
		Box{10, 10, 1, BLUE},
		Box{10, 30, 1, WHITE},
		Box{20, 20, 20, YELLOW}
	}

	fmt.Printf("We have %d boxes in our set\n", len(boxes))
	fmt.Println("The volume of the first one is", boxes[0].Volume(), "cm3")
	fmt.Println("The coolor of the last one is", boxes[len(boxes)-1].color().String())
	fmt.Println("The biggest one is", boxes.BiggestsColor().String())

	fmt.Println("Let's paint them all black")
	boxes.PaintItBlack()
	fmt.Println("The color of the second one is", boxes[1].color.String())

	fmt.Println("Obviously, now, the biggest one is", boxes.BiggestsColor().String())
}
// 自定义类型
// Color作为byte的别名
// 定义了一个struct:Box，含有桑日长宽高字段和一个颜色属性
// 定义了一个slice:BoxList，含有Box

// 然后以上面的自定义类型为接收者定义了一些method
// Volume()定义了接收者为Box，返回Box的容量
// SetColor(c color)，把Box的颜色改为c
// BiggestsColor()定在BixList上面，返回list里面容量最大的颜色
// PaintItBlack()把BixList里面所有Box的颜色全部变为黑色
// String()定义在Color上面，返回Color的具体颜色(字符串格式)


// 指针作为receiver
// 如果一个method的receiver是*T，可以在一个T类型的实例变量V上面调用这个method，而不需要&V去调用这个method
// 类似的，如果一个method的receiver是T，可以在一个*T类型的变量P上面调用这个method，而不需要*P调用这个method
// 所以，不用担心是调用的指针的method还是不是指针的method

// method继承
// 字段是可以继承的，method也是可以继承的。如果匿名字段实现了一个method，那么包含这个匿名字段的struct也能调用该method
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
}

type Employee struct {
	Human
	company string
}

// 在human上面定义一个method
func (h *human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

func main() {
	mark := Student{Human{"Mark", 25, "222-222-YYYY", "MIT"}}
	sam := Employee{Human{"Sam", 45, "111-888-XXXX", "Golang Inc"}}
	mark.SayHi()
	sam.SayHi()
}

// method重写
// 如果Emplyee想要实现自己的SayHi，怎么办？
// 和匿名字段冲突一样的道理，可以在Employee上面定义一个method，重写了匿名字段的方法
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
}

type Employee struct {
	Human
	company string
}

// Human定义method
func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}
// Employee的method重写Human的method
func (e *Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name, e.company, e.phone)
}

func main() {
	mark := Student{Human{"Mark", 25, "222-222-YYYY", "MIT"}}
	sam := Employee{Human{"Sam", 45, "111-888-XXXX", "Golang Inc"}}

	mark.SayHi()
	sam.SayHi()
}

// go里面的面向对象是如此的简单，没有任何的私有、公有关键字，通过大小写来实现(大写开头的为公有，小写开头的为私有)，方法也同样适用这个原则


















