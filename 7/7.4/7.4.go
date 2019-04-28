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
































