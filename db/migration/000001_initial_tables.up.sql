CREATE TABLE "users"(
    "id" SERIAL PRIMARY KEY NOT NULL,
    "username" VARCHAR NOT NULL,
    "password" VARCHAR NOT NULL,
    "username" VARCHAR UNIQUE NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT (NOW())
)