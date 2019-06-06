package main

import (
    "log"
    "net"
    "net/rpc"
    "net/rpc/jsonrpc"
	"rpclisting/rpcdemo"
)

func main() {
    // 注册服务
    rpc.Register(rpcdemo.DemoService{})
    // 启动 server
    listener, err := net.Listen("tcp", "127.0.0.1:1234")

    if err != nil {
        panic(err)
    }

    for {
        // 不断连接服务
        conn, err := listener.Accept()
        if err != nil {
            log.Printf("accept error, %v", err)
            continue
        }
        // 使用 Goroutine：ServeConn runs the JSON-RPC server on a single connection.
        go jsonrpc.ServeConn(conn)
    }
}
