CREATE TABLE "schema_migration" (
"version" TEXT NOT NULL
);
CREATE UNIQUE INDEX "version_idx" ON "schema_migration" (version);
CREATE TABLE "users" (
"id" TEXT PRIMARY KEY,
"name" TEXT NOT NULL,
"email" TEXT NOT NULL,
"bio" text,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
