package models

import (
	"votesystem/utils"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

type CandidateT struct {
	CandidateId int `orm:"description(Primary Key);pk;auto"`
	Name        string
	Sexy        int
	Age         int
	EC          *ElectionCampaignT `orm:"rel(fk)"`
	Description string             `orm:"null"`
	TimeModel
}

func (u *CandidateT) TableName() string {
	return "vote_candidate"
}

func NewCandidate(name, description string, ecId, sexy, age int) error {

	newCandidate := &CandidateT{
		Name:        name,
		Sexy:        sexy,
		Age:         age,
		EC:          &ElectionCampaignT{EcId: ecId},
		Description: description,
	}

	o := orm.NewOrm()
	sucCount, err := o.Insert(newCandidate)
	if err != nil || sucCount == 0 {
		return utils.ErrDbErr
	}

	return nil
}

func NewMultiCandidate(candidates []CandidateT) error {

	o := orm.NewOrm()
	sucCount, err := o.InsertMulti(len(candidates), candidates)
	if err != nil || sucCount == 0 {
		return utils.ErrDbErr
	}

	return nil
}

func GetAllCandidate(ecId int) ([]*CandidateT, error) {

	result := []*CandidateT{}
	o := orm.NewOrm()
	_, err := o.QueryTable(new(CandidateT)).Filter("EC", ecId).All(result)
	if err != nil {
		logs.Error("query table ec[%v] failed:%v", ecId, err)
		return nil, utils.ErrDbErr
	}

	return result, nil
}
