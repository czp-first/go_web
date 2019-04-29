// 7.4 模版处理

// 什么是模版
// MVC设计模式，Model处理数据，View展现结果，Controller控制用户请求，至于View层的处理，在很多动态语言里面都是通过在静态HTML中插入动态语言生成的数据，
// web应用反馈给客户端的信息中的大部分内容是静态的，不变的，而另外少部分是根据用户的请求来动态生成的，例如要显示用户的访问记录列表，用户之间只有记录数据是不同的，而列表的样式则是固定的，此时采用模版可以服用很多静态代码

// go模版使用
// 在go中，使用template包来进行模版处理，使用类似Parse、ParseFile、Execute等方法从文件或者字符串加载模版，然后执行模版的merge操作，
// 例
func handler(w http.ResponseWriter, r *http.Request) {
	t := template.New("some template")  // 创建一个模版
	t, _ = t.ParseFiles("tmpl/welcome.html", nil)  // 解析模版文件
	user := GetUser()  // 获取当前用户信息
	t.Execute(w, user)  // 执行模版的merge操作
}
// 通过上面的例子，可以看到go语言的模版操作非常方便，和其他语言的模版处理类似，都是先获取数据，然后渲染数据
// 为了掩饰和测试代码的方便，在接下来的例子中采用如下格式的代码
//	使用Parse代替ParseFiles，因为Parse可以直接测试一个字符串，而不需要额外的文件
//	不实用handler来写演示代码，而是每个测试一个main，方便测试
//	使用os.Stdout代替http.ResponseWriter, 因为os.Stdout实现了io.Writer接口

// 模版中如何插入数据
// 上面演示了如何解析并渲染模版，接下来来更加详细的了解如何把数据渲染出来。一个模版都是应用在一个go的对象之上，go对象的字段如何插入到模版中？

// 字段操作
// go的模版通过{{}}来包含需要在渲染时被替换的字段，{{.}}表示当前的对象，这和Java或者C++中的this类似，如果要访问当前对象的字段通过{{.FieldName}}, 但是需要注意一点：这个字段必须是导出的(字段首字母必须是大写的)，否则在渲染的时候就会报错，
// 例
package main

import (
	"html/template"
	"os"
)

type Person struct {
	UserName string
}

func main() {
	t := template.New("fieldname example")
	t, _ = t.Parse("hello {{.Username}}")
	p := Person{UserName: "Astaxie"}
	t.Execute(t, p)
}
// 上面的代码可以正确的输出hello Adtaxie，但是如果稍微修改一下代码，在模版中含有了未导出的字段，那么就会报错
type Person struct {
	UserName string
	email string  // 未导出的字段，首字母是小写的
}
t, _ = t.Parse("hello {{.UserName}}! {{.email}}")
// 上面的代码就会报错，因为调用了一个未导出的字段，但是如果调用一个不存在的字段是不会报错的，而且输出为空
// 如果模版中输出{{.}}, 这个一般应用于字符串对象，默认会调用fmt包输出字符串的内容

// 输出潜逃字段内容
// 上面例子展示了如何针对一个对象的字段输出，那么如果字段里面还有对象，如何来循环的输出这些内容呢？可以使用{{with ...}}...{{end}}和{{range ...}}{{end}}来进行数据的输出
//	{{range}}这个和go语法里面的range类似，循环操作数据
//	{{with}}操作是指当前对象的值，类似上下文的概念
// 例 fieldname_example.go

// 条件处理
// 在go模版里面如果需要进行条件判断，那么可以使用和go的if-else类似的方式来处理，如果pipeline为空，那么if就认为是false
// 例 1.go
// 通过这个例子知道了if-else语法相当简单，在使用过程中很容易集成到模版代码中。
// 注意：if里面无法使用条件判断，例如.Mail=="astaxie@gmail.com", 这样的判断是不正确的，if里面只能是bool值

