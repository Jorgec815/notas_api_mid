package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/jorgec815/notas_api_mid/models"
)

// EstudianteController operations for Estudiante
type EstudianteController struct {
	beego.Controller
}

// URLMapping ...
func (c *EstudianteController) URLMapping() {
	c.Mapping("CalcularDefinitiva", c.CalcularDefinitiva)
}

// Post ...
// @Title Create
// @Description create Estudiante
// @Param	body		body 	models.Estudiante	true		"body for Estudiante content"
// @Success 201 {object} models.Estudiante
// @Failure 403 body is empty
// @router / [post]
func (c *EstudianteController) Post() {

}

// GetOne ...
// @Title GetOne
// @Description get Estudiante by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Estudiante
// @Failure 403 :id is empty
// @router /:id [get]
func (c *EstudianteController) GetOne() {

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

func (c *EstudianteController) CalcularDefinitiva() (estudiante []models.Estudiante){
	idStr := c.Ctx.Input.Param(":id")
	if err := request.GetJson(beego.AppConfig.String("UrlCrud")+"/estudiante/:"+idStr, &estudiante); err == nil{
		fmt.Println("Retornando estudiante")
	}else{
		fmt.Println("error en Id")
	}
	fmt.Println("estudiante" + estudiante)
	return
}
