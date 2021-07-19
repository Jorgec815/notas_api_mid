package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/utils_oas/request"
	"github.com/jorgec815/notas_api_mid/models"
	//"github.com/jorgec815/notas_api_mid/helpers"
)

// EstudianteController operations for Estudiante
type EstudianteController struct {
	beego.Controller
}

// URLMapping ...
func (c *EstudianteController) URLMapping() {
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("Post", c.Post)
	c.Mapping("Put", c.Put)
}

// Post ...
// @Title Create
// @Description create Estudiante
// @Param	body		body 	models.Estudiante	true		"body for Estudiante content"
// @Success 201 {object} models.Estudiante
// @Failure 403 body is empty
// @router / [post]
func (c *EstudianteController) Post() {

	var estudiante models.Estudiante			//variable de tipo Estudiante
	var env map[string]interface{}				//interface que se va enviar
	var res map[string]interface{}				//interface que se va recibir como respuesta


	json.Unmarshal(c.Ctx.Input.RequestBody, &env) //se coloca el body dentro del interface
	jsonString, _:= json.Marshal(env)				//se pasa el interface a un string
	json.Unmarshal(jsonString, &estudiante)			//se pasa el string a una estructura de tipo estudiante
	
	
	estudiante.NotaDef = calcularDefinitiva(estudiante.Nota1, estudiante.Nota2, estudiante.Nota3) //se realiza el cálculo de la nota
	e,  _:= json.Marshal(estudiante); 		//se usar una variable auxiliar para pasar de una estructura a un []byte
	json.Unmarshal(e, &env)			//se pasa el []byte a la interface que se va enviar

	//se realiza la petición al CRUD

	if err:= request.SendJson(beego.AppConfig.String("UrlCrud")+"/estudiante", "POST", &res, env); err == nil{

		c.Data["json"] = res
	}else{
		logs.Error(err)
		c.Abort("404")
	}	

		
	c.ServeJSON()

}

// GetOne ...
// @Title GetOne
// @Description get Estudiante by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Estudiante
// @Failure 403 :id is empty
// @router /:id [get]
func (c *EstudianteController) GetOne() {
	IdEstudiante := c.Ctx.Input.Param(":id")

	var res map[string]interface{}

	if err := request.GetJson(beego.AppConfig.String("UrlCrud")+"/estudiante/"+IdEstudiante, &res); err == nil{
		c.Data["json"] = res
	}else{
		logs.Error(err)
		c.Abort("404")
	}
	c.ServeJSON()
}

// GetAll ...
// @Title GetAll
// @Description get Estudiante
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Estudiante
// @Failure 403
// @router / [get]
func (c *EstudianteController) GetAll() {

}

// Put ...
// @Title Put
// @Description update the Estudiante
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Estudiante	true		"body for Estudiante content"
// @Success 200 {object} models.Estudiante
// @Failure 403 :id is not int
// @router /:id [put]
func (c *EstudianteController) Put() {
	IdEstudiante := c.Ctx.Input.Param(":id")

	var estudiante models.Estudiante			//variable de tipo Estudiante
	var env map[string]interface{}				//interface que se va enviar
	var res map[string]interface{}				//interface que se va recibir como respuesta

	json.Unmarshal(c.Ctx.Input.RequestBody, &env) //se coloca el body dentro del interface
	jsonString, _:= json.Marshal(env)				//se pasa el interface a un string
	json.Unmarshal(jsonString, &estudiante)			//se pasa el string a una estructura de tipo estudiante
	
	estudiante.NotaDef = calcularDefinitiva(estudiante.Nota1, estudiante.Nota2, estudiante.Nota3) //se realiza el cálculo de la nota
	e,  _:= json.Marshal(estudiante); 		//se usar una variable auxiliar para pasar de una estructura a un []byte
	json.Unmarshal(e, &env)			//se pasa el []byte a la interface que se va enviar

	if err:= request.SendJson(beego.AppConfig.String("UrlCrud")+"/estudiante/"+IdEstudiante, "PUT", &res, env); err == nil{

		c.Data["json"] = res
	}else{
		logs.Error(err)
		c.Abort("404")
	}	

		
	c.ServeJSON()

}

// Delete ...
// @Title Delete
// @Description delete the Estudiante
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *EstudianteController) Delete() {

}

func calcularDefinitiva (nota1 float64, nota2 float64, nota3 float64) float64{
	return ((nota1*0.35)+(nota2*0.35)+(nota3*0.3))
}
