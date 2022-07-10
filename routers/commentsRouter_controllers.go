package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:ActividadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:ActividadController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:ActividadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:ActividadController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:ActividadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:ActividadController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:ActividadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:ActividadController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:ActividadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:ActividadController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:FichaEBIMGAController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:FichaEBIMGAController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:FichaEBIMGAController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:FichaEBIMGAController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:FichaEBIMGAController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:FichaEBIMGAController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:FichaEBIMGAController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:FichaEBIMGAController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:FichaEBIMGAController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:FichaEBIMGAController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:LineamientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:LineamientoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:LineamientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:LineamientoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:LineamientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:LineamientoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:LineamientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:LineamientoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:LineamientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:LineamientoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:MetaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:MetaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:MetaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:MetaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:MetaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:MetaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:MetaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:MetaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:MetaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:MetaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:PlanAdquisicionesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:PlanAdquisicionesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:PlanAdquisicionesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:PlanAdquisicionesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:PlanAdquisicionesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:PlanAdquisicionesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:PlanAdquisicionesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:PlanAdquisicionesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:PlanAdquisicionesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:PlanAdquisicionesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:PlanAdquisicionesMongoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:PlanAdquisicionesMongoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:PlanAdquisicionesMongoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:PlanAdquisicionesMongoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:PlanAdquisicionesMongoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:PlanAdquisicionesMongoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:PlanAdquisicionesMongoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:PlanAdquisicionesMongoController"],
        beego.ControllerComments{
            Method: "GetDiferencia",
            Router: "/diferencia/:idPlanPublicado",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:RegistroFuncionamientoModalidadSeleccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:RegistroFuncionamientoModalidadSeleccionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:RegistroFuncionamientoModalidadSeleccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:RegistroFuncionamientoModalidadSeleccionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:RegistroFuncionamientoModalidadSeleccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:RegistroFuncionamientoModalidadSeleccionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:RegistroFuncionamientoModalidadSeleccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:RegistroFuncionamientoModalidadSeleccionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:RegistroFuncionamientoModalidadSeleccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:RegistroFuncionamientoModalidadSeleccionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:RegistroInversionActividadFuenteFinanciamientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:RegistroInversionActividadFuenteFinanciamientoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:RegistroInversionActividadFuenteFinanciamientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:RegistroInversionActividadFuenteFinanciamientoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:RegistroInversionActividadFuenteFinanciamientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:RegistroInversionActividadFuenteFinanciamientoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:RegistroInversionActividadFuenteFinanciamientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:RegistroInversionActividadFuenteFinanciamientoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:RegistroInversionActividadFuenteFinanciamientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:RegistroInversionActividadFuenteFinanciamientoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:RegistroPlanAdquisicionesActividadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:RegistroPlanAdquisicionesActividadController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:RegistroPlanAdquisicionesActividadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:RegistroPlanAdquisicionesActividadController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:RegistroPlanAdquisicionesActividadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:RegistroPlanAdquisicionesActividadController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:RegistroPlanAdquisicionesActividadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:RegistroPlanAdquisicionesActividadController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:RegistroPlanAdquisicionesActividadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:RegistroPlanAdquisicionesActividadController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:RegistroPlanAdquisicionesCodigoArkaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:RegistroPlanAdquisicionesCodigoArkaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:RegistroPlanAdquisicionesCodigoArkaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:RegistroPlanAdquisicionesCodigoArkaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:RegistroPlanAdquisicionesCodigoArkaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:RegistroPlanAdquisicionesCodigoArkaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:RegistroPlanAdquisicionesCodigoArkaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:RegistroPlanAdquisicionesCodigoArkaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:RegistroPlanAdquisicionesCodigoArkaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:RegistroPlanAdquisicionesCodigoArkaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:RegistroPlanAdquisicionesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:RegistroPlanAdquisicionesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:RegistroPlanAdquisicionesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:RegistroPlanAdquisicionesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:RegistroPlanAdquisicionesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:RegistroPlanAdquisicionesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:RegistroPlanAdquisicionesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:RegistroPlanAdquisicionesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:RegistroPlanAdquisicionesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:RegistroPlanAdquisicionesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:RegistroPlanAdquisicionesMetasAsociadasController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:RegistroPlanAdquisicionesMetasAsociadasController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:RegistroPlanAdquisicionesMetasAsociadasController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:RegistroPlanAdquisicionesMetasAsociadasController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:RegistroPlanAdquisicionesMetasAsociadasController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:RegistroPlanAdquisicionesMetasAsociadasController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:RegistroPlanAdquisicionesMetasAsociadasController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:RegistroPlanAdquisicionesMetasAsociadasController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:RegistroPlanAdquisicionesMetasAsociadasController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:RegistroPlanAdquisicionesMetasAsociadasController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:RegistroPlanAdquisicionesProductosAsociadosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:RegistroPlanAdquisicionesProductosAsociadosController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:RegistroPlanAdquisicionesProductosAsociadosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:RegistroPlanAdquisicionesProductosAsociadosController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:RegistroPlanAdquisicionesProductosAsociadosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:RegistroPlanAdquisicionesProductosAsociadosController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:RegistroPlanAdquisicionesProductosAsociadosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:RegistroPlanAdquisicionesProductosAsociadosController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:RegistroPlanAdquisicionesProductosAsociadosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_crud/controllers:RegistroPlanAdquisicionesProductosAsociadosController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
