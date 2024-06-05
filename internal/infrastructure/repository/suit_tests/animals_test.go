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

type AnimalsTestSuite struct {
	suite.Suite
	CleanUpFunc    func()
	Repository     *repo.AnimalsRepo
	RepositoryType *repo.AnimalTypesRepo
}

func (s *AnimalsTestSuite) SetupTest() {
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
	s.Repository = repo.NewAnimalsRepo(pgPool)
	s.RepositoryType = repo.NewAnimalTypesRepo(pgPool)
	s.CleanUpFunc = pgPool.Close
}

func (s *AnimalsTestSuite) TestAnimalTypesCrud() {
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
	respAnimal, err := s.Repository.CreateAnimals(ctx, animals)
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

	getAnimals, err := s.Repository.GetAnimals(ctx, &entity.FieldValueReq{
		Field:        "id",
		Value:        animals.ID,
		DeleteStatus: false,
	})
	s.Suite.NoError(err)
	s.Suite.NotNil(getAnimals)
	s.Suite.Equal(getAnimals.ID, animals.ID)
	s.Suite.Equal(getAnimals.Type, animals.Type)
	s.Suite.Equal(getAnimals.Name, animals.Name)
	s.Suite.Equal(getAnimals.Gender, animals.Gender)
	s.Suite.Equal(getAnimals.Weight, animals.Weight)
	s.Suite.Equal(getAnimals.LastFedTime.Format("2006-01-02 15:04"), animals.LastFedTime.Format("2006-01-02 15:04"))
	s.Suite.Equal(getAnimals.LastWateredTime.Format("2006-01-02 15:04"), animals.LastWateredTime.Format("2006-01-02 15:04"))
	s.Suite.Equal(getAnimals.Disease, animals.Disease)

	respAll, err := s.Repository.ListAnimals(ctx, &entity.ListReq{
		Page:         1,
		Limit:        10,
		OrderBy:      "",
		DeleteStatus: false,
	})
	s.Suite.NoError(err)
	s.Suite.NotNil(respAll)

	newUpName := "Update Name"
	newUpWeight := 3
	newUpLastWateredTime := time.Now()

	updatedAnimals, err := s.Repository.UpdateAnimals(ctx, &entity.UpdateAnimalReq{
		ID:              animals.ID,
		Name:            newUpName,
		Gender:          animals.Gender,
		Weight:          newUpWeight,
		LastFedTime:     animals.LastFedTime,
		LastWateredTime: newUpLastWateredTime,
		Disease:         animals.Disease,
	})
	s.Suite.NoError(err)
	s.Suite.NotNil(updatedAnimals)
	s.Suite.NotNil(updatedAnimals.UpdatedAt)
	s.Suite.Equal(updatedAnimals.ID, animals.ID)
	s.Suite.Equal(updatedAnimals.Type, animals.Type)
	s.Suite.Equal(newUpName, updatedAnimals.Name)
	s.Suite.Equal(updatedAnimals.Gender, animals.Gender)
	s.Suite.Equal(newUpWeight, updatedAnimals.Weight)
	s.Suite.Equal(updatedAnimals.LastFedTime.Format("2006-01-02 15:04"), animals.LastFedTime.Format("2006-01-02 15:04"))
	s.Suite.Equal(newUpLastWateredTime.Format("2006-01-02 15:04"), animals.LastWateredTime.Format("2006-01-02 15:04"))
	s.Suite.Equal(getAnimals.Disease, animals.Disease)

	deleteAnimal, err := s.Repository.DeleteAnimals(ctx, &entity.FieldValueReq{
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

func (s *AnimalsTestSuite) TearDownTest() {
	s.CleanUpFunc()
}

func TestAnimalsTestSuite(t *testing.T) {
	suite.Run(t, new(AnimalsTestSuite))
}
