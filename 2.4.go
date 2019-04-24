// 2.4 struct类型

// struct
// 可以声明新的类型，作为其他类型的属性或字段的容器。这样的类型，称之为struct
type person struct {
	name string,
	age int
}
// 使用
var P person
P.name = "Astaxie"
P.age = 25
fmt.Println("The person's name is %s", P.name)
// 还有两种声明
// 1. 按照顺序提供初始化值
P := person{"Tome", 25}
// 2. 通过field:value的方式初始化，这样可以任意顺序
P := person{age:24, name:"Tom"}
// 完整使用struct的例子
package main
import "fmt"

type person struct {
	name string,
	age int
}

func Older(p1, p2 person) (person, int) {
	if p1.age > p2.age {
		return p1, p1.age-p2.age
	}
	return p2, p2.age-p1.age
}

func main() {
	var tom person
	tom.name, tom.age = "Tom", 18
	bob := person{age:25, name:"Bob"}
	paul := person{"Paul", 43}

	tb_Older, tb_diff := Older(tom, bob)
	tp_Older, tp_diff := Older(tom, paul)
	bp_Older, bp_diff := Older(bob, paul)
	fmt.Printf("Of %s and %s, %s is older by %d years\n", tom.name, bob.name, tb_Older.name, tb_diff)
	fmt.Printf("Of %s and %s, %s is older by %d years\n", tom.name, paul.name, tp_Older.name, tp_diff)
	fmt.Printf("Of %s and %s, %s is older by %d years\n", bob.name, paul.name, bp_Older.name, bp_diff)
}

// struct的匿名字段
// 实际上go支持只提供类型，而不写字段名的方式，也就是匿名字段，也称为嵌入字段
// 当匿名字段是一个struct的时候，那么这个struct所拥有的全部字段都被隐式地引入了当前定义的这个struct
package main
import "fmt"

type Human struct {
	name string
	age int
	weight int
}

type Student struct {
	Human  // 匿名字段，那么默认Student就包含了Human的所有字段
	speciality string
}

func main() {
	mark := Student{Human{"Mark", 25, 120}, "Computer Science"}
	// 访问相应的字段
	fmt.Println("His name is ", mark.name)
	fmt.Println("His age is ", mark.age)
	fmt.Println("His weight is ", mark.weight)
	fmt.Println("His speciality is ", mark.speciality)
	// 修改对应的备注信息
	mark.speciality = "AI"
	fmt.Println("Mark changed his speciality")
	fmt.Println("His speciality is ", mark.speciality)
	// 修改年龄
	fmt.Println("Mark become old")
	mark.age = 46
	fmt.Println("His age is ", mark.age)
	// 修改体重
	fmt.Println("Mark is not an athlet anymore")
	mark.weight += 60
	fmt.Println("His weight is ", mark.weight)
}
// Student访问属性age和name的时候，就像访问自己所拥有的字段一样。
// Student还能访问Human这个字段作为字段名
mark.Human = Human{"Marcus", 55, 220}
mark.Human.age -= 1
// 所有内置类型和自定义类型都是可以作为匿名字段的
package main
import "fmt"

type Skills []string

type Human struct {
	name string
	age int
	weight int
}

type Student struct {
	Human  // 匿名字段，struct
	Skills  // 匿名字段，自定义的类型string slice
	int  // 内置类型作为匿名字段
	speciality string
}

func main() {
	jane := Student{Human:Human{"Jane", 35, 100}, speciality: "Biology"}
	// 访问
	fmt.Println("Her name is ", jane.name)
	fmt.Println("Her age is ", jane.age)
	fmt.Println("Her weight is ", jane.weight)
	fmt.Println("Her speciality is ", jane.speciality)
	// 修改skill
	jane.Skills = []string{"anatomy"}
	fmt.Println("Her skills are ", jane.Skills)
	fmt.Println("She acquired two new ones ")
	jane.Skills = append(jane.Skills. "pysics", "golang")
	fmt.Println("Her skills now are ", jane.Skills)
	// 修改匿名内置类型字段
	jane,int = 3
	fmt.Println("Her preferred number is ", jane.int)
}
// 如果human里面有一个字段叫做phone，而student也有一个字段叫做phone，怎么办？
// go中最外层的优先访问，也就是当通过student.phone访问的时候，是访问student里面的字段，而不是human里面的字段
// 这样就允许去冲在通过匿名字段继承的一些字段，当然如果像访问冲在后对应匿名类型里面的字段，可以通过匿名字段名来访问
package main
import "fmt"

type Human struct {
	name string
	age int
	phone string
}

type Employee struct {
	Human
	speciality string
	phone string
}

func main() {
	Bob := Employee{Human{"Bob", 34, "777-444-XXXX"}, "Designer", "333-222"}
	fmt.Println("Bob's work phone is: ", Bob.phone)
	fmt.Println("Bob's personal phone is: ", Bob.Human.phone)
}











































