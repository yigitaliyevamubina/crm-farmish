package usecase

import (
	"context"
	"crm-farmish/internal/entity"
	"github.com/google/uuid"
	"time"
)

type FeedingUseCase struct {
	Repo Feeding
}

func NewFeedingUseCase(r Feeding) *FeedingUseCase {
	return &FeedingUseCase{Repo: r}
}

func (r *FeedingUseCase) CreateFeeding(ctx context.Context, req *entity.Feeding) (*entity.Feeding, error) {
	req.ID = uuid.New().String()
	req.FeedingTime = time.Now()

	return r.Repo.CreateFeeding(ctx, req)
}

func (r *FeedingUseCase) GetFeeding(ctx context.Context, req *entity.FieldValueReq) (*entity.Feeding, error) {
	return r.Repo.GetFeeding(ctx, req)
}

func (r *FeedingUseCase) ListFeedings(ctx context.Context, req *entity.ListReq) (*entity.ListFeeding, error) {
	return r.Repo.ListFeedings(ctx, req)
}

func (r *FeedingUseCase) ListFeedingsByAnimalID(ctx context.Context, req *entity.ListFeedingByAnimalIDReq) (*entity.ListFeeding, error) {
	return r.Repo.ListFeedingsByAnimalID(ctx, req)
}

func (r *FeedingUseCase) CreateWatering(ctx context.Context, req *entity.Watering) (*entity.Watering, error) {
	req.ID = uuid.New().String()
	req.WateringTime = time.Now()

	return r.Repo.CreateWatering(ctx, req)
}

func (r *FeedingUseCase) GetWatering(ctx context.Context, req *entity.FieldValueReq) (*entity.Watering, error) {
	return r.Repo.GetWatering(ctx, req)
}

func (r *FeedingUseCase) NotWatered(ctx context.Context) (*entity.ListWatering, error) {
	return r.Repo.NotWatered(ctx)
}

func (r *FeedingUseCase) NotFeedings(ctx context.Context) (*entity.ListFeeding, error) {
	return r.Repo.NotFeedings(ctx)
}
