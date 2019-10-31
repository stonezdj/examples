package main

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", "postgres://postgres:root123@10.4.140.114/ormdb?sslmode=disable")
}
func main() {
	o := orm.NewOrm()
	o.Using("ormdb") // 默认使用 default，你可以指定为其他数据库
	profile := new(Profile)
	profile.Age = 30
	user := new(MyUser)
	user.Profile = profile
	user.Name = "slene"
	fmt.Println(o.Insert(profile))
	fmt.Println(o.Insert(user))
}
