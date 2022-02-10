DROP TABLE IF EXISTS plan_adquisiciones."Metas_Asociadas" CASCADE;
DROP TABLE IF EXISTS plan_adquisiciones."Productos_Asociados" CASCADE;

ALTER TABLE plan_adquisiciones."Registro_plan_adquisiciones"
	DROP COLUMN IF EXISTS fuente_financiamiento_id,