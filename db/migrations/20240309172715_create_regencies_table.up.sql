CREATE TABLE regencies (
    id bigserial PRIMARY KEY,
    province_id bigint,
    name text, 
    udd_code varchar
);

ALTER TABLE regencies ADD CONSTRAINT fk_province_id FOREIGN KEY ("province_id") REFERENCES "provinces" ("id");