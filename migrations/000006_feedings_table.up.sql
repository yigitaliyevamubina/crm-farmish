CREATE TABLE "feedings"(
                           "id" UUID NOT NULL,
                           "animal_id" UUID NOT NULL,
                           "meal_id" UUID NOT NULL,
                           "feeding_time" TIMESTAMP(0) WITHOUT TIME ZONE,
                            PRIMARY KEY ("id")
);
