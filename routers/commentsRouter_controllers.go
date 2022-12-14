package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["votesystem/controllers:TaskController"] = append(beego.GlobalControllerRouter["votesystem/controllers:TaskController"],
        beego.ControllerComments{
            Method: "CreateNewTask",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["votesystem/controllers:VoteController"] = append(beego.GlobalControllerRouter["votesystem/controllers:VoteController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["votesystem/controllers:VoteController"] = append(beego.GlobalControllerRouter["votesystem/controllers:VoteController"],
        beego.ControllerComments{
            Method: "CreateNewEc",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["votesystem/controllers:VoteController"] = append(beego.GlobalControllerRouter["votesystem/controllers:VoteController"],
        beego.ControllerComments{
            Method: "AddNewEcCandidate",
            Router: "/:ecId/candidate",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["votesystem/controllers:VoteController"] = append(beego.GlobalControllerRouter["votesystem/controllers:VoteController"],
        beego.ControllerComments{
            Method: "GetVoters",
            Router: "/:ecId/condidate/:condidateId/votes/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["votesystem/controllers:VoteController"] = append(beego.GlobalControllerRouter["votesystem/controllers:VoteController"],
        beego.ControllerComments{
            Method: "FinishVote",
            Router: "/:ecId/finish",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["votesystem/controllers:VoteController"] = append(beego.GlobalControllerRouter["votesystem/controllers:VoteController"],
        beego.ControllerComments{
            Method: "StartVote",
            Router: "/:ecId/start",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["votesystem/controllers:VoteController"] = append(beego.GlobalControllerRouter["votesystem/controllers:VoteController"],
        beego.ControllerComments{
            Method: "VoteToCondidate",
            Router: "/:ecId/vote",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
