package helpers

import (
	"encoding/json"

	"github.com/udistrital/plan_adquisiciones_crud/models"
)

// Función para entregar la estructura de los movimientos de da comparación entre planes publicados
func DevuelveMovimientos(planPublicado models.PlanPublicado, planAntiguo models.PlanPublicado) (movimientos []models.MovimientosDetalle, err error) {
	// logs.Debug(fmt.Sprintf("planPublicado.Id: %+v, planAntiguo.Id: %+v", planPublicado.Id, planAntiguo.Id))
	rubrosPublicado, err := ExtraeRubros(planPublicado)
	if err != nil {
		return nil, err
	}

	rubrosAntiguo, err := ExtraeRubros(planAntiguo)
	if err != nil {
		return nil, err
	}

	nuevosRubros, err := CalculaDiferencias(rubrosPublicado, rubrosAntiguo)
	if err != nil {
		return nil, err
	}

	movimientos, err = ModelaMovimientos(nuevosRubros, planPublicado.IdPlan)
	if err != nil {
		return nil, err
	}

	// logs.Debug(fmt.Sprintf("rubrosCalculoPublicado: %+v", rubrosPublicado))
	// logs.Debug(fmt.Sprintf("rubrosCalculoAntiguo: %+v", rubrosAntiguo))
	// logs.Debug(fmt.Sprintf("nuevosRubros: %+v", nuevosRubros))
	// logs.Debug(fmt.Sprintf("movimientos: %+v", movimientos))

	return
}

// Función para extraer los rubros de un plan entregado
func ExtraeRubros(plan models.PlanPublicado) (rubros []models.RubroCalculo, err error) {
	for _, datos := range plan.RegistroPlanAdquisiciones {
		for _, rubro := range datos.Datos {
			for _, datosRubro := range rubro.Datos {
				if datosRubro.FuenteFinanciamientoData.Codigo != "" && datosRubro.ValorActividad != 0 {
					rubroTemp := models.RubroCalculo{
						RubroId:                rubro.Rubro,
						DeltaAcum:              datosRubro.ValorActividad,
						FuenteFinanciamientoId: datosRubro.FuenteFinanciamientoData.Codigo,
					}
					rubros = append(rubros, rubroTemp)
				} else {
					for _, actividad := range datosRubro.RegistroPlanAdquisicionesActividad {
						for _, fuente := range actividad.FuentesFinanciamiento {
							if fuente.FuenteFinanciamiento != "" && fuente.ValorAsignado != 0 {
								rubroTemp := models.RubroCalculo{
									RubroId:                rubro.Rubro,
									DeltaAcum:              fuente.ValorAsignado,
									ActividadId:            actividad.RegistroActividadId,
									FuenteFinanciamientoId: fuente.FuenteFinanciamiento,
								}
								rubros = append(rubros, rubroTemp)
							}
						}
					}
				}
			}
		}
	}
	return
}

// Función que calcula la diferencia entre dos grupos de rubros entregados
func CalculaDiferencias(rubrosPublicado []models.RubroCalculo, rubrosAntiguo []models.RubroCalculo) (rubros []models.RubroCalculo, err error) {
	var rubrosDiferencia []models.RubroCalculo
	for _, rubroPublicado := range rubrosPublicado {
		for _, rubroAntiguo := range rubrosAntiguo {
			if rubroPublicado.RubroId == rubroAntiguo.RubroId {
				if rubroPublicado.ActividadId != 0 &&
					rubroAntiguo.ActividadId != 0 &&
					rubroPublicado.ActividadId == rubroAntiguo.ActividadId &&
					rubroPublicado.FuenteFinanciamientoId == rubroAntiguo.FuenteFinanciamientoId &&
					(rubroPublicado.DeltaAcum-rubroAntiguo.DeltaAcum) != 0 {
					rubroTemp := models.RubroCalculo{
						RubroId:                rubroPublicado.RubroId,
						DeltaAcum:              rubroPublicado.DeltaAcum - rubroAntiguo.DeltaAcum,
						ActividadId:            rubroPublicado.ActividadId,
						FuenteFinanciamientoId: rubroPublicado.FuenteFinanciamientoId,
					}
					rubrosDiferencia = append(rubrosDiferencia, rubroTemp)
				}

				if rubroPublicado.ActividadId == 0 &&
					rubroAntiguo.ActividadId == 0 &&
					rubroPublicado.ActividadId == rubroAntiguo.ActividadId &&
					rubroPublicado.FuenteFinanciamientoId == rubroAntiguo.FuenteFinanciamientoId &&
					(rubroPublicado.DeltaAcum-rubroAntiguo.DeltaAcum) != 0 {
					rubroTemp := models.RubroCalculo{
						RubroId:                rubroPublicado.RubroId,
						DeltaAcum:              rubroPublicado.DeltaAcum - rubroAntiguo.DeltaAcum,
						FuenteFinanciamientoId: rubroPublicado.FuenteFinanciamientoId,
					}
					rubrosDiferencia = append(rubrosDiferencia, rubroTemp)
				}
			}
		}
	}

	rubrosNuevos, err := NuevosRubros(rubrosPublicado, rubrosDiferencia, false)
	if err != nil {
		return nil, err
	}

	rubrosDeprecados, err := NuevosRubros(rubrosAntiguo, rubrosPublicado, true)
	if err != nil {
		return nil, err
	}

	rubrosTemp := append(rubrosDiferencia, rubrosNuevos...)
	rubros = append(rubrosTemp, rubrosDeprecados...)

	return
}

