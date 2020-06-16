package day5web

func methodofwebserver() {
	// web工作方式的几个概念
	// Request：用户请求的信息，用来解析用户的请求信息，包括 post、get、cookie、url 等信息
	//
	// Response：服务器需要反馈给客户端的信息
	//
	// Conn：用户的每次请求链接
	//
	// Handler：处理请求和生成返回信息的处理逻辑
	//
	//分析http包运行机制
	// 图示: https://cdn.learnku.com/build-web-application-with-golang/images/3.3.http.png?raw=true
	//1.创建 Listen Socket, 监听指定的端口，等待客户端请求到来。
	//
	//2.Listen Socket 接受客户端的请求，得到 Client Socket, 接下来通过 Client Socket 与客户端通信。
	//
	//3.处理客户端的请求，首先从 Client Socket 读取 HTTP 请求的协议头，如果是 POST 方法，还可能要读取客户端提交的数据，
	//然后交给相应的 handler 处理请求，handler 处理完毕准备好客户端需要的数据，通过 Client Socket 写给客户端。
	//

	//下面代码来自 Go 的 http 包的源码，通过下面的代码我们可以看到整个的 http 处理过程：
	//
	//func (srv *Server) Serve(l net.Listener) error {
	//	defer l.Close()
	//	var tempDelay time.Duration // how long to sleep on accept failure
	//	for {
	//	rw, e := l.Accept()
	//	if e != nil {
	//	if ne, ok := e.(net.Error); ok && ne.Temporary() {
	//	if tempDelay == 0 {
	//	tempDelay = 5 * time.Millisecond
	//} else {
	//	tempDelay *= 2
	//}
	//	if max := 1 * time.Second; tempDelay > max {
	//	tempDelay = max
	//}
	//	log.Printf("http: Accept error: %v; retrying in %v", e, tempDelay)
	//	time.Sleep(tempDelay)
	//	continue
	//}
	//	return e
	//}
	//	tempDelay = 0
	//	c, err := srv.newConn(rw)
	//	if err != nil {
	//	continue
	//}
	//	go c.serve()
	//}
	//}
	//
	//上面代码执行监控端口之后，调用了 srv.Serve(net.Listener) 函数，这个函数就是处理接收客户端的请求信息。
	//这个函数里面起了一个 for{}，首先通过 Listener 接收请求，其次创建一个 Conn，最后单独开了一个 goroutine，
	//把这个请求的数据当做参数扔给这个 conn 去服务：go c.serve()。这个就是高并发体现了，
	//用户的每一次请求都是在一个新的 goroutine 去服务，相互不影响。
	//
	//那么如何具体分配到相应的函数来处理请求呢？conn 首先会解析 request:c.readRequest(),
	//然后获取相应的 handler:handler := c.server.Handler，也就是我们刚才在调用函数 ListenAndServe 时候的第二个参数，
	//我们前面例子传递的是 nil，也就是为空，那么默认获取 handler = DefaultServeMux, 那么这个变量用来做什么的呢？
	//对，这个变量就是一个路由器，它用来匹配 url 跳转到其相应的 handle 函数，那么这个我们有设置过吗？
	//有，我们调用的代码里面第一句不是调用了 http.HandleFunc("/", sayhelloName) 嘛。
	//这个作用就是注册了请求 / 的路由规则，当请求 uri 为 "/"，路由就会转到函数 sayhelloName，DefaultServeMux 会调用 ServeHTTP 方法，
	//这个方法内部其实就是调用 sayhelloName 本身，最后通过写入 response 的信息反馈到客户端。
	//

	//Go的http包详解
	//Go的http包有两个核心功能: Conn、 ServeMux

	//Conn 的goroutine
	// 与我们一般编写的 http 服务器不同，Go 为了实现高并发和高性能，使用了 goroutines 来处理 Conn 的读写事件，
	// 这样每个请求都能保持独立，相互不会阻塞，可以高效的响应网络事件。这是 Go 高效的保证。
	//
	// Go 在等待客户端请求里面是这样写的：
	//c, err := srv.newConn(rw)
	//if err != nil {
	//	continue
	//}
	//go c.serve()
	//这里我们可以看到客户端的每次请求都会创建一个 Conn，这个 Conn 里面保存了该次请求的信息，
	//然后再传递到对应的 handler，该 handler 中便可以读取到相应的 header 信息，这样保证了每个请求的独立性。
	//
	//ServeMux 的自定义
	//我们前面小节讲述 conn.server 的时候，其实内部是调用了 http 包默认的路由器，通过路由器把本次请求的信息传递到了后端的处理函数。那么这个路由器是怎么实现的呢？
	//
	//它的结构如下：
	//
	//
	//
	//type ServeMux struct {
	//	mu sync.RWMutex // 锁，由于请求涉及到并发处理，因此这里需要一个锁机制
	//	m map[string]muxEntry //路由规则, 一个string对应一个mux实体, 这里的string 就是注册的路由表达式
	//	hosts bool // 是否在任意的规则中带有 host 信息
	//}
	// muxEntry
	//type muxEntry struct {
	//	explicit bool //是否精确匹配
	//	h Handler //这个路由表达式对应哪个handler
	//	pattern string //匹配字符串
	//}
	//Handler
	//type Handler interface {
	//	ServeHTTP(ResponseWriter, *Request) //路由实现器
	//}
	// Handler 是一个接口，但是前一小节中的 sayhelloName 函数并没有实现 ServeHTTP 这个接口，为什么能添加呢？
	// 原来在 http 包里面还定义了一个类型 HandlerFunc, 我们定义的函数 sayhelloName 就是这个 HandlerFunc 调用之后的结果，
	// 这个类型默认就实现了 ServeHTTP 这个接口，即我们调用了 HandlerFunc (f), 强制类型转换 f 成为 HandlerFunc 类型，
	// 这样 f 就拥有了 ServeHTTP 方法。
	//
	//type HandlerFunc func(ResponseWriter, *Request)
	// ServeHTTP calls f(w, r)
	//func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	//	f(w, r)
	//}
	//ServeHTTP
	// func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request){
	//		if r.RequestURI == "*" {
	//  		w.Header().Set("Connection", "close")
	//			w.WriteHeader(StatusBadRequest)
	//			return
	// 		}
	//		h, _ := mux.Handler(r)
	//		h.ServeHTTP(w, r)
	//}
	//
	//如上所示路由器接收到请求之后，如果是 * 那么关闭链接，不然调用 mux.Handler(r)
	//返回对应设置路由的处理 Handler，然后执行 h.ServeHTTP(w, r)
	//

	//func (mux *ServeMux) Handler(r *Request) (h Handler, pattern string) {
	//	if r.Method != "CONNECT" {
	//		if p := cleanPath(r.URL.Path); p != r.URL.Path {
	//			_, pattern = mux.handler(r.Host, p)
	//			return RedirectHandler(p, StatusMovedPermanently), pattern
	//		}
	//	}
	//	return mux.handler(r.Host, r.URL.Path)
	//}
	//
	//func (mux *ServeMux) handler(host, path string) (h Handler, pattern string) {
	//	mux.mu.RLock()
	//	defer mux.mu.RUnlock()
	//	// Host-specific pattern takes precedence over generic ones
	//	if mux.hosts {
	//		h, pattern = mux.match(host + path)
	//	}
	//	if h == nil {
	//		h, pattern = mux.match(path)
	//	}
	//	if h == nil {
	//		h, pattern = NotFoundHandler(), ""
	//	}
	//	return
	//}
	// 原来他是根据用户请求的 URL 和路由器里面存储的 map 去匹配的，
	//当匹配到之后返回存储的 handler，调用这个 handler 的 ServeHTTP 接口就可以执行到相应的函数了。

	// Go 其实支持外部实现的路由器 ListenAndServe 的第二个参数就是用以配置外部路由器的，
	// 它是一个 Handler 接口，即外部路由器只要实现了 Handler 接口就可以，
	// 我们可以在自己实现的路由器的 ServeHTTP 里面实现自定义路由功能。
	//
	//
	//package main
	//
	//import (
	//	"fmt"
	//"net/http"
	//)
	//
	//type MyMux struct {
	//}
	//
	//func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//	if r.URL.Path == "/" {
	//		sayhelloName(w, r)
	//		return
	//	}
	//	http.NotFound(w, r)
	//	return
	//}
	//
	//func sayhelloName(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintf(w, "Hello myroute!")
	//}
	//
	//func main() {
	//	mux := &MyMux{}
	//	http.ListenAndServe(":9090", mux)
	//}

	//Go 代码的执行流程
	//通过对http包的分析之后，现在让我们来梳理一下整个的代码执行过程。
	//首先调用 Http.HandleFunc
	//按顺序做了几件事：
	//1 调用了 DefaultServeMux 的 HandleFunc
	//2 调用了 DefaultServeMux 的 Handle
	//3 往 DefaultServeMux 的 map [string] muxEntry 中增加对应的 handler 和路由规则
	//其次调用 http.ListenAndServe (":9090", nil)
	//按顺序做了几件事情：
	//1 实例化 Server
	//
	//2 调用 Server 的 ListenAndServe ()
	//
	//3 调用 net.Listen ("tcp", addr) 监听端口
	//
	//4 启动一个 for 循环，在循环体中 Accept 请求
	//
	//5 对每个请求实例化一个 Conn，并且开启一个 goroutine 为这个请求进行服务 go c.serve ()
	//
	//6 读取每个请求的内容 w, err := c.readRequest ()
	//
	//7 判断 handler 是否为空，如果没有设置 handler（这个例子就没有设置 handler），handler 就设置为 DefaultServeMux
	//
	//8 调用 handler 的 ServeHttp
	//
	//9 在这个例子中，下面就进入到 DefaultServeMux.ServeHttp
	//
	//10 根据 request 选择 handler，并且进入到这个 handler 的 ServeHTTP
	//
	// mux.handler(r).ServeHTTP(w, r)
	//
	//11 选择handler
	//A 判断是否有路由能满足这个 request（循环遍历 ServeMux 的 muxEntry）
	//B 如果有路由满足，调用这个路由 handler 的 ServeHTTP
	//C 如果没有路由满足，调用 NotFoundHandler 的 ServeHTTP
	//

}
