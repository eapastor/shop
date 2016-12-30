DROP SCHEMA IF EXISTS "public" CASCADE;

DROP EXTENSION IF EXISTS "uuid-ossp";

------------------------------------------------------

CREATE SCHEMA IF NOT EXISTS "public";

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS "items" (
  id            UUID            PRIMARY KEY     NOT NULL    DEFAULT uuid_generate_v4(),
  name          VARCHAR(50)     UNIQUE          NOT NULL,
  description   VARCHAR(200),
  price         NUMERIC                         NOT NULL,
  category_id   UUID                            NOT NULL    DEFAULT '22bee679-2bf8-4d88-ae09-311470e7ce08',
  created       TIMESTAMP                                   DEFAULT now(),
  updated       TIMESTAMP                                   DEFAULT now()
);

CREATE TABLE "categories" (
  id            UUID            PRIMARY KEY     NOT NULL    DEFAULT uuid_generate_v4(),
  name          VARCHAR(50)     UNIQUE          NOT NULL,
  created       TIMESTAMP                                   DEFAULT now(),
  updated       TIMESTAMP                                   DEFAULT now()
);

CREATE TABLE IF NOT EXISTS "users" (
  id            UUID            PRIMARY KEY     NOT NULL    DEFAULT uuid_generate_v4(),
  created       TIMESTAMP                                   DEFAULT now(),
  updated       TIMESTAMP                                   DEFAULT now()
);

CREATE TABLE IF NOT EXISTS "carts" (
  id            UUID            PRIMARY KEY     NOT NULL    DEFAULT uuid_generate_v4(),
  created       TIMESTAMP                                   DEFAULT now(),
  updated       TIMESTAMP                                   DEFAULT now()
);

CREATE TABLE IF NOT EXISTS "orders" (
  id            UUID            PRIMARY KEY     NOT NULL    DEFAULT uuid_generate_v4(),
  created       TIMESTAMP                                   DEFAULT now(),
  updated       TIMESTAMP                                   DEFAULT now()
);

------------------------------------------------------

INSERT INTO "categories" (id, name)
VALUES
  ('22bee679-2bf8-4d88-ae09-311470e7ce08', 'No category');