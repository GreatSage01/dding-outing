package main

import (
	"fmt"
	"net/http"

	"app/dingtalk"
)

func main() {

	// 自定义方法
	outgoingFunc := func(args []string) []byte {
		// do what you want to
		fmt.Println("测试")
		DtSend()
		return nil
	}

	// 自定义方法注册到handler
	dingtalk.RegisterCommand("hello", outgoingFunc, 3, true)

	// 启动http服务
	http.Handle("/outgoing", &dingtalk.OutGoingHandler{})
	_ = http.ListenAndServe(":8000", nil)

}
