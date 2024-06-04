ALTER TABLE
    "treatments" ADD CONSTRAINT "treatments_medicine_id_foreign" FOREIGN KEY("medicine_id") REFERENCES "medicine_warehouse"("id");
ALTER TABLE
    "food_warehouse" ADD CONSTRAINT "food_warehouse_animal_type_foreign" FOREIGN KEY("animal_type") REFERENCES "animal_types"("id");
ALTER TABLE
    "treatments" ADD CONSTRAINT "treatments_animal_id_foreign" FOREIGN KEY("animal_id") REFERENCES "animals"("id");
ALTER TABLE
    "animals_purchases" ADD CONSTRAINT "animals_purchases_animal_id_foreign" FOREIGN KEY("animal_id") REFERENCES "animals"("id");
ALTER TABLE
    "feedings" ADD CONSTRAINT "feedings_animal_id_foreign" FOREIGN KEY("animal_id") REFERENCES "animals"("id");
ALTER TABLE
    "animals" ADD CONSTRAINT "animals_type_foreign" FOREIGN KEY("type") REFERENCES "animal_types"("id");
ALTER TABLE
    "food_warehouse" ADD CONSTRAINT "food_warehouse_animal_id_foreign" FOREIGN KEY("animal_id") REFERENCES "animals"("id");
ALTER TABLE
    "feedings" ADD CONSTRAINT "feedings_meal_id_foreign" FOREIGN KEY("meal_id") REFERENCES "food_warehouse"("id");