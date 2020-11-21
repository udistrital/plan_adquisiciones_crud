package main

import (
	_ "github.com/udistrital/plan_adquisiciones_crud/routers"
	apistatus "github.com/udistrital/utils_oas/apiStatusLib"
	"github.com/udistrital/utils_oas/customerror"
	dbMongoManager "github.com/udistrital/plan_adquisiciones_crud/database"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
	_ "github.com/lib/pq"
)

func main() {
	dbMongoManager.InitDB(beego.AppConfig.String("mongoHost"),
		beego.AppConfig.String("mongoPort"),
		beego.AppConfig.String("mongoUser"),
		beego.AppConfig.String("mongoPass"),
		beego.AppConfig.String("mongoAuth"),
		beego.AppConfig.String("mongodb"),
		beego.AppConfig.String("mongoCollection"),
	)
	orm.RegisterDataBase("default", "postgres", "postgres://"+beego.AppConfig.String("PGuser")+":"+beego.AppConfig.String("PGpass")+"@"+beego.AppConfig.String("PGhost")+"/"+beego.AppConfig.String("PGdb")+"?sslmode=disable&search_path"+"="+beego.AppConfig.String("PGschemas"))
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders: []string{"Origin", "x-requested-with",
			"content-type",
			"accept",
			"origin",
			"authorization",
			"x-csrftoken"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	beego.ErrorController(&customerror.CustomErrorController{})
	apistatus.Init()

	beego.Run()
}
