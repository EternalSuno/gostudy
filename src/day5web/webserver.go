package day5web

func main() {
	//http包 建立web服务器
	//package main
	//
	//import (
	//	"fmt"
	//"net/http"
	//"strings"
	//"log"
	//)
	//
	//func sayhelloName(w http.ResponseWriter, r *http.Request) {
	//	r.ParseForm()  // 解析参数，默认是不会解析的
	//	fmt.Println(r.Form)  // 这些信息是输出到服务器端的打印信息
	//	fmt.Println("path", r.URL.Path)
	//	fmt.Println("scheme", r.URL.Scheme)
	//	fmt.Println(r.Form["url_long"])
	//	for k, v := range r.Form {
	//		fmt.Println("key:", k)
	//		fmt.Println("val:", strings.Join(v, ""))
	//	}
	//	fmt.Fprintf(w, "Hello astaxie!") // 这个写入到 w 的是输出到客户端的
	//}
	//
	//func main() {
	//	http.HandleFunc("/", sayhelloName) // 设置访问的路由
	//	err := http.ListenAndServe(":9090", nil) // 设置监听的端口
	//	if err != nil {
	//		log.Fatal("ListenAndServe: ", err)
	//	}
	//}

	//在浏览器输入 http://localhost:9090
	//
	//可以看到浏览器页面输出了 Hello astaxie!
	//
	//可以换一个地址试试：http://localhost:9090/?url_long=111&url_long=222
	//页面输出 Hello astaxie!
	//
	//服务器输出
	//map[url_long:[111 222]]
	//path /
	//scheme
	//[111 222]
	//key: url_long
	//val: 111222
	//map[]
	//path /favicon.ico
	//scheme
	//[]
	//
	//nginx、apache 服务器不需要吗？Go 就是不需要这些，因为他直接就监听 tcp 端口了，做了 nginx 做的事情，
	//然后 sayhelloName 这个其实就是我们写的逻辑函数了，跟 php 里面的控制层（controller）函数类似。
	//
	//
	//
	//
	//
}
