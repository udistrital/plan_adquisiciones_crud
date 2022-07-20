package models

type PlanPublicado struct {
	Id                        string                                 `json:"_id"`
	IdPlan                    int                                    `json:"id"`
	Descripcion               string                                 `json:"descripcion"`
	Vigencia                  int                                    `json:"vigencia"`
	FechaCreacion             string                                 `json:"fechacreacion"`
	FechaModificacion         string                                 `json:"fechamodificacion"`
	Activo                    bool                                   `json:"activo"`
	Publicado                 bool                                   `json:"publicado"`
	FichaEBImga               []FichaEBIMGA                          `json:"fichaebimga"`
	RegistroPlanAdquisiciones []RubrosRegistroPlanAdquisicionesMongo `json:"registroplanadquisiciones"`
}

type VersionPlan struct {
	Activo                    bool                                   `json:"Activo"`
	Descripcion               string                                 `json:"Descripcion"`
	FechaCreacion             string                                 `json:"FechaCreacion"`
	FechaModificacion         string                                 `json:"FechaModificacion"`
	Id                        int                                    `json:"Id"`
	Publicado                 bool                                   `json:"Publicado"`
	Vigencia                  int                                    `json:"Vigencia"`
	FichaEBImga               []FichaEBIMGA                          `json:"ficha_eb_imga"`
	RegistroPlanAdquisiciones []RubrosRegistroPlanAdquisicionesMongo `json:"registro_plan_adquisiciones"`
}

type RubrosRegistroPlanAdquisicionesMongo struct {
	Fuente     string     `json:"Fuente"`
	FuenteData FuenteData `json:"FuenteData"`
	Datos      []Datos    `json:"datos"`
}

type Datos struct {
	Rubro     string       `json:"Rubro"`
	RubroInfo RubroInfo    `json:"RubroInfo"`
	Datos     []DatosRubro `json:"datos"`
}

type DatosRubro struct {
	FechaEstimadaFin                         string                                          `json:"FechaEstimadaFin"`
	FechaModificacion                        string                                          `json:"FechaModificacion"`
	FuenteFinanciamientoData                 FuenteFinanciamientoData                        `json:"FuenteFinanciamientoData"`
	ProductoId                               string                                          `json:"ProductoId"`
	RubroId                                  string                                          `json:"RubroId"`
	ValorActividad                           float64                                         `json:"ValorActividad"`
	ActividadId                              int                                             `json:"ActividadId"`
	FuenteFinanciamientoId                   string                                          `json:"FuenteFinanciamientoId"`
	MetaId                                   string                                          `json:"MetaId"`
	ResponsableNombre                        string                                          `json:"ResponsableNombre"`
	AreaFuncional                            int                                             `json:"AreaFuncional"`
	CentroGestor                             int                                             `json:"CentroGestor"`
	FechaCreacion                            string                                          `json:"FechaCreacion"`
	FechaEstimadaOfertasInicio               string                                          `json:"FechaEstimadaOfertasInicio"`
	FechaEstimadaOfertasFin                  string                                          `json:"FechaEstimadaOfertasFin"`
	Id                                       int                                             `json:"Id"`
	RegistroPlanAdquisicionesCodigoArka      []RegistroPlanAdquisicionesCodigoArkaMongo      `json:"registro_plan_adquisiciones-codigo_arka"`
	ActividadData                            interface{}                                     `json:"ActividadData"`
	Activo                                   bool                                            `json:"Activo"`
	FechaEstimadaInicio                      string                                          `json:"FechaEstimadaInicio"`
	ResponsableId                            int                                             `json:"ResponsableId"`
	RegistroFuncinamientoModalidadSeleccion  []RegistroFuncionamientoModalidadSeleccionMongo `json:"registro_funcionamiento-modalidad_seleccion"`
	RegistroPlanAdquisicionesActividad       []RegistroPlanAdquisicionesActividadMongoNuevo  `json:"registro_plan_adquisiciones-actividad"`
	RegistroFuncionamientoMetasAsociadas     []RegistroFuncionamientoMetasAsociadas          `json:"registro_funcionamiento-metas_asociadas"`
	ValorTotalActividades                    float64                                         `json:"ValorTotalActividades"`
	RegistroFuncionamientoProductosAsociados []RegistroFuncionamientoProductosAsociados      `json:"registro_funcionamiento-productos_asociados"`
}

