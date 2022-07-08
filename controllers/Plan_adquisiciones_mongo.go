package controllers

import (
	"encoding/json"

	"github.com/udistrital/plan_adquisiciones_crud/helpers"
	"github.com/udistrital/plan_adquisiciones_crud/models"

	"github.com/udistrital/utils_oas/formatdata"
	"github.com/udistrital/utils_oas/time_bogota"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"go.mongodb.org/mongo-driver/bson"
)

// PlanAdquisicionesMongoController operations for PlanAdquisicionesMongo
type PlanAdquisicionesMongoController struct {
	beego.Controller
}

// URLMapping ...
func (c *PlanAdquisicionesMongoController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
}

// Post ...
// @Title Post
// @Description create PlanAdquisicionesMongo
// @Param	body		body 	models.PlanAdquisicionesMongo	true		"body for PlanAdquisicionesMongo content"
// @Success 201 {object} models.PlanAdquisicionesMongo
// @Failure 400 the request contains incorrect syntax
// @router / [post]
func (c *PlanAdquisicionesMongoController) Post() {
	var v models.PlanAdquisicionesMongo
	var data models.PlanAdquisicionesMongoID
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		v.FechaCreacion = time_bogota.TiempoBogotaFormato()
		v.FechaModificacion = time_bogota.TiempoBogotaFormato()
		if idInserted, err := models.AddPlanAdquisicionesMongo(&v); err == nil {
			data.IdMongo = idInserted
			data.PlanAdquisiconesMongo = v
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = data
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
// @Description get PlanAdquisicionesMongo by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.PlanAdquisicionesMongo
// @Failure 404 not found resource
// @router /:id [get]
func (c *PlanAdquisicionesMongoController) GetOne() {
	id := c.Ctx.Input.Param(":id")
	v, err := models.GetPlanAdquisicionesMongoById(id)
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
// @Description get PlanAdquisicionesMongo
// @Param	query	query	string	false	"Filter. e.g. {"vigencia":2020}"
// @Success 200 {object} []models.PlanAdquisicionesMongo
// @Failure 404 not found resource
// @router / [get]
func (c *PlanAdquisicionesMongoController) GetAll() {
	var query bson.M
	query = nil
	if v := c.GetString("query"); v != "" {
		err := json.Unmarshal([]byte(v), &query)
		if err != nil {
			logs.Error("json. Unmarshal() ERROR (query):", err)
			c.Data["system"] = err
			c.Abort("400")
		}
	}
	l, err := models.GetAllPlanAdquisicionesMongo(query)
	if err != nil {
		logs.Error(err)
		c.Data["system"] = err
		c.Abort("404")
	} else {
		if l == nil {
			l = []interface{}{}
		}
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// GetDiferencia ...
// @Title Get Diferencia de Planes
// @Description Retorna la diferencia entre dos planes publicados
// @Param	idPlanPublicado	path string	true	"Id de la versión del plan de adquisiciones"
// @Success 200 {object} []models.PlanAdquisicionesMongo
// @Failure 404 not found resource
// @router /diferencia/:idPlanPublicado [get]
func (c *PlanAdquisicionesMongoController) GetDiferencia() {
	idPlanPublicado := c.Ctx.Input.Param(":idPlanPublicado")
	var planPublicado models.PlanPublicado
	var planAntiguo models.PlanPublicado
	v2, err := models.GetPlanAdquisicionesMongoById(idPlanPublicado)
	var movimientos []models.MovimientosDetalle

	if err := formatdata.FillStruct(*v2, &planPublicado); err != nil {
		logs.Error(err)
		c.Data["system"] = err
		c.Abort("502")
	}

	// Query to retrieve registers by Plan Adquisiciones ID
	query := bson.M{"id": planPublicado.IdPlan}
	// Sorting elements by creation date descendent (-1)
	sort := map[string]interface{}{"fechacreacion": -1}
	// Limit consult results (2)
	limit := int64(2)

	l, err := models.GetFilterPlanAdquisicionesMongo(query, sort, limit)
	if len(l) == 0 {
		logs.Error(err)
		c.Data["system"] = err
		c.Abort("404")
	}

	if len(l) > 1 {
		if err := formatdata.FillStruct(l[0], &planPublicado); err != nil {
			logs.Error(err)
			c.Data["system"] = err
			c.Abort("502")
		}
		if err := formatdata.FillStruct(l[1], &planAntiguo); err != nil {
			logs.Error(err)
			c.Data["system"] = err
			c.Abort("502")
		}
	} else {
		if err := formatdata.FillStruct(l[0], &planPublicado); err != nil {
			logs.Error(err)
			c.Data["system"] = err
			c.Abort("502")
		}
	}

	if len(l) == 1 {
		movimientos, err = helpers.PublicarPlan(planPublicado)
	}

	if len(l) == 2 {
		movimientos, err = helpers.DevuelveMovimientos(planPublicado, planAntiguo)
	}

	if err != nil {
		logs.Error(err)
		c.Data["system"] = err
		c.Abort("404")
	} else {
		c.Data["json"] = movimientos
	}
	c.ServeJSON()
}
