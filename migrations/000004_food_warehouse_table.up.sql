CREATE TABLE "food_warehouse"(
                                 "id" UUID NOT NULL,
                                 "name" VARCHAR(100) NOT NULL,
                                 "quantity" INT NOT NULL,
                                 "quantity_type" VARCHAR(100) NOT NULL,
                                 "animal_id" UUID NOT NULL,
                                 "animal_type" UUID NOT NULL,
                                 "group_feeding" BOOLEAN NOT NULL,
                                 "created_at" TIMESTAMP ,
                                 "updated_at" TIMESTAMP ,
                                 "deleted_at" TIMESTAMP,
                                 PRIMARY KEY("id")
);