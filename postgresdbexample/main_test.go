package main

import (
	"fmt"
	"testing"

	"github.com/astaxie/beego/orm"
)

func TestCRUD(t *testing.T) {

	o := orm.NewOrm()
	user := new(MyUser)
	profile := new(Profile)
	profile.Age = 40

	user.Name = "slene"
	user.Profile = profile

	fmt.Println(o.Insert(user))

	user.Name = "Your"
	fmt.Println(o.Update(user))
	fmt.Println(o.Read(user))
	//fmt.Println(o.Delete(user))

}

func TestRead(t *testing.T) {
	o := orm.NewOrm()
	u := &MyUser{Name: "mike"}
	o.Insert(u)

	user := MyUser{Name: "mike"}

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

	user := new(MyUser)
	user.Name = "sample"
	user.Profile = profile

	fmt.Println(o.Insert(user))
	fmt.Println("Insert correct!")

}

func TestInserCustomer(t *testing.T) {
	o := orm.NewOrm()
	customer := new(Customer)
	customer.Name = "mike"
	email := "mike@example.com"
	customer.Email = &email
	fmt.Println(o.Insert(customer))

	cust2 := &Customer{Name: "mike"}
	o.Read(cust2, "name")
	fmt.Printf("customer found %+v, email:%v\n", cust2, *cust2.Email)

	cust3 := &Customer{}
	o.Raw("select id, name, email from customer where name = ? limit 1", "mike").QueryRow(cust3)
	fmt.Printf("customer found %+v\n", cust3)
	fmt.Printf("customer email found %+v\n", cust3.Email)

}
func TestInserPerson(t *testing.T) {
	o := orm.NewOrm()
	p := new(Person)
	p.Name = "mike"
	email := "mike@example.com"
	p.Email.Scan(email)
	fmt.Println(o.Insert(p))

	p2 := &Person{Name: "mike"}
	o.Read(p2, "name")
	fmt.Printf("person found %+v, email:%v\n", p2, p2.Email.String)

	p3 := &Person{}
	o.Raw("select id, name, email from person where name = ? limit 1", "mike").QueryRow(p3)
	fmt.Printf("person found: %+v\n", p3)
	fmt.Printf("person email: %+v\n", p3.Email.String)

}
