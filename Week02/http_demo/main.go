package main

import (
	"fmt"
	"http_demo/api"
	"log"
	"net/http"
)

const HOST string = "127.0.0.1:8080"

func main() {
	startServer()
}

func startServer() {
	mux := http.NewServeMux()

	mux.Handle("/", api.NewIndexHandler())
	mux.Handle("/user/", api.NewUserHandler())

	fmt.Printf("server start http://%s\n", HOST)

	fmt.Printf("case: 找到存在用户请访问: http://%s/user/1001\n", HOST)
	fmt.Printf("case: 用户不存在请访问: http://%s/user/1002\n", HOST)
	fmt.Printf("case: 查找用户server出错: http://%s/user/1003\n", HOST)
	fmt.Printf("case: 查找用户参数出错: http://%s/user/xxx\n", HOST)

	log.Fatal(http.ListenAndServe(HOST, mux))

}