// Función que modela los movimientos que devuelve el helper
func ModelaMovimientos(nuevosRubros []models.RubroCalculo, idPlanPublicado int) (rubros []models.MovimientosDetalle, err error) {
	for _, rubro := range nuevosRubros {
		detalleFormat, err := json.Marshal(models.DetalleMovimiento{
			RubroId:                rubro.RubroId,
			ActividadId:            rubro.ActividadId,
			FuenteFinanciamientoId: rubro.FuenteFinanciamientoId,
			PlanAdquisicionesId:    idPlanPublicado,
		})
		if err != nil {
			return nil, err
		}

		movimiento := models.MovimientosDetalle{
			Valor:       rubro.DeltaAcum,
			Descripcion: "Movimiento de Afectacion Presupuestal",
			Detalle:     string(detalleFormat),
		}

		rubros = append(rubros, movimiento)
	}

	return
}

// Función que se encarga de crear los movimientos para la primera publicación de un plan
func PublicarPlan(planPublicado models.PlanPublicado) (movimientos []models.MovimientosDetalle, err error) {
	rubros, err := ExtraeRubros(planPublicado)
	if err != nil {
		return nil, err
	}

	movimientos, err = ModelaMovimientos(rubros, planPublicado.IdPlan)
	if err != nil {
		return nil, err
	}

	return
}

// Función que busca dentro de unos rubros objetivo y añade movimientos con valor negativo o positivo según sea el caso
func NuevosRubros(rubrosBusqueda []models.RubroCalculo, rubrosObjetivo []models.RubroCalculo, deprecar bool) (rubros []models.RubroCalculo, err error) {
	var contenido bool
	for _, rubroBusqueda := range rubrosBusqueda {
		contenido = false
		for _, rubroObjetivo := range rubrosObjetivo {
			if rubroBusqueda.RubroId == rubroObjetivo.RubroId {
				if rubroBusqueda.ActividadId != 0 &&
					rubroObjetivo.ActividadId != 0 &&
					rubroBusqueda.ActividadId == rubroObjetivo.ActividadId &&
					rubroBusqueda.FuenteFinanciamientoId == rubroObjetivo.FuenteFinanciamientoId {
					contenido = true
				}

				if rubroBusqueda.ActividadId == 0 &&
					rubroObjetivo.ActividadId == 0 &&
					rubroBusqueda.ActividadId == rubroObjetivo.ActividadId &&
					rubroBusqueda.FuenteFinanciamientoId == rubroObjetivo.FuenteFinanciamientoId {
					contenido = true
				}
			}
		}

		if !contenido && rubroBusqueda.DeltaAcum != 0 {
			if deprecar {
				rubroTemp := models.RubroCalculo{
					RubroId:                rubroBusqueda.RubroId,
					DeltaAcum:              -rubroBusqueda.DeltaAcum,
					ActividadId:            rubroBusqueda.ActividadId,
					FuenteFinanciamientoId: rubroBusqueda.FuenteFinanciamientoId,
				}
				rubros = append(rubros, rubroTemp)
			} else {
				rubroTemp := models.RubroCalculo{
					RubroId:                rubroBusqueda.RubroId,
					DeltaAcum:              rubroBusqueda.DeltaAcum,
					ActividadId:            rubroBusqueda.ActividadId,
					FuenteFinanciamientoId: rubroBusqueda.FuenteFinanciamientoId,
				}
				rubros = append(rubros, rubroTemp)
			}
		}
	}
	return
}
