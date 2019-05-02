// 10.1 设置默认地区
// 什么是Locale
// Locale是一组描述世界上某一特定区域文本格式和语言习惯的设置的集合。
// locale名通常由三个部分组成:
// 第一部分，是一个强制性的，表示语言的缩写，例如"en"表示英文或"zh"表示中文。
// 第二部分，跟在一个下划线之后，是一个可选的国家说明符，用于区分讲同一种语言的不同国家，例如"enUS"表示美国英语，而"enUK"表示英国英语。
// 最后一部分，跟在一个句点之后，是可选的字符集说明符，例如"zh_CN.gb1212"表示中国使用gb2312字符集
// go默认采用"UTF-8"编码集，所以实现il8n时不考虑第三部分，接下来都采用locale描述的前面两部分来作为il8n标准的locale名
// 在Linux和Solaris系统中可以通过locale -a 命令列举所有支持的地区名，读者可以看到这些地区名的命名规范。对于BSD等系统，没有locale命令，但是地区信息存储在/usr/share/locale中

// 设置Locale
// 有了上面对locale的定义，那么就需要根据用户的信息(访问信息、个人信息、访问域名等)来设置与之相关的locale，可以通过如下几种方式来设置用户的locale
// 通过域名设置Locale
// 设置Locale的办法之一就是在应用运行的时候采用域名分级的方式，例如，采用www.asta.com当作我们的英文站(默认站)，而把域名www.asta.cn当作中文站。这样通过在应用里面设置域名和相应的locale的对应关系，就可以设置好地区。这样处理由几点好处：
//	通过URL就可以很明显的识别
//	用户可以通过域名真直观的知道将访问哪种语言的站点
//	在go程序中实现非常的简单方便，通过一个map就可以实现
//	有利于搜索引擎抓取，能够提高站点的SEO
// 可以通过下面的代码来实现域名的对应locale
if r.Host == "www.asta.com" {
	il8n.SetLocale("en")
} else if r.Host == "www.asta.cn" {
	il8n.SetLocale("zh-CN")
} else if r.Host == "www.asta.tw" {
	il8n.SetLocale("zh-TW")
}
// 当然除了整域名设置地区之外，还可以通过子域名来设置地区，例如"en.asta.com"表示英文站点，"cn.asta.com"表示中文站点。实现代码如下所示:
prefix := strings.Split(r.Host, ".")

if prefix[0] == "en" {
	il8n.SetLocale("en")
} else if prefix[0] == "cn" {
	il8n.SetLocale("zh-CN")
} else if prefix[0] == "tw" {
	il8n.SetLocale("zh-TW")
}
// 通过域名设置Locale有如上所示的优点，但是一般开发Web应用的时候不会采用这种方式，因为首先域名成本比较高，开发以个Locale就需要一个域名，而且往往统一名称的域名不一定能申请的到，其次，我们不愿意为每个站点去本地化一个配置，而更多的是采用url后面带参数的方式

// 从域名参数设置Locale
// 目前最常用的设置Locale的方式是在URL里面带上参数，例如www.asta.com/hello?locale=zh或者www.asta.com/zh/hello。这样就可以设置地区:il8n.SetLocale(params["locale"])
// 这种设置方式几乎拥有前面讲的通过域名设置Locale的所有优点，它采用RESTful的方式，以使得不需要增加额外的方法来处理。当时这种方式需要在每一个的link里面增加相应的参数locale，这也许优点复杂而且有时候甚至相当的繁琐。不过可以写一个通用的函数url，让所有的link里面增加相应的参数locale，这也许有点负责而且有时候甚至相当的繁琐。不过可以写一个通用的函数url，让所有的link地址都通过这个函数来生成，然后在这个函数里面增加locale=params["locale"]的参数来缓解一下
// 也许希望URL地址看上去更加的RESTful一点，例如www.asta.com/en/books(英文站点)和www.asta.com/zh/books(中文站点)，这种方式的URL更加有利于SEO，而且对于用户也比较友好，能够通过URL直观的知道访问的站点。那么这样的URL地址可以通过router来获取locale(参考REST小节里面介绍的router插件实现)：
mux.Get("/:locale/books", listbook)

// 从客户端设置地区
// 在一些特殊的情况下，需要根据客户端的信息而不是通过URL来设置Locale，这些信息可能来自于客户端设置的喜好语言(浏览器中设置),用户的IP地址，用户在注册的时候填写的所在地信息等，这种方式比较适合Web为基础的应用
//	Accept-Language
// 客户端请求的时候在HTTP头信息里面有Accept-Language，一般的客户端都会设置该信息，下面是go实现的一个简单的根据Accept-Language实现设置地区的代码：
AL := r.Header.Get("Accept-Language")
if AL = "en" {
	il8n.SetLocale("en")
} else if AL == "zh-CN" {
	il8n.SetLocale("zh-CN")
} else if AL == "zh-TW" {
	il8n.SetLocale("zh-TW")
}
// 当晚在实际应用中，可能需要更加严格的判断来进行设置地区 - IP地址
// 另一种根据客户端来设定地区就是用户访问的IP，根据相应的IP库，对应访问的IP地区，目前全球比较常用的就是全球比较常用的就是GeoIP Li？？？？
//	用户profile
//	当然也可以让用户根据你提供的下拉菜单或者别的什么方式的设置相应的locale，然后将用户输入的信息，保存到与它账号相关的profile中，当用户再次登陆的时候把这个设置复写到locale设置中，这样就可以保证该用户每次访问都是基于自己先前设置的locale来获得页面















