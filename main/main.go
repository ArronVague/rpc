package main

import (
	"log"
	"net"
	"reflect"
	"rpc"
	"strings"
	"sync"
)

func startServer(addr chan string) {
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal("network error:", err)
	}
	log.Println("start rpc server on", l.Addr())
	addr <- l.Addr().String()
	rpc.Accept(l)
}

func main() {
	var wg sync.WaitGroup
	typ := reflect.TypeOf(&wg)
	for i := 0; i < typ.NumMethod(); i++ {
		method := typ.Method(i)
		argv := make([]string, 0, method.Type.NumIn())
		returns := make([]string, 0, method.Type.NumOut())

		for j := 1; j < method.Type.NumIn(); j++ {
			argv = append(argv, method.Type.In(j).Name())
		}
		for j := 0; j < method.Type.NumOut(); j++ {
			returns = append(returns, method.Type.Out(j).Name())
		}
		log.Printf("func (w *%s) %s(%s) %s",
			typ.Elem().Name(), method.Name, strings.Join(argv, ","), strings.Join(returns, ","))
	}
	//log.SetFlags(0)
	//addr := make(chan string)
	//go startServer(addr)
	//client, _ := rpc.Dial("tcp", <-addr)
	//defer func() {
	//	_ = client.Close()
	//}()
	//
	//time.Sleep(time.Second)
	//var wg sync.WaitGroup
	//for i := 0; i < 5; i++ {
	//	wg.Add(1)
	//	go func(i int) {
	//		defer wg.Done()
	//		args := fmt.Sprintf("rpc req %d", i)
	//		var reply string
	//		if err := client.Call("Foo.Sum", args, &reply); err != nil {
	//			log.Fatal("call Foo.Sum error: ", err)
	//		}
	//		log.Println("reply: ", reply)
	//	}(i)
	//}
	//wg.Wait()
}
