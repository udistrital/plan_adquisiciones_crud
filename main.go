package main

import (
	"context"
	_ "github.com/udistrital/plan_adquisiciones_crud/routers"
	apistatus "github.com/udistrital/utils_oas/apiStatusLib"
	"github.com/udistrital/utils_oas/customerror"
	dbMongoManager "github.com/udistrital/plan_adquisiciones_crud/database"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
	_ "github.com/lib/pq"

	"fmt"
	hola "github.com/udistrital/plan_adquisiciones_crud/models"
)

func main() {
	dbMongoManager.InitDB(beego.AppConfig.String("mongoHost"),
		beego.AppConfig.String("mongoPort"),
		beego.AppConfig.String("mongoUser"),
		beego.AppConfig.String("mongoPass"),
		beego.AppConfig.String("mongoAuth"),
		beego.AppConfig.String("mongodb"),
	)
	var dat, _ = dbMongoManager.GetDatabase();
	var coll = dat.Collection("plan_adquisiciones_crud_mongo");
	var post = hola.HelloWorld{ID: 2, Descripcion: "Plan de adquisiciones"};
	var insertResult, err = coll.InsertOne(context.TODO(), post);
	if err != nil {
		fmt.Println("hello ERROR-------------------------------------------");
		fmt.Println(err);
		fmt.Println("hello ERROR-------------------------------------------");
	}
	
	fmt.Println("hello world-------------------------------------------");
	//fmt.Println(coll);
	fmt.Println("Inserted post with ID:", insertResult);
	fmt.Println("hello world-------------------------------------------");
	//fmt.Println(result);
	fmt.Println("hello world-------------------------------------------");
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
