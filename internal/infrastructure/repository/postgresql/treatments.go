package repo

import (
	"context"
	"crm-farmish/internal/entity"
	"crm-farmish/internal/pkg/postgres"
)

// Treatment -.
type Treatment struct {
	db      *postgres.PostgresDB
	table   string
	columns []string
}

// NewTreatmentRepo -.
func NewTreatmentRepo(pg *postgres.PostgresDB) *Treatment {
	return &Treatment{
		pg,
		"treatments",
		[]string{"id", "animal_id", "medicine_id", "treatment_time"}}
}

func (r *Treatment) CreateMedicineTreatment(ctx context.Context, req *entity.Treatment) (*entity.Treatment, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, err
	}
	sql, args, err := r.db.Sq.Builder.
		Insert(r.table).
		Columns(r.columns...).
		Values(req.ID, req.AnimalID, req.MedicineID, req.TreatmentTime).
		ToSql()

	if err != nil {
		tx.Rollback(ctx)
		return nil, err
	}

	_, err = r.db.Exec(ctx, sql, args...)

	if err != nil {
		tx.Rollback(ctx)
		return nil, err
	}

	query := `UPDATE medicine_warehouse SET quantity = quantity - 1 WHERE id = $1;`

	if err != nil {
		tx.Rollback(ctx)
		return nil, err
	}

	_, err = r.db.Exec(ctx, query, req.MedicineID)
	if err != nil {
		tx.Rollback(ctx)
		return nil, err
	}

	if err = tx.Commit(ctx); err != nil {
		return nil, err
	}

	return r.GetMedicineTreatment(ctx, &entity.FieldValueReq{
		Field: "id",
		Value: req.ID,
	})
}

func (r *Treatment) GetMedicineTreatment(ctx context.Context, req *entity.FieldValueReq) (*entity.Treatment, error) {

	var res entity.Treatment

	toSql := r.db.Sq.Builder.
		Select(r.columns...).
		From(r.table)

	if req.Field != "" && req.Value != "" {
		toSql = toSql.
			Where(r.db.Sq.Equal(req.Field, req.Value))
	}

	toSqls, args, err := toSql.ToSql()
	if err != nil {
		return nil, err
	}

	row := r.db.QueryRow(ctx, toSqls, args...)

	err = row.Scan(
		&res.ID,
		&res.AnimalID,
		&res.MedicineID,
		&res.TreatmentTime,
	)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *Treatment) ListMedicineTreatment(ctx context.Context, req *entity.ListTreatmentReq) (*entity.ListTreatment, error) {
	var res entity.ListTreatment

	toSql := r.db.Sq.Builder.
		Select(r.columns...).
		From(r.table)

	countBuilder := r.db.Sq.Builder.Select("count(*)").From(r.table)

	if req.Page >= 1 && req.Limit >= 1 {
		toSql = toSql.
			Limit(req.Limit).
			Offset(req.Limit * (req.Page - 1))
	}

	if req.Field != "" && req.Value != "" {
		toSql = toSql.
			Where(r.db.Sq.Equal(req.Field, req.Value))
	}

	toSqls, args, err := toSql.ToSql()

	if err != nil {
		return nil, err
	}

	queryCount, _, err := countBuilder.ToSql()

	if err != nil {
		return nil, err
	}

	err = r.db.QueryRow(ctx, queryCount).Scan(&res.Count)

	if err != nil {
		return nil, err
	}

	rows, err := r.db.Query(ctx, toSqls, args...)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var medicine entity.Treatment

		err = rows.Scan(
			&medicine.ID,
			&medicine.AnimalID,
			&medicine.MedicineID,
			&medicine.TreatmentTime,
		)
		if err != nil {
			return nil, err
		}

		res.Treatment = append(res.Treatment, medicine)
	}
	return &res, nil
}
