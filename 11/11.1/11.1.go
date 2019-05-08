// 11.1 错误处理
// go主要的设计准则是：简洁、明白，简洁是指语法和C类似，相当的简单，明白是指任何语句都是很明显的，不含有任何隐含的东西，在错误处理方案的设计中也贯彻了这一思想。我们知道在C语言里面是通过返回-1或者NULL之类的信息来表示错误，但是对于使用者来说，不查看相应的API说明文档，根本搞不清楚这个返回值究竟代表什么意思，比如：返回0是成功，还是失败，而go定义了一个叫做error的类型，来显式表达错误。在使用时，通过把返回的error变量与nil的比较，来判定操作是否成功。例如os.Open函数在打开文件失败时返回一个不为nil的error变量
func Open(name string) (file *File, err error)
// 下面这个例子通过调用os.Open打开一个文件，如果出现错误，那么就会调用log.Fatal来输出错误信息:
f, err := os.Open("filename.ext")
if err != nil {
	log.Fatal(err)
}
// 类似于os.Open函数，标准包中所有可能出错的API都会返回一个error变量，以方便错误处理，这个小节将详细地介绍error类型的设计，和讨论开发Web应用中如何更好地处理error

// Error类型
// error类型是一个接口，这是它的定义:
type error interface {
	Error() string
}
// error是一个内置的接口类型，可以在/builtin/包下面找到相应的定义。而我们在很多内部包里面用到的error是errors包下面的实现的私有结构errorString
// errorString is a trivial implementation of error
type errorString struct {
	a string
}
func (e *errorString) Error() string {
	return e.s
}
// 可以功过errors.New把一个字符串转化为errorString, 以得到一个满足接口error的对象，其内部实现如下:
// New returns an error that formats as the given text
func New(text string) error {
	return &errorString{text}
}
// 下面这个例子演示了如何使用errors.New:
func Sqrt(f folat64) (float64, err) {
	if f < 0 {
		return 0, errors.New("math: square root of negative number")
	}
	// implementation
}
// 在下面的例子中，在调用Sqrt的时候传递的一个负数，然后就得到了non-nil的error对象，将此对象与nil比较，结果为true，所以fmt.Println(fmt包在处理error时会调用Error方法)被调用，以输出错误，请看下面调用的示例代码
f, err := Sqrt(-1)
if err != nil {
	fmt.Println(err)
}

// 自定义Error
// 通过上面的介绍知道error是一个interface，所以在实现自己的包的时候，通过定义实现此接口的结构，就可以实现自己的错误定义，请看来自Json包的示例：
type SyntaxError struct {
	msg string // 错误描述
	Offset int64 // 错误发生的位置
}
func (e *SyntaxError) Error() string {
	return e.msg
}
// Offset字段在调用Error的时候不会被打印，但是可以通过类型断言获取错误类型，然后可以打印相应的错误信息，请看下面的例子
if err := dec.Decode(&val); err != nil {
	if serr, ok := err.(*json.SyntaxError); ok {
		line, col := findLine(f, serr.Offset)
		return fmt.Errorf("%s:%d:%d: %v", f.Name(), line, col, err)
	}
	return err
}
// 需要注意的是，函数返回自定义错误时，返回值也应设置为error类型，而非自定义错误类型，也不应预声明自定义错误类型的变量。例如：
func Decode() *SyntaxError {  // 错误，将可能导致上层调用者err!=nil的判断永远为true
	var err *SyntaxError // 预声明错误变量
	if 出错条件 {
		err = &SyntaxError{}
	}
	return err // 错误，虽然err变量等于nil，但仍可能导致上层调用者err!=nil的判断为true
}
// 原因见 http://golang.org/doc/fag#nil_error
// 上面的例子简单的演示了如何自定义Error类型。但是如果还需要更复杂的错误处理呢？此时，参考一个net包采用的方法
package net
type Error interface {
	error
	Timeout() bool // is the error a timeout?
	Temporary() bool // is the error temporary?
}
// 在调用的地方，通过类型断言err是不是net.Error,来细化错误的处理，例如下面的例子，如果一个网络发生临时性错误，那么将会sleep1秒之后重试：
if nerr, ok := err.(net.Error); ok && nerr.Temporary() {
	time.Sleep(1e9)
	continue
}
if err != nil {
	log.Fatal(err)
}

