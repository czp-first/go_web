// 6.4 预防session劫持
// session劫持是一种广泛存在的比较严重的安全威胁，在session技术中，客户端和服务端通过session的标识符来维持会话，但这个标识符很容易就能被嗅探到，从而被其他人利用。它是中间人攻击的一种类型

// session劫持过程
// 如下代码来展示一个count计数器：
func count(w http.ResponseWriter, r *http.Request) {
	sess := globalSessions.SessionStart(w, r)
	ct := sess.Get("countnum")
	if ct == nil {
		sess.Set("countnum", 1)
	} else {
		sess.Set("countnum", (ct.(int) + 1))
	}
	t, _ := template.ParseFiles("count.gtpl")
	w.Header().Set("Content-Type", "text/html")
	t.Execute(w, sess.Get("countnum"))
}
// count.gtpl的代码如下所示
// Hi. Now count:{{.}}

// 随着刷新，数字将不断增长，当数字显示为6的时候，打开浏览器(以chrome为例)的cookie管理器，可以看到cookie信息
// 打开另一个浏览器(firefox)，赋值chrome地址栏里的地址到新打开的浏览器的地址栏中，然后打开firefox的cookie模拟插件，新建一个cookie，把上图中的cookie内容原样在firefox中重建一份
// 回车后，将看到Hi. Now count:7
// 劫持seesion成功
// 可以看到虽然欢乐浏览器，但是却获得了sessionID，然后模拟了cookie存储的过程。这个例子是在同一台计算机上做的，不过即使换用两台来做，其结果仍然一样。此时如果交替点击两个浏览器里的链接，会发现它们其实操纵的是同一个计数器。不必惊讶，此处firework盗用了chrome和goserver之间的维持会话的钥匙，即gosessionid，这是一种类型的会话劫持。
// 在goserver看来，它从http请求中得到了一个gosessionid，由于HTTP协议的无状态性，它无法得知这个gosessionid是从chrome那里劫持来的，它依然会去查找相应的session，并执行相关计算。与此同时chrome也无法得知自己的会话已经被劫持

// session劫持防范

// cookieonly和token
// 通过上面劫持的简单演示可以了解到session一旦被其他人劫持，就非常危险，劫持者可以假装成被劫持者进行很多非法操作。那么如何有效的防止session劫持呢。
// 其中一个解决方案就是sessionID的值值允许cookie设置，而不是通过URL重置方式设置，同时设置cookie的httponly为true，这个属性是设置是否可通过客户端脚本访问这个设置的cookie，第一这个可以防止这个cookie被XSS读取从而引起session劫持，第二cookie设置不会像URL重置方式那么容易获取sessionID。
// 第二步就是在每个请求里面加上token，实现类似前面章节里面讲的防止form重复递交类似的功能，在每个请求里面加上一个隐藏的token，然后每次验证这个token，从而保证用户的请求都是唯一性
h := md5.New()
salt := "astaxie%^7&8888"
io.WriteString(h, salt+time.Now().String())
token := fmt.Sprintf("%x", h.Sum(nil))
if r.Form["token"] != token {
	// 提示登录
}
sess.Set("token", token)

// 间接生成新的SID
// 还有一个解决方案就是，给session额外设置一个创建时间的值，一旦过了一定的时间，销毁这个sessionID，重新生成新的sessionID，这个可以一定程度上防止session劫持的问题
createtime := sess.Get("createtime")
if createtime == nil {
	sess.Set("createtime", time.Now().Unix())
} else if (createtime.(int64) + 60) < (time.Now().Unix()) {
	globalSessions.SessionDestroy(w, r)
	sess = globalSessions.SessionStart(w, r)
}
// session启动后，设置了一个值，用于记录生成sessionID的时间。通过判断每次请求是否过期(这里设置了60秒)定期生成新的ID，这样使得攻击者获取有效sessionID的机会大大降低。
// 上面两个手段的组合可以在实践中消除session劫持的风险，一方面，由于sessionID频繁改变，使攻击者难有机会获取有效的sessionID;另一方面，因为sessionID只能在cookie中传递，然后设置了httponly，所以基于URL攻击的可能性为零，同时被XSS获取sessionID也不可能。最后，由于还设置了MaxAge=0，这就相当于session cookie不会留在浏览器的历史记录里面


