type RegistroFuncionamientoProductosAsociados struct {
	FechaModificacion      string       `json:"FechaModificacion"`
	Id                     int          `json:"Id"`
	PorcentajeDistribucion int          `json:"PorcentajeDistribucion"`
	ProductoAsociadoId     string       `json:"ProductoAsociadoId"`
	ProductoData           ProductoData `json:"ProductoData"`
	Activo                 bool         `json:"Activo"`
	FechaCreacion          string       `json:"FechaCreacion"`
}

type ProductoData struct {
	Codigo            int    `json:"Codigo"`
	Descripcion       string `json:"Descripcion"`
	FechaCreacion     string `json:"FechaCreacion"`
	FechaModificacion string `json:"FechaModificacion"`
	Nombre            string `json:"Nombre"`
	Vigencia          int    `json:"Vigencia"`
	Id                string `json:"_id"`
	Activo            bool   `json:"Activo"`
}

type RegistroFuncionamientoMetasAsociadas struct {
	FechaModificacion string `json:"FechaModificacion"`
	Id                int    `json:"Id"`
	MetaId            MetaId `json:"MetaId"`
	Activo            bool   `json:"Activo"`
	FechaCreacion     string `json:"FechaCreacion"`
}

type MetaId struct {
	Id                int    `json:"Id"`
	LineamientoId     string `json:"LineamientoId"`
	Nombre            string `json:"Nombre"`
	Numero            int    `json:"Numero"`
	Rubro             string `json:"Rubro"`
	Activo            bool   `json:"Activo"`
	FechaCreacion     string `json:"FechaCreacion"`
	FechaModificacion string `json:"FechaModificacion"`
}

type FuenteFinanciamientoData struct {
	FechaCreacion     string      `json:"FechaCreacion"`
	NumeroDocumento   string      `json:"NumeroDocumento"`
	Rubros            interface{} `json:"Rubros"`
	ValorInicial      float64     `json:"ValorInicial"`
	Activo            bool        `json:"Activo"`
	TipoDocumento     string      `json:"TipoDocumento"`
	Nombre            string      `json:"Nombre"`
	Movimientos       interface{} `json:"Movimientos"`
	UnidadEjecutora   string      `json:"UnidadEjecutora"`
	FechaModificacion string      `json:"FechaModificacion"`
	Descripcion       string      `json:"Descripcion"`
	Estado            string      `json:"Estado"`
	TipoFuente        interface{} `json:"TipoFuente"`
	ValorActual       float64     `json:"ValorActual"`
	Vigencia          int         `json:"Vigencia"`
	Codigo            string      `json:"Codigo"`
}

type FuenteData struct {
	UnidadEjecutora string `json:"UnidadEjecutora"`
	Codigo          string `json:"Codigo"`
	Nombre          string `json:"Nombre"`
}

type RubroInfo struct {
	Apropiaciones     bool          `json:"Apropiaciones"`
	Hijos             []interface{} `json:"Hijos"`
	Nombre            string        `json:"Nombre"`
	Codigo            string        `json:"Codigo"`
	Estado            string        `json:"Estado"`
	FechaCreacion     string        `json:"FechaCreacion"`
	FechaModificacion string        `json:"FechaModificacion"`
	Activo            bool          `json:"Activo"`
	Descripcion       string        `json:"Descripcion"`
	Productos         interface{}   `json:"Productos"`
	UnidadEjecutora   string        `json:"UnidadEjecutora"`
	Vigencia          int           `json:"Vigencia"`
	Bloqueado         bool          `json:"Bloqueado"`
	Padre             string        `json:"Padre"`
	ValorInicial      float64       `json:"ValorInicial"`
}

type RegistroPlanAdquisicionesActividadMongoNuevo struct {
	NumeroMeta                  int                     `json:"NumeroMeta"`
	RegistroActividadId         int                     `json:"RegistroActividadId"`
	RegistroPlanAdquisicionesId string                  `json:"RegistroPlanAdquisicionesId"`
	Valor                       float64                 `json:"Valor"`
	Actividad                   Actividad               `json:"actividad"`
	Activo                      bool                    `json:"Activo"`
	FechaCreacion               string                  `json:"FechaCreacion"`
	FechaModificacion           string                  `json:"FechaModificacion"`
	FuentesFinanciamiento       []FuentesFinanciamiento `json:"FuentesFinanciamiento"`
}

type FuentesFinanciamiento struct {
	Activo               bool    `json:"Activo"`
	FechaCreacion        string  `json:"FechaCreacion"`
	FechaModificacion    string  `json:"FechaModificacion"`
	FuenteFinanciamiento string  `json:"FuenteFinanciamiento"`
	Id                   int     `json:"Id"`
	Nombre               string  `json:"Nombre"`
	ValorAsignado        float64 `json:"ValorAsignado"`
}
