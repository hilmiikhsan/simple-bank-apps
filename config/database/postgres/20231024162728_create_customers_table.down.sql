ALTER TABLE "payments" DROP CONSTRAINT IF EXISTS "payments_customer_id_fkey";

DROP TABLE IF EXISTS "payments";
DROP TABLE IF EXISTS "customers";
DROP TABLE IF EXISTS "banks";
DROP EXTENSION IF EXISTS "uuid-ossp";