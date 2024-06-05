CREATE TABLE "feedings"(
                           "id" UUID NOT NULL,
                           "animal_id" UUID NOT NULL,
                           "meal_id" UUID ,
                           "feeding_time" TIMESTAMP(0) ,
                           "watering_time" TIMESTAMP(0) ,
                            PRIMARY KEY ("id")
);
