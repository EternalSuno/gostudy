package day4

import (
	"fmt"
)

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human  //匿名字段 Human
	school string
	loan   float32
}

type Employee struct {
	Human   //匿名字段 Human
	company string
	money   float32
}

// Human 对象实现 Sayhi 方法
func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

// Human 对象实现 Sing 方法
func (h *Human) Sing(Iyrics string) {
	fmt.Println("La la, lalalal...", Iyrics)
}

// Human 对象实现 Guzzle 方法
func (h *Human) Guzzle(beerStein string) {
	fmt.Println("Guzzle Guzzle ...", beerStein)
}

//Employee 重载 Human 的 Sayhi 方法
func (e *Employee) Sayhi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone)
}

//Student 实现 BorrowMoney 方法
func (s *Student) BorrowMoney(amount float32) {
	s.loan += amount
}

func (e *Employee) SpendSalary(amount float32) {
	e.money -= amount
}

type Men interface {
	SayHi()
	Sing(Iyrics string)
	Guzzle(beerStein string)
}

type YoungChap interface {
	SayHi()
	Sing(song string)
	BorrowMoney(amount float32)
}

type ElderlyGent interface {
	SayHi()
	Sing(song string)
	SpendSalary(amount float32)
}

func interfacelearn() {
	//	interface 是一组 method 签名的组合，我们通过 interface 来定义对象的一组行为。
	//	interface 类型
	//	interface 类型定义了一组方法，如果某个对象实现了某个接口的所有方法，则此对象就实现了此接口。详细的语法参考下面这个例子
	//	 interface 就是一组抽象方法的集合，它必须由其他非 interface 类型实现，而不能自我实现， Go 通过 interface 实现了 duck-typing: 即 "当看到一只鸟走起来像鸭子、游泳起来像鸭子、叫起来也像鸭子，那么这只鸟就可以被称为鸭子"。
	//

	//空interface
	//空 interface (interface {}) 不包含任何的 method，正因为如此，所有的类型都实现了空 interface。
	//空 interface 对于描述起不到任何的作用 (因为它不包含任何的 method），
	//但是空 interface 在我们需要存储任意类型的数值的时候相当有用，因为它可以存储任意类型的数值。它有点类似于 C 语言的 void* 类型。
	//

	//定义为a为空接口
	var a interface{}
	var i int = 5
	s := "Hello world"
	//a 可以储存任意类型的数值
	a = i
	fmt.Println(a) // 5
	a = s
	fmt.Println(a) // Hello world

	//interface 函数参数
	//interface 的变量可以持有任意实现该 interface 类型的对象
	//fmt.Println 是我们常用的一个函数，但是你是否注意到它可以接受任意类型的数据。打开 fmt 的源码文件，你会看到这样一个定义:
	//type Stringer interface {
	//	String() string
	//}
	//也就是说，任何实现了 String 方法的类型都能作为参数被 fmt.Println 调用
	//
	//package main
	//import (
	//	"fmt"
	//"strconv"
	//)
	//
	//type Human struct {
	//	name string
	//	age int
	//	phone string
	//}
	//
	//// 通过这个方法 Human 实现了 fmt.Stringer
	//func (h Human) String() string {
	//	return "❰"+h.name+" - "+strconv.Itoa(h.age)+" years -  ✆ " +h.phone+"❱"
	//}
	//
	//func main() {
	//	Bob := Human{"Bob", 39, "000-7777-XXX"}
	//	fmt.Println("This Human is : ", Bob)
	//}

	//interface 变量存储的类型
	//我们知道 interface 的变量里面可以存储任意类型的数值 (该类型实现了 interface)。
	//那么我们怎么反向知道这个变量里面实际保存了的是哪个类型的对象呢？目前常用的有两种方法：
	//Comma-ok 断言
	//		Go 语言里面有一个语法，可以直接判断是否是该类型的变量： value, ok = element.(T)，
	//		这里 value 就是变量的值，ok 是一个 bool 类型，element 是 interface 变量，T 是断言的类型。
	//	package main
	//
	//	import (
	//		"fmt"
	//	"strconv"
	//	)
	//
	//	type Element interface{}
	//	type List [] Element
	//
	//	type Person struct {
	//		name string
	//		age int
	//	}
	//
	//	// 定义了 String 方法，实现了 fmt.Stringer
	//	func (p Person) String() string {
	//		return "(name: " + p.name + " - age: "+strconv.Itoa(p.age)+ " years)"
	//	}
	//
	//	func main() {
	//		list := make(List, 3)
	//		list[0] = 1 // an int
	//		list[1] = "Hello" // a string
	//		list[2] = Person{"Dennis", 70}
	//
	//		for index, element := range list {
	//			if value, ok := element.(int); ok {
	//				fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
	//			} else if value, ok := element.(string); ok {
	//				fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
	//			} else if value, ok := element.(Person); ok {
	//				fmt.Printf("list[%d] is a Person and its value is %s\n", index, value)
	//			} else {
	//				fmt.Printf("list[%d] is of a different type\n", index)
	//			}
	//		}
	//	}

	//switch 测试
	//
	//package main
	//
	//import (
	//	"fmt"
	//"strconv"
	//)
	//
	//type Element interface{}
	//type List [] Element
	//
	//type Person struct {
	//	name string
	//	age int
	//}
	//
	//// 打印
	//func (p Person) String() string {
	//	return "(name: " + p.name + " - age: "+strconv.Itoa(p.age)+ " years)"
	//}
	//
	//func main() {
	//	list := make(List, 3)
	//	list[0] = 1 // an int
	//	list[1] = "Hello" // a string
	//	list[2] = Person{"Dennis", 70}
	//
	//	for index, element := range list{
	//		switch value,ok := element.(type) {
	//		case int:
	//			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
	//		case string:
	//			fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
	//		case Person:
	//			fmt.Printf("list[%d] is a Person and its value is %s\n", index, value)
	//		default:
	//			fmt.Printf("list[%d] is of a different type", index)
	//		}
	//	}
	//}

	// 嵌入interface
	// 如果一个 interface1 作为 interface2 的一个嵌入字段，那么 interface2 隐式的包含了 interface1 里面的 method。
	// 源码包 container/heap 里面有这样的一个定义

	//type Interface interface {
	//	sort.Interface // 嵌入字段 sort.Interface
	//	Push(x interface{}) // a Push method to push elements into the heap
	//	Pop() interface{} // a Pop elements that pops elements from the heap
	//}

	//
	//type Interface interface {
	//	// Len is the number of elements in the collection.
	//	Len() int
	//	// Less returns whether the element with index i should sort
	//	// before the element with index j.
	//	Less(i, j int) bool
	//	// Swap swaps the elements with indexes i and j.
	//	Swap(i, j int)
	//}

	//io 包下面的 io.ReadWriter ，它包含了 io 包下面的 Reader 和 Writer 两个 interface：
	//type ReadWriter interface {
	//	Reader
	//	Writer
	//}

	//反射
	// Go 语言实现了反射，所谓反射就是能检查程序在运行时的状态。我们一般用到的包是 reflect 包。如何运用 reflect 包
	//使用 reflect 一般分成三步，下面简要的讲解一下：要去反射是一个类型的值 (这些值都实现了空 interface)，
	//首先需要把它转化成 reflect 对象 (reflect.Type 或者 reflect.Value，根据不同的情况调用不同的函数)。这两种获取方式如下：
	//t := reflect.TypeOf(i)    // 得到类型的元数据,通过t我们能获取类型定义里面的所有元素
	//v := reflect.ValueOf(i)   // 得到实际的值，通过 v 我们获取存储在里面的值，还可以去改变值
	//转化为 reflect 对象之后我们就可以进行一些操作了，也就是将 reflect 对象转化成相应的值，例如
	//	name := t.Elem().Field(0).Name  // 获取定义在 struct 里面第一个字段的字段名
	//	value := v.Elem().Field(0).String()  // 获取存储在第一个字段里面的值
	// 获取反射值能返回相应的类型和数值
	//	var x float64 = 3.4
	//	v := reflect.ValueOf(x)
	//	fmt.Println("type:", v.Type())
	//	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	//	fmt.Println("value:", v.Float())
	//最后，反射的话，那么反射的字段必须是可修改的，我们前面学习过传值和传引用，这个里面也是一样的道理。
	//反射的字段必须是可读写的意思是，如果下面这样写，那么会发生错误

	//var x float64 = 3.4
	//v := reflect.ValueOf(x)
	//v.SetFloat(7.1)
	//如果要修改相应的值，必须这样写
	//var x float64 = 3.4
	//p := reflect.ValueOf(&x)
	//v := p.Elem()
	//v.SetFloat(7.1)

}
