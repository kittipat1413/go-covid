CREATE TABLE patient (
    id       	             SERIAL PRIMARY KEY,
    first_name               VARCHAR NOT NULL,
    last_name                VARCHAR NOT NULL,
    gender    	             VARCHAR NOT NULL,
    age                      SMALLINT NOT NULL CHECK (age > 0),
    latitude                 VARCHAR NOT NULL,
    longitude                VARCHAR NOT NULL,
    city                     VARCHAR NOT NULL,
    district                 VARCHAR NOT NULL,
    street                   VARCHAR NOT NULL,
    created_at 	             TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);