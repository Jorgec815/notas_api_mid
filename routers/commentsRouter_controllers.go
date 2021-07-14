package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init(){

	beego.GlobalControllerRouter["github.com/jorgec815/notas_api_mid/controllers:EstudianteController"] = append(beego.GlobalControllerRouter["github.com/jorgec815/notas_api_crud/controllers:EstudianteController"],
        beego.ControllerComments{
            Method: "CalcularDefinitiva",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})
}