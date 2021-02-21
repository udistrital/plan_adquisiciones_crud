package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AgregarValorActividadYDistribucion_20210209_170227 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AgregarValorActividadYDistribucion_20210209_170227{}
	m.Created = "20210209_170227"

	migration.Register("AgregarValorActividadYDistribucion_20210209_170227", m)
}

// Run the migrations
func (m *AgregarValorActividadYDistribucion_20210209_170227) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/agregar_valor_actividad_y_distribucion.up.sql")

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
func (m *AgregarValorActividadYDistribucion_20210209_170227) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/agregar_valor_actividad_y_distribucion.down.sql")

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
