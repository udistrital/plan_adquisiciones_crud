ALTER TABLE plan_adquisiciones."Registro_plan_adquisiciones"
	DROP COLUMN IF EXISTS actividad_id,
    DROP COLUMN IF EXISTS valor_actividad;
	
ALTER TABLE plan_adquisiciones."Registro_plan_adquisiciones-Productos_Asociados"
	DROP COLUMN IF EXISTS porcentaje_distribucion;
