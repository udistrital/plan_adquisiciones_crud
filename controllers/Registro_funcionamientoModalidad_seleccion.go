package controllers

import (
	"plan_cuentas_crud/models"
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

// RegistroFuncionamientoModalidadSeleccionController operations for RegistroFuncionamientoModalidadSeleccion
type RegistroFuncionamientoModalidadSeleccionController struct {
	beego.Controller
}

// URLMapping ...
func (c *RegistroFuncionamientoModalidadSeleccionController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create RegistroFuncionamientoModalidadSeleccion
// @Param	body		body 	models.RegistroFuncionamientoModalidadSeleccion	true		"body for RegistroFuncionamientoModalidadSeleccion content"
// @Success 201 {int} models.RegistroFuncionamientoModalidadSeleccion
// @Failure 400 the request contains incorrect syntax
// @router / [post]
func (c *RegistroFuncionamientoModalidadSeleccionController) Post() {
	var v models.RegistroFuncionamientoModalidadSeleccion
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddRegistroFuncionamientoModalidadSeleccion(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
logs.Error(err)
 //c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
c.Data["system"] = err
 c.Abort("400")
		}
	} else {
logs.Error(err)
 //c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
c.Data["system"] = err
 c.Abort("400")
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get RegistroFuncionamientoModalidadSeleccion by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.RegistroFuncionamientoModalidadSeleccion
// @Failure 404 not found resource
// @router /:id [get]
func (c *RegistroFuncionamientoModalidadSeleccionController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetRegistroFuncionamientoModalidadSeleccionById(id)
	if err != nil {
logs.Error(err)
 //c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
c.Data["system"] = err
 c.Abort("404")
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get RegistroFuncionamientoModalidadSeleccion
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.RegistroFuncionamientoModalidadSeleccion
// @Failure 404 not found resource
// @router / [get]
func (c *RegistroFuncionamientoModalidadSeleccionController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllRegistroFuncionamientoModalidadSeleccion(query, fields, sortby, order, offset, limit)
	if err != nil {
logs.Error(err)
 //c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
c.Data["system"] = err
 c.Abort("404")
	} else {
if l == nil {
 l = append(l, map[string]interface{}{})
 }
 c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the RegistroFuncionamientoModalidadSeleccion
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.RegistroFuncionamientoModalidadSeleccion	true		"body for RegistroFuncionamientoModalidadSeleccion content"
// @Success 200 {object} models.RegistroFuncionamientoModalidadSeleccion
// @Failure 400 the request contains incorrect syntax
// @router /:id [put]
func (c *RegistroFuncionamientoModalidadSeleccionController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.RegistroFuncionamientoModalidadSeleccion{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateRegistroFuncionamientoModalidadSeleccionById(&v); err == nil {
c.Data["json"] = v
		} else {
logs.Error(err)
 //c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
c.Data["system"] = err
 c.Abort("400")
		}
	} else {
logs.Error(err)
 //c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
c.Data["system"] = err
 c.Abort("400")
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the RegistroFuncionamientoModalidadSeleccion
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 404 not found resource
// @router /:id [delete]
func (c *RegistroFuncionamientoModalidadSeleccionController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteRegistroFuncionamientoModalidadSeleccion(id); err == nil {
c.Data["json"] = map[string]interface{}{"Id": id}
	} else {
logs.Error(err)
 //c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
c.Data["system"] = err
 c.Abort("404")
	}
	c.ServeJSON()
}
