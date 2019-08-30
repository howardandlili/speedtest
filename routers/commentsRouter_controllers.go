package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["speedtest/controllers:OrdersController"] = append(beego.GlobalControllerRouter["speedtest/controllers:OrdersController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["speedtest/controllers:OrdersController"] = append(beego.GlobalControllerRouter["speedtest/controllers:OrdersController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["speedtest/controllers:OrdersController"] = append(beego.GlobalControllerRouter["speedtest/controllers:OrdersController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["speedtest/controllers:OrdersController"] = append(beego.GlobalControllerRouter["speedtest/controllers:OrdersController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["speedtest/controllers:OrdersController"] = append(beego.GlobalControllerRouter["speedtest/controllers:OrdersController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["speedtest/controllers:UsersController"] = append(beego.GlobalControllerRouter["speedtest/controllers:UsersController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/user`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["speedtest/controllers:UsersController"] = append(beego.GlobalControllerRouter["speedtest/controllers:UsersController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/user`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["speedtest/controllers:UsersController"] = append(beego.GlobalControllerRouter["speedtest/controllers:UsersController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/user/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["speedtest/controllers:UsersController"] = append(beego.GlobalControllerRouter["speedtest/controllers:UsersController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/user/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["speedtest/controllers:UsersController"] = append(beego.GlobalControllerRouter["speedtest/controllers:UsersController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/user/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}