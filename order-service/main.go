package main

import (
	"log"
	"net"

	"kitex-nacos-test/common"
	"kitex-nacos-test/kitex_gen/order/orderservice"
	"github.com/cloudwego/kitex/server"
)

func main() {
	// Create nacos registry
	r, err := common.NewNacosRegistry()
	if err != nil {
		log.Fatal(err)
	}

	// Create service address
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:9002")
	if err != nil {
		log.Fatal(err)
	}

	svr := orderservice.NewServer(
		new(OrderServiceImpl),
		server.WithRegistry(r),
		server.WithServiceAddr(addr),
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
