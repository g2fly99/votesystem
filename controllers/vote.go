package controllers

import (
	"encoding/json"
	"errors"
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

// @Title GetAll Election Campaign
// @Description get all Election Campaigns
// @Success 200 {object} models.ElectionCampaignT
// @router / [get]
func (v *VoteController) GetAll() {

	logs.Debug("get all election campaign:%v", v.Ctx.Request.URL.String())

	offset, err := v.GetInt("offset", 0)
	if err != nil {
		logs.Error("get ecid wrong:%v", err)
		v.Data["json"] = response.ErrParamWrong
		v.ServeJSON()
		return
	}

	limit, err := v.GetInt("limit", 10)
	if err != nil {
		logs.Error("get ecid wrong:%v", err)
		v.Data["json"] = response.ErrParamWrong
		v.ServeJSON()
		return
	}

	ecs, err := models.ListEcs(limit, offset)
	if err != nil {
		if errors.Is(err, utils.ErrDbErr) {
			v.Data["json"] = response.ErrSystem
			v.ServeJSON()
			return
		}

		v.Data["json"] = response.ErrSystem
		v.ServeJSON()
		return
	}

	result := []VoteDetailT{}
	for _, ec := range ecs {

		tmpEc := VoteDetailT{
			EcId:        ec.EcId,
			FinishTime:  ec.FinishTime,
			Finished:    ec.Finished,
			Expire:      ec.Expire,
			StartTime:   ec.StartTime,
			Description: ec.Description,
		}

		candidates, err := models.GetAllCandidate(tmpEc.EcId)
		if err != nil {

		}

		cdts := []*CandidateT{}
		for _, candidate := range candidates {

			//get vote number
			number, _ := models.CountVotes(tmpEc.EcId, candidate.CandidateId)
			cdt := &CandidateT{
				CandidateId: candidate.CandidateId,
				Name:        candidate.Name,
				Sexy:        candidate.Sexy,
				Age:         candidate.Age,
				Description: candidate.Description,
				VoteNumber:  number,
			}

			cdts = append(cdts, cdt)
		}

		tmpEc.Candidates = cdts

		result = append(result, tmpEc)
	}

	v.Data["json"] = response.SuccessHandle(result)
	v.ServeJSON()

	return
}

// @Title create election campaign
// @Description get all Election Campaigns
// @Success 200 {object} models.ElectionCampaignT
// @router / [post]
func (v *VoteController) CreateNewEc() {

	logs.Debug("create an election candidate:%v", v.Ctx.Request.URL.String())
	//expair time.Time, description string

	description := v.GetString("description")

	expire := v.GetString("expire")

	expireTime, err := time.Parse("20060102150405", expire)
	if err != nil {
		logs.Error("time format not match:%v", err)
		v.Data["json"] = utils.ErrInvalid
		v.ServeJSON()
		return
	}

	newEc, err := models.CreateNewEC(expireTime, description)
	if err != nil {
		logs.Error("create new election campaign failed:%v", err)
		v.Data["json"] = utils.ErrDbErr
		v.ServeJSON()
		return
	}

	//save to cache
	newVoteDetail := VoteDetailT{
		EcId:        newEc.EcId,
		FinishTime:  newEc.FinishTime,
		Expire:      newEc.Expire,
		StartTime:   newEc.StartTime,
		CreateTime:  newEc.Created,
		Description: newEc.Description,
	}

	v.Data["json"] = response.SuccessHandle(newVoteDetail)
	v.ServeJSON()
	return
}

type CandidateAddParamT struct {
	Candidate []CandidateT `json:"candidate"`
}

// @Title add candidates to election campaign
// @Description add Candidates to Election Campaigns
// @Success 200 {object} response.ResponseT
// @router /{ecId}/candidate [post]
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

	ec, err := models.GetEcInfo(ecId)
	if err != nil {
		logs.Error("get info from db failed:%v", err)
		v.Data["json"] = response.ErrSystem
		v.ServeJSON()
		return
	}

	if ec.EcId == 0 {
		logs.Error("ec is not exist:%v", err)
		v.Data["json"] = response.ErrEcIsFinish
		v.ServeJSON()
		return
	}

	if ec.Finished {
		logs.Debug("ec is finished [ %v ]", ecId)
		v.Data["json"] = response.ErrParamHandler(err.Error())
		v.ServeJSON()
		return
	}

	args := make([]models.CandidateT, 0)
	for _, person := range input.Candidate {
		candidate := models.CandidateT{
			EC:          ec,
			Name:        person.Name,
			Sexy:        person.Sexy,
			Age:         person.Age,
			Description: person.Description,
		}
		args = append(args, candidate)
	}

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
