package main

import (
	"github.com/udistrital/utils_oas/customerror"
	_ "github.com/udistrital/plan_adquisiciones_crud/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
	_ "github.com/lib/pq"
)

func main() {
	orm.RegisterDataBase("default", "postgres",
	"postgres://"+beego.AppConfig.String("PGuser")
	+":"+beego.AppConfig.String("PGpass")+"@"
	+beego.AppConfig.String("PGhost")+"/"+beego.AppConfig.String("PGdb")
	+"?search_path"+"="+beego.AppConfig.String("PGschemas"))
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
	beego.Run()
}