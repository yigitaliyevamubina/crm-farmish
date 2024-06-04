CREATE TABLE "animals_purchases"(
                                    "id" UUID NOT NULL,
                                    "animal_id" UUID NOT NULL,
                                    "name" VARCHAR(255) NOT NULL,
                                    "total_price" FLOAT NOT NULL,
                                    "count" INT NOT NULL,
                                    "created_at" TIMESTAMP ,
                                    "updated_at" TIMESTAMP ,
                                    "deleted_at" TIMESTAMP,
                                    PRIMARY KEY("id")
);
