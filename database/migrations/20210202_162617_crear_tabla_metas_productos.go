package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaMetasProductos_20210202_162617 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaMetasProductos_20210202_162617{}
	m.Created = "20210202_162617"

	migration.Register("CrearTablaMetasProductos_20210202_162617", m)
}

// Run the migrations
func (m *CrearTablaMetasProductos_20210202_162617) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/crear_tabla_metas_productos.up.sql")

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
func (m *CrearTablaMetasProductos_20210202_162617) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/crear_tabla_metas_productos.down.sql")

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
