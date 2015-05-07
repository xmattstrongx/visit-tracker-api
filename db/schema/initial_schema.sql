CREATE TABLE IF NOT EXISTS states (
  id SERIAL PRIMARY KEY,
  name character varying(50) UNIQUE NOT NULL,
  abbreviation character varying(2) UNIQUE NOT NULL,
  dateAdded date NOT NULL,
  dateTimeAdded date NOT NULL,
  lastUpdated date NOT NULL
);

COPY states(name,abbreviation,dateAdded,dateTimeAdded,lastUpdated) FROM '/data/State.csv' DELIMITER ',' CSV;

CREATE TABLE IF NOT EXISTS tmp_cities (
  id SERIAL PRIMARY KEY,
  name character varying(100) NOT NULL,
  stateid integer references states(id) NOT NULL,
  county character varying(100) NOT NULL,
  status character varying(50) NOT NULL,
  latitude character varying(50) NOT NULL,
  longitude character varying(50) NOT NULL,
  dateAdded date NOT NULL,
  dateTimeAdded date NOT NULL,
  lastUpdated date NOT NULL
);

COPY tmp_cities(latitude,longitude,name,stateid,county,dateAdded,dateTimeAdded,lastUpdated,status) FROM '/data/Cities_Extended.csv' DELIMITER ',' CSV;

CREATE EXTENSION citext;

CREATE TABLE IF NOT EXISTS cities (
  id SERIAL PRIMARY KEY,
  name CITEXT UNIQUE NOT NULL,
  stateid integer references states(id) NOT NULL,
  county CITEXT NOT NULL,
  status character varying(50) NOT NULL,
  latitude character varying(50) NOT NULL,
  longitude character varying(50) NOT NULL,
  dateAdded date NOT NULL,
  dateTimeAdded date NOT NULL,
  lastUpdated date NOT NULL
);


-- mimicing an upsert to get unique cities from the tmp_cities table
INSERT INTO cities
SELECT * FROM tmp_cities
ON CONFLICT (name) DO NOTHING;
DROP TABLE tmp_cities;


CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users (
  id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v1(),
  firstName character varying(50) NOT NULL,
  lastName character varying(50) NOT NULL,
  dateAdded date NOT NULL,
  dateTimeAdded date NOT NULL,
  lastUpdated date NOT NULL
);

COPY users(firstName,lastName,dateAdded,dateTimeAdded,lastUpdated) FROM '/data/User.csv' DELIMITER ',' CSV;

CREATE TABLE IF NOT EXISTS visits (
  id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v1(),
  userid UUID references users(id) NOT NULL,
  city CITEXT references cities(name) NOT NULL,
  abbreviation character varying(2) references states(abbreviation) NOT NULL,
  dateTimeAdded date NOT NULL DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO visits ( userid, city, abbreviation)
SELECT id, 'Abernant', 'AL'
FROM users 
WHERE firstname = 'Henry';