package models

import (
	"context"
	dbMongoManager "github.com/udistrital/plan_adquisiciones_crud/database"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/astaxie/beego/logs"
	"encoding/json"
)

type HelloWorld struct {
	ID                int    `json:"id"`
	Descripcion       string `json:"descripcion"`
}

type PlanAdquisicionesMongo struct {
	ID                int    `json:"id"`
	Descripcion       string `json:"descripcion"`
	Vigencia          int    `json:"vigencia"`
	FechaCreacion     string `json:"fecha_creacion"`
	FechaModificacion string `json:"fecha_modificacion"`
	Activo            bool   `json:"activo"`
	FichaEbImga       []struct {
		ID                int    `json:"id"`
		MetaID            int    `json:"meta_id"`
		Proceso           string `json:"proceso"`
		Magnitud          int    `json:"magnitud"`
		UnidadMedida      string `json:"unidad_medida"`
		Descripcion       string `json:"descripcion"`
		FechaModificacion string `json:"fecha_modificacion"`
		FechaCreacion     string `json:"fecha_creacion"`
		Activo            bool   `json:"activo"`
		Rubro             string `json:"rubro"`
	} `json:"ficha_eb_imga"`
	RegistroPlanAdquisiciones []struct {
		ID                                  int    `json:"id"`
		MetaID                              string `json:"meta_id"`
		CentroGestor                        int    `json:"centro_gestor"`
		FechaCreacion                       string `json:"fecha_creacion"`
		FechaModificacion                   string `json:"fecha_modificacion"`
		ResponsableID                       int    `json:"responsable_id"`
		Activo                              bool   `json:"activo"`
		ProductoID                          string `json:"producto_id"`
		RegistroPlanAdquisicionesCodigoArka []struct {
			ID                int    `json:"id"`
			CodigoArka        string `json:"codigo_arka"`
			FechaModificacion string `json:"fecha_modificacion"`
			Activo            bool   `json:"activo"`
			FechaCreacion     string `json:"fecha_creacion"`
		} `json:"registro_plan_adquisiciones-codigo_arka"`
		RegistroFuncionamientoModalidadSeleccion []struct {
			ID                   int    `json:"id"`
			IDModalidadSeleccion int    `json:"id_modalidad_seleccion"`
			FechaModificacion    string `json:"fecha_modificacion"`
			Activo               bool   `json:"activo"`
			FechaCreacion        string `json:"fecha_creacion"`
		} `json:"registro_funcionamiento-modalidad_seleccion"`
		RegistroPlanAdquisicionesActividad []struct {
			ID                                             int    `json:"id"`
			Valor                                          int    `json:"valor"`
			FechaCreacion                                  string `json:"fecha_creacion"`
			FechaModificacion                              string `json:"fecha_modificacion"`
			Activo                                         bool   `json:"activo"`
			RegistroInversionActividadFuenteFinanciamiento []struct {
				ID                     int    `json:"id"`
				FuenteFinanciamientoID string `json:"fuente_financiamiento_id"`
				ValorAsignado          int    `json:"valor_asignado"`
				FechaModificacion      string `json:"fecha_modificacion"`
				Activo                 bool   `json:"activo"`
				FechaCreacion          string `json:"fecha_creacion"`
			} `json:"registro_inversion_actividad-fuente_financiamiento"`
			Actividad struct {
				ID                int    `json:"id"`
				Numero            int    `json:"numero"`
				Nombre            string `json:"nombre"`
				FechaCreacion     string `json:"fecha_creacion"`
				FechaModificacion string `json:"fecha_modificacion"`
				Activo            bool   `json:"activo"`
				Meta              struct {
					ID                int    `json:"id"`
					Numero            int    `json:"numero"`
					Nombre            string `json:"nombre"`
					FechaCreacion     string `json:"fecha_creacion"`
					FechaModificacion string `json:"fecha_modificacion"`
					Activo            bool   `json:"activo"`
					Rubro             string `json:"rubro"`
					Lineamiento       struct {
						ID                int    `json:"id"`
						Numero            int    `json:"numero"`
						Nombre            string `json:"nombre"`
						Objetivo          string `json:"objetivo"`
						FuenteRecursoID   int    `json:"fuente_recurso_id"`
						CentroGestor      int    `json:"centro_gestor"`
						AreaFuncionalID   int    `json:"area_funcional_id"`
						Vigencia          int    `json:"vigencia"`
						FechaCreacion     string `json:"fecha_creacion"`
						FechaModificacion string `json:"fecha_modificacion"`
						Activo            bool   `json:"activo"`
					} `json:"lineamiento"`
				} `json:"meta"`
			} `json:"Actividad"`
		} `json:"registro_plan_adquisiciones-actividad"`
	} `json:"registro_plan_adquisiciones"`
}

// GetAllPlanAdquisiciones retrieves all PlanAdquisiciones matches certain condition. Returns empty list if
// no records exist
func GetAllPlanAdquisicionesMongo() (ml []interface{}, err error) {
	db,_ := dbMongoManager.GetDatabase();
	collection := db.Collection("plan_adquisiciones_crud_mongo");
	cursor, err := collection.Find(context.TODO(), bson.M{});
	var ml2 []PlanAdquisicionesMongo;
	if err = cursor.All(context.TODO(), &ml2);
	err != nil {
		logs.Error(err);
		return nil, err;
	}
	for i, v := range ml2 {
		ml = append(ml, &PlanAdquisicionesMongo{});
	}
	var test = &ml2;
	return ml, nil;
}