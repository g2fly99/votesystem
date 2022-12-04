package models

import (
	"errors"
	"time"
	"votesystem/utils"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

type VoteT struct {
	VoteId    int                `orm:"description(Primary Key);pk;auto"`
	Ec        *ElectionCampaignT `orm:"rel(fk);description(Election Campaign)"`
	Candidate *CandidateT        `orm:"rel(fk);description(vote to which Candidate)"`
	Voter     *User              `orm:"rel(fk);description(who votes)"`
	Created   time.Time          `orm:"auto_now_add; type(datetime)"`
}

func (this *VoteT) TableName() string {
	return "vote_votes_detail"
}

func AddNewVote(ec *ElectionCampaignT, candidate *CandidateT, voter *User) error {

	newVote := &VoteT{
		Ec:        ec,
		Candidate: candidate,
		Voter:     voter,
	}

	o := orm.NewOrm()
	_, err := o.Insert(newVote)
	if err != nil {
		logs.Error("add new vote failed:%v", err)
		return utils.ErrDbErr
	}

	return nil
}

func CountVotes(ecId, candidateId int) (int, error) {

	o := orm.NewOrm()
	num, err := o.QueryTable(new(VoteT)).
		Filter("Ec", ecId).
		Filter("Candidate", candidateId).
		Count()

	if err != nil {
		logs.Error("add new vote failed:%v", err)
		return 0, utils.ErrDbErr
	}

	return int(num), nil
}

func GetVoteWithUser(ecId, userId int) (int, error) {

	o := orm.NewOrm()
	count, err := o.QueryTable(new(VoteT)).
		Filter("Ec", ecId).
		Filter("Voter", userId).
		Count()
	if err != nil {
		logs.Error("get vote user failed:%v", err)
		return 0, utils.ErrDbErr
	}
	return int(count), nil
}

func ListAllVote(ecId, candidateId, limit, offset int) ([]*User, error) {

	o := orm.NewOrm()
	result := make([]*VoteT, limit)
	count, err := o.QueryTable(new(VoteT)).
		Filter("Ec", ecId).
		Filter("Candidate", candidateId).
		Limit(limit, offset*limit).
		RelatedSel("Voter").
		All(&result)
	if err != nil {
		logs.Error("query from db failed:%v", err)
		return nil, utils.ErrDbErr
	} else if count == 0 {
		return nil, utils.ErrEmpty
	}

	users := make([]*User, 0)
	for _, v := range result {
		//logs.Debug("voter:%v", v.Voter)
		users = append(users, v.Voter)
	}

	return users, nil
}

func ListAllEcVoters(ecId, limit, offset int) ([]*User, error) {

	o := orm.NewOrm()
	result := make([]*VoteT, limit)
	count, err := o.QueryTable(new(VoteT)).
		Filter("Ec", ecId).
		Limit(limit, offset*limit).
		RelatedSel("Voter").
		All(&result)
	if err != nil {
		logs.Error("query from db failed:%v", err)
		return nil, utils.ErrDbErr
	} else if count == 0 {
		return nil, utils.ErrEmpty
	}

	users := make([]*User, 0)
	for _, v := range result {
		//logs.Debug("voter:%v", v.Voter)
		users = append(users, v.Voter)
	}

	return users, nil
}

type ElectionCampaignT struct {
	EcId        int       `orm:"description(Primary Key);auto"`
	Finished    bool      `orm:"default(false)"`
	Expire      time.Time `orm:"type(datetime)"`
	StartTime   time.Time `orm:"type(datetime);null"`
	FinishTime  time.Time `orm:"type(datetime);null"`
	Description string    `orm:"null"`
	Created     time.Time `orm:"auto_now_add; type(datetime)"`
	Updated     time.Time `orm:"type(datetime);null"`
}

func (this *ElectionCampaignT) TableName() string {
	return "vote_election_campaign"
}

func CreateNewEC(expire time.Time, description string) (*ElectionCampaignT, error) {

	newEc := &ElectionCampaignT{
		Expire:      expire,
		Description: description,
	}

	o := orm.NewOrm()
	sucCount, err := o.Insert(newEc)
	if err != nil || sucCount == 0 {
		return nil, utils.ErrDbErr
	}

	return newEc, nil
}

func ECStart(ecId int) error {

	newEc := &ElectionCampaignT{
		EcId:      ecId,
		StartTime: time.Now(),
	}

	o := orm.NewOrm()
	_, err := o.Update(newEc, "StartTime")
	if err != nil {
		return utils.ErrDbErr
	}

	return nil
}

func ECSetFinish(ecId int) error {

	newEc := &ElectionCampaignT{
		EcId:       ecId,
		FinishTime: time.Now(),
		Finished:   true,
	}

	o := orm.NewOrm()
	_, err := o.Update(newEc, "FinishTime", "Finished")
	if err != nil {
		return utils.ErrDbErr
	}

	return nil
}

func GetEcInfo(ecId int) (*ElectionCampaignT, error) {

	newEc := &ElectionCampaignT{
		EcId: ecId,
	}

	o := orm.NewOrm()
	err := o.Read(newEc, "EcId")
	if err != nil {
		logs.Error("read from db failed:%v")
		return nil, utils.ErrDbErr
	}

	return newEc, nil
}

func ListEcs(size, offset int) ([]*ElectionCampaignT, error) {

	o := orm.NewOrm()
	result := make([]*ElectionCampaignT, size)
	count, err := o.QueryTable(new(ElectionCampaignT)).
		Limit(size, offset).
		All(&result)
	if err != nil {
		logs.Error("query from db failed:%v", err)
		return nil, errors.New("dberr")
	} else if count == 0 {
		return nil, utils.ErrDbErr
	}

	return result, nil
}

func ListActiveEcs(size, offset int) ([]*ElectionCampaignT, error) {

	o := orm.NewOrm()
	result := make([]*ElectionCampaignT, size)
	count, err := o.QueryTable(new(ElectionCampaignT)).
		Filter("Finished", false).
		Limit(size, offset).
		All(&result)
	if err != nil {
		logs.Error("query from db failed:%v", err)
		return nil, errors.New("dberr")
	} else if count == 0 {
		return nil, utils.ErrDbErr
	}

	return result, nil
}

func registerVotes() {
	orm.RegisterModel(new(ElectionCampaignT), new(VoteT))
}
