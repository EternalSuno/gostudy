package day3

func learn() {
	//oop
	//method 函数的另一种形态，带有接收者的函数

	//example:
	//package main
	//
	//import "fmt"
	//
	//type Rectangle struct {
	//    width, height float64
	//}
	//
	//func area(r Rectangle) float64 {
	//    return r.width*r.height
	//}
	//
	//func main() {
	//    r1 := Rectangle{12, 2}
	//    r2 := Rectangle{9, 4}
	//    fmt.Println("Area of r1 is: ", area(r1))
	//		Area of r1 is:  24
	//    fmt.Println("Area of r2 is: ", area(r2))
	//		Area of r2 is:  36
	//}
	//这段代码可以计算出来长方形的面积，但是area()不是作为Rectangle的方法实现的（类似面向对象里面的方法），
	//而是将Rectangle的对象（如r1,r2）作为参数传入函数计算面积的。

	//这样实现当然没有问题咯，但是当需要增加圆形、正方形、五边形甚至其它多边形的时候，
	//你想计算他们的面积的时候怎么办啊？ 那就只能增加新的函数咯，
	//但是函数名你就必须要跟着换了，变成area_rectangle, area_circle, area_triangle...

	//这些函数并不从属于struct(或者以面向对象的术语来说，并不属于class)，
	//他们是单独存在于struct外围，而非在概念上属于某个struct的。

	//很显然，这样的实现并不优雅，并且从概念上来说"面积"是"形状"的一个属性，它是属于这个特定的形状的，就像长方形的长和宽一样。
	//
	//基于上面的原因所以就有了method的概念，method是附属在一个给定的类型上的，
	//他的语法和函数的声明语法几乎一样， 只是在func后面增加了一个receiver(也就是method所依从的主体)。

	//用上面提到的形状的例子来说，method area() 是依赖于某个形状(比如说Rectangle)来发生作用的。
	//Rectangle.area()的发出者是Rectangle， area()是属于Rectangle的方法，而非一个外围函数。
	//更具体地说，Rectangle存在字段length 和 width, 同时存在方法area(), 这些字段和方法都属于Rectangle。

	//method的语法如下：
	//func (r ReceiverType) funcName(parameters) (results)
	//example:
	//package main
	//import (
	//    "fmt"
	//    "math"
	//)
	//
	//type Rectangle struct {
	//    width, height float64
	//}
	//
	//type Circle struct {
	//    radius float64
	//}
	//
	//func (r Rectangle) area() float64 {
	//    return r.width*r.height
	//}
	//
	//func (c Circle) area() float64 {
	//    return c.radius * c.radius * math.Pi
	//}
	//
	//func main() {
	//    r1 := Rectangle{12, 2}
	//    r2 := Rectangle{9, 4}
	//    c1 := Circle{10}
	//    c2 := Circle{25}
	//
	//    fmt.Println("Area of r1 is: ", r1.area())
	//    fmt.Println("Area of r2 is: ", r2.area())
	//    fmt.Println("Area of c1 is: ", c1.area())
	//    fmt.Println("Area of c2 is: ", c2.area())
	//Area of r1 is:  24
	//Area of r2 is:  36
	//Area of c1 is:  314.1592653589793
	//Area of c2 is:  1963.4954084936207
	//}

	//在使用method的时候重要注意几点
	//
	//虽然method的名字一模一样，但是如果接收者不一样，那么method就不一样
	//method里面可以访问接收者的字段
	//调用method通过.访问，就像struct里面访问字段一样

	//method area() 分别属于Rectangle和Circle， 于是他们的 Receiver 就变成了Rectangle 和 Circle,
	//或者说，这个area()方法 是由 Rectangle/Circle 发出的。

	//那是不是method只能作用在struct上面呢？当然不是咯，他可以定义在任何你自定义的类型、内置类型、struct等各种类型上面。
	//这里你是不是有点迷糊了，什么叫自定义类型，自定义类型不就是struct嘛，不是这样的哦，
	//struct只是自定义类型里面一种比较特殊的类型而已，
	//还有其他自定义类型申明，可以通过如下这样的申明来实现。
	//type typeName typeLiteral

	//请看下面这个申明自定义类型的代码
	//
	//type ages int
	//type money float32
	//type months map[string]int
	//m := months {
	//    "January":31,
	//    "February":28,
	//    ...
	//    "December":31,
	//}

	//package main
	//import "fmt"
	//
	//const (
	//    WHITE = iota
	//    BLACK
	//    BLUE
	//    RED
	//    YELLOW
	//)
	//
	//type Color byte
	//
	//type Box struct {
	//    width, height, depth float64
	//    color Color
	//}
	//
	//type BoxList []Box //a slice of boxes
	//
	//func (b Box) Volume() float64 {
	//    return b.width * b.height * b.depth
	//}
	//
	//func (b *Box) SetColor(c Color) {
	//    b.color = c
	//}
	//
	//func (bl BoxList) BiggestColor() Color {
	//    v := 0.00
	//    k := Color(WHITE)
	//    for _, b := range bl {
	//        if bv := b.Volume(); bv > v {
	//            v = bv
	//            k = b.color
	//        }
	//    }
	//    return k
	//}
	//
	//func (bl BoxList) PaintItBlack() {
	//    for i, _ := range bl {
	//        bl[i].SetColor(BLACK)
	//    }
	//}
	//
	//func (c Color) String() string {
	//    strings := []string {"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
	//    return strings[c]
	//}
	//
	//func main() {
	//    boxes := BoxList {
	//        Box{4, 4, 4, RED},
	//        Box{10, 10, 1, YELLOW},
	//        Box{1, 1, 20, BLACK},
	//        Box{10, 10, 1, BLUE},
	//        Box{10, 30, 1, WHITE},
	//        Box{20, 20, 20, YELLOW},
	//    }
	//
	//    fmt.Printf("We have %d boxes in our set\n", len(boxes))
	//		We have 6 boxes in our set
	//    fmt.Println("The volume of the first one is", boxes[0].Volume(), "cm³")
	//		The volume of the first one is 64 cm³
	//    fmt.Println("The color of the last one is",boxes[len(boxes) - 1].color.String())
	//		The color of the last one is YELLOW
	//    fmt.Println("The biggest one is", boxes.BiggestColor().String())
	//		The biggest one is YELLOW
	//    fmt.Println("Let's paint them all black")
	//		Let's paint them all black
	//    boxes.PaintItBlack()
	//    fmt.Println("The color of the second one is", boxes[1].color.String())
	//		The color of the second one is BLACK
	//    fmt.Println("Obviously, now, the biggest one is", boxes.BiggestColor().String())
	//		Obviously, now, the biggest one is BLACK
	//}

	//上面代码通过const定义了一些常量，然后定义了一些自定义类型
	//Color作为byte的别名
	//定义了一个struct:Box，含有三个长宽高字段和一个颜色属性
	//定义了一个slice:BoxList，含有Box

	//以上面的自定义类型为接收者定义了一些method
	//
	//Volume()定义了接收者为Box，返回Box的容量
	//SetColor(c Color)，把Box的颜色改为c
	//BiggestColor()定在在BoxList上面，返回list里面容量最大的颜色
	//PaintItBlack()把BoxList里面所有Box的颜色全部变成黑色
	//String()定义在Color上面，返回Color的具体颜色(字符串格式)

	//指针作为 receiver

	//现在让我们回过头来看看 SetColor 这个 method，它的 receiver 是一个指向 Box 的指针，是的，
	//你可以使用 *Box。想想为啥要使用指针而不是 Box 本身呢？
	//我们定义 SetColor 的真正目的是想改变这个 Box 的颜色，如果不传 Box 的指针，
	//那么 SetColor 接受的其实是 Box 的一个 copy，也就是说 method 内对于颜色值的修改，
	//其实只作用于 Box 的 copy，而不是真正的 Box。所以我们需要传入指针。
	//这里可以把 receiver 当作 method 的第一个参数来看，然后结合前面函数讲解的传值和传引用就不难理解
	//这里你也许会问了那 SetColor 函数里面应该这样定义 *b.Color=c, 而不是 b.Color=c, 因为我们需要读取到指针相应的值。
	//其实 Go 里面这两种方式都是正确的，当你用指针去访问相应的字段时 (虽然指针没有任何的字段)，Go 知道你要通过指针去获取这个值

	//如果一个 method 的 receiver 是 *T, 你可以在一个 T 类型的实例变量 V 上面调用这个 method，而不需要 &V 去调用这个 method
	//如果一个 method 的 receiver 是 T，你可以在一个 T 类型的变量 P 上面调用这个 method，而不需要 P 去调用这个 method

	//method 继承
	//如果匿名字段实现了一个 method，那么包含这个匿名字段的 struct 也能调用该 method。

}
