// 7.6 字符串处理
// 字符串在平常的web开发中经常用到，包括用户的输入，数据库读取的数据等，经常需要对字符串进行分隔、连接、转换等操作，本小节通过go标准库中的strings和strconv两个包中的函数来讲解如何进行有效快速的操作

// 字符串操作
// 下面这些函数来自于strings包，这里介绍一些作者平常用到的函数，更详细的请参考官方的文档
//	func Contains(s, substr string) bool
//	字符串s中是否包含substr，返回bool值
fmt.Println(strings.Contains("seafood", "foo"))
fmt.Println(strings.Contains("seafood", "bar"))
fmt.Println(strings.Contains("seafood", ""))
ftm.Println(strings.Contains("", ""))
// Output:
// true
// false
// true
// true

//	func Join(a []string, sep string) string
//	字符串连接，把slice a通过sep连接起来
s := []string{"foo", "bar", "baz"}
fmt.Println(strings.Join(s, ", "))
// Output: foo, bar, baz

//	func Index(s, sep string) int
//	在字符串s中查找sep所在的位置，返回位置值，找不到返回-1
fmt.Println(strings.Index("chicken", "ken"))
fmt.Println(strings.Index("chicken", "dmr"))
// Output:4
// -1

//	func Repeat(s string, count int) string
//	重复s字符串count次，最后返回重复的字符串
fmt.Pritnln("ba" + strings.Repeat("na", 2))
// Output: banana

//	func Replace(s, old, new string, n int) string
//	在s字符串中，把old字符串替换为new字符串，n表示替换的次数，小于0表示全部替换
fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))
// Output:oinky oinky oinky
// moo moo moo

// func Split(s, sep string) []string
//	把s字符串按照sep分隔，返回slice
fmt.Printf("%q\n", strings.Split("a,b,c", ","))
fmt.Printf("%q\n", strings.Split("a man a plan a panama", "a "))
fmt.Printf("%q\n", strings.Split(" xyz", ""))
fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))
// Output: ["a" "b" "c"]
// ["" "man " "plan " "canal panama"]
// [" " "x" "y" "z" " "]
// [""]

//	func Trim(s string, cutset string) string
//	在s字符串中去除cutset指定的字符串
fmt.Printf("[%q]", strings.Trim(" !!! Achtung !!! ", "! "))
// Output: ["Achtung"]

//	func Fields(s string) []string
//	去除s字符串的空格符，并且按照空格分割返回slice
fmt.Printf("Fields are: %q", strings.Fields("  foo bar  baz   "))
// Output: Fields are: ["foo" "bar" "baz"]


// 字符串转换
// 字符串转化的函数在strconv中，如下也只是列表一些常用的
//	Append 系列函数将整数等转换为字符串后，添加到现有的字节数组中
//	示例：1.go
//	Format 系列函数把其他类型的转换为字符串
//	示例：2.go
//	Parse 系列函数把字符串转换为其他类型
//	示例：3.go














































