BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS "users" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"user_id"	integer,
	"user_name"	text,
	"user_password"	text,
	"user_type"	text,
	"name"	text,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "zones" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"zone_id"	integer,
	"zone_name"	text,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "beds" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"bed_id"	integer,
	PRIMARY KEY("id")
);
CREATE INDEX IF NOT EXISTS "idx_users_deleted_at" ON "users" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_zones_deleted_at" ON "zones" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_beds_deleted_at" ON "beds" (
	"deleted_at"
);
COMMIT;
