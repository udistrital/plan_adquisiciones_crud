// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"plan_cuentas_crud/controllers"

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

		beego.NSNamespace("/Registro_plan_adquisiciones",
			beego.NSInclude(
				&controllers.RegistroPlanAdquisicionesController{},
			),
		),

		beego.NSNamespace("/Registro_plan_adquisiciones-Actividad",
			beego.NSInclude(
				&controllers.RegistroPlanAdquisiciones-ActividadController{},
			),
		),

		beego.NSNamespace("/Registro_inversion_actividad-Fuente_financiamiento",
			beego.NSInclude(
				&controllers.RegistroInversionActividad-FuenteFinanciamientoController{},
			),
		),

		beego.NSNamespace("/Ficha_EB_IMGA",
			beego.NSInclude(
				&controllers.FichaEBIMGAController{},
			),
		),

		beego.NSNamespace("/Registro_plan_adquisiciones-Codigo_arka",
			beego.NSInclude(
				&controllers.RegistroPlanAdquisiciones-CodigoArkaController{},
			),
		),

		beego.NSNamespace("/Registro_funcionamiento-Modalidad_seleccion",
			beego.NSInclude(
				&controllers.RegistroFuncionamiento-ModalidadSeleccionController{},
			),
		),

		beego.NSNamespace("/Lineamiento",
			beego.NSInclude(
				&controllers.LineamientoController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
