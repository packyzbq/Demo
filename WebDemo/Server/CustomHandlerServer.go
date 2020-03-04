// 这里使用3种方式启动一个服务，handlefunc 、 ServeMux 以及 Server 结构体

package Server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// 1.HandleFunc 实际上是实现了 Handler 接口，这里可以自定义实现 Handler 接口
type helloHandler struct {
}

func (_ *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}

// 2.通过创建mux 服务路由器实现访问，ServeMux 也实现了 handler 接口
func getMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/", &helloHandler{})
	return mux
}

// 3. 更加底层的用法，直接使用 Server
func getServer(addr string) *http.Server {
	mux := getMux()
	// 增加一个超时api
	mux.HandleFunc("/timeout", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second)
		w.Write([]byte("Timeout"))
	})
	server := &http.Server{
		Addr:    addr,
		Handler: mux,
		// 调用更底层的server 可以多定义一些特性，比如超时时间
		WriteTimeout: 2 * time.Second,
	}
	return server
}

func StartCustomHandlerServer() {
	server := getServer("localhost:8080")

	// 优雅的停止服务
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit

		if err := server.Shutdown(context.Background()); err != nil {
			log.Fatal("SHUTDOWN server:", err)
		}
	}()

	// 1. 自定义handler
	// log.Fatal(http.ListenAndServe("localhost:8080", &helloHandler{}))

	// 2. 自定义 mux 实现路由，是大部分web框架的底层用法
	//log.Fatal(http.ListenAndServe("localhost:8080", getMux()))

	// 3. 使用自定义 server
	log.Fatal(server.ListenAndServe())

}
