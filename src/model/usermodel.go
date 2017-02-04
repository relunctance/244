package model

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type User struct {
	Id       int    `pk:"auto"`
	Username string `orm:"size(50)"`
	Password string `orm:"size(32)"`
	Phone    string `orm:"size(11)"`
	Sex      int
	Ctime    time.Time
}

type UserModel struct {
	BaseModel
}

func (this *UserModel) Demo() {
	var r orm.RawSeter
	r = this.O.Raw("select * from user where id>=?", 1)
	/*
		var userdata User
		r.QueryRow(&userdata)
	*/
	/*
		r = this.O.Raw("select username from user where id>=?", 1)
			type Userdata struct {
				Username string
			}
			var userdata []Userdata
			r.QueryRows(&userdata)
	*/

	/*
		var userdata []orm.Params
		r.Values(&userdata)
	*/

	var userdata []orm.ParamsList
	r.ValuesList(&userdata)

	//fmt.Printf("%#v", userdata)

	field := []string{
		"username",
		"password",
		"phone",
		"sex",
	}
	items := make(BatchInsertMap)
	for key, val := range userdata {
		items[key] = []string{
			fmt.Sprintf("%s", val[1]),
			fmt.Sprintf("%s", val[2]),
			fmt.Sprintf("%s", val[3]),
			fmt.Sprintf("%s", val[4]),
		}
	}
	//fmt.Printf("%#v", items)
	//fmt.Println(field)

	err := this.BatchInsert(field, items, false)

	if err != nil {
		fmt.Println("error ###############\n")
		fmt.Println(err)
		fmt.Println("error ###############\n")
	}
}
