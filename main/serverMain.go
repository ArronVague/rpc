package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"rpc"
)

var (
	serverIp   string
	serverPort string
)

func main() {
	flag.StringVar(&serverIp, "l", "0.0.0.0", "server listen ip   usage:-l <ipv4/ipv6>")
	flag.StringVar(&serverPort, "p", "", "server listen port   usage:-p <port>")
	flag.Parse()

	if serverPort == "" {
		log.Fatalln("server port is required")
	}
	addr := fmt.Sprintf("%s:%s", serverIp, serverPort)

	var foo rpc.Foo
	l, _ := net.Listen("tcp", addr)
	server := rpc.NewServer()
	_ = server.Register(&foo)
	server.Accept(l)
}
