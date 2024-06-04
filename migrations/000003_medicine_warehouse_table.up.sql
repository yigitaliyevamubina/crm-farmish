CREATE TABLE "medicine_warehouse" (
                                      "id" UUID NOT NULL,
                                      "name" VARCHAR(150) NOT NULL,
                                      "quantity" INT NOT NULL,
                                      "quantity_type" VARCHAR(100) NOT NULL,
                                      "created_at" TIMESTAMP ,
                                      "updated_at" TIMESTAMP,
                                      "deleted_at" TIMESTAMP,
                                      PRIMARY KEY ("id")
);