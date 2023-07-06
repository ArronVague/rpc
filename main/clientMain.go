package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"rpc"
)

var (
	clientIp   string
	clientPort string
)

func main() {
	flag.StringVar(&clientIp, "i", "", "client connect ip   usage:-i <ipv4/ipv6>")
	flag.StringVar(&clientPort, "p", "", "client connect port   usage:-p <port>")
	flag.Parse()

	if clientIp == "" {
		log.Fatalln("client's connect ip is required")
	}
	if clientPort == "" {
		log.Fatalln("client's connect port is required")
	}

	addr := fmt.Sprintf("%s:%s", clientIp, clientPort)
	client, _ := rpc.Dial("tcp", addr)
	var reply int
	err := client.Call(context.Background(), "Foo.Sum", rpc.Args{
		Num1: 1,
		Num2: 2,
	}, &reply)
	//time.Sleep(time.Second)
	//call(addr)
	if err != nil {
		log.Println(err)
	} else {
		log.Printf("1 + 2 = %d", reply)
	}
	// 测试服务调用，支持客户端并发
	//wg := new(sync.WaitGroup)
	//wg.Add(2)
	//go func(wg *sync.WaitGroup) {
	//	defer wg.Done()
	//	res1 := client.Call("Add", 2, 2)
	//	logger.Infoln(logger.InfoMsg(fmt.Sprintf("Add远程调用的响应消息：%v", res1)))
	//}(wg)
	//
	//go func(wg *sync.WaitGroup) {
	//	defer wg.Done()
	//	res2 := client.Call("Substract", 2, 2)
	//	logger.Infoln(logger.InfoMsg(fmt.Sprintf("Substract远程调用的响应消息：%v", res2)))
	//}(wg)
	//wg.Wait()
}
