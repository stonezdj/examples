package main

import (
	"fmt"
	"testing"

	"github.com/astaxie/beego/orm"
)

func TestCRUD(t *testing.T) {

	o := orm.NewOrm()
	user := new(User)
	profile := new(Profile)
	profile.Age = 40

	user.Name = "slene"
	user.Profile = profile

	fmt.Println(o.Insert(user))

	user.Name = "Your"
	fmt.Println(o.Update(user))
	fmt.Println(o.Read(user))
	fmt.Println(o.Delete(user))

}

func TestRead(t *testing.T) {
	o := orm.NewOrm()
	user := User{Id: 1}

	err := o.Read(&user)

	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	} else {
		fmt.Println(user.Id, user.Name)
	}
}

func TestInsert(t *testing.T) {
	o := orm.NewOrm()
	profile := new(Profile)
	profile.Age = 39

	user := new(User)
	user.Name = "sample"
	user.Profile = profile

	fmt.Println(o.Insert(user))
	fmt.Println("Insert correct!")

}
