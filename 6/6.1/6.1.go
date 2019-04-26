// 6.1 session和cookie

// cookie
// cookie是有时间限制的，根据生命期不同分成两种:会话cookie和持久cookie
// 如果不设置过期时间，则表示这个cookie生命周期为从创建到浏览器关闭，只要关闭浏览器窗口，cookie就消失了。这种生命期为会话期的cookie被称为会话cookie。会话cookie哟版不保存在硬盘上而是保存在内存里。
// 如果设置了过期时间(setMaxAge(606024))，浏览器就会把cookie保存到硬盘上，关闭后再次打开浏览器，这些cookie依然有效知道超过设定的过期时间。存储在硬盘上的cookie可以在不同的浏览器进程间共享，比如两个IE窗口。二回羽保存在内存的cookie，不同的浏览器有不同的处理方式

// go设置cookie
// go通过net/http包中的SetCookie来设置
// http.SetCookie(w, http.ResponseWriter, cookie *Cookie)
// w表示需要写入的response，cookie是一个struct
type Cookie struct {
	Name string
	Value string
	Path string
	Domain string
	Expires time.Time
	RawExpires string
	// MaxAge = 0 means no 'Max-Age' attribute specified
	// MaxAge < 0 means delete cookie now, equivalently 'Max-Age: 0'
	// MaxAge > 0 means Max-Age attribute present and given in seconds
	MaxAge int
	Secure bool
	HttpOnly bool
	Raw string
	Unparsed []string // Raw text of unparsed attribute-value pairs
}
// 例子，设置cookie
expiration := *time.LocalTime()
expiration.Year += 1
cookie := http.Cookie{Name: "username", value:"astaxie", Expires: expiration}
http.SetCookie(w, &cookie)

// go读取cookie
cookie, _ := r.Cookie("username")
fmt.Fprint(w, cookie)
// 另一种读取方式
for _, cookie range r.Cookies() {
	fmt.Fprint(w, cookie.Name)
}

// session
// session机制是一种服务器端的机制，服务器使用一种类似于散列表的结构(也就是使用散列表)来保存信息。
// 但程序需要为某个客户端的请求创建一个session的时候，服务器首先检查这个客户端的请求里是否包含了一个session标识-称为session id，如果已经包含一个session id则说明疫情已经为此用户创建过session，服务器就按照session id把这个session检索出来使用(如果检索不到，可能会新建一个，这种情况可能出现在服务器端已经删除了该用户对应的session对象，但用户人为地在请求的URL后面附加上一个JESSION的参数)。如果客户请求不包含session id，则为此客户端创建一个session，并且同时生成一个与此session相关联的session id，这个session将在本次响应中返回给客户端保存
