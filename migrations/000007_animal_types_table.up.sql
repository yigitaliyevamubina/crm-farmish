CREATE TABLE "animal_types"(
                               "id" UUID NOT NULL,
                               "type" VARCHAR(50) NOT NULL,
                               "feeding_interval" INT NOT NULL,
                               "watering_interval" INT NOT NULL,
                               "created_at" TIMESTAMP ,
                               "updated_at" TIMESTAMP,
                               "deleted_at" TIMESTAMP,
                               PRIMARY KEY("id")
);
