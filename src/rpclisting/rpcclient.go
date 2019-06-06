package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
	"rpclisting/rpcdemo"
)

func main() {
    // 连接server并创建jsonrpc-client
    client, err := jsonrpc.Dial("tcp", "127.0.0.1:1234")
    if err != nil {
        log.Printf("create jsonrpc client error")
    }

    var result float64
    // 发起调用
    err = client.Call("DemoService.DIV", rpcdemo.Args{3, 4}, &result)
    if err != nil {
        log.Printf("call error")
    }
    fmt.Println(result)
}
