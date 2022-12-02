package models

import (
	"time"
	"votesystem/utils"

	"github.com/beego/beego/v2/core/logs"

	"github.com/beego/beego/v2/client/orm"
)

const (
	RoleAdmin  = "admin"
	RoleNormal = "normal"
)

type RoleT struct {
	RoleId      int `orm:"description(Primary Key);pk;auto"`
	RoleName    string
	Description string `orm:"null"`
	TimeModel
}

func (this *RoleT) TableName() string {
	return "vote_role"
}

type RoleRihgtsT struct {
	Id      int
	Role    *RoleT      `orm:"rel(fk)"`
	Right   *AuthorityT `orm:"rel(fk)"`
	Created time.Time   `orm:"auto_now_add; type(datetime)"`
}

func (this *RoleRihgtsT) TableName() string {
	return "vote_role_rights"
}

func (this *RoleT) TableUnique() [][]string {
	return [][]string{
		[]string{"RoleName"},
	}
}

func InsertNewRole(roleName, description string) (*RoleT, error) {

	newRole := &RoleT{
		RoleName:    roleName,
		Description: description,
	}

	o := orm.NewOrm()
	sucCount, err := o.Insert(newRole)
	if err != nil || sucCount == 0 {
		logs.Error("insert role:%v failed:%v", roleName, err)
		return nil, utils.ErrDbErr
	}
	return newRole, nil
}

func GetRole(name string) (*RoleT, error) {

	result := &RoleT{}

	o := orm.NewOrm()
	err := o.QueryTable(new(RoleT)).
		Filter("RoleName", name).
		One(result)

	if err != nil {
		logs.Error("query table failed:%v", err)
		return nil, utils.ErrDbErr
	}

	return result, nil
}

func (this *RoleT) AddRight(right *AuthorityT) error {

	newRight := &RoleRihgtsT{
		Role:  this,
		Right: right,
	}

	o := orm.NewOrm()
	_, err := o.Insert(newRight)
	if err != nil {
		logs.Error("add right for role failed:%v", err)
		return utils.ErrDbErr
	}

	return nil
}

func GetRoleRights(name string) ([]*AuthorityT, error) {

	role := &RoleT{}

	o := orm.NewOrm()
	err := o.QueryTable(new(RoleT)).
		Filter("RoleName", name).
		One(role)

	if err != nil {
		logs.Error("query table failed:%v", err)
		return nil, utils.ErrDbErr
	}

	result := []*AuthorityT{}
	_, err = o.QueryTable(new(RoleRihgtsT)).
		Filter("RoleId", role.RoleId).
		RelatedSel("Right").
		All(&result)
	if err != nil {
		logs.Error("query table failed:%v", err)
		return nil, utils.ErrDbErr
	}

	return result, nil
}

func InitBaseRole() {

	vote, _ := NewAuthority(AuthorityVote, "vote right")
	cec, _ := NewAuthority(AuthorityCreateEc, "right for create election campaign")
	ac, _ := NewAuthority(AuthorityAddCandidate, "right for add new candidate for election campaign")

	admin, _ := InsertNewRole(RoleAdmin, "admin for vote system")
	admin.AddRight(cec)
	admin.AddRight(ac)

	normal, _ := InsertNewRole(RoleNormal, "normal user")
	normal.AddRight(vote)
}

func registerRole() {
	orm.RegisterModel(new(RoleT), new(RoleRihgtsT))
}
