package main

import (
	"fmt"
	//	"logic"
	"model"
)

func init() {
}

func main() {
	user := model.UserModel{} // ./main.go:13: cannot use userStr (type string) as type model.BaseModel in field value
	user.Tabname = "user"

	var userdata []model.User
	fmt.Println(user.GetAllData(&userdata)) //	fmt.Println(model.D)
	fmt.Println(userdata)

	myuser := model.MyuserModel{BaseModel: model.BaseModel{Tabname: "myuser"}}

	var myuserdata []model.Myuser
	fmt.Println(myuser.GetAllData(&myuserdata))
	fmt.Println(myuserdata)

}
