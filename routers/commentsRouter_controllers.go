package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["bbb-api-meetings/controllers:MeetingController"] = append(beego.GlobalControllerRouter["bbb-api-meetings/controllers:MeetingController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bbb-api-meetings/controllers:MeetingController"] = append(beego.GlobalControllerRouter["bbb-api-meetings/controllers:MeetingController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bbb-api-meetings/controllers:MeetingController"] = append(beego.GlobalControllerRouter["bbb-api-meetings/controllers:MeetingController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:meetingId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bbb-api-meetings/controllers:MeetingController"] = append(beego.GlobalControllerRouter["bbb-api-meetings/controllers:MeetingController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:meetingId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bbb-api-meetings/controllers:MeetingController"] = append(beego.GlobalControllerRouter["bbb-api-meetings/controllers:MeetingController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:meetingId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bbb-api-meetings/controllers:ObjectController"] = append(beego.GlobalControllerRouter["bbb-api-meetings/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bbb-api-meetings/controllers:ObjectController"] = append(beego.GlobalControllerRouter["bbb-api-meetings/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bbb-api-meetings/controllers:ObjectController"] = append(beego.GlobalControllerRouter["bbb-api-meetings/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bbb-api-meetings/controllers:ObjectController"] = append(beego.GlobalControllerRouter["bbb-api-meetings/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bbb-api-meetings/controllers:ObjectController"] = append(beego.GlobalControllerRouter["bbb-api-meetings/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bbb-api-meetings/controllers:UserController"] = append(beego.GlobalControllerRouter["bbb-api-meetings/controllers:UserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bbb-api-meetings/controllers:UserController"] = append(beego.GlobalControllerRouter["bbb-api-meetings/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bbb-api-meetings/controllers:UserController"] = append(beego.GlobalControllerRouter["bbb-api-meetings/controllers:UserController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bbb-api-meetings/controllers:UserController"] = append(beego.GlobalControllerRouter["bbb-api-meetings/controllers:UserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bbb-api-meetings/controllers:UserController"] = append(beego.GlobalControllerRouter["bbb-api-meetings/controllers:UserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bbb-api-meetings/controllers:UserController"] = append(beego.GlobalControllerRouter["bbb-api-meetings/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bbb-api-meetings/controllers:UserController"] = append(beego.GlobalControllerRouter["bbb-api-meetings/controllers:UserController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
