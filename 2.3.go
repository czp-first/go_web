// 流程和函数

// 流程控制
// 流程控制包含分三大类：条件判断，循环控制和无条件跳转

// if
if x > 10 {
	fmt.Println("x is greater than 10")
} else {
	fmt.Println("x is less than 10")
}
// go的if里面允许声明一个变量，这个变量的作用域只能在该条件逻辑块内，其他地方就不起作用了
if x := computedValue(); x > 10 {
	fmt.Println("x is greater than 10")
} else {
	fmt.Println("x is less than 10")
}
fmt.Println(x)

// 多个条件的时候
if integer == 3 {
	fmt.Println("The integer is equal to 3")
} else if integer < 3 {
	fmt.Println("The integer is less than 3")
} else {
	fmt.Println("The integer is greater than 3")
}

// goto
// 用goto跳转到必须在当前函数内定义的标签。
func myFunc() {
	i := 0
Here:	// 这行的第一个词，以冒号结束作为标签
	println(i)
	i++
	goto Here	// 跳转到Here去
}
// 标签名是大小写敏感的

// for
// 它既可以用来循环读取数据，又可以当作while来控制逻辑，还能迭代操作
for expression1; expression2; expression3 {
	// ...
}
// expression1、expression2和expression3，其中expression1和expression3是变量声明或者函数调用返回值之类的，expression2是用来条件判断，expression1在循环开始之前调用，expression3在每轮循环结束之时调用
package main
import "fmt"
func main() {
	sum := -;
	for index := 0; index < 10; index++ {
		sum += index
	}
	fmt.Println("sum is equal to ", sum)
}
// 有些时候需要进行多个赋值操作，由于go里面没有,操作，那么可以使用平行赋值i,j = i+1, j-1
// 有些时候可以忽略expression1和expression3
sum := 1
for ; sum < 100; {
	sum += sum
}
// 其中；也可以省略
sum := 1
for sum < 1000 {
	sum += sum
}
// 循环操作，break是跳出当前循环，continue是跳出本次循环。当嵌套过深的时候，break可以配合标签使用，即跳转至标签所指定的位置
for index := 10; index>0; index-- {
	if index == 5{
		break // 或者continue
	}
	fmt.Println(index)
}
// break和continue还可以跟着标号，用来跳到多重循环中的外层循环
// for配合range可以用于读取slice和map的数据
for k,v := range map {
	fmt.Println("map's key:", k)
	fmt.Println("map's val:", v)
}
// go支持"多值返回"，而对于"声明而未被调用"的变量，编译器会报错，在这种情况下，可以使用_来丢弃不需要的返回值
for _, v := range map{
	fmt.Println("map's val:", v)
}

// switch
switch sExpr {
case expr1:
	some intructions
case expr2:
	some other intructions
case expr3:
	some other intructions
default:
	other code
}
// sExpr和expr1、expr2、expr3的类型必须一致。go的switch非常灵活，表达式不必是常量或整数，执行的过长从上至下，直到找到匹配项；如果switch没有表达式，它会匹配true
i := 10
switch i {
case 1:
	fmt.Println("1")
case 2, 3, 4:
	fmt.Println("i is equal 2, 3 or 4")
case 10:
	fmt.Println("10")
default:
	fmt.Println("All I know is that i is an integer")
}
// 在第5行中，可以把很多值聚合在一个case里面，同时，go里面switch默认相当于每个case最后带有break，匹配成功后不会自动向下执行其他case，而是跳出整个switch，但是可以使用fallthrough强制执行后面的case代码
integer := 6
switch integer {
case 4:
	fmt.Println("<= 4")
	fallthrough
case 5:
	fmt.Println("<= 5")
	fallthrough
case 6:
	fmt.Println("<= 6")
	fallthrough
case 7:
	fmt.Println("<= 7")
	fallthrough
case 8:
	fmt.Println("<= 8")
	fallthrough
default:
	fmt.Println("default case")
}




















