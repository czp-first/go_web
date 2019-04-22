package main

// 定义变量
// var variableName type
// var vname1, vname2, vname3 type
// var variableName type = value
// var vname1, vname2, vname3 type = v1, v2, v3
// var vname1, vname2, vname3 = v1, v2, v3
// vname1, vname2, vname3 := v1, v2, v3
// := 只能用在函数内部；在函数外部使用则会无法编译通过，所以一般用var方式来定义全局变量
// _是个特殊的变量名，任何赋予它的值都会被丢弃
// go对于已声明但未使用的变量会在编译阶段报错。

// 常量
// const constName = value
// const Pi float32 = 3.1415926
// const Pi = 3.1415926
// const i = 1000
// const MaxThread = 10
// const prefix = "astaxie_"

// 内置基础类型
// Boolean
// go中，布尔值的类型为bool，值为true或false，默认为false
var isActive bool
var enabled, disabled = true, false
func test() {
	var available bool
	valid := false
	available = true
}

// 数值类型
// 整数
// 类型有无符号和带符号两种。go同时支持int和unit，这两种类型的长度相同，但具体长度取决于不同编译器的实现。
// go也有直接定义好位数的类型：rune,int8,int16,int32,int64和byte,unit8,unit16,unit32,unit64.
// 其中rune是int32的别称，byte是unit8的别称
// 这些类型的变量之间不允许互相赋值或操作，不然会在编译时引起编译器报错
// 尽管int的长度是32bit，但int与int32并不可以互用
// 浮点数
// 类型有float32和float64两种(没有float类型)，默认为float64
// 复数
// 默认类型为complex128(64位实数+64位虚数)。还有complex64(32位实数+32位虚数)
// 复数形式位RE+IMi，RE是实数部分，IM是虚数部分，最后的i是虚数单位
var c complex64 = 5 + 5i
// output: (5 + 5i)
fmt.Printf("Value is: %v", c)

// 字符串
// 字符串是用一对双引号("")或反引号(``)括起来定义的，它的类型是string
var frenchHello string
var emptyString string = ""
func test() {
	no, yes, maybe := "no", "yes", "maybe"
	japaneseHello := "ohaiou"
	frenchHello = "Bonjour"
}
// go中字符串是不可变的
// 但如果真的想修改，下面代码可以实现
s := "hello"
c := []byte(s)
c[0] = "c"
s2 = string(c)
fmt.Printf("%s\n", s2)
// go中可以使用+来连接两个字符串
s := "hello"
m := "world"
a := s + m
fmt.Printf("%s\n", a)
// 修改字符串也可以写为
s := "hello"
s := "c" + s[1:]
fmt.Printf("%s\n", s)
// 声明一个多行的字符串，通过``来声明
m := `hello 
	world`
// ``括起来的字符串为Raw字符串，即字符串在代码中的形式就是打印时的形式，它没有字符转义，换行也将原样输出

// 错误类型
// go内置有一个error类型，专门用来处理错误信息，go的package里还有一个包errors来处理粗错误
err := errors.New("emit macho dwarf: elf header corrupted")
if err != nil {
	fmt.Print(err)
}

// go数据底层的存储
i := 1234  //type:int
j := int32(1)  //type:int32
f := float32(3.14)  //type:float32
bytes := [5]byte{'h', 'e', 'l', 'l', 'o'}  //type:[5]byte
primes := [4]int{2, 3, 5, 7}  //type:[4]int

// 分组声明
// go语言中，同时声明多个常量、变量，或者导入多个包时，可采用分组的方式进行声明
// 例如下面代码
import "fmt"
import "os"

const i = 100
const pi = 3.1415
const prefix = "Go_"

var i int
var pi float32
var prefix string
// 可以分组写成如下形式
import (
	"fmt"
	"os"
)

const (
	i = 100
	pi = 3.1415
	prefix = "Go_"
)
var (
	i int
	pi float32
	prefix string
)
// 除非被显示设置为其他值或iota，每个const分组的第一个常量被默认设置为它的0值，第二及后续的常量被默认设置为它前面那个常量的值，如果前面那个常量的值是iota，则它也被设置为iota

//iota枚举
//关键字iota用来声明enum的时候采用，它默认开始值是0，每调用一次加1
const (
	x = iota  //x == 0
	y = iota  //y == 1
	z = iota  //z == 2
	w  // 常量声明省略值时，默认和之前一个值的字面相同。这里隐式地说w = iota, 因此w == 3
)
const v = iota  // 每遇到一个const关键字，iota就会重置，此时v == 0