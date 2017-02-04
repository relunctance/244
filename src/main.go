package main

import (
	//"fmt"
	"model"
)

func main() {
	user := model.UserModel{} // ./main.go:13: cannot use userStr (type string) as type model.BaseModel in field value
	user.Tabname = "user"
	user.Selectdb = "default"
	user.Initialize()
	user.Demo()

	/*
		data := make(model.BatchInsertMap)
			item := []string{
				"1",
				"gaoqilin1",
				"pass11",
			}
			item2 := []string{
				"2",
				"gaoqilin2",
				"pass22",
			}
			data[0] = item
			data[1] = item2
			fields := []string{
				"id",
				"username",
				"password",
			}

			//fmt.Printf("%#v", data)
			//fmt.Printf("%#v", fields)
			err := user.BatchInsert(fields, data, true)
			if err != nil {
				fmt.Println("--------------------- error -------------------\n")
				fmt.Println(err)
				fmt.Println("--------------------- error -------------------\n")
			}
	*/

}
