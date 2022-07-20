ALTER TABLE plan_adquisiciones."Registro_plan_adquisiciones"
ADD COLUMN IF NOT EXISTS fecha_estimada_ofertas_inicio TIMESTAMP WITHOUT TIME ZONE;

ALTER TABLE plan_adquisiciones."Registro_plan_adquisiciones"
ADD COLUMN IF NOT EXISTS fecha_estimada_ofertas_fin TIMESTAMP WITHOUT TIME ZONE;

COMMENT ON COLUMN plan_adquisiciones."Registro_plan_adquisiciones".fecha_estimada_ofertas_inicio IS E'Fecha de inicio del proceso de entrega de ofertas.';
COMMENT ON COLUMN plan_adquisiciones."Registro_plan_adquisiciones".fecha_estimada_ofertas_fin IS E'Fecha de fin del proceso de entrega de ofertas.';
