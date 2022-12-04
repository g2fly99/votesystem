package models

import (
	"strings"
	"votesystem/utils"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

type User struct {
	UserId     int    `orm:"description(Primary Key);auto"`
	Username   string `orm:"null"`
	Email      string
	IdentityNo string
	Role       *RoleT `orm:"null;rel(fk);on_delete(set_null)"`
}

func (this *User) TableName() string {
	return "vote_user"
}

func (this *User) TableUnique() [][]string {
	return [][]string{
		[]string{"IdentityNo", "Email"},
	}
}

func GetNormalUser(idNo, email string) ([]*User, error) {

	u := []*User{}

	cond := orm.NewCondition()
	emailCondition := cond.Or("Email", email)
	iddentCondition := cond.Or("IdentityNo", idNo).OrCond(emailCondition)

	o := orm.NewOrm()
	qs := o.QueryTable(&User{})

	_, err := qs.SetCond(iddentCondition).All(&u)
	if err != nil {
		logs.Error("select user from db failed:%v,id[%v],email[%v]", err, idNo, email)
		return nil, utils.ErrDbErr
	}
	return u, nil
}

func AddNormalUser(name, identityNo, email string) (*User, error) {

	role, err := GetRole(RoleNormal)
	if err != nil {
		logs.Error("get normal role failed:%v", err)
		return nil, utils.ErrDbErr
	}

	newUser := &User{
		Username:   name,
		Email:      email,
		IdentityNo: identityNo,
		Role:       role,
	}

	o := orm.NewOrm()
	Id, err := o.Insert(newUser)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate") {
			return nil, utils.ErrExist
		} else {
			logs.Error("add normal user failed:%v", err)
			return nil, utils.ErrDbErr
		}
	}

	newUser.UserId = int(Id)
	return newUser, nil
}

func registerUser() {
	orm.RegisterModel(new(User))
}
