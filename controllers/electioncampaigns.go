package controllers

import (
	"encoding/json"
	"errors"
	"time"
	"votesystem/cache"
	"votesystem/models"
	"votesystem/response"
	"votesystem/utils"

	"github.com/beego/beego/v2/core/logs"
)

const (
	gECNeedCandidateNumberMin = 2
	gECdetectFrequency        = 1 // minute
)

var (
	gDetectRoutineRun = false
)

// @Title start an election campaigns
// @Description start in an Election Campaigns
// @Success 200 {object} response.ResponseT
// @router /:ecId/start [put]
func (v *VoteController) StartVote() {

	ecId, err := v.GetInt(":ecId", 0)
	if err != nil {
		logs.Error("ecId decode:%v", err)
		v.Data["json"] = response.ErrParamHandler(err.Error())
		v.ServeJSON()
		return
	}

	res, err := models.GetAllCandidate(ecId)
	if err != nil {
		logs.Debug("get candidate failed")
		v.Data["json"] = response.ErrSystem
		v.ServeJSON()
		return
	}

	if len(res) >= gECNeedCandidateNumberMin {

		err = models.ECStart(ecId)
		if err != nil {
			logs.Debug("get candidate failed")
			v.Data["json"] = response.ErrSystem
		} else {
			v.Data["json"] = response.Success
		}

		v.ServeJSON()
	} else {
		v.Data["json"] = response.ErrCancdidateNotEnough
		v.ServeJSON()
	}

	detectEcIsFinishedRoutine()

	return
}

// @Title finish an election campaigns
// @Description finish an Election Campaigns
// @Success 200 {object} response.ResponseT
// @router /:ecId/finish [put]
func (v *VoteController) FinishVote() {

	ecId, err := v.GetInt(":ecId", 0)
	if err != nil {
		logs.Error("ecId decode:%v", err)
		v.Data["json"] = response.ErrParamHandler(err.Error())
		v.ServeJSON()
		return
	}

	err = models.ECSetFinish(ecId)
	if err != nil {
		logs.Debug("get candidate failed")
		v.Data["json"] = response.ErrSystem
	} else {
		v.Data["json"] = response.Success
	}

	v.ServeJSON()

	//send email to voters in background
	go finishToSendEmail(ecId)
	return
}

func getVoteDetail(ec *models.ElectionCampaignT) (VoteDetailT, error) {

	res := VoteDetailT{
		EcId:        ec.EcId,
		FinishTime:  ec.FinishTime,
		Finished:    ec.Finished,
		Expire:      ec.Expire,
		StartTime:   ec.StartTime,
		Description: ec.Description,
	}

	candidates, err := models.ListCandidate(res.EcId, 1000, 0)
	if err != nil {

	}

	cdts := []*CandidateT{}
	for _, candidate := range candidates {

		//get vote number
		number, _ := models.CountVotes(res.EcId, candidate.CandidateId)
		cdt := &CandidateT{
			CandidateId: candidate.CandidateId,
			Name:        candidate.Name,
			Sex:         candidate.Sex,
			Age:         candidate.Age,
			Description: candidate.Description,
			VoteNumber:  number,
		}

		cdts = append(cdts, cdt)
	}

	res.Candidates = cdts
	return res, nil
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

	logs.Debug("get ecs number:%v", len(ecs))
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

		candidates, err := models.ListCandidate(tmpEc.EcId, 1000, 0)
		if err != nil {

		}

		cdts := []*CandidateT{}
		for _, candidate := range candidates {

			//get vote number
			number, _ := models.CountVotes(tmpEc.EcId, candidate.CandidateId)
			cdt := &CandidateT{
				CandidateId: candidate.CandidateId,
				Name:        candidate.Name,
				Sex:         candidate.Sex,
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

	expireTime, err := time.Parse("200601021504", expire)
	if err != nil {
		logs.Error("time format not match:%v", err)
		v.Data["json"] = response.ErrParamWrong
		v.ServeJSON()
		return
	}

	newEc, err := models.CreateNewEC(expireTime, description)
	if err != nil {
		logs.Error("create new election campaign failed:%v", err)
		v.Data["json"] = response.ErrSystem
		v.ServeJSON()
		return
	}

	newVoteDetail := VoteDetailT{
		EcId:        newEc.EcId,
		FinishTime:  newEc.FinishTime,
		Expire:      newEc.Expire,
		StartTime:   newEc.StartTime,
		CreateTime:  newEc.Created,
		Description: newEc.Description,
	}
	data, _ := json.Marshal(newVoteDetail)

	//save to cache
	err = cache.Save(newEc, string(data))
	if err != nil {
		logs.Error("save to cache failed:%v", err)
	}

	v.Data["json"] = response.SuccessHandle(newVoteDetail)
	v.ServeJSON()
	return
}

func finishToSendEmail(ecId int) {

	for i := 0; ; i++ {
		users, err := models.ListAllEcVoters(ecId, 1000, 1000*i)
		if err != nil {
			logs.Error("get all candidate failed:%v", err)
			return
		}
		if len(users) == 0 {
			return
		}

		for _, u := range users {

			logs.Debug("send result of election campaigns to voters:%v,email:%v", u.Username, u.Email)
		}
	}
}

func detectEcIsFinishedRoutine() {

	if gDetectRoutineRun {
		return
	} else {
		gDetectRoutineRun = true
	}

	go func() {
		for i := 0; ; i++ {
			ecs, err := models.ListActiveEcs(100, 100*i)
			if err != nil {
				logs.Error("get all candidate failed:%v", err)
				time.Sleep(1 * time.Minute)
			}

			now := time.Now()
			for _, ec := range ecs {
				if now.After(ec.Expire) {
					models.ECSetFinish(ec.EcId)

					//send email to voters in background
					go finishToSendEmail(ec.EcId)
				}
			}

			time.Sleep(gECdetectFrequency * time.Minute)
		}
	}()
}
