package main

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq" // register pgsql driver
)

func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	info := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		`10.162.17.149`, `5432`, `postgres`, `root123`, `registry`, `disable`)
	orm.RegisterDataBase("default", "postgres", info)
}
func main() {
	o := orm.NewOrm()
	o.Using("registry") // 默认使用 default，你可以指定为其他数据库
	var users []*HarborUser

	for index := 0; index < 1000; index++ {
		user := &HarborUser{
			Username: fmt.Sprintf("stonezdj%04d", index),
			Email:    fmt.Sprintf("stonezdj%04d@163.com", index),
			Realname: fmt.Sprintf("stonezdj%04d", index),
			Comment:  fmt.Sprintf("example user: stonezdj%04d", index),
		}
		users = append(users, user)

	}
	fmt.Println(o.InsertMulti(1000, users))

	fmt.Println("Done!")
}
