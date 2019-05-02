// 本地化资源
// 设置好Locale之后需要解决的问题就是如何存储相应的Locale对应的信息？这里面的信息包括:文本信息、时间和日期、货币值、图片、包含文件以及视图等资源。那么接下来将对这些信息一一进行介绍，go中把这些格式信息存储在JSON中，然后通过合适的方式展现出来。(接下来以中文和英文两种语言对比举例，存储格式文件en.json和zh-CN.json)

// 本地化文本消息
// 本信息是编写Web应用中最常用到的，也是本地化资源中最多的信息，想要以适合本地语言的方式来显示文本信息，可行的一种方案是:建立需要的语言相应的map来维护一个key-value的关系，在输出之前按需从适合的map中去获取相应的文本，如下是一个简单的示例：
// 1.go
// 上面示例演示了不同locale的文本翻译，实现了中文和英文对于同一个key显示不同语言的实现，上面实现了中文的文本消息，如果想切换到英文版本，只需要把lang设置为en即可
// 有些时候仅是key-value替换是不能满足需要的，例如"I am 30 years old", 中文表达是"我今年30岁了", 而此处的30是一个变量，该怎么办呢？这个时候，可以结合fmt.Printf函数来实现，请看下面的代码：
en["how old"] = "I am %d years old"
cn["how old"] = "我今年%d岁了"
fmt.Pringf(msg(lang, "how old"), 30)
// 上面的示例代码仅用以演示内部的实现方案，而实际数据是存储在JSON里面的，所以可以通过json.Unmarshal来为相应的map填充数据

// 本地化日期和时间
// 因为时区的关系，同一时刻，在不同的地区，表示是不一样的，而且因为Locale的关系，时间格式也不尽相同，例如中文环境下可能显示:2012年10月24日 星期三 23时11分13秒 CST，而在英文环境下可能显示Wed Oct 24 23:11:13 CST 2012。这里需要解决两点：
//	1. 时区问题
//	2. 格式问题
// $GOROOT/lib/time包中的timeinfo.zip含有locale对应的时区的定义，为了获得对应于当前locale的时间，应首先使用time.LoadLocation(name string)获取相应于地区的locale,比如Asia/Shanghia或America/Chicago对应的时区信息，然后再利用此信息与调用time.Now获得的Time对象协作来获得最终的时间
// 详细的请看下面的例子(该例子采用上面例子的一些变量):
en["time_zone"] = "America/Chicago"
cn["time_zone"] = "Asia/Shanghai"
loc, _ := LoadLocation(msg(lang, "time_zone"))
t := time.Now()
t = t.In(loc)
fmt.Println(t.Format(time.RFC3339))
// 可以通过类似处理文本格式的方式来解决时间格式的问题，举例如下：
en["date_format"] = "%Y-%m-%d %H:%M:%S"
en["date_format"] = "%Y年%m月%d日 %H时%M分%S秒"
fmt.Println(date(msg(lang, "date_format"), t))

func date(formate string, t time.Time) string {
	year, month, day = t.Date()
	hour, min, sec = t.Closk()
	// 解析相应的%Y %m %d %H %M %S然后返回信息
	// %Y 替换成2012
	// %m 替换成10
	// %d 替换成24
}

// 本地化货币值
// 各个地区货币表示也不一样，处理方式也与日期差不多，细节请看下面代码
en["money"] = "USD %d"
cn["money"] = "¥ %d元"
fmt.Println(date(msg(lang, "date_format"), 100))
func money_format(fomate string, money int64) string {
	return fmt.Sprintf(fmate, money)
}

// 本地化视图和资源
// 可能会根据Locale的不同来展示视图，这些视图包含不同的图片、css、js等各种静态资源。那么如何来处理这些信息呢？首先应按locale来组织文件信息，请看下面的文件目录安排
views
|--en	// 英文模版
	|--images	// 存储图片信息
	|--js	// 存储JS文件
	|--css	// 存储css文件
	index.tpl	// 用户首页
	login.tpl	// 登陆首页
|--zh-CN
	|--images
	|--js
	|--css
	index.tpl
	login.tpl
// 有了这个目录结构后就可以在渲染的地方这样来实现代码：
s1, _ := templae.ParseFiles("views"+lang+"index.tpl")
vv.Lang = lang
s1.Execute(os.Stdout, VV)
// 而对于里面的index.tpl里面的资源设置如下:
// js文件
<script type="text/javascript" src="views/{{.VV.Lang}}/js/jqery/jquery-1.8.0.min.js"></script>
// css文件
<link href="views/{{.VV.Lang}}/css/bootstrap-responsive.min.css" rel="stylesheet">
// 图片文件
<img src="views/{{.VV.Lang}}/images/btn.png">
// 采用这种方式来本地化视图以及资源时，就可以很容易的进行扩展了















