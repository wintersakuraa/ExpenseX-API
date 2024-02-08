CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create "currencies" table
CREATE TABLE "currencies"
(
  "id"         uuid              NOT NULL DEFAULT uuid_generate_v4(),
  "code"       character varying NOT NULL,
  "name"       character varying NOT NULL,
  "created_at" timestamptz       NOT NULL DEFAULT now(),
  "updated_at" timestamptz       NOT NULL DEFAULT now(),
  PRIMARY KEY ("id")
);

-- Create "users" table
CREATE TABLE "users"
(
  "id"          character varying(36) NOT NULL,
  "currency_id" uuid                  NOT NULL,
  "created_at"  timestamptz           NOT NULL DEFAULT now(),
  "updated_at"  timestamptz           NOT NULL DEFAULT now(),
  PRIMARY KEY ("id"),
  CONSTRAINT "currency_id" FOREIGN KEY ("currency_id") REFERENCES "currencies" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);

-- Create "categories" table
CREATE TABLE "categories"
(
  "id"         uuid                  NOT NULL DEFAULT uuid_generate_v4(),
  "type"       character varying     NOT NULL,
  "name"       character varying     NOT NULL,
  "user_id"    character varying(36) NOT NULL,
  "created_at" timestamptz           NOT NULL DEFAULT now(),
  "updated_at" timestamptz           NOT NULL DEFAULT now(),
  PRIMARY KEY ("id"),
  CONSTRAINT "user_id" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);

-- Create "transactions" table
CREATE TABLE "transactions"
(
  "id"          uuid                  NOT NULL DEFAULT uuid_generate_v4(),
  "type"        character varying     NOT NULL,
  "note"        text                  NOT NULL,
  "amount"      numeric(10, 2)        NOT NULL DEFAULT 0.00,
  "date"        timestamptz           NOT NULL,
  "user_id"     character varying(36) NOT NULL,
  "currency_id" uuid                  NOT NULL,
  "category_id" uuid                  NOT NULL,
  "created_at"  timestamptz           NOT NULL DEFAULT now(),
  "updated_at"  timestamptz           NOT NULL DEFAULT now(),
  PRIMARY KEY ("id"),
  CONSTRAINT "category_id" FOREIGN KEY ("category_id") REFERENCES "categories" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "currency_id" FOREIGN KEY ("currency_id") REFERENCES "currencies" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "user_id" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);


