package controllers

import (
	"encoding/json"
	"strconv"
	"votesystem/models"
	"votesystem/response"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

type CandidateT struct {
	CandidateId int    `json:"candidateId"`
	Name        string `json:"name"`
	Sex         int    `json:"sex"`
	Age         int    `json:"age"`
	Description string `json:"description"`
	VoteNumber  int    `json:"voteNumber"`
}

// Operations about CondidateController
type CondidateController struct {
	beego.Controller
}

type VoteUserT struct {
	Username    string `json:"username"`
	Email       string `json:"email"`
	IdentityNo  string `json:"identityNo"`
	CandidateId int    `json:"candidateId"`
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

// @Title Get candidate's votes
// @Description get all Election Campaigns
// @Success 200 {object} controllers.VoteUserT
// @router /:ecId/condidate/:condidateId/votes/ [get]
func (c *VoteController) GetVoters() {

	logs.Debug("list votes:%v", c.Ctx.Input.URI())
	condidateIdArg := c.Ctx.Input.Param(":condidateId")
	condidateId, err := strconv.Atoi(condidateIdArg)
	if err != nil {
		logs.Error("condidateId decode:%v", err)
		c.Data["json"] = response.ErrParamHandler("condidateId")
		c.ServeJSON()
		return
	}

	ec := c.Ctx.Input.Param(":ecId")
	ecId, err := strconv.Atoi(ec)
	if err != nil {
		logs.Error("ecId decode:%v", err)
		c.Data["json"] = response.ErrParamHandler(err.Error())
		c.ServeJSON()
		return
	}

	offset, err := c.GetInt("offset", 0)
	limit, err := c.GetInt("limit", 10)

	logs.Debug("get limit:%v,offset:%v", limit, offset)
	result := make([]VoteUserT, 0)
	voters, err := models.ListAllVote(ecId, condidateId, limit, offset)
	for _, user := range voters {

		u := VoteUserT{
			CandidateId: condidateId,
			Username:    user.Username,
			Email:       user.Email,
			IdentityNo:  user.IdentityNo,
		}
		result = append(result, u)
	}

	c.Data["json"] = response.SuccessHandle(result)
	c.ServeJSON()
	return
}
