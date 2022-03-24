package models

import (
	"context"

	"github.com/astaxie/beego/logs"
	dbMongoManager "github.com/udistrital/plan_adquisiciones_crud/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//PlanAdquisicionesMongo con ID
type PlanAdquisicionesMongoID struct {
	IdMongo               interface{}
	PlanAdquisiconesMongo PlanAdquisicionesMongo
}

// PlanAdquisicionesMongo ...
type PlanAdquisicionesMongo struct {
	ID                        int                      `json:"id"`
	Descripcion               string                   `json:"descripcion"`
	Vigencia                  int                      `json:"vigencia"`
	FechaCreacion             string                   `json:"FechaCreacion"`
	FechaModificacion         string                   `json:"FechaModificacion"`
	Activo                    bool                     `json:"activo"`
	Publicado                 bool                     `json:"publicado"`
	FichaEbImga               []FichaEBIMGA            `json:"ficha_eb_imga"`
	RegistroPlanAdquisiciones []map[string]interface{} `json:"registro_plan_adquisiciones"`
}

// FichaEbImga ...
type FichaEbImga struct {
	ID                int    `json:"id"`
	MetaID            int    `json:"meta_id"`
	Proceso           string `json:"proceso"`
	Magnitud          int    `json:"magnitud"`
	UnidadMedida      string `json:"unidad_medida"`
	Descripcion       string `json:"descripcion"`
	FechaModificacion string `json:"fecha_modificacion"`
	FechaCreacion     string `json:"FechaModificacion"`
	Activo            bool   `json:"activo"`
	Rubro             string `json:"rubro"`
}

// RegistroPlanAdquisicionesMongo ...
type RegistroPlanAdquisicionesMongo struct {
	ID                                       int                                             `json:"id"`
	MetaID                                   string                                          `json:"MetaId"`
	AreaFuncional                            int                                             `json:"AreaFuncional"`
	MetaNombre                               string                                          `json:"MetaNombre"`
	CentroGestor                             int                                             `json:"CentroGestor"`
	FechaCreacion                            string                                          `json:"FechaCreacion"`
	FechaModificacion                        string                                          `json:"FechaModificacion"`
	ResponsableID                            int                                             `json:"ResponsableId"`
	ResponsableNombre                        string                                          `json:"ResponsableNombre"`
	Activo                                   bool                                            `json:"activo"`
	ProductoID                               string                                          `json:"ProductoId"`
	ProductoNombre                           string                                          `json:"ProductoNombre"`
	RubroID                                  string                                          `json:"RubroId"`
	RubroNombre                              string                                          `json:"RubroNombre"`
	FuenteRecursosNombre                     string                                          `json:"FuenteRecursosNombre"`
	ValorTotalActividades                    float64                                         `json:"ValorTotalActividades"`
	FechaEstimadaInicio                      string                                          `json:"FechaEstimadaInicio"`
	FechaEstimadaFin                         string                                          `json:"FechaEstimadaFin"`
	RegistroPlanAdquisicionesCodigoArka      []RegistroPlanAdquisicionesCodigoArkaMongo      `json:"registro_plan_adquisiciones-codigo_arka"`
	RegistroFuncionamientoModalidadSeleccion []RegistroFuncionamientoModalidadSeleccionMongo `json:"registro_funcionamiento-modalidad_seleccion"`
	RegistroPlanAdquisicionesActividad       []RegistroPlanAdquisicionesActividadMongo       `json:"registro_plan_adquisiciones-actividad"`
}

// RegistroPlanAdquisicionesCodigoArkaMongo ...
type RegistroPlanAdquisicionesCodigoArkaMongo struct {
	ID                int    `json:"id"`
	CodigoArka        string `json:"CodigoArka"`
	FechaModificacion string `json:"FechaModificacion"`
	Activo            bool   `json:"activo"`
	FechaCreacion     string `json:"FechaCreacion"`
	Descripcion       string `json:"Descripcion"`
}

// RegistroFuncionamientoModalidadSeleccionMongo ...
type RegistroFuncionamientoModalidadSeleccionMongo struct {
	ID                   int    `json:"id"`
	IDModalidadSeleccion string `json:"IdModalidadSeleccion"`
	FechaModificacion    string `json:"FechaModificacion"`
	Activo               bool   `json:"activo"`
	FechaCreacion        string `json:"FechaCreacion"`
	Nombre               string `json:"Nombre"`
}

// RegistroPlanAdquisicionesActividadMongo ...
type RegistroPlanAdquisicionesActividadMongo struct {
	ID                                             int                                                   `json:"RegistroActividadId"`
	Valor                                          int                                                   `json:"valor"`
	FechaCreacion                                  string                                                `json:"FechaCreacion"`
	FechaModificacion                              string                                                `json:"FechaModificacion"`
	Activo                                         bool                                                  `json:"activo"`
	RegistroInversionActividadFuenteFinanciamiento []RegistroInversionActividadFuenteFinanciamientoMongo `json:"FuentesFinanciamiento"`
	Actividad                                      ActividadMongo                                        `json:"actividad"`
}

// RegistroInversionActividadFuenteFinanciamientoMongo ...
type RegistroInversionActividadFuenteFinanciamientoMongo struct {
	ID                     int    `json:"id"`
	FuenteFinanciamientoID string `json:"FuenteFinanciamiento"`
	ValorAsignado          int    `json:"ValorAsignado"`
	FechaModificacion      string `json:"FechaModificacion"`
	Activo                 bool   `json:"activo"`
	FechaCreacion          string `json:"FechaCreacion"`
	Nombre                 string `json:"Nombre"`
}

// ActividadMongo ...
type ActividadMongo struct {
	ID                int       `json:"id"`
	Numero            int       `json:"numero"`
	Nombre            string    `json:"nombre"`
	FechaCreacion     string    `json:"FechaCreacion"`
	FechaModificacion string    `json:"FechaModificacion"`
	Activo            bool      `json:"activo"`
	Meta              MetaMongo `json:"MetaId"`
}

// MetaMongo ...
type MetaMongo struct {
	ID                int              `json:"id"`
	Numero            int              `json:"numero"`
	Nombre            string           `json:"nombre"`
	FechaCreacion     string           `json:"FechaCreacion"`
	FechaModificacion string           `json:"FechaModificacion"`
	Activo            bool             `json:"activo"`
	Rubro             string           `json:"rubro"`
	Lineamiento       LineamientoMongo `json:"LineamientoId"`
}

// LineamientoMongo ...
type LineamientoMongo struct {
	ID                int    `json:"id"`
	Numero            int    `json:"numero"`
	Nombre            string `json:"nombre"`
	Objetivo          string `json:"objetivo"`
	FuenteRecursoID   string `json:"FuenteRecursoId"`
	CentroGestor      int    `json:"CentroGestor"`
	AreaFuncionalID   int    `json:"AreaFuncionalId"`
	Vigencia          int    `json:"vigencia"`
	FechaCreacion     string `json:"FechaCreacion"`
	FechaModificacion string `json:"FechaModificacion"`
	Activo            bool   `json:"activo"`
}

// AddPlanAdquisicionesMongo insert a new PlanAdquisiciones into database and returns
// last inserted Id on success.
func AddPlanAdquisicionesMongo(m *PlanAdquisicionesMongo) (id interface{}, err error) {
	collection, err := dbMongoManager.GetMainCollection()
	insertResult, err := collection.InsertOne(context.TODO(), m)
	return insertResult.InsertedID, err
}

// GetPlanAdquisicionesMongoById retrieves PlanAdquisiciones by Id. Returns error if
// Id doesn't exist
func GetPlanAdquisicionesMongoById(id string) (v *bson.M, err error) {
	collection, err := dbMongoManager.GetMainCollection()
	objId, err := primitive.ObjectIDFromHex(id)
	collection.FindOne(context.TODO(), bson.M{"_id": objId}).Decode(&v)
	if err != nil {
		logs.Error(err)
	} else {
		return v, nil
	}
	return nil, err
}

// GetAllPlanAdquisicionesMongo retrieves all PlanAdquisiciones matches certain condition. Returns empty list if
// no records exist
func GetAllPlanAdquisicionesMongo(query bson.M) (ml []interface{}, err error) {
	collection, err := dbMongoManager.GetMainCollection()
	if query == nil {
		query = bson.M{}
	}
	cursor, err := collection.Find(context.TODO(), query)
	var ml2 []bson.M
	if err = cursor.All(context.TODO(), &ml2); err != nil {
		logs.Error(err)
		return nil, err
	}
	for _, m := range ml2 {
		ml = append(ml, m)
	}
	return ml, nil
}
