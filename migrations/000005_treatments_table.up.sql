CREATE TABLE "treatments"(
                             "id" UUID NOT NULL,
                             "animal_id" UUID NOT NULL,
                             "medicine_id" UUID NOT NULL,
                             "treatment_time" TIMESTAMP(0) WITHOUT TIME ZONE,
                             PRiMARY KEY("id")
);
