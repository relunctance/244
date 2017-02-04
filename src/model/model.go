package model

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

type BaseModel struct {
	Tabname  string //
	Selectdb string
	O        orm.Ormer
}
type BatchInsertMap map[int][]string

func init() {
	//mysql -uuser_abc -h127.0.0.1 -ppass_123456 mytestdb
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", "user_abc", "pass_1234", "127.0.0.1", 3306, "mytestdb", "utf8"))
	orm.Debug = true // 需要放到配置文件里

}

//注册model , 可选择不注册, 主要是为了 orm的其他操作
//func (this *BaseModel) RegisterModel(model interface{}) {
//orm.RegisterModel(new(model))
//}
func (this *BaseModel) Initialize() {
	this.O = orm.NewOrm()       //初始化
	this.O.Using(this.Selectdb) //选择数据库
}

//批量更新
func (this *BaseModel) BatchInsert(field []string, data BatchInsertMap, needUpdate bool) (int, error) {

	dataLength := len(data)
	fieldLength := len(field)
	if fieldLength == 0 {
		return 0, fmt.Errorf("%s", "batchInset field is empty")
	}

	if dataLength == 0 {
		return 0, fmt.Errorf("%s", "batchInset values is empty")
	}

	if fieldLength != len(currentMapInt(data)) {
		return 0, fmt.Errorf("%s", "field length not same BatchInsertMap Item. ")
	}

	sql := fmt.Sprintf("INSERT INTO `%s`(%s) ", this.Tabname, "`"+strings.Join(field, "`,`")+"`")
	fillItem := make([]string, fieldLength)
	for i := 0; i < fieldLength; i++ {
		fillItem[i] = "?"
	}
	valQmarks, valueItems := []string{}, []string{}
	fillItemStr := "(" + strings.Join(fillItem, ",") + ")"
	for _, vals := range data {
		valQmarks = append(valQmarks, fillItemStr)
		valueItems = append(valueItems, vals...)
	}

	sql = fmt.Sprintf("%s VALUES%s", sql, strings.Join(valQmarks, ","))
	if needUpdate == true {
		sql = fmt.Sprintf("%s %s", sql, duplicateUpdate(field))
	}
	res, err := this.O.Raw(sql, valueItems).Exec()
	if err != nil {
		return 0, err
	}
	return res.RowsAffected(), nil

}

func duplicateUpdate(field []string) string {
	updateStr := ""
	for _, v := range field {
		updateStr += fmt.Sprintf("`%s` = VALUES(`%s`),", v, v)
	}
	updateStr = strings.TrimRight(updateStr, ",")
	return fmt.Sprintf(" ON DUPLICATE KEY UPDATE %s", updateStr)
}

func currentMapInt(maps map[int][]string) (value []string) {
	for _, val := range maps {
		value = val
		return value
	}
	return value
}
