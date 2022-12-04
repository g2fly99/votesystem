package models

import (
	"votesystem/utils"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

const (
	AuthorityVote         = "rightForVote"
	AuthorityCreateEc     = "rightForCreateEc"
	AuthorityAddCandidate = "rightForaddAcndidate"
)

type AuthorityT struct {
	AuthorityId   int    `orm:"auto"`
	AuthorityName string `orm:"unique"`
	Description   string `orm:"null"`
	TimeModel
}

func (this *AuthorityT) TableName() string {
	return "vote_authority"
}

func GetAllAuthority() (result []*AuthorityT, err error) {

	o := orm.NewOrm()
	_, err = o.QueryTable(&AuthorityT{}).All(&result)
	if err != nil {
		logs.Error("query table failed:%v", err)
		err = utils.ErrDbErr
		return
	}
	return
}

func NewAuthority(authorityName, description string) (*AuthorityT, error) {

	newAuthority := &AuthorityT{
		AuthorityName: authorityName,
		Description:   description,
	}

	o := orm.NewOrm()
	_, sucCount, err := o.ReadOrCreate(newAuthority, "AuthorityName")
	if err != nil || sucCount == 0 {
		return nil, utils.ErrDbErr
	}
	return newAuthority, nil
}

func registerAuthority() {
	orm.RegisterModel(new(AuthorityT))
}
