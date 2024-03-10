CREATE TABLE transfusion_units (
    id bigserial PRIMARY KEY,
    province_id bigint,
    name text, 
    code varchar
);

ALTER TABLE transfusion_units ADD CONSTRAINT fk_province_id FOREIGN KEY ("province_id") REFERENCES "provinces" ("id");