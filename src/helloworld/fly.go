package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {

	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Print("Hello World")
	fmt.Println(math.Pi)
	fmt.Println(add(1, 2))
	fmt.Println(add2(3, 4))
	a, b := swap("Hello", "World")
	fmt.Println(a, b)
	fmt.Println(split(16))
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
