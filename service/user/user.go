package user

import (
	"context"
	"fmt"

	userCtr "github.com/Sahil-Mandaliya/OrderService/controller/user"
	"github.com/Sahil-Mandaliya/OrderService/pkg/json"
	userpb "github.com/Sahil-Mandaliya/OrderService/proto/user"
)

type OrderService struct {
	userpb.UnimplementedOrderServiceServer
}

func (rpc OrderService) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	reqUser := userCtr.UserPbToModel(req.User)
	users, err := userCtr.CreateUser(ctx, reqUser)
	if err != nil {
		return nil, err
	}
	res := &userpb.CreateUserResponse{
		User: userCtr.UserModelToPb(users),
	}
	return res, nil
}

func (rpc OrderService) GetAllUsers(ctx context.Context, req *userpb.GetAllUsersRequest) (*userpb.GetAllUsersResponse, error) {
	users, err := userCtr.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}
	usersPb := userCtr.UsersModelToPb(users)
	res := &userpb.GetAllUsersResponse{
		Users: usersPb,
	}
	return res, nil
}
