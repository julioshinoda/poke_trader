CREATE SEQUENCE IF NOT EXISTS trade_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	CACHE 1
	NO CYCLE;

CREATE TABLE IF NOT EXISTS trade (
	id bigint NOT NULL DEFAULT nextval('trade_id_seq'::regclass),
	trainerOne jsonb NULL,
	trainerTwo jsonb NULL,
	created_at date NULL,
	fair boolean NULL DEFAULT false,
	CONSTRAINT trade_pk PRIMARY KEY (id)
);

