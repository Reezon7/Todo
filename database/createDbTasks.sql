CREATE DATABASE IF NOT EXISTS apptasks

CREATE TABLE IF NOT EXISTS public.taskrepo
(
    id serial NOT NULL,
    name character varying NOT NULL,
    description character varying,
    "start" date NOT NULL,
    "end" date,
    active boolean NOT NULL,
    PRIMARY KEY (id)
);