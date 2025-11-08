package main

import (
	"log"

	"kitex-nacos-test/common"
	"kitex-nacos-test/kitex_gen/order/orderservice"
)

func main() {
	// Create nacos registry
	r, err := common.NewNacosRegistry()
	if err != nil {
		log.Fatal(err)
	}

	svr := orderservice.NewServer(
		new(OrderServiceImpl),
		orderservice.WithRegistry(r),
		orderservice.WithServiceAddr(&orderservice.Address{Host: "127.0.0.1", Port: "9002"}),
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
