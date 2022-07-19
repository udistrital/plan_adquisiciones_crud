ALTER TABLE plan_adquisiciones."Registro_plan_adquisiciones"
ADD COLUMN IF NOT EXISTS fecha_estimada_ofertas_inicio TIMESTAMP WITH TIME ZONE;

ALTER TABLE plan_adquisiciones."Registro_plan_adquisiciones"
ADD COLUMN IF NOT EXISTS fecha_estimada_ofertas_fin TIMESTAMP WITH TIME ZONE;