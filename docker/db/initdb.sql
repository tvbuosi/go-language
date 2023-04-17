CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE SCHEMA hex_golang;

CREATE TABLE public.vehicles (
	id uuid NOT NULL,
	"name" text NOT NULL,
	"description" text NOT NULL,
	"type" text NOT NULL,
	created_at timestamp NOT NULL,
	created_by timestamp NOT NULL,
	updated_at timestamp NOT NULL,
	updated_by timestamp NOT NULL,
	rate numeric NOT NULL,

    PRIMARY KEY (id)
);