package main

import (
	"database/sql"

	"github.com/astaxie/beego/orm"
)

type MyUser struct {
	Id      int
	Name    string
	Profile *Profile `orm:"rel(one)"`      // OneToOne relation
	Post    []*Post  `orm:"reverse(many)"` // 设置一对多的反向关系
}
type Profile struct {
	Id   int
	Age  int16
	User *MyUser `orm:"reverse(one)"` // 设置一对一反向关系(可选)
}
type Post struct {
	Id    int
	Title string
	User  *MyUser `orm:"rel(fk)"` //设置一对多关系
	Tags  []*Tag  `orm:"rel(m2m)"`
}
type Tag struct {
	Id    int
	Name  string
	Posts []*Post `orm:"reverse(many)"`
}

type Customer struct {
	ID    int     `orm:"pk;auto;column(id)" json:"id,omitempty"`
	Name  string  `orm:"column(name)" json:"name,omitempty"`
	Email *string `orm:"null;column(email)" json:"email,omitempty"`
}

type Person struct {
	ID    int            `orm:"pk;auto;column(id)" json:"id,omitempty"`
	Name  string         `orm:"column(name)" json:"name,omitempty"`
	Email sql.NullString `orm:"null;column(email)" json:"email,omitempty"`
}

func init() {
	orm.RegisterModel(new(MyUser), new(Post), new(Profile), new(Tag), new(Customer), new(Person))
	orm.RunSyncdb("default", true, false)
}
