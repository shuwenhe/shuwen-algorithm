package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc/jsonrpc"
)

func runClient() {
	client, err := net.Dial("tcp", "127.0.0.1:4321")
	if err != nil {
		fmt.Println("dialing", err)
	}

	// 第一阶段：请求票
	args := &Args{"127.0.0.1:4321:require_ticket:34"}
	var reply Reply
	c := jsonrpc.NewClient(client)
	err = c.Call("PaxosService.Process", args, &reply)
	if err != nil {
		log.Fatal("PaxosService error:", err)
	}
	fmt.Println("阶段1 response:", reply.StrResult)

	// 第二阶段：propose
	args = &Args{"127.0.0.1:4321:require_propose:34:akp"}
	err = c.Call("PaxosService.Process", args, &reply)
	if err != nil {
		log.Fatal("PaxosService error:", err)
	}
	fmt.Println("阶段2 response:", reply.StrResult)

	// 第三阶段：commit
	args = &Args{"127.0.0.1:4321:require_propose:34:akp"}
	err = c.Call("PaxosService.Process", args, &reply)
	if err != nil {
		log.Fatal("PaxosService error:", err)
	}
	fmt.Println("阶段3 response:", reply.StrResult)
}
