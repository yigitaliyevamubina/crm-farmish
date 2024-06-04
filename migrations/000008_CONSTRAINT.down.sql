ALTER TABLE "treatments" DROP CONSTRAINT "treatments_medicine_id_foreign";
ALTER TABLE "food_warehouse" DROP CONSTRAINT "food_warehouse_animal_type_foreign";
ALTER TABLE "treatments" DROP CONSTRAINT "treatments_animal_id_foreign";
ALTER TABLE "animals_purchases" DROP CONSTRAINT "animals_purchases_animal_id_foreign";
ALTER TABLE "feedings" DROP CONSTRAINT "feedings_animal_id_foreign";
ALTER TABLE "animals" DROP CONSTRAINT "animals_type_foreign";
ALTER TABLE "food_warehouse" DROP CONSTRAINT "food_warehouse_animal_id_foreign";
ALTER TABLE "feedings" DROP CONSTRAINT "feedings_meal_id_foreign";