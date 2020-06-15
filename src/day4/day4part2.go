package day4

func concurrent() {
	//goroutine
	//goroutine 是 Go 并行设计的核心。
	//goroutine 说到底其实就是协程，但是它比线程更小，十几个 goroutine 可能体现在底层就是五六个线程，
	//Go 语言内部帮你实现了这些 goroutine 之间的内存共享。
	//执行 goroutine 只需极少的栈内存 (大概是 4~5 KB)，当然会根据相应的数据伸缩。
	//也正因为如此，可同时运行成千上万个并发任务。goroutine 比 thread 更易用、更高效、更轻便。
	//goroutine 是通过 Go 的 runtime 管理的一个线程管理器。goroutine 通过 go 关键字实现了，其实就是一个普通的函数。
	//go hello(a, b, c)
	//通过关键字 go 就启动了一个 goroutine。
	//package main
	//
	//import (
	//    "fmt"
	//    "runtime"
	//)
	//
	//func say(s string) {
	//    for i := 0; i < 5; i++ {
	//        runtime.Gosched()
	//        fmt.Println(s)
	//    }
	//}
	//
	//func main() {
	//    go say("world") // 开一个新的 Goroutines 执行
	//    say("hello") // 当前 Goroutines 执行
	//}
	//
	//// 以上程序执行后将输出：
	//// hello
	//// world
	//// hello
	//// world
	//// hello
	//// world
	//// hello
	//// world
	//// hello
	//// world

	//我们可以看到 go 关键字很方便的就实现了并发编程。
	//上面的多个 goroutine 运行在同一个进程里面，共享内存数据，不过设计上我们要遵循：
	//不要通过共享来通信，而要通过通信来共享。
	//runtime.Gosched () 表示让 CPU 把时间片让给别人，下次某个时候继续恢复执行该 goroutine。
	//默认情况下，在 Go 1.5 将标识并发系统线程个数的 runtime.GOMAXPROCS 的初始值由 1 改为了运行环境的 CPU 核数。
	//但在 Go 1.5 以前调度器仅使用单线程，也就是说只实现了并发。想要发挥多核处理器的并行，
	//需要在我们的程序中显式调用 runtime.GOMAXPROCS (n) 告诉调度器同时使用多个线程。
	//GOMAXPROCS 设置了同时运行逻辑代码的系统线程的最大数量，并返回之前的设置。如果 n < 1，不会改变当前设置。
	//channels
	//goroutine 运行在相同的地址空间，因此访问共享内存必须做好同步。
	//那么 goroutine 之间如何进行数据的通信呢，Go 提供了一个很好的通信机制 channel。
	//channel 可以与 Unix shell 中的双向管道做类比：
	//可以通过它发送或者接收值。这些值只能是特定的类型： channel 类型。
	//定义一个 channel 时，也需要定义发送到 channel 的值的类型。
	//注意，必须使用 make 创建 channel：
	//

	//ci := make(chan int)
	//cs := make(chan string)
	//cf := make(chan interface{})
	// ch < -v // 发送 v 到channel ch.
	// v:= <-ch //从ch中接收数据, 并赋值给v

	//example
	//
	//package main
	//
	//import "fmt"
	//
	//func sum(a []int, c chan int) {
	//	total := 0
	//	for _, v := range a {
	//		total += v
	//	}
	//	c <- total  // send total to c
	//}
	//
	//func main() {
	//	a := []int{7, 2, 8, -9, 4, 0}
	//
	//	c := make(chan int)
	//	go sum(a[:len(a)/2], c)
	//	go sum(a[len(a)/2:], c)
	//	x, y := <-c, <-c  // receive from c
	//
	//	fmt.Println(x, y, x + y)
	//}
	//

	//补 : 冒号用法
	//1. 声明
	// a:= 1
	//2. 取值
	// s := []int{7, 2, 8, -9, 4, 0}
	//fmt.Println(s[3:5])     //  表示取第3 至 4位，不包括5位
	//输出 [-9 4]
	//如果冒号前不写数字，则默认从0开始
	//s := []int{7, 2, 8, -9, 4, 0}
	//fmt.Println(s[:5])      //  表示取0至4位
	// 输出 [7 2 8 -9 4]
	//如果冒号后不写数字，则默认一直到结束
	//s := []int{7, 2, 8, -9, 4, 0}
	//fmt.Println(s[3:])      //  表示取3到最后一位
	// 输出[-9 4 0]

	//Buffered Channels
	//上面我们介绍了默认的非缓存类型的 channel，不过 Go 也允许指定 channel 的缓冲大小，很简单，就是 channel 可以存储多少元素。
	//ch:= make (chan bool, 4)，创建了可以存储 4 个元素的 bool 型 channel。
	//在这个 channel 中，前 4 个元素可以无阻塞的写入。
	//当写入第 5 个元素时，代码将会阻塞，直到其他 goroutine 从 channel 中读取一些元素，腾出空间。
	//ch := make(chan type, value)
	//当 value = 0 时，channel 是无缓冲阻塞读写的，
	//当 value > 0 时，channel 有缓冲、是非阻塞的，直到写满 value 个元素才阻塞写入。

	/*	package main

		import "fmt"

		func main() {
			c := make(chan int, 2) // 修改 2 为 1 就报错，修改 2 为 3 可以正常运行
			c <- 1
			c <- 2
			fmt.Println(<-c)
			fmt.Println(<-c)
		}*/
	// 修改为 1 报如下的错误:
	// fatal error: all goroutines are asleep - deadlock!

	//Range 和 Close

}
