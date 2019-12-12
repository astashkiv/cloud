-- Database: rivers_database

-- DROP DATABASE rivers_database;

CREATE DATABASE rivers_database
  WITH OWNER = postgres
       ENCODING = 'UTF8'
       TABLESPACE = pg_default
       CONNECTION LIMIT = -1;

-- Table: rivers

-- DROP TABLE rivers;

CREATE TABLE rivers
(
 id serial NOT NULL,
 name character varying NOT NULL,
 groups character varying,
 rating integer,
 arrear integer,
 CONSTRAINT pk_rivers PRIMARY KEY (id )
)
WITH (
 OIDS=FALSE
);
ALTER TABLE rivers
 OWNER TO postgres;
