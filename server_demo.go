package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
)

// 编写一个 HTTP 服务器：
// 接收客户端 request，并将 request 中带的 header 写入 response header
// 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
// Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
// 当访问 localhost/healthz 时，应返回 200

func main() {
	HttpServerStart(8080)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func HttpServerStart(port int) {

	http.HandleFunc("/", httpAccessFunc)
	http.HandleFunc("/healthz", healthzFunc)

	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		log.Fatal(err)
	}
}

func healthzFunc(w http.ResponseWriter, r *http.Request) {
	HealthzCode := "200"
	w.Write([]byte(HealthzCode))
}

func httpAccessFunc(w http.ResponseWriter, r *http.Request) {
	if len(r.Header) > 0 {
		for k, v := range r.Header {
			log.Printf("%s=%s", k, v[0])
			w.Header().Set(k, v[0])
		}
	}

	r.ParseForm()
	if len(r.Form) > 0 {
		for k, v := range r.Form {
			log.Printf("%s=%s", k, v[0])
		}
	}

	name := os.Getenv("VERSION")
	w.Header().Set("Version", name)

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		fmt.Println("err:", err)
	}

	if net.ParseIP(ip) != nil {
		log.Println(ip)
	}

	w.WriteHeader(http.StatusOK)

	w.Write([]byte("Server Access,Success!"))
}
