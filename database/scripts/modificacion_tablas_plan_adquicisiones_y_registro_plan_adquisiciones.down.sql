ALTER TABLE plan_adquisiciones."Plan_adquisiciones"
	DROP COLUMN IF EXISTS publicado;
	
ALTER TABLE plan_adquisiciones."Registro_plan_adquisiciones"
	DROP COLUMN IF EXISTS rubro_id,
	DROP COLUMN IF EXISTS fecha_estimada_inicio,
	DROP COLUMN IF EXISTS fecha_estimada_fin;
