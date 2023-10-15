package user

import (
	userModel "github.com/Sahil-Mandaliya/OrderService/models/user"
	userpb "github.com/Sahil-Mandaliya/OrderService/proto/user"
)

func UsersModelToPb(users []*userModel.User) []*userpb.User {
	usersPb := make([]*userpb.User, 0)
	for _, user := range users {
		usersPb = append(usersPb, UserModelToPb(user))
	}
	return usersPb
}

func UserModelToPb(user *userModel.User) *userpb.User {
	userpb := &userpb.User{
		UserId:       user.ID,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Email:        user.Email,
		MobileNumber: user.MobileNumber,
		Address:      user.Address,
	}
	return userpb
}

func UserPbToModel(user *userpb.User) *userModel.User {
	userpb := &userModel.User{
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Email:        user.Email,
		MobileNumber: user.MobileNumber,
		Address:      user.Address,
	}
	return userpb
}
