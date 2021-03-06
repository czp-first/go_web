// 9.1 预防CSRF攻击

// 什么是CSRF
// CSRF(Cross-sit request foregery), 中文名称：跨站请求伪造，也被称为: one click attack/session riding, 缩写为：CSRF/XSRF
// CSRF到底能够干嘛呢？可以这样简单理解：攻击者可以调用你的登陆信息，以你的身份模拟发送各种请求。攻击者只要借助少许的社会工程学的诡计，例如通过QQ等聊天软件发送的链接(有些还伪装成短域名，用户无法分辨)，攻击者就能迫使Web应用的用户去执行攻击者预设的操作。例如，当用户登录网络银行去查看其存款余额，在他没有退出时，就点击了一个QQ好友发来的链接，那么该用户银行账户中的资金就有可能被转移到攻击者指定的账户中
// 所以遇到CSRF攻击时，将对终端用户的数据和操作指令构成严重的威胁；当受攻击的终端用户具有管理员账户的时候，CSRF攻击将危机整个Web应用程序
// CSRF的原理
// 要完成一次CSRF攻击，受害者必须依次完成两个步骤
//	1. 登陆受信任网站A，并在本地生成Cookie
//	2. 在不退出A的情况下，访问危险网站B
// 看到这里，也许会问:"也许不满足以上两个条件中的任意一个，就不会收到CSRF的攻击"，是的，确实如此，但你不能保证以下情况不会发生：
//	你不能保证你登陆了一个网站后，不再打开一个tab页面并访问另外的网站，特别现在浏览器都是支持多tab的
//	你不能保证你关闭浏览器了后，你本地的Cookie立刻过期，你上次的会话已经结束
//	上图中所谓的攻击网站，可能是一个存在其他漏洞的可信任的经常被人访问的网站
// 因此对于用户来说很难避免在登陆一个网站之后不惦记一些链接进行其他操作，所以随时可能成为CSRF的受害者
// CSRF攻击主要是因为Web的隐式身份验证机制，Web的身份验证机制虽然可以保证一个请求是来自于某个用户的浏览器，但却无法保证该请求是用户批准发送的

// 如何预防CSRF
// CSRF的防御可以从服务端和客户端两方面着手，防御效果是从服务端着手效果比较好，现在一般的CSRF防御也都在服务端进行
// 服务端的预防CSRF攻击的方式方法有多种，但思想上都差不多，主要从以下两个方面入手
//	1. 正确使用GET、POST和Cookie
//	2. 在非GET请求中增加伪随机数
// 上一章介绍过REST方式的Web应用，一般而言，普通的Web应用都是以GET、POST为主，还有一种请求是Cookie方式。
// 一般都是按照如下方式设计应用：
//	1， GET常用在查看，列举，展示等不需要改变资源属性的时候；
//	2. POST常用在下达订单，改变一个资源的属性或者做其他一些事情
// 接下来就以go语言来举例说明，如何限制对资源的访问方法:
mux.GET("/user/:uid", getuser)
mux.POST("/user/:uid", modifyuser)
// 这样处理后，因为限定了修改只能使用POST，当GET方式请求时就拒绝响应，所以GET方式的CSRF攻击就可以防止了，但这样就能全部解决问题了么？当然不是，因为POST也是可以模拟的
// 因此需要实施第二步，在非GET方式的请求中增加随机数，这个大概有三种方式来进行：
//	为每个用户生成一个唯一的cookie token，所有表单都包含同一个伪随机值，这种方案最简单，因为攻击者不能获得第三方的Cookie(理论上),所以表单中的数据也就构造失败，但是由于用户的Cookie很容易由于网站的XSS漏洞而被盗取，所以这个方案必须要在没有XSS的情况下才安全
//	每个请求使用验证码，这个方案是完美的，因为要多次输入验证码，所以用户友好性很差，所以不适合实际运用
//	不同的表单包含一个不同的伪随机值，在4.4小节介绍"如何防止表单多次递交"时介绍过此方案，服用相关代码，实现如下：
// 生成随机数token
h := md5.New()
io.WriteString(h, strconv.FormatInt(crutime, 10))
io.WriteString(h, "ganraomaxxxxxxxxxx")
token := fmt.Sprintf("%x", h.Sum(nil))
t, _ := template.ParseFiles("login.gtpl")
t.Execute(w, token)
// 输出token
<input type="hidden" name="token" value={{.}}>
// 验证token
r.ParseForm()
token := r.Form.Get("token")
if token != "" {
	// 验证token的合法性
} else {
	// 不存在token报错
}
// 这样基本就实现了安全的POST，但是如果破解了token的算法呢，按照理论上说是，但是实际上破解是基本不可能的，因为有人曾计算过，暴力破解该串大概需要2的11次方时间


























