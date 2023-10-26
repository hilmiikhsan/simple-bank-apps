CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS "customers" (
  "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  "username" VARCHAR(255) UNIQUE NOT NULL,
  "password" VARCHAR(255) NOT NULL,
  "amount" bigint NOT NULL DEFAULT 0,
  "account_number" VARCHAR(255) UNIQUE NOT NULL,
  "account_name" VARCHAR(255) NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NULL
);

CREATE TABLE IF NOT EXISTS "banks" (
  "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  "name" VARCHAR(255) UNIQUE NOT NULL,
  "account_number" VARCHAR(255) UNIQUE NOT NULL,
  "account_name" VARCHAR(255) NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NULL
);

CREATE TABLE IF NOT EXISTS "payments" (
  "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  "customer_id" UUID NOT NULL,
  "amount" bigint NOT NULL DEFAULT 0,
  "account_number" VARCHAR(255) NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NULL
);

ALTER TABLE "payments" ADD CONSTRAINT "payments_customer_id_fkey" FOREIGN KEY ("customer_id") REFERENCES "customers" ("id") ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE "payments" ADD CONSTRAINT "payments_account_number_fkey" FOREIGN KEY ("account_number") REFERENCES "banks" ("account_number") ON DELETE CASCADE ON UPDATE CASCADE;

INSERT INTO "banks" ("name", "account_number", "account_name") VALUES ('BNI', '1234567890', 'Admin'), ('BCA', '0987654321', 'Admin 2');
