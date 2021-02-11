ALTER TABLE plan_adquisiciones."Registro_plan_adquisiciones"
	ADD COLUMN actividad_id integer,
	ADD COLUMN valor_actividad integer;

-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Registro_plan_adquisiciones".actividad_id IS E'Actividad Relacionada al registro del plan de adquisiciones de funcionamiento';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Registro_plan_adquisiciones".valor_actividad IS E'Valor de la actividad relacionada del registro del plan de adquisiciones de funcionamiento';


ALTER TABLE plan_adquisiciones."Registro_plan_adquisiciones-Productos_Asociados"
	ADD COLUMN porcentaje_distribucion integer;

-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Registro_plan_adquisiciones-Productos_Asociados".porcentaje_distribucion IS E'porcentaje de distribucion del producto asociado al registro del plan de adquisiciones de inversion';
