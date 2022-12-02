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
	Sexy        int    `json:"sexy"`
	Age         int    `json:"age"`
	Description string `json:"description"`
	VoteNumber  int    `json:"voteNumber"`
}

// Operations about CondidateController
type CondidateController struct {
	beego.Controller
}

type VoteUserT struct {
	Username   string `json:"username"`
	Email      string `json:"email"`
	IdentityNo string `json:"identityNo"`
}

// @Title Get candidate's votes
// @Description get all Election Campaigns
// @Success 200 {object} VoteUserT
// @router /:condidateId/votes/ [get]
func (c *CondidateController) GetVoters() {

	condidateIdArg := c.Ctx.Input.Param(":condidateId")
	condidateId, err := strconv.Atoi(condidateIdArg)
	if err != nil {
		logs.Error("condidateId decode:%v", err)
		c.Data["json"] = response.ErrParamHandler("condidateId")
		c.ServeJSON()
		return
	}

	ecId, err := c.GetInt("ecId", 0)
	if err != nil {
		logs.Error("ecId decode:%v", err)
		c.Data["json"] = response.ErrParamHandler(err.Error())
		c.ServeJSON()
		return
	}

	/*ec, err := models.GetEcInfo(ecId)
	if err != nil {
		logs.Error("get info from db failed:%v", err)
		c.Data["json"] = response.ErrSystem
		c.ServeJSON()
		return
	}*/

	offset, err := c.GetInt(":offset", 0)
	limit, err := c.GetInt(":limit", 10)

	result := make([]VoteUserT, 0)
	voters, err := models.ListAllVote(ecId, condidateId, limit, offset)
	for _, user := range voters {

		u := VoteUserT{
			Username:   user.Username,
			Email:      user.Email,
			IdentityNo: user.IdentityNo,
		}
		result = append(result, u)
	}

	c.Data["json"] = response.SuccessHandle(result)
	c.ServeJSON()
	return
}
