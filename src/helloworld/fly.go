package main

//分组声明
//import "fmt"
//import "math"
import (
	"fmt"
	"math"
	"net/http"
	"runtime"
	"strconv"
)

type testInt func(int) bool //声明一个函数类型

func isOdd(integer int) bool {
	if integer%2 == 0 {
		return false
	}
	return true
}

func isEven(integer int) bool {
	if integer%2 == 0 {
		return true
	}
	return false
}

func filter(slice []int, f testInt) []int {
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

//声明一个新的类型
type person struct {
	name string
	age  int
}

//比较两个人的年龄, 返回年龄大的那个人, 并且返回年龄差
//struct也是传值的
func Older(p1, p2 person) (person, int) {
	if p1.age > p2.age { //比较p1和p2这两个人的年龄
		return p1, p1.age - p2.age
	}
	return p2, p2.age - p1.age
}

type Skills []string

//type Human struct {
//	name   string
//	age    int
//	weight int
//	phone  string
//}
//
//type Student struct {
//	Human      //匿名字段, struct
//	Skills     //匿名字段, 自定义的类型string slice
//	int        //内置类型作为匿名字段
//	speciality string
//}
//
//type Employee struct {
//	Human      //匿名字段Human
//	speciality string
//	phone      string //雇员的phone字段
//}

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func area(r Rectangle) float64 {
	return r.width * r.height
}

const (
	WHITE  = iota //0
	BLACK         //1
	BLUE          //2
	RED           //3
	YELLOW        //4
)

type Color byte

type Box struct {
	width, height, depth float64
	color                Color
}

type BoxList []Box // a slice of boxes

func (b Box) Volume() float64 {
	return b.width * b.height * b.depth
}

func (b *Box) SetColor(c Color) {
	b.color = c
}

func (bl BoxList) BiggestColor() Color {
	v := 0.00
	k := Color(WHITE)
	for _, b := range bl {
		if bv := b.Volume(); bv > v {
			v = bv
			k = b.color
		}
	}
	return k
}

func (bl BoxList) PaintItBlack() {
	for i, _ := range bl {
		bl[i].SetColor(BLACK)
	}
}

func (c Color) String() string {
	strings := []string{"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
	return strings[c]
}

//type Human struct {
//	name  string
//	age   int
//	phone string
//}
//
//type Student struct {
//	Human  //匿名字段
//	school string
//}
//
//type Employee struct {
//	Human   //匿名字段
//	company string
//}
//
//// 在Human 上面定义了一个method
//func (h *Human) sayHi() {
//	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
//}
//
////Employee 的method 重写 Human 的method
//func (e *Employee) sayHi() {
//	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name, e.company, e.phone)
//}

type Human struct {
	name  string
	age   int
	phone string
}

func (h Human) String() string {
	return "<" + h.name + "-" + strconv.Itoa(h.age) + "years - phone:" + h.phone + ">"
}

type Element interface{}
type List []Element
type Person struct {
	name string
	age  int
}

//定义了 String 方法,实现了 fmt.Stringer
func (p Person) String() string {
	return "(name: " + p.name + " - age: " + strconv.Itoa(p.age) + " years)"
}

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total //send total to c
}
func fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

//func sayhelloName(w http.ResponseWriter, r *http.Request) {
//	r.ParseForm() //解析参数, 默认是不会解析
//	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
//	fmt.Println("path", r.URL.Path)
//	fmt.Println("scheme", r.URL.Scheme)
//	fmt.Println(r.Form["url_long"])
//	for k, v := range r.Form {
//		fmt.Println("key:", k)
//		fmt.Println("val:", strings.Join(v, ""))
//	}
//	fmt.Fprint(w, "Hello astaxie!") // 这个写入到 w 的是输出到客户端的
//
//}

type MyMux struct {
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayhelloName(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello myroute!")
}

func main() {

	mux := &MyMux{}
	http.ListenAndServe(":9090", mux)

	//http.HandleFunc("/", sayhelloName) // 设置访问的路由
	//err := http.ListenAndServe(":9090", nil) //设置监听的端口
	//if err != nil {
	//	log.Fatal("ListenAndServe:", err)
	//}
	//nginx、apache 服务器不需要吗？Go 就是不需要这些，因为他直接就监听 tcp 端口了，做了 nginx 做的事情，
	//然后 sayhelloName 这个其实就是我们写的逻辑函数了，跟 php 里面的控制层（controller）函数类似。
	//

	//goroutine超时
	//c := make(chan int)
	//o := make(chan bool)
	//go func() {
	//	for {
	//		select {
	//		case v := <-c:
	//			println(v)
	//		case <-time.After(5 * time.Second):
	//			println("timeout")
	//			o <- true
	//			break
	//		}
	//	}
	//}()
	//<-o

	//c := make(chan int, 10)
	//go fibonacci(cap(c), c)
	//for i := range c {
	//	fmt.Println(i)
	//}
	//输出 1	1 2	3 5 8 13 21 34 55

	//ch:= make (chan bool, 4)，创建了可以存储 4 个元素的 bool 型 channel。
	//在这个 channel 中，前 4 个元素可以无阻塞的写入。
	//当写入第 5 个元素时，代码将会阻塞，直到其他 goroutine 从 channel 中读取一些元素，腾出空间。
	//ch := make(chan type, value)
	//当 value = 0 时，channel 是无缓冲阻塞读写的，
	//当 value > 0 时，channel 有缓冲、是非阻塞的，直到写满 value 个元素才阻塞写入。
	//c := make(chan int, 1) //// 修改 2 为 1 就报错，修改 2 为 3 可以正常运行
	//c <- 1
	//c <- 2
	//fmt.Println(<-c)
	//fmt.Println(<-c)
	// 修改为 1 报如下的错误:
	// fatal error: all goroutines are asleep - deadlock!

	//a := []int{7, 2, 8, -9, 4, 0}
	//c := make(chan int)
	//go sum(a[:len(a)/2], c)
	//go sum(a[len(a)/2:], c)
	//x, y := <-c, <-c  //receive from c
	//
	//fmt.Println(x, y, x+y)
	//-5 17 12
	//默认情况下，channel 接收和发送数据都是阻塞的，除非另一端已经准备好，
	//这样就使得 Goroutines 同步变的更加的简单，而不需要显式的 lock。所谓阻塞，
	//也就是如果读取（value := <-ch）它将会被阻塞，直到有数据接收。
	//其次，任何发送（ch<-5）将会被阻塞，直到数据被读出。
	//无缓冲 channel 是在多个 goroutine 之间同步很棒的工具。

	//go say("world") // 开了一个新的Goroutines 执行
	//say("hello") // 当前 Goroutines 执行

	//list := make(List, 3)
	//list[0] = 1 // an int
	//list[1] = "Hello"
	//list[2] = Person{"Dennis", 70}

	//for index, element := range list {
	//	if value, ok := element.(int); ok {
	//		fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
	//	} else if value, ok := element.(string); ok {
	//		fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
	//	} else if value, ok := element.(Person); ok {
	//		fmt.Printf("list[%d] is a Person and its value is %s\n", index, value)
	//	} else {
	//		fmt.Printf("list[%d] is of a different type \n", index)
	//	}
	//}

	//for index, element := range list {
	//	switch value := element.(type) {
	//	case int:
	//		fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
	//	case string:
	//		fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
	//	case Person:
	//		fmt.Printf("list[%d] is a Person and its value is %s\n", index, value)
	//	default:
	//		fmt.Printf("list[%d] is of a different type \n", index)
	//	}
	//}

	//Bob := Human{"Bob", 39, "000-7777-xxx"}
	//fmt.Println("This Human is : ", Bob)
	//空interfce
	//定义为a为空接口
	//var a interface{}
	//var i int = 5
	//s := "Hello world"
	////a 可以储存任意类型的数值
	//a = i
	//fmt.Println(a)
	//a = s
	//fmt.Println(a)

	//mark := Student{Human{"Mark", 25, "222-222-YYYY"}, "MIT"}
	//sam := Employee{Human{"Sam", 45, "111-888-XXXX"}, "Golang Inc"}
	//
	//mark.sayHi()
	//sam.sayHi()

	//user := new(struct{ Username, Password string })
	//user.Username = "test1"
	//user.Password = "test2"
	//fmt.Println(user.Username)
	//fmt.Println(user.Password)

	//mark := Student{Human{"Mark", 25, 120}, "Computer Science"}
	//
	//fmt.Println("His name is ", mark.name)
	//fmt.Println("His age is ", mark.age)
	//fmt.Println("His weight is ", mark.weight)
	//fmt.Println("His speciality is ", mark.speciality)
	//
	//mark.speciality = "AI"
	//fmt.Println("Mark changed his speciality")
	//fmt.Println("His speciality is ", mark.speciality)
	//
	//fmt.Println("Mark become old")
	//mark.age = 46
	//fmt.Println("His age is", mark.age)
	//
	//fmt.Println("Mark is not an athlet anymore")
	//
	//mark.weight += 60
	//fmt.Println("His weight is", mark.weight)

	//var tom person
	////赋值初始化
	//tom.name, tom.age = "Tom", 18
	////两个字段都写清楚的初始化
	//bob := person{age: 25, name:"Bob"}
	////按照struct定义顺序初始化
	//paul := person{"Paul", 43}
	//
	//tb_Older, tb_diff := Older(tom, bob)
	//tp_Older, tp_diff := Older(tom, paul)
	//bp_Older, bp_diff := Older(bob, paul)
	//
	//fmt.Printf("Of %s and %s, %s is older by %d years\n", tom.name, bob.name, tb_Older.name, tb_diff)
	//fmt.Printf("Of %s and %s, %s is older by %d years\n", tom.name, paul.name, tp_Older.name, tp_diff)
	//fmt.Printf("Of %s and %s, %s is older by %d years\n", bob.name, paul.name, bp_Older.name, bp_diff)
	//

	//slice := []int{1, 2, 3, 4, 5, 7}
	//fmt.Println("slice = ", slice)
	//odd := filter(slice, isOdd) //函数当做值来传递
	//fmt.Println("Odd elements of slice are: ", odd)
	//even := filter(slice, isEven) //函数当做值来传递
	//fmt.Println("Even elements of slice are: ", even)
	//var rating = map[string]float32{}
	//fmt.Println(rating)
	//var m = make(map[string]string)
	//fmt.Println(m)
	//fmt.Println("My favorite number is", rand.Intn(10))
	//fmt.Print("Hello World")
	//fmt.Println(math.Pi)
	//fmt.Println(add(1, 2))
	//fmt.Println(add2(3, 4))
	//a, b := swap("Hello", "World")
	//fmt.Println(a, b)
	//fmt.Println(split(16))
}

func add3(a *int) int {
	*a = *a + 1
	return *a
}

func add1(a int) int {
	a = a + 1
	return a
}

func SumAndProduct(A, B int) (int, int) {
	return A + B, A * B
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func add(x int, y int) int {
	return x + y
}

//当连续两个或多个函数的已命名形参类型相同时，除最后一个类型以外，其它都可以省略。
func add2(x, y int) int {
	return x + y
}

//函数可以返回任意数量的返回值, swap 函数返回了两个字符串。
func swap(x, y string) (string, string) {
	return y, x
}

//Go 的返回值可被命名，它们会被视作定义在函数顶部的变量。
//返回值的名称应当具有一定的意义，它可以作为文档使用。
//没有参数的 return 语句返回已命名的返回值。也就是 直接 返回。
//直接返回语句应当仅用在下面这样的短函数中。在长的函数中它们会影响代码的可读性。
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var isActive bool                   //全局变量声明
var enabled, disabled = true, false //忽略类型声明
func base() {
	//Println 与Printf 都是fmt 包中的公共方法，在需要打印信息时需要用到这二个函数，那么这二个函数有什么区别呢？
	//Println :可以打印出字符串，和变量
	//Printf : 只可以打印出格式化的字符串,可以输出字符串类型的变量，不可以输出整形变量和整形

	//Boolean
	//在Go中，布尔值的类型为bool，值是true或false，默认为false。
	var available bool //一般声明
	valid := false     //简短声明
	available = true   //赋值操作
	fmt.Println(available)
	fmt.Println(valid)

	//数值类型
	//整数类型有无符号和带符号两种。Go同时支持int和uint，这两种类型的长度相同，但具体长度取决于不同编译器的实现。
	//Go里面也有直接定义好位数的类型：
	// rune, int8, int16, int32, int64和byte, uint8, uint16, uint32,uint64。
	// 其中rune是int32的别称，byte是uint8的别称。
	//需要注意的一点是，这些类型的变量之间不允许互相赋值或操作，不然会在编译时引起编译器报错。
	//如下的代码会产生错误：invalid operation: a + b (mismatched types int8 and int32)
	//var a int8
	//var b int32
	//c:=a + b
	//另外，尽管int的长度是32 bit, 但int 与 int32并不可以互用。

	//浮点数的类型有float32和float64两种（没有float类型），默认是float64。

	//复数 它的默认类型是complex128（64位实数+64位虚数）。complex64(32位实数+32位虚数)。

	var c complex64 = 5 + 5i
	//output: (5+5i)
	fmt.Printf("Value is: %v", c)

	//String
	//Go中的字符串都是采用UTF-8
	//字符串是用一对双引号（""）或反引号（）括起来定义，它的类型是string
	var frenchHello string //声明变量为字符串的一般方法
	var emptyString string = ""
	fmt.Println(frenchHello)
	fmt.Println(emptyString)

	//在Go中字符串是不可变的，例如下面的代码编译时会报错：cannot assign to s[0]
	//var s string = "hello"
	//s[0] = 'c'
	//
	// 实现方法
	//s := "hello"
	//c := []byte(s)  // 将字符串 s 转换为 []byte 类型
	//c[0] = 'c'
	//s2 := string(c)  // 再转换回 string 类型
	//fmt.Printf("%s\n", s2)

	//Go中可以使用+操作符来连接两个字符串
	s := "hello"
	m := " world"
	a := s + m
	fmt.Printf(a)

	d := "hello"
	d = "c" + d[1:] // 字符串虽不能更改，但可进行切片操作
	fmt.Printf("%s\n", d)

	//声明多行字符串
	n := `hello
    world`
	fmt.Printf(n)

}

func extraBase() {
	//const i = 100
	//const pi = 3.1415
	//const prefix = "Go_"
	//分组声明
	//import(
	//	"fmt"
	//   "os"
	//)

	const (
		ii     = 100
		pi     = 3.1415
		prefix = "Go_"
	)

	var (
		jj      int
		ff      float32
		prefixs string
	)

	fmt.Println(jj)
	fmt.Println(ff)
	fmt.Printf(prefixs)

	//iota 枚举
	//用来声明enum 默认开始值是0，const中每增加一行加1
	const (
		x = iota // x == 0
		y = iota // y == 1
		z = iota // z == 2
		w        // 常量声明省略值时，默认和之前一个值的字面相同。这里隐式地说w = iota，因此w == 3。其实上面y和z可同样不用"= iota"
	)

	const v = iota // 每遇到一个const关键字，iota就会重置，此时v == 0

	const (
		h, i, j = iota, iota, iota //h=0,i=0,j=0 iota在同一行值相同
	)

	const (
		a       = iota //a=0
		b       = "B"
		c       = iota             //c=2
		d, e, f = iota, iota, iota //d=3,e=3,f=3
		g       = iota             //g = 4
	)

	fmt.Println(a, b, c, d, e, f, g, h, i, j, x, y, z, w, v)
}

func test() {
	no, yes, maybe := "no", "yes", "maybe" //简短声明，同时声明多个变量
	japaneseHello := "Konichiwa"
	var frenchHello string  //声明变量为字符串的一般方法
	frenchHello = "Bonjour" //常规赋值
	fmt.Println(no)
	fmt.Println(yes)
	fmt.Println(maybe)
	fmt.Println(japaneseHello)
	fmt.Println(frenchHello)
}
func defineVar() {
	//Go对于已声明但未使用的变量会在编译阶段报错
	//定义一个名称为“variableName”，类型为"type"的变量 var variableName type
	var a int
	fmt.Println(a)
	//定义三个类型都是“type”的变量
	var v1, v2, v3 int
	fmt.Println(v1, v2, v3)
	/*
		    定义变量并初始化值
			初始化“variableName”的变量为“value”值，类型是“type”
	*/
	var variableName int = 1
	fmt.Println(variableName)
	/*
	   同时初始化多个变量
	   定义三个类型都是"type"的变量,并且分别初始化为相应的值
	   vname1为1，vname2为2，vname3为3
	*/
	var vname1, vname2, vname3 int = 1, 2, 3
	fmt.Println(vname1, vname2, vname3)
	/*
	   定义三个变量，它们分别初始化为相应的值
	   vname1为v1，vname2为v2，vname3为v3
	   然后Go会根据其相应值的类型来帮你初始化它们
	*/
	var vnamee1, vnamee2, vnamee3 = 2, "v2", "3"
	fmt.Println(vnamee1, vnamee2, vnamee3)
	/*
	   := 简短声明 只能用在函数内部 在函数外部使用则会无法编译通过
	   一般用var方式来定义全局变量
	   定义三个变量，它们分别初始化为相应的值
	   vname1为v1，vname2为v2，vname3为v3
	   编译器会根据初始化的值自动推导出相应的类型
	*/
	vnameee1, vnameee2, vnameee3 := 1, "abc", "3"
	fmt.Println(vnameee1, vnameee2, vnameee3)
	/*
		_（下划线）是个特殊的变量名，任何赋予它的值都会被丢弃。
	*/
	_, b := 34, 35
	fmt.Println(b)
}

func defineConst() {
	/*
		常量
		在Go程序中，常量可定义为数值、布尔值或字符串等类型。
	*/
	const constantName = 333
	const Pi float32 = 3.1415926
	const str = "string"
	//Go 常量和一般程序语言不同的是，可以指定相当多的小数位数(例如200位)，
	// 若指定給float32自动缩短为32bit，指定给float64自动缩短为64bit
}

func ArraySliceMap() {
	//数组 array
	//var arr [n]type //在[n]type中，n表示数组的长度，type表示存储元素的类型。
	var arr [10]int                                 // 声明了一个int类型的数组
	arr[0] = 42                                     // 数组下标是从0开始的
	arr[1] = 13                                     // 赋值操作
	fmt.Printf("The first element is %d\n", arr[0]) // 获取数据，返回42
	fmt.Printf("The last element is %d\n", arr[9])  //返回未赋值的最后一个元素，默认返回0

	//由于长度也是数组类型的一部分，因此[3]int与[4]int是不同的类型，数组也就不能改变长度。
	//数组之间的赋值是值的赋值，即当把一个数组作为参数传入函数的时候，传入的其实是该数组的副本，而不是它的指针。
	//如果要使用指针，那么就需要用到后面介绍的slice类型了。
	//数组可以使用另一种:=来声明
	a := [3]int{1, 2, 3} // 声明了一个长度为3的int数组
	fmt.Printf("%v\n", a)
	b := [10]int{1, 2, 3} // 声明了一个长度为10的int数组，其中前三个元素初始化为1、2、3，其它默认为0
	fmt.Printf("%v\n", b)
	c := [...]int{4, 5, 6} // 可以省略长度而采用`...`的方式，Go会自动根据元素个数来计算长度
	fmt.Printf("%v\n", c)
	//二位数组
	// 声明了一个二维数组，该数组以两个数组作为元素，其中每个数组中又有4个int类型的元素
	doubleArray := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}
	fmt.Printf("%v\n", doubleArray)
	// 上面的声明可以简化，直接忽略内部的类型
	easyArray := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	fmt.Printf("%v\n", easyArray)

	//slice
	//在很多应用场景中，数组并不能满足我们的需求。
	// 在初始定义数组时，我们并不知道需要多大的数组，因此我们就需要“动态数组”。
	// 在Go里面这种数据结构叫slice
	// slice并不是真正意义上的动态数组，而是一个引用类型。
	// slice总是指向一个底层array，slice的声明也可以像array一样，只是不需要长度。

	// 和声明array一样，只是少了长度
	//var fslice []int
	//slice := []byte {'a', 'b', 'c', 'd'}
	//slice可以从一个数组或一个已经存在的slice中再次声明。
	//slice通过array[i:j]来获取，其中i是数组的开始位置，j是结束位置，但不包含array[j]，它的长度是j-i。

	// 声明一个含有10个元素元素类型为byte的数组
	var ar = [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// 声明两个含有byte的slice
	var aa, bb []byte

	// a指向数组的第3个元素开始，并到第五个元素结束，
	aa = ar[2:5]
	//现在a含有的元素: ar[2]、ar[3]和ar[4]

	// b是数组ar的另一个slice
	bb = ar[3:5]
	fmt.Printf("%v\n", aa)
	fmt.Printf("%v\n", bb)
	//注意slice和数组在声明时的区别：声明数组时，
	//方括号内写明了数组的长度或使用...自动计算长度，而声明slice时，方括号内没有任何字符。

	//对于slice有几个有用的内置函数：
	//len 获取slice的长度
	//cap 获取slice的最大容量
	//append 向slice里面追加一个或者多个元素，然后返回一个和slice一样类型的slice
	//copy 函数copy从源slice的src中复制元素到目标dst，并且返回复制的元素的个数
	//注：append函数会改变slice所引用的数组的内容，从而影响到引用同一数组的其它slice。
	//但当slice中没有剩余空间（即(cap-len) == 0）时，此时将动态分配新的数组空间。
	// 返回的slice数组指针将指向这个空间，而原数组的内容将保持不变；其它引用此数组的slice则不受影响。

	//从Go1.2开始slice支持了三个参数的slice，之前我们一直采用这种方式在slice或者array基础上来获取一个slice
	//var array [10]int
	//slice := array[2:4]
	//这个例子里面slice的容量是8，新版本里面可以指定这个容量
	//slice = array[2:4:7]
	//上面这个的容量就是7-2，即5。这样这个产生的新的slice就没办法访问最后的三个元素。
	//如果slice是这样的形式array[:i:j]，即第一个参数为空，默认值就是0。

	//map
	// map也就是Python中字典的概念，它的格式为map[keyType]valueType
	//map的读取和设置也类似slice一样，通过key来操作，
	//只是slice的index只能是｀int｀类型，
	//而map多了很多类型，可以是int，可以是string及所有完全定义了==与!=操作的类型。

	/// 声明一个key是字符串，值为int的字典,这种方式的声明需要在使用之前使用make初始化
	//var numbers map[string]int
	//// 另一种map的声明方式
	var numbers = make(map[string]int)
	numbers["one"] = 1  //赋值
	numbers["ten"] = 10 //赋值
	numbers["three"] = 3
	fmt.Println("第三个数字是: ", numbers["three"]) // 读取数据
	//// 打印出来如:第三个数字是: 3

	//	这个map就像我们平常看到的表格一样，左边列是key，右边列是值
	//	使用map过程中需要注意的几点：
	//	map是无序的，每次打印出来的map都会不一样，它不能通过index获取，而必须通过key获取
	//	map的长度是不固定的，也就是和slice一样，也是一种引用类型
	//	内置的len函数同样适用于map，返回map拥有的key的数量
	//	map的值可以很方便的修改，通过numbers["one"]=11可以很容易的把key为one的字典值改为11
	//	map和其他基本型别不同，它不是thread-safe，在多个go-routine存取时，必须使用mutex lock机制
	//  map的初始化可以通过key:val的方式初始化值，同时map内置有判断是否存在key的方式

	//通过delete删除map的元素：
	rating := map[string]float32{"C": 5, "Go": 4.5, "Python": 4.5, "C++": 2}
	// map有两个返回值，第二个返回值，如果不存在key，那么ok为false，如果存在ok为true
	csharpRating, ok := rating["C#"]
	if ok {
		fmt.Println("C# is in the map and its rating is ", csharpRating)
	} else {
		fmt.Println("We have no rating associated with C# in the map")
	}
	delete(rating, "C") // 删除key为C的元素

	//map是一种引用类型，如果两个map同时指向一个底层，那么一个改变，另一个也相应的改变
	m := make(map[string]string)
	m["Hello"] = "Bonjour"
	m1 := m
	m1["Hello"] = "Salut"    // 现在m["hello"]的值已经是Salut了
	fmt.Println(m["Hello"])  // Salut
	fmt.Println(m1["Hello"]) // Salut

	// make用于内建类型（map、slice 和channel）的内存分配。
	// new用于各种类型的内存分配。
	// 内建函数new本质上说跟其它语言中的同名函数功能一样：new(T)分配了零值填充的T类型的内存空间，
	// 并且返回其地址，即一个*T类型的值。
	// 用Go的术语说，它返回了一个指针，指向新分配的类型T的零值。有一点非常重要：
	// new返回指针。

	// 内建函数make(T, args)与new(T)有着不同的功能，make只能创建slice、map和channel，
	// 并且返回一个有初始值(非零)的T类型，而不是*T。
	// 本质来讲，导致这三个类型有所不同的原因是指向数据结构的引用在使用前必须被初始化。
	// 例如，一个slice，是一个包含指向数据（内部array）的指针、长度和容量的三项描述符；
	// 在这些项目被初始化之前，slice为nil。对于slice、map和channel来说，make初始化了内部的数据结构，填充适当的值。
	// make返回初始化后的（非零）值。

	//零值
	//关于“零值”，所指并非是空值，而是一种“变量未填充前”的默认值，通常为0。 此处罗列 部分类型 的 “零值”
	//int     0
	//int8    0
	//int32   0
	//int64   0
	//uint    0x0
	//rune    0 //rune的实际类型是 int32
	//byte    0x0 // byte的实际类型是 uint8
	//float32 0 //长度为 4 byte
	//float64 0 //长度为 8 byte
	//bool    false
	//string  ""

}
