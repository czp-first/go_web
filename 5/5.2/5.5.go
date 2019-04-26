// 5.5 使用beedb库进行ORM开发

// 安装
// go get github.com/astaxie/beedb

// 如何初始化
import (
	"database/sql"
	"github.com/astaxie/beedb"
	_ "github.com/ziutek/mymysql/godrv"
)
// 导入必须的package之后，需要打开数据库的链接，然后创建一个beedb对象(以MySQL为例)
db, err := sql.Open("mymysql", "test/xiemengjun/123456")
if err != nil {
	panic(err)
}
orm := beedb.New(db)
// beede的New函数实际上应该有两个参数，第一个参数是标准接口的db，第二个参数是使用的数据库引擎，如果你使用的数据库引擎是MySQL/Sqlite，那么第二个参数可以省略
// 如果使用的数据库是SQLServer，初始化需要：
orm = beedb.New(db, "mssql")
// 如果使用了PostgreSQL，初始化需要：
orm = beedb.New(db, "pg")
// beedb支持打印调试，可以通过如下代码实现调试
beedb.OnDebug = true
// 使用前面的数据库表Userinfo，建立相应的struct
type Userinfo struct {
	Uid int `pk` // 如果表的主键不是id，那么需要加上pk注释，显式的说这个字段是主键
	Username string
	Departname string
	Created time.time
}
// 注意，beedb针对驼峰命名会自动帮你转化成下划线字段，例如你定义了Struct名字为UserInfo，那么转化成底层实现的时候是user_info，字段命名也遵循该规则

// 插入数据
// 操作的是struct对象，而不是原生sql语句，最后通过调用Save接口将数据保存到数据库
var saveone Userinfo
saveone.Username = "Test Add User"
saveone.Departname = "Test Add Departname"
saveone.Created = time.Now()
orm.Save(&saveone)
// 插入之后save.Uid就是插入成功之后的自增ID。Save接口会自动帮你存进去
// beedb接口提供了另外一种插入的方式，map数据插入
add := make(map[string]interface{})
add["username"] = "astaxie"
add["departname"] = "cloud_develop"
add["created"] = "2012-12-02"
orm.SetTable("userinfo").Insert(add)
// 插入多条数据
addSlice := make([]map[string]interface{})
add := make(map[string]interface{})
add2 := make(map[string]interface{})
add["username"] = "astaxie"
add["departname"] = "cloud develop"
add["created"] = "2012-12-02"
add2["username"] = "astaxie2"
add2["departname"] = "cloud develop2"
add2["created"] = "2012-12-02"
addslice.append(addslice, add, add2)
orm.SetTable("userinfo").Insert(addslice)
// 上面调用SetTable函数是显式的告诉ORM，要执行的map对应的数据表是userinfo

//更新数据
// 继续上面的例子来演示更新操作，现在saveone的主键已经有值了，此时调用save接口，beedb内部会自动调用update
saveone.Username = "Update Username"
saveone.Departname = "Update Departname"
saveone.Created = time.Now()
orm.Save(&saveone)  // 现在saveone有了主键值，就执行更新操作
// 更新数据也支持直接使用map操作
t := make(map[string]interface{})
t["username"] = "astxie"
orm.SetTable("userinfo").SetPK("uid").Where(2).Update(t)
// SetPK: 显式的告诉ORM，数据库表userinfo的主键是uid
// Where: 用来设置条件，支持多个参数，第一个参数如果为整数，相当于调用了Where("主键=?",值)。Update函数直接接收map类型的数据，执行更新数据

// 查询数据
// beedb的查询接口比较灵活
// 例1:根据主键获取数据
var user Userinfo
orm.Where("uid=?", 37).Find(&user)
// 例2:
var user2 Userinfo
orm.Where(3).Find(&user2)
// 例3:不是主键类型的条件
var user3 Userinfo
orm.Where("name=?", "john").Find(&user3)
// 例4:更加复杂的条件
var user4 Userinfo
orm.Where("name=? and age<?", "jojn", 88).Find(&user4)

// 可以通过如下接口获取多条数据
// 例1:根据条件id>3,获取20位置开始的10条数据的数据
var allusers []Userinfo
err := orm.Where("id>?", "3").Limit(10, 20).FindAll(&allusers)
// 例2:省略limit的第二个参数，默认从0开始，获取10条数据
var tenusers []Userinfo
err := orm.Where("id>?", 3).Limit(10).FindAll(&tenusers)
// 例3:获取全部数据
var everyone []Userinfo
err := orm.OrderBy("uid desc, Username asc").FindAll(&everyone)
// Limit函数，是用来控制查询结构条数的，第一个参数表示查询的条数，第二个参数表示读取数据的起始位置，默认为0
// OrderBy: 这个函数用来进行查询排序，参数是需要排序的条件
// 上面的例子都是将获取的数据直接映射成struct对象，如果只是想获取一些数据到map，以下方式可以实现
a, _ := orm.SetTable("userinfo").SetPK("uid").Where(2).Select("uid, username").FindMap()
// Select函数用来指定需要查询多少个字段。默认为全部字段*。
// FindMap()函数返回的是[]map[string][]byte类型，需要自己作类型转换

// 删除数据
// 例1，删除单条数据
// saveone就是上面示例中的那个saveone
orm.Delete(&saveone)
// 例2，删除多条数据
// alluser就是上面定义的获取多条数据的slice
orm.DeleteAll(&alluser)
// 例3:根据sql删除数据
orm.SetTable("userinfo").Where("uid>?", 3).DeleteRow()











