// 错误处理
// go在错误处理上采用了与C类似的检查返回值的方式，而不是其他多数主流语言采用的一场方式，这造成了代码编写上的一个很大的缺点:错误处理代码的冗余，对于这种情况时通过复用监测函数来减少类似的代码
// 请看下面这个例子代码：
func init() {
	http.HandleFunc("/view", viewRecord)
}
func viewRecord(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	key := datastore.NewKey(c, "Record", r.FormValue("id"), 0, nil)
	record := new(Record)
	if err := datastore.Get(c, key, record); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if err != viewTemplate.Execute(w, record); err != nil {
		http.Error(w, err.Error(), 500)
	}
}
// 上面的例子中获取数据和模版展示调用时都有检测错误，当错误发生时，调用了统一的处理函数http.Error,返回给客户端500错误码，并显示相应的错误数据。但是当越来越多的HandleFunc加入之后，这样的错误处理逻辑代码就会越来越多，其实可以通过自定义路由器来缩减代码(实现的思路可以参考第三章的HTTP详解)
type appHandler func(http.ResponseWriter, *http.Request) error 
func (fn appHandler) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	if err := fn(w, r) != nil {
		http.Error(w, err.Error(), 500)
	}
// 上面定义了自定义的路由器，然后可以通过如下方式来注册函数:
func init() {
	http.Handle("/view", appHandler(viewRecord))
}
// 当请求/view的时候，逻辑处理可以变成如下代码，和第一种实现方式相比较已经简单了很多
func viewRecord(w http.ResponseWriter, r *http.Request) error {
	c := appengine.NewContext(r)
	key := datastore.NewKey(c, "Record", r.FormValue("id"), 0, nil)
	record := new(Record)
	if err := datastore.Get(c, key, record); err != nil {
		return err
	}
	return viewTemplate.Execute(w, record)
}
// 上面的例子错误处理的时候所有的错误返回给用户的都是500错误代码，然后打印出来相应的错误发麻，其实可以把这个错误信息定义的更加友好，调试的时候也方便定位问题，可以自定义返回的错误类型
type appError struct {
	Error error
	Message string
	Code int
}
// 这样自定义路由器可以改成如下方式
type appHandler func(http.ResponseWriter, *http.Request) *appError 
func (fn appHandler) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	if e := fn(w, r); e != nil { // e is *appError, not os.Error
		c := appengine.NewContext(r)
		c.Errorf("%v", e.Error)
		http.Error(w, e.Message, e.Code)
	}
}
// 这样修改完自定义错误之后，逻辑处理可以改成如下方式
func viewRecord(w http.ResponseWriter. r * http.Request) *appError {
	c := appengine.NewContext(r)
	key := datastore.NewKey(c, "Record", r.FormValue("id"), 0, nil)
	record := new(Record)
	if err := datastore.Get(c, key, record); err != nil {
		return &appError{err, "Record not fount", 404}
	}
	if err := viewTemplate.Execute(w, record); err != nil {
		return &appError{err, "Can't display record", 500}
	}
	return nil
}
// 如上所示，在访问view的时候可以根据不同的情况获取不同的错误码和错误信息，虽然这个和第一个版本的代码量差不多，但是这个现实的错误更加明显，提示的错误信息更加友好扩展性也比第一个更好

// 总结
// 在程序设计中，容错是相当重要的一部分工作，在go中它是通过错误处理来实现的，error虽然只是一个接口，但是其变化却可以有很多，可以根据自己的需求来实现不同的处理，最后介绍的错误处理方案，希望能给大家在如何设计更好Web错误处理方案上带来一点思路









