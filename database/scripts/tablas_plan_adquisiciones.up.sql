-- Database generated with pgModeler (PostgreSQL Database Modeler).
-- pgModeler  version: 0.9.2
-- PostgreSQL version: 9.5
-- Project Site: pgmodeler.io
-- Model Author: ---


-- Database creation must be done outside a multicommand file.
-- These commands were put in this file only as a convenience.
-- -- object: new_database | type: DATABASE --
-- -- DROP DATABASE IF EXISTS new_database;
-- CREATE DATABASE new_database;
-- -- ddl-end --
-- 

-- object: plan_adquisiciones | type: SCHEMA --
-- DROP SCHEMA IF EXISTS plan_adquisiciones CASCADE;
CREATE SCHEMA plan_adquisiciones;
-- ddl-end --
-- ALTER SCHEMA plan_adquisiciones OWNER TO postgres;
-- ddl-end --

SET search_path TO pg_catalog,public,plan_adquisiciones;
-- ddl-end --

-- object: plan_adquisiciones."Meta" | type: TABLE --
-- DROP TABLE IF EXISTS plan_adquisiciones."Meta" CASCADE;
CREATE TABLE plan_adquisiciones."Meta" (
	id serial NOT NULL,
	numero integer NOT NULL,
	nombre varchar(250) NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	activo boolean NOT NULL,
	rubro varchar NOT NULL,
	"Lineamiento_id" integer,
	CONSTRAINT "Meta_pk" PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE plan_adquisiciones."Meta" IS E'Tabla para almacenar las metas creadas en el plan de adquisición.';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Meta".id IS E'Identificador de la meta.';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Meta".numero IS E'Número asociado a la meta.';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Meta".nombre IS E'Nombre de la meta';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Meta".fecha_creacion IS E'Fecha en la que se creo la meta.';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Meta".fecha_modificacion IS E'Última fecha en la que se modifico la meta.';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Meta".activo IS E'Indica si el registro se encuentra activo dentro de la base de datos.';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Meta".rubro IS E'Rubro al que se encuentra asociada la meta';
-- ddl-end --
-- ALTER TABLE plan_adquisiciones."Meta" OWNER TO postgres;
-- ddl-end --

-- object: plan_adquisiciones."Lineamiento" | type: TABLE --
-- DROP TABLE IF EXISTS plan_adquisiciones."Lineamiento" CASCADE;
CREATE TABLE plan_adquisiciones."Lineamiento" (
	id serial NOT NULL,
	numero integer NOT NULL,
	nombre varchar(120) NOT NULL,
	objetivo text NOT NULL,
	fuente_recurso_id varchar(5) NOT NULL,
	centro_gestor integer NOT NULL,
	area_funcional_id integer NOT NULL,
	vigencia integer NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	activo boolean NOT NULL,
	CONSTRAINT "Lineamiento_pk" PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE plan_adquisiciones."Lineamiento" IS E'Tabla para almacenar los lineamientos del plan de adquisiciones.';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Lineamiento".id IS E'ID del lineamiento';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Lineamiento".numero IS E'Número asignado al lineamiento.';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Lineamiento".nombre IS E'Nombre asignado al lineamiento.';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Lineamiento".objetivo IS E'Objetivo asignado al lineamiento.';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Lineamiento".fuente_recurso_id IS E'Fuente de recurso asignado al lineamiento.';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Lineamiento".centro_gestor IS E'Centro gestor asignado al lineamiento.';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Lineamiento".area_funcional_id IS E'Área funcional a la que se realiza el lineamiento.';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Lineamiento".vigencia IS E'Vigencia a la que pertenece el lineamiento.';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Lineamiento".fecha_creacion IS E'Fecha en la que ha sido creado el lineamiento.';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Lineamiento".fecha_modificacion IS E'Última fecha en la que se modificó el lineamiento.';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Lineamiento".activo IS E'Indica si el registro se encuentra activo en la base de datos';
-- ddl-end --
-- ALTER TABLE plan_adquisiciones."Lineamiento" OWNER TO postgres;
-- ddl-end --

-- object: plan_adquisiciones."Actividad" | type: TABLE --
-- DROP TABLE IF EXISTS plan_adquisiciones."Actividad" CASCADE;
CREATE TABLE plan_adquisiciones."Actividad" (
	id serial NOT NULL,
	numero integer NOT NULL,
	nombre varchar(250) NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	activo boolean NOT NULL,
	"Meta_id" integer,
	CONSTRAINT "Actividad_pk" PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE plan_adquisiciones."Actividad" IS E'Actividades creadas dentro de un plan de adquisiciones asociadas a una meta.';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Actividad".id IS E'Identificador de la actividad.';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Actividad".numero IS E'Número asociado a la actividad.';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Actividad".nombre IS E'Nombre asociado a la actividad.';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Actividad".fecha_creacion IS E'Fecha de creación de la actividad.';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Actividad".fecha_modificacion IS E'Última fecha en la que se modificó la actividad.';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Actividad".activo IS E'Indica si el registro se encuentra activo dentro de la base de datos.';
-- ddl-end --
-- ALTER TABLE plan_adquisiciones."Actividad" OWNER TO postgres;
-- ddl-end --

-- object: "Meta_fk" | type: CONSTRAINT --
-- ALTER TABLE plan_adquisiciones."Actividad" DROP CONSTRAINT IF EXISTS "Meta_fk" CASCADE;
ALTER TABLE plan_adquisiciones."Actividad" ADD CONSTRAINT "Meta_fk" FOREIGN KEY ("Meta_id")
REFERENCES plan_adquisiciones."Meta" (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
-- ddl-end --

-- object: plan_adquisiciones."Registro_plan_adquisiciones" | type: TABLE --
-- DROP TABLE IF EXISTS plan_adquisiciones."Registro_plan_adquisiciones" CASCADE;
CREATE TABLE plan_adquisiciones."Registro_plan_adquisiciones" (
	id serial NOT NULL,
	area_funcional integer NOT NULL,
	centro_gestor integer NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	responsable_id integer NOT NULL,
	activo boolean NOT NULL,
	meta_id varchar,
	producto_id varchar,
	"Plan_adquisiciones_id" integer,
	CONSTRAINT "Registro_plan_adquisiciones_pk" PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE plan_adquisiciones."Registro_plan_adquisiciones" IS E'En esta tabla se guardan los campos comúnes de los renglos de inversión y funcionamiento del plan de adquisiciones.';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Registro_plan_adquisiciones".id IS E'Identificador de la tabla.';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Registro_plan_adquisiciones".area_funcional IS E'Área funcional a la que se le crea el registro.';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Registro_plan_adquisiciones".centro_gestor IS E'Centro gestor que crea el registro del plan de adquisiciones.';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Registro_plan_adquisiciones".fecha_creacion IS E'Fecha de la creación del registro del plan de adquisiciones.';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Registro_plan_adquisiciones".fecha_modificacion IS E'Fecha de la última modificación hecha en el registro.';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Registro_plan_adquisiciones".responsable_id IS E'Identificador del responsable del registro del plan de adquisiciones.';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Registro_plan_adquisiciones".activo IS E'Campo que indica si el registro se encuentra activo dentro de la base de datos.';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Registro_plan_adquisiciones".meta_id IS E'Meta a la que se encuentra asociada el registro';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Registro_plan_adquisiciones".producto_id IS E'Identificador del producto al que está asociado el registro.';
-- ddl-end --
-- ALTER TABLE plan_adquisiciones."Registro_plan_adquisiciones" OWNER TO postgres;
-- ddl-end --

-- object: plan_adquisiciones."Plan_adquisiciones" | type: TABLE --
-- DROP TABLE IF EXISTS plan_adquisiciones."Plan_adquisiciones" CASCADE;
CREATE TABLE plan_adquisiciones."Plan_adquisiciones" (
	id serial NOT NULL,
	descripcion varchar(250) NOT NULL,
	vigencia integer NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	activo boolean NOT NULL,
	CONSTRAINT "Plan_adquisiciones_pk" PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE plan_adquisiciones."Plan_adquisiciones" IS E'Tabla para guardar el plan de adquisiciones.';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Plan_adquisiciones".id IS E'Identificador del plan de adquisiciones.';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Plan_adquisiciones".descripcion IS E'Descripción del plan de adquisiciones';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Plan_adquisiciones".vigencia IS E'Vigencia para la que se desarrolla el plan de adquisiciones.';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Plan_adquisiciones".fecha_creacion IS E'Fecha en la que se creo el registro.';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Plan_adquisiciones".fecha_modificacion IS E'Fecha de la última modificación que se hizo en el registro.';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Plan_adquisiciones".activo IS E'Indica si el registro se encuentra activo en la base de datos.';
-- ddl-end --
-- ALTER TABLE plan_adquisiciones."Plan_adquisiciones" OWNER TO postgres;
-- ddl-end --

-- object: "Plan_adquisiciones_fk" | type: CONSTRAINT --
-- ALTER TABLE plan_adquisiciones."Registro_plan_adquisiciones" DROP CONSTRAINT IF EXISTS "Plan_adquisiciones_fk" CASCADE;
ALTER TABLE plan_adquisiciones."Registro_plan_adquisiciones" ADD CONSTRAINT "Plan_adquisiciones_fk" FOREIGN KEY ("Plan_adquisiciones_id")
REFERENCES plan_adquisiciones."Plan_adquisiciones" (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
-- ddl-end --

-- object: plan_adquisiciones."Registro_funcionamiento-Modalidad_seleccion" | type: TABLE --
-- DROP TABLE IF EXISTS plan_adquisiciones."Registro_funcionamiento-Modalidad_seleccion" CASCADE;
CREATE TABLE plan_adquisiciones."Registro_funcionamiento-Modalidad_seleccion" (
	id serial NOT NULL,
	id_modalidad_seleccion varchar NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL,
	"Registro_plan_adquisiciones_id" integer,
	CONSTRAINT "Registro_funcionamiento-Modalidad_seleccion_pk" PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE plan_adquisiciones."Registro_funcionamiento-Modalidad_seleccion" IS E'Tabla de rompimiento entre el registro de funcionamiento y las posibles modalidades de selección.';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Registro_funcionamiento-Modalidad_seleccion".id IS E'Identificador de la tabla';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Registro_funcionamiento-Modalidad_seleccion".id_modalidad_seleccion IS E'Identificador de la modalidad de selección relacionada al registro de funcionamiento.';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Registro_funcionamiento-Modalidad_seleccion".fecha_modificacion IS E'Fecha de la última modificación realizada al registro';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Registro_funcionamiento-Modalidad_seleccion".activo IS E'Indica si el registro se encuentra activo dentro de la base de datos.';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Registro_funcionamiento-Modalidad_seleccion".fecha_creacion IS E'Fecha en la que se creo el registro.';
-- ddl-end --
-- ALTER TABLE plan_adquisiciones."Registro_funcionamiento-Modalidad_seleccion" OWNER TO postgres;
-- ddl-end --

-- object: plan_adquisiciones."Registro_plan_adquisiciones-Actividad" | type: TABLE --
-- DROP TABLE IF EXISTS plan_adquisiciones."Registro_plan_adquisiciones-Actividad" CASCADE;
CREATE TABLE plan_adquisiciones."Registro_plan_adquisiciones-Actividad" (
	id serial NOT NULL,
	valor integer NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	activo boolean NOT NULL,
	"Actividad_id" integer,
	"Registro_plan_adquisiciones_id" integer,
	CONSTRAINT "Registro_inversion-Actividad_pk" PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE plan_adquisiciones."Registro_plan_adquisiciones-Actividad" IS E'Tabla de rompimiento para las actividades asociadas a un registro del plan de inversion.';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Registro_plan_adquisiciones-Actividad".id IS E'Identificador de la tabla';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Registro_plan_adquisiciones-Actividad".valor IS E'Valor de la actividad';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Registro_plan_adquisiciones-Actividad".fecha_creacion IS E'Fecha en la que se creo el registro';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Registro_plan_adquisiciones-Actividad".fecha_modificacion IS E'Fecha de la última modificación realizada al registro.';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Registro_plan_adquisiciones-Actividad".activo IS E'Indica si el registro se encuentra activo dentro de la base de datos.';
-- ddl-end --
-- ALTER TABLE plan_adquisiciones."Registro_plan_adquisiciones-Actividad" OWNER TO postgres;
-- ddl-end --

-- object: "Actividad_fk" | type: CONSTRAINT --
-- ALTER TABLE plan_adquisiciones."Registro_plan_adquisiciones-Actividad" DROP CONSTRAINT IF EXISTS "Actividad_fk" CASCADE;
ALTER TABLE plan_adquisiciones."Registro_plan_adquisiciones-Actividad" ADD CONSTRAINT "Actividad_fk" FOREIGN KEY ("Actividad_id")
REFERENCES plan_adquisiciones."Actividad" (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
-- ddl-end --

-- object: plan_adquisiciones."Registro_inversion_actividad-Fuente_financiamiento" | type: TABLE --
-- DROP TABLE IF EXISTS plan_adquisiciones."Registro_inversion_actividad-Fuente_financiamiento" CASCADE;
CREATE TABLE plan_adquisiciones."Registro_inversion_actividad-Fuente_financiamiento" (
	id serial NOT NULL,
	fuente_financiamiento_id varchar(20) NOT NULL,
	valor_asignado integer NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL,
	"Registro_plan_adquisiciones-Actividad_id" integer,
	CONSTRAINT "Registro_inversion_actividad-Fuente_financiamiento_pk" PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE plan_adquisiciones."Registro_inversion_actividad-Fuente_financiamiento" IS E'Almacena las fuentes de financiamiento relacionadas con las actividades de un plan de inversión.';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Registro_inversion_actividad-Fuente_financiamiento".id IS E'Identificadro de la tabla';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Registro_inversion_actividad-Fuente_financiamiento".fuente_financiamiento_id IS E'Fuente de financiamiento asociada a al actividad';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Registro_inversion_actividad-Fuente_financiamiento".valor_asignado IS E'Porcentaje de la participación de la fuente de financiamiento dentro del valor total de la actividad.';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Registro_inversion_actividad-Fuente_financiamiento".fecha_modificacion IS E'Fecha de la última modificación realizada al regsitro.';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Registro_inversion_actividad-Fuente_financiamiento".activo IS E'Indica si el regsitro se encuentra activo dentro de la base de datos.';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Registro_inversion_actividad-Fuente_financiamiento".fecha_creacion IS E'Fecha en la que se creo el registro';
-- ddl-end --
-- ALTER TABLE plan_adquisiciones."Registro_inversion_actividad-Fuente_financiamiento" OWNER TO postgres;
-- ddl-end --

-- object: "Registro_plan_adquisiciones-Actividad_fk" | type: CONSTRAINT --
-- ALTER TABLE plan_adquisiciones."Registro_inversion_actividad-Fuente_financiamiento" DROP CONSTRAINT IF EXISTS "Registro_plan_adquisiciones-Actividad_fk" CASCADE;
ALTER TABLE plan_adquisiciones."Registro_inversion_actividad-Fuente_financiamiento" ADD CONSTRAINT "Registro_plan_adquisiciones-Actividad_fk" FOREIGN KEY ("Registro_plan_adquisiciones-Actividad_id")
REFERENCES plan_adquisiciones."Registro_plan_adquisiciones-Actividad" (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
-- ddl-end --

-- object: plan_adquisiciones."Ficha_EB_IMGA" | type: TABLE --
-- DROP TABLE IF EXISTS plan_adquisiciones."Ficha_EB_IMGA" CASCADE;
CREATE TABLE plan_adquisiciones."Ficha_EB_IMGA" (
	id serial NOT NULL,
	meta_id integer NOT NULL,
	proceso varchar(100) NOT NULL,
	magnitud integer NOT NULL,
	unidad_medida varchar(120) NOT NULL,
	descripcion text NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	fecha_creacion timestamp NOT NULL,
	activo boolean NOT NULL,
	rubro varchar NOT NULL,
	"Plan_adquisiciones_id" integer,
	CONSTRAINT "Ficha_EB_IMGA_pk" PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE plan_adquisiciones."Ficha_EB_IMGA" IS E'Se guarda la Ficha estadística básica y de inversión y de metodología general ajustada.';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Ficha_EB_IMGA".id IS E'Identificador de la tabla';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Ficha_EB_IMGA".meta_id IS E'Número de la meta asociada';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Ficha_EB_IMGA".proceso IS E'Verbo que describe el proceso';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Ficha_EB_IMGA".magnitud IS E'Número de unidades a adquirir.(Unidades de la unidad de medida)';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Ficha_EB_IMGA".unidad_medida IS E'Unidad usada como medida en la meta.';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Ficha_EB_IMGA".descripcion IS E'Descripción de la meta.';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Ficha_EB_IMGA".fecha_modificacion IS E'Fecha de la última modificación realizada al registro';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Ficha_EB_IMGA".fecha_creacion IS E'Fecha en la que se creo el registro';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Ficha_EB_IMGA".activo IS E'Indica si el registro se encuentra activo dentro de la base de datos.';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Ficha_EB_IMGA".rubro IS E'Rubro de financiamiento al cual se encuentra asociada la ficha';
-- ddl-end --
-- ALTER TABLE plan_adquisiciones."Ficha_EB_IMGA" OWNER TO postgres;
-- ddl-end --

-- object: "Plan_adquisiciones_fk" | type: CONSTRAINT --
-- ALTER TABLE plan_adquisiciones."Ficha_EB_IMGA" DROP CONSTRAINT IF EXISTS "Plan_adquisiciones_fk" CASCADE;
ALTER TABLE plan_adquisiciones."Ficha_EB_IMGA" ADD CONSTRAINT "Plan_adquisiciones_fk" FOREIGN KEY ("Plan_adquisiciones_id")
REFERENCES plan_adquisiciones."Plan_adquisiciones" (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
-- ddl-end --

-- object: plan_adquisiciones."Registro_plan_adquisiciones-Codigo_arka" | type: TABLE --
-- DROP TABLE IF EXISTS plan_adquisiciones."Registro_plan_adquisiciones-Codigo_arka" CASCADE;
CREATE TABLE plan_adquisiciones."Registro_plan_adquisiciones-Codigo_arka" (
	id serial NOT NULL,
	codigo_arka varchar NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL,
	"Registro_plan_adquisiciones_id" integer,
	CONSTRAINT "Registro_plan_adquisiciones-Codigo_arka_pk" PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE plan_adquisiciones."Registro_plan_adquisiciones-Codigo_arka" IS E'Códigos de Arka asociados al plan de inversion o de funcionamiento.';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Registro_plan_adquisiciones-Codigo_arka".id IS E'Identificador de la tabla';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Registro_plan_adquisiciones-Codigo_arka".codigo_arka IS E'Código de producto que viene desde el subsistema Arka';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Registro_plan_adquisiciones-Codigo_arka".fecha_modificacion IS E'Fecha de la última modificación realizada al registro.';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Registro_plan_adquisiciones-Codigo_arka".activo IS E'Indica si el registro se encuentra activo dentro de la base de datos';
-- ddl-end --
COMMENT ON COLUMN plan_adquisiciones."Registro_plan_adquisiciones-Codigo_arka".fecha_creacion IS E'Fecha en la que se creo el registro';
-- ddl-end --
-- ALTER TABLE plan_adquisiciones."Registro_plan_adquisiciones-Codigo_arka" OWNER TO postgres;
-- ddl-end --

-- object: "Registro_plan_adquisiciones_fk" | type: CONSTRAINT --
-- ALTER TABLE plan_adquisiciones."Registro_plan_adquisiciones-Codigo_arka" DROP CONSTRAINT IF EXISTS "Registro_plan_adquisiciones_fk" CASCADE;
ALTER TABLE plan_adquisiciones."Registro_plan_adquisiciones-Codigo_arka" ADD CONSTRAINT "Registro_plan_adquisiciones_fk" FOREIGN KEY ("Registro_plan_adquisiciones_id")
REFERENCES plan_adquisiciones."Registro_plan_adquisiciones" (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
-- ddl-end --

-- object: "Registro_plan_adquisiciones_fk" | type: CONSTRAINT --
-- ALTER TABLE plan_adquisiciones."Registro_funcionamiento-Modalidad_seleccion" DROP CONSTRAINT IF EXISTS "Registro_plan_adquisiciones_fk" CASCADE;
ALTER TABLE plan_adquisiciones."Registro_funcionamiento-Modalidad_seleccion" ADD CONSTRAINT "Registro_plan_adquisiciones_fk" FOREIGN KEY ("Registro_plan_adquisiciones_id")
REFERENCES plan_adquisiciones."Registro_plan_adquisiciones" (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
-- ddl-end --

-- object: "Registro_plan_adquisiciones_fk" | type: CONSTRAINT --
-- ALTER TABLE plan_adquisiciones."Registro_plan_adquisiciones-Actividad" DROP CONSTRAINT IF EXISTS "Registro_plan_adquisiciones_fk" CASCADE;
ALTER TABLE plan_adquisiciones."Registro_plan_adquisiciones-Actividad" ADD CONSTRAINT "Registro_plan_adquisiciones_fk" FOREIGN KEY ("Registro_plan_adquisiciones_id")
REFERENCES plan_adquisiciones."Registro_plan_adquisiciones" (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
-- ddl-end --

-- object: "Lineamiento_fk" | type: CONSTRAINT --
-- ALTER TABLE plan_adquisiciones."Meta" DROP CONSTRAINT IF EXISTS "Lineamiento_fk" CASCADE;
ALTER TABLE plan_adquisiciones."Meta" ADD CONSTRAINT "Lineamiento_fk" FOREIGN KEY ("Lineamiento_id")
REFERENCES plan_adquisiciones."Lineamiento" (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
-- ddl-end --


