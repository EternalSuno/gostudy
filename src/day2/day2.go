package day2

import (
	"fmt"
)

func init() {
	// struct
	// Go语言中，也和C或者其他语言一样，我们可以声明新的类型，作为其它类型的属性或字段的容器。
	// 例如，我们可以创建一个自定义类型person代表一个人的实体。
	// 这个实体拥有属性：姓名和年龄。这样的类型我们称之struct。
	type preson struct {
		name string
		age  int
	}

	var P preson
	P.name = "Astaxie"                            // 赋值"Astaxie" 给P的name属性
	P.age = 25                                    //赋值"25"给变量P的age属性
	fmt.Printf("The person's name is %s", P.name) //访问P的属性
	//输出 The person's name is Astaxieslice

	//除了上面这种P的声明使用之外，还有另外几种声明使用方式：
	//1.按照顺序提供初始化值
	//P := person{"Tom", 25}
	//2.通过field:value的方式初始化，这样可以任意顺序
	//P:=persion{age:24, name:"Tom"}
	//3.当然也可以通过new函数分配一个指针，此处P的类型为*person
	//P := new(person)

	// example
	//package main
	//
	//import "fmt"
	//
	//// 声明一个新的类型
	//type person struct {
	//	name string
	//	age int
	//}
	//
	//// 比较两个人的年龄，返回年龄大的那个人，并且返回年龄差
	//// struct也是传值的
	//func Older(p1, p2 person) (person, int) {
	//	if p1.age>p2.age {  // 比较p1和p2这两个人的年龄
	//		return p1, p1.age-p2.age
	//	}
	//	return p2, p2.age-p1.age
	//}
	//
	//func main() {
	//	var tom person
	//
	//	// 赋值初始化
	//	tom.name, tom.age = "Tom", 18
	//
	//	// 两个字段都写清楚的初始化
	//	bob := person{age:25, name:"Bob"}
	//
	//	// 按照struct定义顺序初始化值
	//	paul := person{"Paul", 43}
	//
	//	tb_Older, tb_diff := Older(tom, bob)
	//	tp_Older, tp_diff := Older(tom, paul)
	//	bp_Older, bp_diff := Older(bob, paul)
	//
	//	fmt.Printf("Of %s and %s, %s is older by %d years\n",
	//		tom.name, bob.name, tb_Older.name, tb_diff)
	//  //输出: Of Tom and Bob, Bob is older by 7 years
	//	fmt.Printf("Of %s and %s, %s is older by %d years\n",
	//		tom.name, paul.name, tp_Older.name, tp_diff)
	//  //输出: Of Tom and Paul, Paul is older by 25 years
	//	fmt.Printf("Of %s and %s, %s is older by %d years\n",
	//		bob.name, paul.name, bp_Older.name, bp_diff)
	//  //输出:Of Bob and Paul, Paul is older by 18 years
	//}

	//struct的匿名字段
	//实际上Go支持只提供类型，而不写字段名的方式，也就是匿名字段，也称为嵌入字段。
	//当匿名字段是一个struct的时候，那么这个struct所拥有的全部字段都被隐式地引入了当前定义的这个struct。

	//example:
	//package main
	//
	//import "fmt"
	//
	//type Human struct {
	//	name string
	//	age int
	//	weight int
	//}
	//
	//type Student struct {
	//	Human  // 匿名字段，那么默认Student就包含了Human的所有字段
	//	speciality string
	//}
	//
	//func main() {
	//	// 我们初始化一个学生
	//	mark := Student{Human{"Mark", 25, 120}, "Computer Science"}
	//
	//	// 我们访问相应的字段
	//	fmt.Println("His name is ", mark.name)
	//	fmt.Println("His age is ", mark.age)
	//	fmt.Println("His weight is ", mark.weight)
	//	fmt.Println("His speciality is ", mark.speciality)
	//	// 修改对应的备注信息
	//	mark.speciality = "AI"
	//	fmt.Println("Mark changed his speciality")
	//	fmt.Println("His speciality is ", mark.speciality)
	//	// 修改他的年龄信息
	//	fmt.Println("Mark become old")
	//	mark.age = 46
	//	fmt.Println("His age is", mark.age)
	//	// 修改他的体重信息
	//	fmt.Println("Mark is not an athlet anymore")
	//	mark.weight += 60
	//	fmt.Println("His weight is", mark.weight)
	//
	//	mark.Human = Human{"Marcus", 55, 220}
	//	mark.Human.age -= 1
	//
	//}

	//通过匿名访问和修改字段相当的有用，但是不仅仅是struct字段哦，
	//所有的内置类型和自定义类型都是可以作为匿名字段的。
	//example:
	//package main
	//
	//import "fmt"
	//
	//type Skills []string
	//
	//type Human struct {
	//    name   string
	//    age    int
	//    weight int
	//}
	//
	//type Student struct {
	//    Human  // 匿名字段，struct
	//    Skills // 匿名字段，自定义的类型string slice
	//    int    // 内置类型作为匿名字段
	//    speciality string
	//}
	//
	//func main() {
	//    // 初始化学生Jane
	//    jane := Student{ Human : Human{"Jane", 35, 100}, speciality : "Biology"}
	//    // 现在我们来访问相应的字段
	//    fmt.Println("Her name is ", jane.name)
	//    fmt.Println("Her age is ", jane.age)
	//    fmt.Println("Her weight is ", jane.weight)
	//    fmt.Println("Her speciality is ", jane.speciality)
	//    // 我们来修改他的skill技能字段
	//    jane.Skills = []string{"anatomy"}
	//    fmt.Println("Her skills are ", jane.Skills)
	//    fmt.Println("She acquired two new ones ")
	//    jane.Skills = append(jane.Skills, "physics", "golang")
	//    fmt.Println("Her skills now are ", jane.Skills)
	//    // 修改匿名内置类型字段
	//    jane.int = 3
	//    fmt.Println("Her preferred number is", jane.int)
	//}
	// output:
	//Her name is  Jane
	//Her age is  35
	//Her weight is  100
	//Her speciality is  Biology
	//Her skills are  [anatomy]
	//She acquired two new ones
	//Her skills now are  [anatomy physics golang]
	//Her preferred number is 3
	//struct不仅仅能够将struct作为匿名字段、自定义类型、内置类型都可以作为匿名字段，
	//而且可以在相应的字段上面进行函数操作（如例子中的append）。

	//这里有一个问题：如果human里面有一个字段叫做phone，而student也有一个字段叫做phone，那么该怎么办呢？
	//
	//Go里面很简单的解决了这个问题，最外层的优先访问，也就是当你通过student.phone访问的时候，
	//是访问student里面的字段，而不是human里面的字段。
	//
	//这样就允许我们去重载通过匿名字段继承的一些字段，
	//当然如果我们想访问重载后对应匿名类型里面的字段， 可以通过匿名字段名来访问。
	//package main
	//import "fmt"
	//
	//type Human struct {
	//    name  string
	//    age   int
	//    phone string  // Human类型拥有的字段
	//}
	//
	//type Employee struct {
	//    Human              // 匿名字段Human
	//    speciality string
	//    phone      string  // 雇员的phone字段
	//}
	//
	//func main() {
	//Bob := Employee{Human{"Bob", 34, 75,"777-444-XXXX"}, "Designer", "333-222"}
	//fmt.Println("Bob's work phone is:", Bob.phone)
	////Bob's work phone is: 333-222
	////如果我们要访问Human的phone字段
	//fmt.Println("Bob's personal phone is:", Bob.Human.phone)
	//Bob's personal phone is: 777-444-XXXX
	//}

	//匿名结构体
	//上面用到的结构体都是先定义后使用的。但是，还可以直接定义并使用结构体，在某些时候还是比较有用的。
	//方法一
	//var user struct{Username, Password string}
	//user.Username = "jack"
	//user.Password = "admin"

	//方法二
	//user := struct{Username, Password string}{"test1", "test2"}
	//user := &struct{Username, Password string}{"test1", "test2"}

	//方法三
	//user := new(struct{Username, Password string})
	//user.Username = "test1"
	//user.Password = "test2"

}
