package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type TablasPlanAdquisiciones_20201105_225039 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &TablasPlanAdquisiciones_20201105_225039{}
	m.Created = "20201105_225039"

	migration.Register("TablasPlanAdquisiciones_20201105_225039", m)
}

// Run the migrations
func (m *TablasPlanAdquisiciones_20201105_225039) Up() {
	file, err := ioutil.ReadFile("../scripts/tablas_plan_adquisiciones.up.sql")

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
func (m *TablasPlanAdquisiciones_20201105_225039) Down() {
	file, err := ioutil.ReadFile("../scripts/tablas_plan_adquisiciones.down.sql")

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
