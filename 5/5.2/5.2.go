// 5.2 使用MySql数据库

// MySQL驱动
// Go中支持MySQL的驱动目前比较多，有如下几种
// https://github.com/Go-SQL-Driver/MySQL 支持database/sql，全部采用go写
// https://github.com/ziutek/mymysql 支持database/sql，也支持自定义的接口，全部采用go写
// https://github.com/Phillo/Gomysql 不支持database/sql，自定义接口，全部采用go写

// 接下来的例子主要以第一个驱动为例，理由
// 这个驱动比较新，维护的比较好
// 完全支持database/sql接口
// 支持keepalive，保持长连接，这个从底层就支持keepalive

// 示例代码
// mysql.go

// 关键的几个函数
// sql.Open()函数用来打开一个注册过的数据库驱动，Go-MySQL-Driver中注册了mysql这个数据库驱动，第二个参数是DNS(Data Source Name)，它是Go-MySQL-Driver定义的一些数据库链接和配置信息
// 支持如下格式
// user@unix(/path/to/socket)/dbname?charset=utf8
// user:password@tcp(localhost:5555)/dbname?charset=utf8
// user:password@/dbname
// user:password@tcp/([de:ab:be:ef::ca:fe]:80)/dbname
// db.Parse()函数用来返回准备要执行的sql操作，然后返回准备完毕的执行状态
// db.Query()函数用来直接执行sql返回Rows结果
// stmt.Exec()函数用来执行stmt准备好的SQL语句
// 传入的参数都是=？对应的数据，这样做的方式可以一定程度上防止sql注入】