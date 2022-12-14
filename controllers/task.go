package controllers

import (
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

// Operations about VoteController
type TaskController struct {
	beego.Controller
}

// @Title create election campaign
// @Description get all Election Campaigns
// @Param    expire    query     string     true     "expire time eg:202212042008"
// @Param    description    query     string     false     "description of the election campaign"
// @Success 200 {object} response.ResponseT
// @router / [post]
func (v *TaskController) CreateNewTask() {
	logs.Debug("new Task:%v")

}
