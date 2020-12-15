package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type ModificacionTablasPlanAdquicisionesYRegistroPlanAdquisiciones_20201123_092348 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &ModificacionTablasPlanAdquicisionesYRegistroPlanAdquisiciones_20201123_092348{}
	m.Created = "20201123_092348"

	migration.Register("ModificacionTablasPlanAdquicisionesYRegistroPlanAdquisiciones_20201123_092348", m)
}

// Run the migrations
func (m *ModificacionTablasPlanAdquicisionesYRegistroPlanAdquisiciones_20201123_092348) Up() {

	file, err := ioutil.ReadFile("../scripts/modificacion_tablas_plan_adquicisiones_y_registro_plan_adquisiciones.up.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}
}

// Reverse the migrations
func (m *ModificacionTablasPlanAdquicisionesYRegistroPlanAdquisiciones_20201123_092348) Down() {
	file, err := ioutil.ReadFile("../scripts/modificacion_tablas_plan_adquicisiones_y_registro_plan_adquisiciones.down.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}

}
