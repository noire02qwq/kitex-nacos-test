package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"kitex-nacos-test/common"
	"kitex-nacos-test/kitex_gen/order"
	"kitex-nacos-test/kitex_gen/user"
	"kitex-nacos-test/kitex_gen/user/userservice"
)

// OrderServiceImpl implements the last service interface defined in the IDL.
type OrderServiceImpl struct{}

// CreateOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) CreateOrder(ctx context.Context, req *order.CreateOrderRequest) (resp *order.CreateOrderResponse, err error) {
	log.Printf("Received request to create order for user ID: %d, product: %s", req.UserID, req.Product)

	// Create user service client with Nacos resolver
	r, err := common.NewNacosResolver()
	if err != nil {
		log.Fatal(err)
	}

	client, err := userservice.NewClient("UserService", userservice.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}

	// Call user service to get user info
	userReq := &user.GetUserRequest{UserID: req.UserID}
	userResp, err := client.GetUserInfo(ctx, userReq)
	if err != nil {
		resp = &order.CreateOrderResponse{
			Message: fmt.Sprintf("Failed to get user info: %v", err),
		}
		return
	}

	// Create mock order
	rand.Seed(time.Now().UnixNano())
	orderInfo := &order.OrderInfo{
		OrderID:  rand.Int63n(1000000),
		UserID:   req.UserID,
		Product:  req.Product,
		UserName: userResp.User.Name,
	}

	resp = &order.CreateOrderResponse{
		Order:   orderInfo,
		Message: "Success",
	}

	log.Printf("Created order ID: %d for user: %s", orderInfo.OrderID, userResp.User.Name)
	return
}
