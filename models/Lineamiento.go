package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type Lineamiento struct {
	Id                int    `orm:"column(id);pk;auto"`
	Numero            int    `orm:"column(numero)"`
	Nombre            string `orm:"column(nombre)"`
	Objetivo          string `orm:"column(objetivo)"`
	FuenteRecursoId   string `orm:"column(fuente_recurso_id)"`
	CentroGestor      int    `orm:"column(centro_gestor)"`
	AreaFuncionalId   int    `orm:"column(area_funcional_id)"`
	Vigencia          int    `orm:"column(vigencia)"`
	FechaCreacion     string `orm:"column(fecha_creacion);type(timestamp without time zone)"`
	FechaModificacion string `orm:"column(fecha_modificacion);type(timestamp without time zone)"`
	Activo            bool   `orm:"column(activo)"`
}

func (t *Lineamiento) TableName() string {
	return "Lineamiento"
}

func init() {
	orm.RegisterModel(new(Lineamiento))
}

// AddLineamiento insert a new Lineamiento into database and returns
// last inserted Id on success.
func AddLineamiento(m *Lineamiento) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetLineamientoById retrieves Lineamiento by Id. Returns error if
// Id doesn't exist
func GetLineamientoById(id int) (v *Lineamiento, err error) {
	o := orm.NewOrm()
	v = &Lineamiento{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllLineamiento retrieves all Lineamiento matches certain condition. Returns empty list if
// no records exist
func GetAllLineamiento(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Lineamiento))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Lineamiento
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateLineamiento updates Lineamiento by Id and returns error if
// the record to be updated doesn't exist
func UpdateLineamientoById(m *Lineamiento) (err error) {
	o := orm.NewOrm()
	v := Lineamiento{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteLineamiento deletes Lineamiento by Id and returns error if
// the record to be deleted doesn't exist
func DeleteLineamiento(id int) (err error) {
	o := orm.NewOrm()
	v := Lineamiento{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Lineamiento{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
