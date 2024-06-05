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

type AnimalTypesTestSuite struct {
	suite.Suite
	CleanUpFunc func()
	Repository  *repo.AnimalTypesRepo
}

func (s *AnimalTypesTestSuite) SetupTest() {
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
	s.Repository = repo.NewAnimalTypesRepo(pgPool)
	s.CleanUpFunc = pgPool.Close
}

func (s *AnimalTypesTestSuite) TestAnimalTypesCrud() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(2))
	defer cancel()

	animalTypes := &entity.AnimalTypeCreate{
		ID:               uuid.NewString(),
		Type:             "type",
		FeedingInterval:  1,
		WateringInterval: 1,
	}
	resp, err := s.Repository.CreateAnimalTypes(ctx, animalTypes)
	s.Suite.NoError(err)
	s.Suite.NotNil(resp)
	s.Suite.Equal(resp.ID, animalTypes.ID)
	s.Suite.Equal(resp.Type, animalTypes.Type)
	s.Suite.Equal(resp.FeedingInterval, animalTypes.FeedingInterval)
	s.Suite.Equal(resp.WateringInterval, animalTypes.WateringInterval)

	getAnimalTypes, err := s.Repository.GetAnimalTypes(ctx, &entity.FieldValueReq{
		Field:        "id",
		Value:        animalTypes.ID,
		DeleteStatus: false,
	})
	s.Suite.NoError(err)
	s.Suite.NotNil(getAnimalTypes)
	s.Suite.Equal(getAnimalTypes.ID, animalTypes.ID)
	s.Suite.Equal(getAnimalTypes.Type, animalTypes.Type)
	s.Suite.Equal(getAnimalTypes.FeedingInterval, animalTypes.FeedingInterval)
	s.Suite.Equal(getAnimalTypes.FeedingInterval, animalTypes.FeedingInterval)

	respAll, err := s.Repository.ListAnimalTypes(ctx, &entity.ListReq{
		Page:         1,
		Limit:        10,
		OrderBy:      "",
		DeleteStatus: false,
	})
	s.Suite.NoError(err)
	s.Suite.NotNil(respAll)

	newUpType := "Update Type"
	newUpFeedingInterval := 3

	updatedAnimalTypes, err := s.Repository.UpdateAnimalsType(ctx, &entity.UpdateAnimalTypeReq{
		ID:               animalTypes.ID,
		Type:             newUpType,
		FeedingInterval:  newUpFeedingInterval,
		WateringInterval: 1,
	})
	s.Suite.NoError(err)
	s.Suite.NotNil(updatedAnimalTypes)
	s.Suite.NotNil(updatedAnimalTypes.UpdatedAt)
	s.Suite.Equal(updatedAnimalTypes.ID, animalTypes.ID)
	s.Suite.Equal(newUpType, updatedAnimalTypes.Type)
	s.Suite.Equal(newUpFeedingInterval, updatedAnimalTypes.FeedingInterval)
	s.Suite.Equal(updatedAnimalTypes.WateringInterval, animalTypes.WateringInterval)

	deleteDep, err := s.Repository.DeleteAnimalTypes(ctx, &entity.FieldValueReq{
		Field:        "id",
		Value:        animalTypes.ID,
		DeleteStatus: true,
	})
	s.Suite.NotNil(deleteDep)
	s.Suite.NoError(err)
}

func (s *AnimalTypesTestSuite) TearDownTest() {
	s.CleanUpFunc()
}

func TestAnimalTypesTestSuite(t *testing.T) {
	suite.Run(t, new(AnimalTypesTestSuite))
}
