package main

import (
	"log"

	"kitex-nacos-test/common"
	"kitex-nacos-test/kitex_gen/user/userservice"
)

func main() {
	// Create nacos registry
	r, err := common.NewNacosRegistry()
	if err != nil {
		log.Fatal(err)
	}

	svr := userservice.NewServer(
		new(UserServiceImpl),
		userservice.WithRegistry(r),
		userservice.WithServiceAddr(&userservice.Address{Host: "127.0.0.1", Port: "9001"}),
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
