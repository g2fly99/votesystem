package models

import (
	"votesystem/utils"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

type User struct {
	UserId     int    `orm:"description(Primary Key);auto"`
	Username   string `orm:"null"`
	Email      string
	IdentityNo string
	Role       *RoleT `orm:"null;rel(one);on_delete(set_null)"`
}

func (this *User) TableName() string {
	return "vote_user"
}

func (this *User) TableUnique() [][]string {
	return [][]string{
		[]string{"IdentityNo", "Email"},
	}
}

func AddNormalUser(name, email, identityNo string) error {

	role, err := GetRole(RoleNormal)
	if err != nil {
		logs.Error("get normal role failed:%v", err)
		return utils.ErrDbErr
	}

	newUser := &User{
		Username:   name,
		Email:      email,
		IdentityNo: identityNo,
		Role:       role,
	}

	o := orm.NewOrm()
	_, err = o.Insert(newUser)
	if err != nil {
		logs.Error("add normal user failed:%v", err)
		return utils.ErrDbErr
	}

	return nil
}

func GetUser(uid int) (u *User, err error) {

	return nil, utils.ErrNotExist
}

func registerUser() {
	orm.RegisterModel(new(User))
}
