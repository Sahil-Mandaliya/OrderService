package user

import (
	"context"

	userModel "github.com/Sahil-Mandaliya/OrderService/models/user"
)

type Response struct {
	Status   bool
	Message  string
	Response interface{}
}

func CreateUser(ctx context.Context, reqUser *userModel.User) (*userModel.User, error) {
	user, err := userModel.CreateUser(ctx, reqUser)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetAllUsers(ctx context.Context) ([]*userModel.User, error) {
	users, err := userModel.GetUsers(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}
