-- ALTER TABLE plan_adquisiciones."Registro_plan_adquisiciones" DROP COLUMN IF EXISTS "publicado","fecha_estimada_inicio","fecha_estimada_fin";
ALTER TABLE plan_adquisiciones."Registro_plan_adquisiciones"
	ADD COLUMN fuente_financiamiento_id varchar(20);

-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Registro_plan_adquisiciones".rubro_id IS E'Cuando el registro es de funcionamiento, es necesario guardar la fuente de financiamiento de forma directa';
-- ddl-end --


-- object: plan_adquisiciones."Registro_plan_adquisiciones-Metas_Asociadas" | type: TABLE --
-- DROP TABLE IF EXISTS plan_adquisiciones."Registro_plan_adquisiciones-Metas_Asociadas" CASCADE;
CREATE TABLE plan_adquisiciones."Registro_plan_adquisiciones-Metas_Asociadas" (
	id serial NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL,
	"Registro_plan_adquisiciones_id" integer,
	"Meta_id" integer,
	CONSTRAINT "Registro_plan_adquisiciones-Metas_Asociadas_pk" PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE plan_adquisiciones."Registro_plan_adquisiciones-Metas_Asociadas" IS E'Códigos de Arka asociados al plan de inversion o de funcionamiento.';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Registro_plan_adquisiciones-Metas_Asociadas".id IS E'Identificador de la tabla';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Registro_plan_adquisiciones-Metas_Asociadas".fecha_modificacion IS E'Fecha de la última modificación realizada al registro.';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Registro_plan_adquisiciones-Metas_Asociadas".activo IS E'Indica si el registro se encuentra activo dentro de la base de datos';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Registro_plan_adquisiciones-Metas_Asociadas".fecha_creacion IS E'Fecha en la que se creo el registro';
-- ddl-end --
-- ALTER TABLE plan_adquisiciones."Registro_plan_adquisiciones-Metas_Asociadas" OWNER TO postgres;
-- ddl-end --

-- object: "Registro_plan_adquisiciones_fk" | type: CONSTRAINT --
-- ALTER TABLE plan_adquisiciones."Registro_plan_adquisiciones-Metas_Asociadas" DROP CONSTRAINT IF EXISTS "Registro_plan_adquisiciones_fk" CASCADE;
ALTER TABLE plan_adquisiciones."Registro_plan_adquisiciones-Metas_Asociadas" ADD CONSTRAINT "Registro_plan_adquisiciones_fk" FOREIGN KEY ("Registro_plan_adquisiciones_id")
REFERENCES plan_adquisiciones."Registro_plan_adquisiciones" (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
-- ddl-end --

-- object: plan_adquisiciones."Registro_plan_adquisiciones-Productos_Asociados" | type: TABLE --
-- DROP TABLE IF EXISTS plan_adquisiciones."Registro_plan_adquisiciones-Productos_Asociados" CASCADE;
CREATE TABLE plan_adquisiciones."Registro_plan_adquisiciones-Productos_Asociados" (
	id serial NOT NULL,
	producto_asociado_id varchar NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL,
	"Registro_plan_adquisiciones_id" integer,
	CONSTRAINT "Registro_plan_adquisiciones-Productos_Asociados_pk" PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE plan_adquisiciones."Registro_plan_adquisiciones-Productos_Asociados" IS E'Códigos de Arka asociados al plan de inversion o de funcionamiento.';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Registro_plan_adquisiciones-Productos_Asociados".id IS E'Identificador de la tabla';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Registro_plan_adquisiciones-Productos_Asociados".producto_asociado_id IS E'Codigo de Meta Asociada';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Registro_plan_adquisiciones-Productos_Asociados".fecha_modificacion IS E'Fecha de la última modificación realizada al registro.';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Registro_plan_adquisiciones-Productos_Asociados".activo IS E'Indica si el registro se encuentra activo dentro de la base de datos';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Registro_plan_adquisiciones-Productos_Asociados".fecha_creacion IS E'Fecha en la que se creo el registro';
-- ddl-end --
-- ALTER TABLE plan_adquisiciones."Registro_plan_adquisiciones-Productos_Asociados" OWNER TO postgres;
-- ddl-end --

-- object: "Registro_plan_adquisiciones_fk" | type: CONSTRAINT --
-- ALTER TABLE plan_adquisiciones."Registro_plan_adquisiciones-Productos_Asociados" DROP CONSTRAINT IF EXISTS "Registro_plan_adquisiciones_fk" CASCADE;
ALTER TABLE plan_adquisiciones."Registro_plan_adquisiciones-Productos_Asociados" ADD CONSTRAINT "Registro_plan_adquisiciones_fk" FOREIGN KEY ("Registro_plan_adquisiciones_id")
REFERENCES plan_adquisiciones."Registro_plan_adquisiciones" (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
-- ddl-end --

-- object: "Meta_fk" | type: CONSTRAINT --
-- ALTER TABLE plan_adquisiciones."Registro_plan_adquisiciones-Metas_Asociadas" DROP CONSTRAINT IF EXISTS "Meta_fk" CASCADE;
ALTER TABLE plan_adquisiciones."Registro_plan_adquisiciones-Metas_Asociadas" ADD CONSTRAINT "Meta_fk" FOREIGN KEY ("Meta_id")
REFERENCES plan_adquisiciones."Meta" (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
-- ddl-end --

