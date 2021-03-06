// 11.3 go怎么写测试用例
// 开发程序其中很重要的一点是测试，如何保证代码的质量，如何保证每个函数是可运行，运行结果是正确的，又如何保证写出来的代码性能是好的，单元测试的重点在于发现程序设计或实现的逻辑错误，使问题及早暴露，便于问题的定位解决，而性能测试的重点在于发现程序设计上的一些问题，让线上的程序能够在高并发的情况下还能保持稳定。
// go中自带有一个轻量级的测试框架testing和自带的go test命令来实现单元测试和性能测试，testing框架和其他语言中的测试框架类似，可以基于这个框架写针对相应函数的测试用例，也可以基于该框架写相应的压力测试用例

// 如何编写测试用例
// 由于go test命令只能在一个相应的目录下执行所有文件，所以接下来新建一个项目目录gotest，这样所有的代码和测试代码都在这个目录下
// 接下来在该目录下面创建两个文件: gotest.go和gotest_test.go
//	1. gotest.go:这个文件里面是创建了一个包，里面有一个函数实现了除法运算
//	2. gotest_test.go: 这是单元测试文件，但是记住下面的这些原则
//		文件名必须是`_test.go`结尾的，这样在执行`go test`的时候才会执行到相应的代码
//		必须import `testing`这个包
//		所有的测试用例函数必须是`Test`开头
//		测试用例会按照源代码中写的顺序依次执行
//		测试函数`TestXxxx()`的参数是`testing.T`，可以使用该类型来记录错误或者测试状态
//		测试格式：`func TestXxx (t *testing.T)`, `Xxx`部分可以为任意的字母数字的组合，但是首字母不能是小写字母[a-z]
//		函数中通过调用`testing.T`的`Error`, `Errorf`, `FailNow`, `Fatal`, `FatalIf`方法，说明测试不通过，调用`L。。。
// 测试用例代码: gotest_test.go
// 在项目目录下面执行`go test` 或`go test -v`

// 如何编写压力测试
// 压力测试用来监测函数(方法)的性能，和编写单元功能测试的方法类似，但需要注意以下几点
//	压力测试用例必须遵循如下格式，其中XXX可以是任意字母数字的组合，但是首字母不能是小写字母
//	func BenchmarkXXX(b *testing.B) { ... }
//	go test 不会默认执行压力测试的函数，如果要执行压力测试需要带上参数-test.bench,语法：-test.bench="test_name_regex", 例如go test -test.bench=".*"表示测试全部的压力测试函数
//	在压力测试用例中，请记得在循环体内使用testing.B.N, 以使测试可以正常的运行
//	文件名也必须以_test.go结尾






























