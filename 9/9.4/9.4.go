// 9.4 避免SQL注入

// 什么是SQL注入
// SQL注入攻击(SQL Injection)，简称注入攻击，是Web开发中最常见的一种安全漏洞。可以用它来从数据库获取敏感信息，或者利用数据库的特性执行添加用户，导出文件等一系列恶意操作，甚至有可能获取数据库乃至系统用户最高权限
// 而造成SQL注入的原因是因为程序没有有效过滤用户的输入，使攻击者成功的向服务器提交恶意的SQL查询代码，程序在接收后错误的将攻击者的输入作为查询语句的一部分执行，导致原始的查询逻辑被改变，额外的执行了攻击者精心构造的恶意代码

// SQL注入实例
// 很多web开发者没有意识到SQL查询是可以被篡改的，从而把查询当作可信任的命令。殊不知，SQL查询是可以绕开控制访问，从而绕过身份验证和权限检查的。更有甚者，有可能通过SQL查询去运行主机系统级的命令
// 下面通过一些真实的例子详细讲解SQL注入的方式
// 考虑以下简单的登录表单:
<form action="/login" method="POST">
<p>Username: <input type="text" name="username" /></p>
<p>Password: <input type="password" name="password" /></p>
<p><input type="submit" value="登陆"/></p>
</form>
// 处理里面的SQL可能是这样的
username := r.Form.Get("username")
password := r.Form.Get("password")
sql := "SELECT * FROM user WHERE username='"+username"' AND password='"+password"'" 
// 如果用户的输入的用户名如下，密码任意
myuser' or 'foo' = 'foo' --
// 那么SQL变成了如下所示
SELECT * FROM user WHERE username='username' or 'foo' == 'foo' -- '' AND password = 'xxx'
// 在SQL里面--是注释标记，所以查询语句会在此中断。这就让攻击者在不知道任何合法用户名和密码的情况下成功登录了
// 对于MSSQL还有更加危险的一种SQL注入，就是控制系统，下面这个可怕的例子将演示如何在某些版本的MSSQL数据库上执行系统命令
sql := "SELECT * FROM products WHERE name LIKE '%"+product"%'"
Db.Exec(sql)
// 如果攻击提交a%' exec master..xp_cmdshell 'net user test testpass /ADD' -- 作为变量prod的值，那么sql将会变成
sql := "SELECT * FROM products WHERE name LIKE '%a%' exec master..xp_cmdshell 'net user test testpass /ADD'"
// MSSQL服务器会执行这条SQL语句，包括它后面那个用于向系统添加新用户的命令。如果这个程序是以sa运行而MSSQLSERVER又有足够的权限的话，攻击者就可以获得一个系统账号来访问主机了
// 虽然以上的例子是针对某一特定的数据库系统的，但是这并不代表不能对其他数据库系统实施类似的攻击。针对这种安全漏洞，只要使用不同方法，各种数据库都有可能遭殃

// 如何预防SQL注入
// 也许你会说攻击者要知道数据库结构的信息才能实施SQL注入攻击。确实如此，但没人能保证攻击者一定拿不到这些信息，一旦他们拿到了，数据库就存在泄漏的危险。如果你在用开放源代码的软件包来访问数据库，比如论坛程序，攻击者就很容易得到相关的代码。如果这些代码设计不良的话，风险就更大了。目前Discuz、phpwind、phpcms等这些流行的开源程序都有被SQL注入攻击的先例
// 这些攻击总是发生在安全性不高的代码上。所以永远不要相信外界输入的数据，特别是来自于用户的数据，包括选择框、表单隐藏域和cookie就如上面的第一个例子那样，就算是正常的查询也有可能造成灾难
// SQL注入攻击的危害这么大，那么该如何来防治呢？下面这些建议或许对防治SQL注入有一定的帮助
//	1. 严格限制Web应用的数据库的操作权限，给此用户提供仅仅能够满足其工作的最低权限，从而最大限度的减少注入攻击对数据库的危害
//	2. 检查输入的数据是否具有所期望的数据格式，严格限制变量的类型，例如使用regexp包进行一些匹配处理。或者使用strconv包对字符串转化成其他基本类型的数据进行判断
//	3. 对进入数据库的特殊字符('"\尖括号&*;等)进行转义处理，或编码转换。go的text/template包里面的HTMLEscapeString函数可以对字符串进行转义处理
//	4. 所有的查询语句建议使用数据库提供的参数化查询接口，参数化的语句使用参数而不是将用户输入变量嵌入到SQL语句中，即不要直接拼接SQL，例如使用database/sql里面的查询函数Prepare和Query，或者Exec(query string, args ...interface{})
//	5. 在应用发布之前见识使用专业的SQL注入检测工具进行检测，以及时修补被发现的SQL注入漏洞。网上有很多这方面的开源工具，例如sqlmap、SQLninja等
//	6. 避免网站打印出SQL错误信息，比如类型错误、字段不匹配等，把代码里的SQL语句暴露出来，以防止攻击者利用这些错误信息进行SQL注入