// pipeline
// Unix用户已经很熟悉pipe了，ls | grep "beego",过滤当前目录下面的文件，显示含有"beego"的数据，表达的意思就是前面的输出可以当作后面的输入，最后显示想要的数据，而go语言模版最强大的一点就是支持pipe数据，在go里面任何{{}}里面的都是pipeline数据，例如上面输出的email里面如果还有一些可能引起XSS注入的，如何来进行转化
{{. | html}}
// 在email输出的地方可以采用如上方式可以把输出全部转化为html的实体，上面的这种方式和平常写Unix的方式一模一样，操作起来相当方便，调用其他函数也是类似的方式

// 模版变量
// 有时候，在模版使用过程中需要定义一些局部变量，可以在一些操作中申明局部变量，例如withrangeif过程中申明局部变量，这个变量的作用域是{{end}}之前，go通过申明的局部变量格式如下：
$variable := pipeline
// 详细例子
{{with $x := "output" | printf "%q"}}{{$x}}{{end}}
{{with $x := "output" | printf "%q" $x}}{{end}}
{{with $x := "output" | printf}}{{$x | printf "%q"}}{{end}}

// 模版函数
// 模版在输出对象的字段值时，采用了fmt包把对对象转化成了字符串。但是有时候需求可能不是这样的，例如有时候为了防止垃圾邮件发送者通过采集网页的方式来发送给我们的邮箱信息，我们希望把@替换成at例如：
// astaxie at beego.me, 如果要实现这样的功能，就需要自定义函数来做这个功能
type FuncMap map[string]interface{}
// 例如，如果想要的email函数模版函数名时emailDeal，它关联的go函数名称是EmailDealWith，那么需要通过下面的方式来注册这个函数
t = t.Funcs(template.FuncMap("emailDeal": EmailDealWith))
// EmailDearWith这个函数的参数和返回值定义如下
func EmailDealWith(args ...interface{}) string
// 例2.go
// 上面演示了如何自定义函数，其实，在模版包内部已经有内置的实现函数，下面代码截取自模版包里面
var builtins = FuncMap{
	"and": and,
	"call": call,
	"html": HTMLEscaper,
	"index": index,
	"js": JSEscaper,
	"len": length,
	"not": not,
	"or": or,
	"print": fmt.Sprint,
	"printf": fmt.Sprintf,
	"println": fmt.Sprintln,
	"urlquery": URLQueryEscaper
}

// Must操作
// 模版包里面有一个函数Must，它的作用是检测模版是否正确，例如大括号是否匹配，注释是否正确的关闭，变量是否正确的书写。
// 接下来演示一个例子，用Must来判断模版是否正确
// 3.go
// 输出：
// The first one parsed OK.
// The second one parsed OK.
// The next one ought to fail.
// panic: template: check parse error with Must:1: unexpected "}" in operand

// 嵌套模版
// 平常开发Web应用的时候，经常会遇到一些模版有些部分是固定不变的，然后可以抽取出来作为一个独立的部分，例如一个博客的头部和尾部是不变的，而唯一改变的是中间的内容部分。所以可以定义成header、content、footer三个部分。go中通过如下的语法来申明
{{define "子模版名称"}}内容{{end}}
// 通过如下方式来调用
{{template "子模版名称"}}
// 接下来演示如何使用嵌套模版，定义三个文件，header.tmpl, content.tmpl, footer.tmpl文件
// 内容如下
// header.tmpl
// content.tmpl
// footer.tmpl
// 例：4.go
// 通过上面的例子，可以看到通过template.ParseFiles把所有的嵌套模版全部解析到模版里面，其实每一个定义的{{define}}都是一个独立的模板，它们相互独立，是并行存在的关系，内部其实存储的是类似map的一种关系(key是模板的名称，value是模板的内容)，然后通过ExecuteTemplate来执行相应子模版内容，可以看到header、footer都是相对独立的，都能输出内容，content中因为潜逃了header和footer的内容，就会同时输出三个的内容。但是当执行s1.Execute，没有任何的输出，因为在默认的情况下，没有默认的子模版，所以不会输出任何的东西
// 同一个集合类的模版是互相知晓的，如果同一模板被多个集合使用，则它需要在多个集合中分别解析


























