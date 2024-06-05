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
