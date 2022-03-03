package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CambiarTipoDatoValorIntegerToNumeric_20220302_190101 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CambiarTipoDatoValorIntegerToNumeric_20220302_190101{}
	m.Created = "20220302_190101"

	migration.Register("CambiarTipoDatoValorIntegerToNumeric_20220302_190101", m)
}

// Run the migrations
func (m *CambiarTipoDatoValorIntegerToNumeric_20220302_190101) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20220302_190101_cambiar_tipo_dato_valor_integer_to_numeric.up.sql")

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
func (m *CambiarTipoDatoValorIntegerToNumeric_20220302_190101) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20220302_190101_cambiar_tipo_dato_valor_integer_to_numeric.down.sql")

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
