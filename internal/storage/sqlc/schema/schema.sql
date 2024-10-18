CREATE TABLE IF NOT EXISTS "url" (
    "id" INTEGER PRIMARY KEY,
    "alias" varchar(255) NOT NULL UNIQUE,
    "url" varchar(255) NOT NULL
    );
CREATE INDEX IF NOT EXISTS idx_alias ON "url" ("alias");
