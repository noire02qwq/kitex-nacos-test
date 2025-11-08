package main

import (
	"context"
	"fmt"
	"log"

	"kitex-nacos-test/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// GetUserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserInfo(ctx context.Context, req *user.GetUserRequest) (resp *user.GetUserResponse, err error) {
	log.Printf("Received request for user ID: %d", req.UserID)

	// Mock user data
	users := map[int64]*user.User{
		1: {UserID: 1, Name: "Alice", Email: "alice@example.com"},
		2: {UserID: 2, Name: "Bob", Email: "bob@example.com"},
	}

	if u, ok := users[req.UserID]; ok {
		resp = &user.GetUserResponse{
			User:    u,
			Message: "Success",
		}
	} else {
		resp = &user.GetUserResponse{
			Message: fmt.Sprintf("User with ID %d not found", req.UserID),
		}
	}

	log.Printf("Returning response for user ID %d: %s", req.UserID, resp.Message)
	return
}
