package controllers

import (
	"encoding/json"
	"regexp"
	"strconv"
	"time"
	"votesystem/models"
	"votesystem/response"
	"votesystem/utils"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

// Operations about VoteController
type VoteController struct {
	beego.Controller
}

type VoteDetailT struct {
	EcId        int           `json:"ecId"`
	Finished    bool          `json:"finished"`
	Expire      time.Time     `json:"expire"`
	StartTime   time.Time     `json:"startTime"`
	FinishTime  time.Time     `json:"finishTime"`
	Description string        `json:"description"`
	CreateTime  time.Time     `json:"created"`
	VoteNumber  int           `json:"voteNumber"`
	Candidates  []*CandidateT `json:"candidates"`
}

type CandidateAddParamT struct {
	Candidate []CandidateT `json:"candidate"`
}

func checkEcVoteValid(ecId int) (bool, response.ResponseT) {

	ec, err := models.GetEcInfo(ecId)
	if err != nil {
		logs.Error("get info from db failed:%v", err)
		return false, response.ErrSystem
	}

	if ec.EcId == 0 {
		logs.Error("ec is not exist:%v", err)
		return false, response.ErrNotFound
	}

	if ec.Finished {
		logs.Debug("ec is finished [ %v ]", ecId)
		return false, response.ErrEcIsFinish
	}

	return true, response.Success
}

// @Title add candidates to election campaign
// @Description add Candidates to Election Campaigns
// @Param    candidate    {object}     CandidateAddParamT     true
// @Success 200 {object} response.ResponseT
// @router /:ecId/candidate [post]
func (v *VoteController) AddNewEcCandidate() {

	ecId, err := v.GetInt(":ecId", 0)
	if err != nil {
		logs.Error("ecId decode:%v", err)
		v.Data["json"] = response.ErrParamHandler(err.Error())
		v.ServeJSON()
		return
	}

	input := new(CandidateAddParamT)
	err = json.Unmarshal(v.Ctx.Input.RequestBody, &input)
	if err != nil {
		logs.Error("json decode:%v", err)
		v.Data["json"] = response.ErrParamHandler(err.Error())
		v.ServeJSON()
		return
	}

	valid, resp := checkEcVoteValid(ecId)
	if valid == false {
		v.Data["json"] = resp
		v.ServeJSON()
		return
	}

	args := make([]models.CandidateT, 0)
	for _, person := range input.Candidate {
		candidate := models.CandidateT{
			ECId:        ecId,
			Name:        person.Name,
			Sex:         person.Sex,
			Age:         person.Age,
			Description: person.Description,
		}
		args = append(args, candidate)
	}

	logs.Debug("add candidate:%v", args)
	err = models.NewMultiCandidate(args)
	if err != nil {
		logs.Error("json decode failed:%v", err)
		v.Data["json"] = response.ErrSystem
		v.ServeJSON()
		return
	}

	v.Data["json"] = response.Success
	v.ServeJSON()
	return
}

func checkIdentityNo(IdNo string) bool {

	if len(IdNo) != 10 {
		logs.Error("ident len not match:%v", len(IdNo))
		return false
	}

	if IdNo[0] < 'A' || IdNo[0] > 'Z' {
		logs.Error("first letter is invalid:%v", IdNo[0])
		return false
	}

	_, err := strconv.Atoi(IdNo[1:7])
	if err != nil {
		logs.Error("2-6 not all numbers:%v", IdNo)
		return false
	}

	if IdNo[7] != '(' || IdNo[9] != ')' {
		logs.Error("brackets not match:%v", IdNo[7:9])
		return false
	}

	if IdNo[8] < '0' || IdNo[8] > '9' {
		logs.Error("last number not match:%v", IdNo[7:9])
		return false
	}

	return true
}

func emailCheckIsValid(email string) bool {

	emailRegex := regexp.MustCompile(`^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`)

	if len(email) < 3 || len(email) > 250 {
		return false
	}
	return emailRegex.MatchString(email)
}

// @Title vote
// @Description vote in an Election Campaigns
// @Success 200 {object} response.ResponseT
// @router /:ecId/vote [post]
func (v *VoteController) VoteToCondidate() {

	logs.Debug("new vote:%v", string(v.Ctx.Input.RequestBody))
	ecId, err := v.GetInt(":ecId", 0)
	if err != nil {
		logs.Error("ecId decode:%v", err)
		v.Data["json"] = response.ErrParamHandler(err.Error())
		v.ServeJSON()
		return
	}

	input := new(VoteUserT)
	err = json.Unmarshal(v.Ctx.Input.RequestBody, &input)
	if err != nil {
		logs.Error("json decode:%v", err)
		v.Data["json"] = response.ErrParamHandler(err.Error())
		v.ServeJSON()
		return
	}

	ec, err := models.GetEcInfo(ecId)
	if err != nil {
		logs.Error("get info from db failed:%v", err)
		v.Data["json"] = response.ErrSystem
		v.ServeJSON()
		return
	}

	if ec.EcId == 0 {
		logs.Error("ec is not exist:%v", err)
		v.Data["json"] = response.ErrNotFound
		v.ServeJSON()
		return
	}

	if ec.StartTime.IsZero() {
		logs.Debug("ec is not start [ %v ]", ecId)
		v.Data["json"] = response.ErrEcIsFinish
		v.ServeJSON()
		return
	}

	if ec.Finished {
		logs.Debug("ec is finished [ %v ]", ecId)
		v.Data["json"] = response.ErrEcIsFinish
		v.ServeJSON()
		return
	}

	candidates, err := models.GetAllCandidate(ecId)
	if err != nil {
		if utils.ErrEmpty == err {
			logs.Debug("there is no candidate")
		} else {
			logs.Debug("get candidate failed")
		}
		v.Data["json"] = response.ErrSystem
		v.ServeJSON()
		return
	}

	//check vote candidate is valid
	candidateToVote := &models.CandidateT{CandidateId: -1}
	for _, candidate := range candidates {
		if candidate.CandidateId == input.CandidateId {
			candidateToVote = candidate
			break
		}
	}

	//candidate is invalid
	if candidateToVote.CandidateId == -1 {
		logs.Debug("candidate not in Election Campaigns[%v/%v]", input.CandidateId, ecId)
		v.Data["json"] = response.ErrVoteInvalid
		v.ServeJSON()
		return
	}

	if checkIdentityNo(input.IdentityNo) == false {

		logs.Debug("identity is invalid [ %v ]", input.IdentityNo)
		v.Data["json"] = response.ErrParamIdentityNoWrong
		v.ServeJSON()
		return
	}

	if emailCheckIsValid(input.Email) == false {
		logs.Debug("email is invalid [ %v ]", input.Email)
		v.Data["json"] = response.ErrParamEmailWrong
		v.ServeJSON()
		return
	}

	//get the user
	users, err := models.GetNormalUser(input.IdentityNo, input.Email)
	if err != nil {
		logs.Debug("get vote user failed")
		v.Data["json"] = response.ErrSystem
		v.ServeJSON()
		return
	}

	logs.Debug("get voters:%v", len(users))
	// maybe identno and email belong the deferent person
	if len(users) > 1 {
		logs.Debug("user email or iddent have used")
		v.Data["json"] = response.ErrIsExist
		v.ServeJSON()
		return
	}

	// new voted
	if len(users) == 0 {
		voter, err := models.AddNormalUser(input.Username, input.IdentityNo, input.Email)
		if err != nil {
			logs.Debug("add vote user failed")
			v.Data["json"] = response.ErrSystem
			v.ServeJSON()
			return
		}

		err = models.AddNewVote(ec, candidateToVote, voter)
		if err != nil {
			logs.Debug("vote failed :%v,voter:%v;candidate:%v ", ec.EcId, users[0].Username, candidateToVote.Name)
			v.Data["json"] = response.ErrSystem
			v.ServeJSON()
			return
		}

		detail, _ := getVoteDetail(ec)
		v.Data["json"] = response.SuccessHandle(detail)
		v.ServeJSON()
		return
	}

	//check voter have been voted
	count, err := models.GetVoteWithUser(ecId, users[0].UserId)
	if err != nil {
		logs.Debug("get count of voter[%v] failed", users[0].IdentityNo)
		v.Data["json"] = response.ErrSystem
		v.ServeJSON()
		return
	}

	//have voted
	if count != 0 {
		logs.Info("user [%v] have voted in this ec [%v]", users[0].IdentityNo, ecId)
		v.Data["json"] = response.ErrHaveVoted
		v.ServeJSON()
		return
	}

	//never voted
	err = models.AddNewVote(ec, candidateToVote, users[0])
	if err != nil {
		logs.Debug("vote failed :%v,voter:%v;candidate:%v ", ec.EcId, users[0].Username, candidateToVote.Name)
		v.Data["json"] = response.ErrSystem
		v.ServeJSON()
		return
	}

	detail, _ := getVoteDetail(ec)
	v.Data["json"] = response.SuccessHandle(detail)
	v.ServeJSON()

	return
}
