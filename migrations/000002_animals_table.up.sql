CREATE TABLE "animals" (
                           "id" UUID NOT NULL,
                           "type" UUID NOT NULL,
                           "name" VARCHAR(255) NOT NULL,
                           "gender" VARCHAR(255) CHECK ("gender" IN ('male', 'female')) NOT NULL,
                           "weight" FLOAT NOT NULL,
                           "last_fed_time" TIMESTAMP(0) WITHOUT TIME ZONE,
                           "last_watered_time" TIMESTAMP(0) WITHOUT TIME ZONE,
                           "disease" VARCHAR(255) DEFAULT '' NOT NULL,
                           "created_at" TIMESTAMP ,
                           "updated_at" TIMESTAMP,
                           "deleted_at" TIMESTAMP,
                           PRIMARY KEY ("id")
);
