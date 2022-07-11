package models

type RubroCalculo struct {
	RubroId                string
	DeltaAcum              float64
	ActividadId            int
	FuenteFinanciamientoId string
}

type DetalleMovimiento struct {
	RubroId                string
	ActividadId            int
	FuenteFinanciamientoId string
	PlanAdquisicionesId    int
}
