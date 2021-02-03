// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/udistrital/plan_adquisiciones_crud/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/Meta",
			beego.NSInclude(
				&controllers.MetaController{},
			),
		),

		beego.NSNamespace("/Actividad",
			beego.NSInclude(
				&controllers.ActividadController{},
			),
		),

		beego.NSNamespace("/Plan_adquisiciones",
			beego.NSInclude(
				&controllers.PlanAdquisicionesController{},
			),
		),

		beego.NSNamespace("/Plan_adquisiciones_mongo",
			beego.NSInclude(
				&controllers.PlanAdquisicionesMongoController{},
			),
		),

		beego.NSNamespace("/Registro_plan_adquisiciones",
			beego.NSInclude(
				&controllers.RegistroPlanAdquisicionesController{},
			),
		),

		beego.NSNamespace("/Registro_plan_adquisiciones-Actividad",
			beego.NSInclude(
				&controllers.RegistroPlanAdquisicionesActividadController{},
			),
		),

		beego.NSNamespace("/Registro_inversion_actividad-Fuente_financiamiento",
			beego.NSInclude(
				&controllers.RegistroInversionActividadFuenteFinanciamientoController{},
			),
		),

		beego.NSNamespace("/Ficha_EB_IMGA",
			beego.NSInclude(
				&controllers.FichaEBIMGAController{},
			),
		),

		beego.NSNamespace("/Registro_plan_adquisiciones-Codigo_arka",
			beego.NSInclude(
				&controllers.RegistroPlanAdquisicionesCodigoArkaController{},
			),
		),

		beego.NSNamespace("/Registro_funcionamiento-Modalidad_seleccion",
			beego.NSInclude(
				&controllers.RegistroFuncionamientoModalidadSeleccionController{},
			),
		),

		beego.NSNamespace("/Lineamiento",
			beego.NSInclude(
				&controllers.LineamientoController{},
			),
		),
		
		beego.NSNamespace("/Registro_plan_adquisiciones-Metas_Asociadas",
			beego.NSInclude(
				&controllers.RegistroPlanAdquisicionesMetasAsociadasController{},
			),
		),

		beego.NSNamespace("/Registro_plan_adquisiciones-Productos_Asociados",
			beego.NSInclude(
				&controllers.RegistroPlanAdquisicionesProductosAsociadosController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
