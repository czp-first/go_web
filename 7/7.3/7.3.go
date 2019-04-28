// 7.3 正则处理
// go通过regexp标准包为正则表达式提供了官方支持，go实现的是RE2标准

// 通过正则判断是否匹配
// regexp包中含有三个函数用来判断是否匹配，如果匹配返回true，否则返回false
//	func Match(pattern string, b []byte) (matched bool, error error)
//	func MatchReader(pattern string, r io.RuneReader) (mached bool, error error)
//	func MatchString(pattern string, s string) (matched bool, error error)
// 上面三个函数实现了同一个功能，就是判断pattern是否和输入源匹配，匹配的话就返回true，如果解析正则出错则返回error。三个函数的输入源分别是byte slice、RuneReader和string
// 验证一个输入是不是IP地址
func IsIP(ip string) (b bool) {
	if m, _ := regexp.MatchString("^[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3}$", ip); !m {
		return false
	}
	return true
}
// regexp的pattern和平常使用的正则一摸一样。
// 再看一个例子，当用户输入一个字符串，判断是不是一次合法的输入
func main() {
	if len(os.Args) == 1 {
		fmt.Println("Usage: regexp [string]")
		os.Exit(1)
	} else if m, _ := regexp.MatchString("^[0-9]+$", os.Args[1]); m {
		fmt.Println("数字")
	} else {
		fmt.Println("不是数字")
	}
}

// 通过正则获取内容
// Match模式只能用来对字符串的判断，而无法街区字符串的一部分、过滤字符串、或者提取出符合条件的一批字符串。如果想要满足这些需求，那就需要使用正则表达式的复杂模式
// 经常需要一些爬虫程序，下面就以爬虫为例来说明如何使用正则来过滤或截取抓取到的数据
// regex1.go
// 从这个示例中可以看出，使用复杂的正则首先是Compile，它会解析正则表达式是否合法，如果正确，那么就会返回一个Regexp，然后就可以利用返回的Regexp在任意的字符串上面执行需要的操作
// 解析正则表达式的有如下几个方法：
//	func Compile(expr string) (*Regexp, error)
//	func CompilePOSIX(expr string) (*Regexp, error)
//	func MustCompile(str string) *Regexp
//	func MustCompilePOSIX(str string) *Regexp
// CompilePOSIX和Compile的不同点在于POSIX必须使用POSIX语法，它使用最左最长方式搜索，而Compile是采用的则只采用最左方式搜索(例如[a-z]{2,4}这样一个正则表达式，应用于"aa09aaa88aaaa"这个文本串时，CompilePOSIX返回了aaaa，而Compile的返回的是aa)。前缀有Must的函数表示，在解析正则语法的时候，如果匹配模式串不满足正确的语法则直接panic，而不加Must的则只是返回错误。
// 在了解了如何新建一个Regexp之后，再来看一下这个struct提供了哪些方法来辅助操作字符串，首先来看下面这些用来搜索的函数
//	func (re *Regexp) Find(b []byte) []byte
//	func (re *Regexp) FindAll(b []byte, n int) [][]byte
//	func (re *Regexp) FindAllIndex(b []byte, n int) [][]int 
//	func (re *Regexp) FindAllString(s string, n int) []string
//	func (re *Regexp) FindAllStringIndex(s string, n int) [][]int
//	func (re *Regexp) FindAllStringSubmatch(s string, n int) [][]string
//	func (re *Regexp) FIndAllStringSubmatchIndex(s string, n int) [][]int
//	func (re *Regexp) FindAllSubmatch(b []byte, n int) [][]byte
//	func (re *Regexp) FindAllSubmatchIndex(b []byte, n int) [][]int
//	func (re *Regexp) FIndIndex(b []byte) (loc []int)
//	func (re *Regexp) FindReaderIndex(r io.RunReader) (loc []int)
//	func (re *Regexp) FindReaderSubmatchIndex(r io.RunReader) []int
//	func (re *Regexp) FindString(s string) string
//	func (re *Regexp) FindStringIndex(s string) (loc []int)
//	func (re *Regexp) FindStringSubmatch(s string) []string
//	func (re *Regexp) FindStringSubmatchIndex(s string) []int
//	func (re *Regexp) FindSubmatch(b []byte) [][]byte
//	func (re *Regexp) FindSubmatchIndex(b []byte) []int
// 上面这18个函数根据输入源(byte slice,string和io.RuneReader)不同还可以继续简化成如下几个，其他的只是输入源不一样，其他功能基本是一样的
//	func (re *Regexp) Find(b []byte) []byte
//	func (re *Regexp) FindAll(b []byte, n int) []byte
//	func (re *Regexp) FindAllIndex(b []byte, n int) [][]int
//	func (re *Regexp) FindAllSubmatch(b []byte, n int) [][][]byte
//	func (re *Regexp) FindAllSubmatchIndex(b []byte, n int) [][]int
//	func (re *Regexp) FindIndex(b []byte) (loc []int)
//	func (re *Regexp) FindSubmatch(b []byte) [][]byte
//	func (re *Regexp) FindSubmatchIndex(b []byte) []int

// 对于这些函数的使用看regex2.go

// 前面介绍过匹配函数，Regexp也定义了三个函数，它们和同名的外部函数功能一摸一样，其实外部函数就是调用了这Regexp的三个函数来实现的
//	func (re *Regexp) Match(b []byte) bool
//	func (re *Regexp) MatchReader(r io.RuneReader) bool
//	func (re *Regexp) MatchString(s string) bool

// 接下来了解替换函数是怎么操作的？
//	func (re *Regexp) ReplaceAll(src, repl []byte) []byte
//	func (re *Regexp) ReplaceAllFunc(src []byte, repl func([]byte) []byte) []byte
//	func (re *Regexp) ReplaceAllLiteral(src, repl string) string
//	func (re *Regexp) ReplaceAllString(src, repl string) string
//	func (re *Regexp) ReplaceAllStringFunc(src string, repl func(string) string) string
// 这些替换函数，在上面的抓网页的例子有详细应用示例
// 接下来看一下Expand的解释
//	func (re *Regexp) Expand(dst []byte, template []byte, src []byte, match []int) []byte
//	func (re *Regexp) ExpandString(dst []byte, template string, src string, match []int) []byte
// 那么Expand到底用来干嘛？看例子



































