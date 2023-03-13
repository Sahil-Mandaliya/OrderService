package user

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	userModel "github.com/Sahil-Mandaliya/OrderService/models/user"

	"github.com/Sahil-Mandaliya/OrderService/pkg/util"

	jsonEncode "github.com/Sahil-Mandaliya/OrderService/pkg/json"
)

type Response struct {
	Status   bool
	Message  string
	Response interface{}
}

// Create user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := &userModel.User{}

	jsonEncode.HttpRequestUnmarshal(r, user)

	err := userModel.CreateUser(ctx, user)
	if err != nil {
		response := &Response{
			Status:  false,
			Message: err.Error(),
		}
		jsonEncode.JsonResposeEncoder(w, r, response)
		return
	}
	response := &Response{
		Status:   true,
		Response: user,
	}
	jsonEncode.JsonResposeEncoder(w, r, response)
	return
}

// Get users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	users, err := userModel.GetUsers(ctx)
	if err != nil {
		response := &Response{
			Status:  false,
			Message: err.Error(),
		}
		jsonEncode.JsonResposeEncoder(w, r, response)
		return
	}
	response := &Response{
		Status:   true,
		Response: users,
	}
	jsonEncode.JsonResposeEncoder(w, r, response)
	return
}

// Get users
func GetUserById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	queryParams := r.URL.Query()
	userId, exists := queryParams["user_id"]
	if !exists {
		response := &Response{
			Status:  false,
			Message: "Invalid UserId",
		}
		jsonEncode.JsonResposeEncoder(w, r, response)
		return
	}
	users, err := userModel.GetUserById(ctx, util.ParseUint64(userId[0]))
	if err != nil {
		response := &Response{
			Status:  false,
			Message: err.Error(),
		}
		jsonEncode.JsonResposeEncoder(w, r, response)
		return
	}
	response := &Response{
		Status:   true,
		Response: users,
	}
	jsonEncode.JsonResposeEncoder(w, r, response)
	return
}
