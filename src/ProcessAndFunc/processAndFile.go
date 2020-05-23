package processfile

import "fmt"

func process() {
	//流程控制
	//if Go里面if条件判断语句中不需要括号
	if s > 10 {
		fmt.Println("x is greater than 10")
	} else {
		fmt.Println("x is less than 10")
	}
	//Go的if还有一个强大的地方就是条件判断语句里面允许声明一个变量，
	// 这个变量的作用域只能在该条件逻辑块内，其他地方就不起作用了

	// 计算获取值x,然后根据x返回的大小，判断是否大于10。
	if x := 20; x > 10 {
		fmt.Println("x is greater than 10")
	} else {
		fmt.Println("x is less than 10")
	}
	//这个地方如果这样调用就编译出错了，因为x是条件里面的变量
	//fmt.Println(x)

	//多条件
	if integer == 3 {
		fmt.Println("The integer is equal to 3")
	} else if integer < 3 {
		fmt.Println("The integer is less than 3")
	} else {
		fmt.Println("The integer is greater than 3")
	}

	//goto
	//用goto跳转到必须在当前函数内定义的标签。
	//func myFunc() {
	//	i := 0
	//Here:   //这行的第一个词，以冒号结束作为标签
	//	println(i)
	//	i++
	//	goto Here   //跳转到Here去
	//}
	//标签名是大小写敏感的。

	//for
	// 既可以用来循环读取数据，又可以当作while来控制逻辑，还能迭代操作。
	// for expression1; expression2; expression3 {
	//	...
	//}

	//expression1、expression2和expression3都是表达式，
	// 其中expression1和expression3是变量声明或者函数调用返回值之类的，
	// expression2是用来条件判断，expression1在循环开始之前调用，expression3在每轮循环结束之时调用。

	sum := 0
	for index := 0; index < 10; index++ {
		sum += index
	}
	fmt.Println("sum is equal to ", sum)
	// 输出：sum is equal to 45

	//有些时候需要进行多个赋值操作，由于Go里面没有,操作符，那么可以使用平行赋值i, j = i+1, j-1
	//可以省略 expression1 和 expression3 和 ; 就变成while 功能
	sum1 := 1
	for sum1 < 1000 {
		sum1 += sum1
	}

	//在循环里面有两个关键操作break和continue ,
	// break操作是跳出当前循环，continue是跳过本次循环。
	// 当嵌套过深的时候，break可以配合标签使用，即跳转至标签所指定的位置，详细参考如下例子：
	for index := 10; index>0; index-- {
		if index == 5{
			break // 或者continue
		}
		fmt.Println(index)
	}
	// break打印出来10、9、8、7、6
	// continue打印出来10、9、8、7、6、4、3、2、1

	//break和continue还可以跟着标号，用来跳到多重循环中的外层循环
	//for配合range可以用于读取slice和map的数据：

	//待解决问题
	//-----------------
	for k,v:=range map {
		fmt.Println("map's key:", k)
		fmt.Println("map's val:", v)
	}
	//-----------------



}
