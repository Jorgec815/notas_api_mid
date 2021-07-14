package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/jorgec815/notas_api_mid/controllers:EstudianteController"] = append(beego.GlobalControllerRouter["github.com/jorgec815/notas_api_mid/controllers:EstudianteController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/jorgec815/notas_api_mid/controllers:EstudianteController"] = append(beego.GlobalControllerRouter["github.com/jorgec815/notas_api_mid/controllers:EstudianteController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/jorgec815/notas_api_mid/controllers:EstudianteController"] = append(beego.GlobalControllerRouter["github.com/jorgec815/notas_api_mid/controllers:EstudianteController"],
        beego.ControllerComments{
            Method: "CalcularDefinitiva",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/jorgec815/notas_api_mid/controllers:EstudianteController"] = append(beego.GlobalControllerRouter["github.com/jorgec815/notas_api_mid/controllers:EstudianteController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/jorgec815/notas_api_mid/controllers:EstudianteController"] = append(beego.GlobalControllerRouter["github.com/jorgec815/notas_api_mid/controllers:EstudianteController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
