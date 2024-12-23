-- Create "posts" table
CREATE TABLE "public"."posts" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "title" text NOT NULL,
  "description" text NULL,
  "phrase" text NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_posts_deleted_at" to table: "posts"
CREATE INDEX "idx_posts_deleted_at" ON "public"."posts" ("deleted_at");
-- Create "post_locations" table
CREATE TABLE "public"."post_locations" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "post_id" bigint NOT NULL,
  "comment" text NULL,
  "country" text NULL,
  "city" text NULL,
  "address" text NULL,
  "latitude" numeric NULL,
  "longitude" numeric NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_posts_location" FOREIGN KEY ("post_id") REFERENCES "public"."posts" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_post_locations_deleted_at" to table: "post_locations"
CREATE INDEX "idx_post_locations_deleted_at" ON "public"."post_locations" ("deleted_at");
-- Create index "idx_post_locations_post_id" to table: "post_locations"
CREATE UNIQUE INDEX "idx_post_locations_post_id" ON "public"."post_locations" ("post_id");
-- Create "records" table
CREATE TABLE "public"."records" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "post_id" bigint NOT NULL,
  "user_id" text NOT NULL,
  "timestamp" timestamptz NULL,
  "method" bigint NULL DEFAULT 0,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_posts_records" FOREIGN KEY ("post_id") REFERENCES "public"."posts" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_records_deleted_at" to table: "records"
CREATE INDEX "idx_records_deleted_at" ON "public"."records" ("deleted_at");
-- Create index "idx_records_post_id" to table: "records"
CREATE INDEX "idx_records_post_id" ON "public"."records" ("post_id");
-- Create "routes" table
CREATE TABLE "public"."routes" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "title" text NOT NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_routes_deleted_at" to table: "routes"
CREATE INDEX "idx_routes_deleted_at" ON "public"."routes" ("deleted_at");
-- Create "route_milestones" table
CREATE TABLE "public"."route_milestones" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "route_id" bigint NOT NULL,
  "post_id" bigint NOT NULL,
  "priority" bigint NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_routes_milestones" FOREIGN KEY ("route_id") REFERENCES "public"."routes" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_route_milestones_deleted_at" to table: "route_milestones"
CREATE INDEX "idx_route_milestones_deleted_at" ON "public"."route_milestones" ("deleted_at");
-- Create index "idx_route_milestones_post_id" to table: "route_milestones"
CREATE INDEX "idx_route_milestones_post_id" ON "public"."route_milestones" ("post_id");
-- Create index "idx_route_milestones_route_id" to table: "route_milestones"
CREATE INDEX "idx_route_milestones_route_id" ON "public"."route_milestones" ("route_id");
