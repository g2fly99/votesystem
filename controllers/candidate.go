package controllers

import (
	"strconv"
	"votesystem/models"
	"votesystem/response"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

type CandidateT struct {
	CandidateId int `json:"candidateId"`
	EcId        string
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

// @Title Get candidate's votes
// @Description get all Election Campaigns
// @Success 200 {object} VoteUserT
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
