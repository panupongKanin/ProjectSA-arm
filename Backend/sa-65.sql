BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS "companies" (
	"company_id"	integer,
	"code"	text,
	"name"	text
);
CREATE TABLE IF NOT EXISTS "users" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"name"	text,
	"company_id"	integer,
	CONSTRAINT "fk_users_company" FOREIGN KEY("company_id") REFERENCES "companies"("company_id"),
	PRIMARY KEY("id")
);
CREATE INDEX IF NOT EXISTS "idx_users_deleted_at" ON "users" (
	"deleted_at"
);
COMMIT;
