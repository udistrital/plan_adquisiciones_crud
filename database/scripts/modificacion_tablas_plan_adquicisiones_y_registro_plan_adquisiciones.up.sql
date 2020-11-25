-- ALTER TABLE plan_adquisiciones."Plan_adquisiciones" DROP COLUMN IF EXISTS "publicado";
ALTER TABLE plan_adquisiciones."Plan_adquisiciones"
	ADD COLUMN publicado bool NOT NULL;

-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Plan_adquisiciones".publicado IS E'Verificar si estado de publicacion del plan de adquisicion';



-- ALTER TABLE plan_adquisiciones."Registro_plan_adquisiciones" DROP COLUMN IF EXISTS "publicado","fecha_estimada_inicio","fecha_estimada_fin";
ALTER TABLE plan_adquisiciones."Registro_plan_adquisiciones"
	ADD COLUMN rubro_id text NOT NULL,
	ADD COLUMN fecha_estimada_inicio timestamp NOT NULL,
	ADD COLUMN fecha_estimada_fin timestamp NOT NULL;

-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Registro_plan_adquisiciones".rubro_id IS E'Conocer a que rubro pertenece el registro y poder usarse para ordenar los registros';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Registro_plan_adquisiciones".fecha_estimada_inicio IS E'Fecha estimada de inicio del proceso';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Registro_plan_adquisiciones".fecha_estimada_fin IS E'Fecha estimada de finalizacion del proceso';
