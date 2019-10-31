package main

import (
	"github.com/astaxie/beego/orm"
)

// HarborUser holds the details of a user.
type HarborUser struct {
	UserID          int    `orm:"pk;auto;column(user_id)" json:"user_id"`
	Username        string `orm:"column(username)" json:"username"`
	Email           string `orm:"column(email)" json:"email"`
	Password        string `orm:"column(password)" json:"password"`
	Realname        string `orm:"column(realname)" json:"realname"`
	Comment         string `orm:"column(comment)" json:"comment"`
	Deleted         bool   `orm:"column(deleted)" json:"deleted"`
	Rolename        string `orm:"-" json:"role_name"`
	// if this field is named as "RoleID", beego orm can not map role_id
	// to it.
	Role int `orm:"-" json:"role_id"`
	//	RoleList     []Role `json:"role_list"`
	HasAdminRole bool   `orm:"column(sysadmin_flag)" json:"has_admin_role"`
	ResetUUID    string `orm:"column(reset_uuid)" json:"reset_uuid"`
	Salt         string `orm:"column(salt)" json:"-"`
}

func init() {
	orm.RegisterModel(new(HarborUser))
	orm.RunSyncdb("default", true, false)
}
