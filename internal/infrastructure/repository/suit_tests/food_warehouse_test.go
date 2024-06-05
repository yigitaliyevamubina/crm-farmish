package suit_tests

import (
	"context"
	"crm-farmish/internal/entity"
	repo "crm-farmish/internal/infrastructure/repository/postgresql"
	"crm-farmish/internal/pkg/config"
	db "crm-farmish/internal/pkg/postgres"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"log"
	"testing"
	"time"
)

type FoodWareHouseTestSuite struct {
	suite.Suite
	CleanUpFunc      func()
	Repository       *repo.FoodWarehouseRepo
	RepositoryType   *repo.AnimalTypesRepo
	RepositoryAnimal *repo.AnimalsRepo
}

func (s *FoodWareHouseTestSuite) SetupTest() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Println(err)
		return
	}
	pgPool, err := db.New(cfg)
	if err != nil {
		log.Fatal(err)
		return
	}
	s.Repository = repo.NewFoodWarehouseRepo(pgPool)
	s.RepositoryType = repo.NewAnimalTypesRepo(pgPool)
	s.RepositoryAnimal = repo.NewAnimalsRepo(pgPool)
	s.CleanUpFunc = pgPool.Close
}

func (s *FoodWareHouseTestSuite) TestAnimalTypesCrud() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(2))
	defer cancel()

	animalTypes := &entity.AnimalTypeCreate{
		ID:               uuid.NewString(),
		Type:             "type",
		FeedingInterval:  1,
		WateringInterval: 1,
	}

	resp, err := s.RepositoryType.CreateAnimalTypes(ctx, animalTypes)
	s.Suite.NoError(err)
	s.Suite.NotNil(resp)
	s.Suite.Equal(resp.ID, animalTypes.ID)
	s.Suite.Equal(resp.Type, animalTypes.Type)
	s.Suite.Equal(resp.FeedingInterval, animalTypes.FeedingInterval)
	s.Suite.Equal(resp.WateringInterval, animalTypes.WateringInterval)

	animals := &entity.AnimalsCreate{
		ID:              uuid.NewString(),
		Type:            animalTypes.ID,
		Name:            "test name",
		Gender:          "male",
		Weight:          2,
		LastFedTime:     time.Now(),
		LastWateredTime: time.Now(),
		Disease:         "test disease",
	}
	respAnimal, err := s.RepositoryAnimal.CreateAnimals(ctx, animals)
	s.Suite.NoError(err)
	s.Suite.NotNil(respAnimal)
	s.Suite.Equal(respAnimal.ID, animals.ID)
	s.Suite.Equal(respAnimal.Type, animals.Type)
	s.Suite.Equal(respAnimal.Name, animals.Name)
	s.Suite.Equal(respAnimal.Gender, animals.Gender)
	s.Suite.Equal(respAnimal.Weight, animals.Weight)
	s.Suite.Equal(respAnimal.LastFedTime.Format("2006-01-02 15:04"), animals.LastFedTime.Format("2006-01-02 15:04"))
	s.Suite.Equal(respAnimal.LastWateredTime.Format("2006-01-02 15:04"), animals.LastWateredTime.Format("2006-01-02 15:04"))
	s.Suite.Equal(respAnimal.Disease, animals.Disease)

	foodWareHouse := &entity.FoodWarehouseCreate{
		ID:           uuid.NewString(),
		Name:         "Name",
		Quantity:     2,
		QuantityType: "quantity type",
		AnimalID:     animals.ID,
		AnimalType:   animalTypes.ID,
		GroupFeeding: false,
	}
	respFoodWareHouse, err := s.Repository.CreateFoodWarehouse(ctx, foodWareHouse)
	s.Suite.NoError(err)
	s.Suite.NotNil(respFoodWareHouse)
	s.Suite.Equal(respFoodWareHouse.ID, foodWareHouse.ID)
	s.Suite.Equal(respFoodWareHouse.Name, foodWareHouse.Name)
	s.Suite.Equal(respFoodWareHouse.Quantity, foodWareHouse.Quantity)
	s.Suite.Equal(respFoodWareHouse.QuantityType, foodWareHouse.QuantityType)
	s.Suite.Equal(respFoodWareHouse.AnimalID, foodWareHouse.AnimalID)
	s.Suite.Equal(respFoodWareHouse.AnimalType, foodWareHouse.AnimalType)
	s.Suite.Equal(respFoodWareHouse.GroupFeeding, foodWareHouse.GroupFeeding)

	getAnimals, err := s.Repository.GetFoodWarehouse(ctx, &entity.FieldValueReq{
		Field:        "id",
		Value:        foodWareHouse.ID,
		DeleteStatus: false,
	})
	s.Suite.NoError(err)
	s.Suite.NotNil(getAnimals)
	s.Suite.Equal(getAnimals.ID, foodWareHouse.ID)
	s.Suite.Equal(getAnimals.Name, foodWareHouse.Name)
	s.Suite.Equal(getAnimals.Quantity, foodWareHouse.Quantity)
	s.Suite.Equal(getAnimals.QuantityType, foodWareHouse.QuantityType)
	s.Suite.Equal(getAnimals.AnimalID, foodWareHouse.AnimalID)
	s.Suite.Equal(getAnimals.AnimalType, foodWareHouse.AnimalType)
	s.Suite.Equal(getAnimals.GroupFeeding, foodWareHouse.GroupFeeding)

	respAll, err := s.Repository.ListFoodWarehouse(ctx, &entity.ListReq{
		Page:         1,
		Limit:        10,
		OrderBy:      "",
		DeleteStatus: false,
	})
	s.Suite.NoError(err)
	s.Suite.NotNil(respAll)

	newUpName := "Update Name"
	newUpQuantityType := "Update Quantity Type"

	updatedFoodWareHouse, err := s.Repository.UpdateFoodWarehouse(ctx, &entity.UpdateFoodWarehouseReq{
		ID:           foodWareHouse.ID,
		Name:         newUpName,
		Quantity:     foodWareHouse.Quantity,
		QuantityType: newUpQuantityType,
		AnimalID:     animals.ID,
		AnimalType:   animalTypes.ID,
		GroupFeeding: foodWareHouse.GroupFeeding,
	})
	s.Suite.NoError(err)
	s.Suite.NotNil(updatedFoodWareHouse)
	s.Suite.Equal(updatedFoodWareHouse.ID, foodWareHouse.ID)
	s.Suite.Equal(updatedFoodWareHouse.Name, newUpName)
	s.Suite.Equal(updatedFoodWareHouse.Quantity, foodWareHouse.Quantity)
	s.Suite.Equal(updatedFoodWareHouse.QuantityType, newUpQuantityType)
	s.Suite.Equal(updatedFoodWareHouse.AnimalID, foodWareHouse.AnimalID)
	s.Suite.Equal(updatedFoodWareHouse.AnimalType, foodWareHouse.AnimalType)
	s.Suite.Equal(updatedFoodWareHouse.GroupFeeding, foodWareHouse.GroupFeeding)

	deleteFoodWareHouse, err := s.Repository.DeleteFoodWarehouse(ctx, &entity.FieldValueReq{
		Field:        "id",
		Value:        foodWareHouse.ID,
		DeleteStatus: true,
	})
	s.Suite.NotNil(deleteFoodWareHouse)
	s.Suite.NoError(err)

	deleteAnimal, err := s.RepositoryAnimal.DeleteAnimals(ctx, &entity.FieldValueReq{
		Field:        "id",
		Value:        animals.ID,
		DeleteStatus: true,
	})
	s.Suite.NotNil(deleteAnimal)
	s.Suite.NoError(err)

	deleteAnimalType, err := s.RepositoryType.DeleteAnimalTypes(ctx, &entity.FieldValueReq{
		Field:        "id",
		Value:        animalTypes.ID,
		DeleteStatus: true,
	})
	s.Suite.NotNil(deleteAnimalType)
	s.Suite.NoError(err)
}

func (s *FoodWareHouseTestSuite) TearDownTest() {
	s.CleanUpFunc()
}

func TestFoodWareHouseTestSuite(t *testing.T) {
	suite.Run(t, new(FoodWareHouseTestSuite))
}
