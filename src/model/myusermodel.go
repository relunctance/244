package model

import (
	//"fmt"
	"time"
)

type Myuser struct {
	Id       int    `pk:"auto"`
	Username string `orm:"size(50)"`
	Password string `orm:"size(32)"`
	Phone    string `orm:"size(11)"`
	Sex      int
	Ctime    time.Time
}

type MyuserModel struct {
	BaseModel
}
