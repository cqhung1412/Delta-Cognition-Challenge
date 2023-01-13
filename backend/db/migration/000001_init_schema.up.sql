CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "email" varchar UNIQUE NOT NULL,
  "hashed_password" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "dogs" (
  "id" bigserial PRIMARY KEY,
  "owner_id" bigint,
  "image" bytea NOT NULL,
  "name" varchar NOT NULL,
  "breed" varchar NOT NULL,
  "birth_year" integer NOT NULL,
  "message" varchar,
  "labels" varchar[],
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "favorite" (
  "user_id" bigint,
  "dog_id" bigint,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  PRIMARY KEY ("user_id", "dog_id")
);

CREATE INDEX ON "dogs" ("owner_id");

CREATE INDEX ON "dogs" ("breed");

CREATE INDEX ON "dogs" ("labels");

CREATE INDEX ON "dogs" ("labels", "breed");

CREATE INDEX ON "favorite" ("user_id", "dog_id");

COMMENT ON COLUMN "dogs"."labels" IS 'get from rekognition';

ALTER TABLE "dogs" ADD FOREIGN KEY ("owner_id") REFERENCES "users" ("id");

ALTER TABLE "favorite" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "favorite" ADD FOREIGN KEY ("dog_id") REFERENCES "dogs" ("id");
