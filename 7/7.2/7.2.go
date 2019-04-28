// 7.2 JSON处理
// json串：{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}

// 解析JSON

// 解析到结构体
// go的JSON包中有如下函数
// func Unmarshal(data []byte, v interface{}) error

// 例子代码：parse_json_to_struct.go

// 在例子中，首先定义了与json数据相应的结构体，数组对应slice，字段名对应JSON里面的KEY，在解析的时候，如何将json数据与struct字段相匹配呢？例如JSON的key是Foo，那么怎么找对应的字段呢？
//	首先查找tag含有Foo的可导出的struct字段(首字母大写)
//	其次查找字段名为Foo的导出字段
//	最后查找类似FOO或者FoO这样的处理首字母之外其他大小写不敏感的导出字段
// 能够被赋值的字段必须是可导出字段(即首字母大写)。同时JSON解析的时候只会解析能找得到的字段，如果找不到的字段会被忽略，这样的一个好处是：当你接收到一个很大的JSON数据结构而你想获取其中的部分数据的时候，你只需将你想要的数据对应的字段名大写，即可轻松解决这个问题

// 解析道interface
// 上面那种解析方式是在知晓被解析的JSON数据的结构的前提下采取的方案，如果不知道被解析的数据的格式，又应该如何来解析？

// interface{}可以用来存储任意数据类型的对象，这种数据结构正好用于存储解析的未知结构的json数据的结果。JSON包中采用map[string]interface{}和[]interface{}结构来存储任意的JSON对象和数组。go类型和JSON类型的对应关系如下
//	bool 代表 JSON booleans
//	float64 代表 JSON numbers
//	string 代表 JSON strings
//	nil 代表 JSON null
// 假设有如下的JSON数据
b := []byte{`{"Name": "Wednesday", "Age": 6, "Parents": ["Gomez", "Morticia"]}`}
// 如果在不知道它的结构的情况下，把它解析到interface{}里面
var f interface{}
err := json.Unmarshal(b, f)
// 这个时候f里面存储了一个map类似，他们的key是string，值存储在空的interface{}里
f = map[string]interface{}{
	"Name": "Wednesday",
	"Age": 6,
	"Parents": []interface{}{
		"Gomez",
		"Morticia",
	},
}
// 那么如何来访问这些数据呢？通过断言的方式
m := f.(map[string]interface{})
// 通过断言之后，就可以通过如下方式来访问里面的数据了
for k, v := range m {
	switch vv := v.(type) {
	case string:
		fmt.Println(k, "is string", vv)
	case int:
		fmt.Println(k, "is int", vv)
	case []interface{}:
		fmt.Println(k, "is an array:")
		for i, u := range vv {
			fmt.Println(i, u)
		}
	default:
		fmt.Println(k, "is of a type I don't know how to handle")
	}
}
// 通过上面的示例可以看到，通过interface{}与type assert的配合，就可以解析未知结构的JSON数了
// 上面这个是官方提供的解决方案，其实很多时候通过类型断言，操作起来不是很方便，目前bitly公司开园了一个叫做simplejson的包，在处理未知结构体的JSON时相当方便，详细例子如下所示：
js, err := NewJson([]byte(`{
	"test": {
		"array": [1, "2", 3],
		"int": 10,
		"float": 5.150,
		"bignum": 9223372036954775807,
		"string": "simplejson",
		"bool": true
	}
}`))
arr, _ := js.Get("test").Get("array").Array()
i, _ := js.Get("test").Get("int").Int()
ms := js.Get("test").Get("string").MustString()
// 使用这个哭操作JSON比起官方包来说，简单的多，详细参考如下地址：http://github.com/bilty/go-simplejson

// 生成JSON
// 输出JSON数据传，通过JSON包里面的Marshal函数来处理
// 函数定义如下
// func Marchsal(v interface{}) ([]byte, error)
// output_json.go输出上面的服务器列表信息
// 上面的输出字段名都是大写的，如果想用小写的怎么办？JSON输出的时候必须注意，只有导出的字段才会被输出，如果修改字段名，那么就会发现什么都不会输出，所以必须通过struct tag定义来实现
type Server struct {
	ServerName string `json:"serverName"`
	ServerIP string `json:"serverIP"`
}

type Serverslice struct {
	Servers []Server `json:"servers"`
}
// 通过修改上面的结构题定义，输出的JSON串就和最开始定义的JSON串保持一致了
// 针对JSON的输出，在定义struct tag的时候需要注意的几点是：
//	字段的tag是"-"，那么这个字段不会输出到JSON
//	tag中带有自定义名称，那么这个自定义名称会出现在JSON的字段名中，例如上面例子中serverName
//	tag中如果带有"omitempty"选项，那么如果该字段值为空，就不会输出到JSON串中
//	如果字段类型是bool，string，int，int64，而tag中带有",string"选项，那么这个字段在输出到JSON的时候会把该字段对应的值转换成JSON字符串
// 举例来说
type Serve struct {
	// ID 不会到处到JSON
	ID int `json:"-"`
	// ServerName的值会进行二次JSON编码
	ServerName string `json:"serverName"`
	ServerName2 string `json:"serverName2, string"`
	// 如果ServerIP为空，则不输出到JSON串中
	ServerIP string `json:"serverIP,omitempty"`
}

s := Server {
	ID : 3,
	ServerName: `Go "1.0" `,
	ServerName2: `Go "1.0" `,
	ServerIP: ``,
}
b, _ := json.Marshal(s)
os.Stdout.Write(b)
// 会输出一下内容:
// {"serverName":"Go \"1.0\" ","serverName2":"\"Go \\\"1.0\\\" \""}
// Marshal函数只有在转换成功的时候才会返回数据，在转换的过程中需要注意几点：
//	JSON对象只支持string作为key，所以要编码一个map，那么必须是map[string]T这种类型(T是go中的任意类型)
//	Channel，complex和function是不能被编码成JSON的
//	嵌套的数据是不能编码的，不然会让JSON编码进入死循环
//	指针在编码的时候会输出指针指向的内容，而空指针会输出null












































