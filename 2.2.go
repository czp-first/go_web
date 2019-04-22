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


