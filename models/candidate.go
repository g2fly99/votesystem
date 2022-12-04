package models

import (
	"votesystem/utils"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

type CandidateT struct {
	CandidateId int `orm:"description(Primary Key);pk;auto"`
	Name        string
	Sex         int
	Age         int
	ECId        int    `orm:"column(ec_id)"`
	Description string `orm:"null"`
	TimeModel
}

func (u *CandidateT) TableName() string {
	return "vote_candidate"
}

func registerCondidate() {
	orm.RegisterModel(new(CandidateT))
}

func NewCandidate(name, description string, ecId, sexy, age int) error {

	newCandidate := &CandidateT{
		Name:        name,
		Sex:         sexy,
		Age:         age,
		ECId:        ecId,
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

func ListCandidate(ecId, size, offset int) ([]*CandidateT, error) {

	o := orm.NewOrm()
	result := make([]*CandidateT, size)
	count, err := o.QueryTable(new(CandidateT)).
		Filter("ECId", ecId).
		Limit(size, offset).
		All(&result)
	if err != nil {
		logs.Error("query from db failed:%v", err)
		return nil, utils.ErrDbErr
	} else if count == 0 {
		return nil, utils.ErrEmpty
	}

	return result, nil
}

func GetAllCandidate(ecId int) ([]*CandidateT, error) {

	o := orm.NewOrm()
	result := make([]*CandidateT, 0)
	count, err := o.QueryTable(new(CandidateT)).
		Filter("ECId", ecId).
		All(&result)
	if err != nil {
		logs.Error("query from db failed:%v", err)
		return nil, utils.ErrDbErr
	} else if count == 0 {

		return nil, utils.ErrEmpty
	}

	return result, nil
}

func _GetAllCandidate(ecId int) ([]*CandidateT, error) {

	result := make([]*CandidateT, 0)
	//result := &CandidateT{}
	o := orm.NewOrm()
	count, err := o.QueryTable(new(CandidateT)).
		Filter("ECId", ecId).
		All(result)

	if err != nil {
		logs.Error("query table ec[%v] failed:%v", ecId, err)
		return nil, utils.ErrDbErr
	} else if count == 0 {
		logs.Error("query table ec[%v] no condidate:%v", ecId)
		return nil, utils.ErrEmpty
	}

	logs.Info("get result :%v", result)
	return nil, nil
}
