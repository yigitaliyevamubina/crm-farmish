package usecase

import (
	"context"
	"crm-farmish/internal/entity"
	"github.com/google/uuid"
	"time"
)

type TreatmentUseCase struct {
	Repo Treatment
}

func NewTreatmentUseCase(r Treatment) *TreatmentUseCase {
	return &TreatmentUseCase{Repo: r}
}

func (r *TreatmentUseCase) CreateMedicineTreatment(ctx context.Context, req *entity.Treatment) (*entity.Treatment, error) {
	req.ID = uuid.New().String()
	req.TreatmentTime = time.Now()

	return r.Repo.CreateMedicineTreatment(ctx, req)
}

func (r *TreatmentUseCase) GetMedicineTreatment(ctx context.Context, req *entity.FieldValueReq) (*entity.Treatment, error) {
	return r.Repo.GetMedicineTreatment(ctx, req)
}

func (r *TreatmentUseCase) ListMedicineTreatment(ctx context.Context, req *entity.ListTreatmentReq) (*entity.ListTreatment, error) {
	return r.Repo.ListMedicineTreatment(ctx, req)
}
