package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["votesystem/controllers:CondidateController"] = append(beego.GlobalControllerRouter["votesystem/controllers:CondidateController"],
        beego.ControllerComments{
            Method: "GetVoters",
            Router: "/:condidateId/votes/",
            AllowHTTPMethods: []string{"get"},
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
            Router: "/{ecId}/candidate",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
