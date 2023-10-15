package user

import (
	"context"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	// mysqlDb "github.com/Sahil-Mandaliya/Order/backend/pkg/mysql"

	mysqlDb "github.com/Sahil-Mandaliya/OrderService/pkg/mysql"
)

/***************************************************/

type User struct {
	ID           uint64 `json:"id"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Email        string `json:"email"`
	MobileNumber string `json:"mobileNumber"`
	Address      string `json:"address"`
}

// Get all users
func GetUsers(ctx context.Context) ([]*User, error) {
	db := mysqlDb.DBConnection()
	users := []*User{}
	err := db.Model(&User{}).Find(&users).Error
	if err != nil {
		return users, err
	}
	return users, nil
}

// Get user by id
func GetUserById(ctx context.Context, userId uint64) (*User, error) {
	db := mysqlDb.DBConnection()
	user := &User{}
	err := db.Model(&User{}).Where("id=?", userId).Find(user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

// Create user
func CreateUser(ctx context.Context, user *User) (*User,error) {
	db := mysqlDb.DBConnection()
	gormObj := db.Create(user)
	if gormObj.Error != nil {
		return nil, gormObj.Error
	}
	return user, nil
}

// Get user by ID
func GetUser(ctx context.Context, userId uint64) (*User, error) {
	db := mysqlDb.DBConnection()
	user := &User{}
	err := db.Model(&User{}).Where("id = ?", userId).Take(user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

// // Update user
// func UpdateUser(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	stmt, err := mysqlDb.DBConnection().Prepare("UPDATE users SET first_name = ?," +
// 		"last_name= ?, email=? WHERE id = ?")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	keyVal := make(map[string]string)
// 	json.Unmarshal(body, &keyVal)
// 	first_name := keyVal["firstName"]
// 	last_name := keyVal["lastName"]
// 	email := keyVal["email"]
// 	_, err = stmt.Exec(first_name, last_name, email,
// 		params["id"])
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	fmt.Fprintf(w, "User with ID = %s was updated",
// 		params["id"])
// }

// // Delete User
// func DeleteUser(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	stmt, err := mysqlDb.DBConnection().Prepare("DELETE FROM users WHERE id = ?")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	_, err = stmt.Exec(params["id"])
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	fmt.Fprintf(w, "User with ID = %s was deleted",
// 		params["id"])
// }
