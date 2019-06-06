package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	service := ":7778"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		var buf [512]byte
		conn.Read(buf[0:])
		fmt.Println(string(buf[0:]))
		daytime := time.Now().String()
		conn.Write([]byte(daytime)) // don't care about return value
		conn.Close()                // we're finished with this client
	}
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
