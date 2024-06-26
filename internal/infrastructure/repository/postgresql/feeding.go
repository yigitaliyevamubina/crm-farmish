package repo

import (
	"context"
	"crm-farmish/internal/entity"
	"crm-farmish/internal/pkg/postgres"
	"time"
)

// FeedingRepo -.
type FeedingRepo struct {
	db      *postgres.PostgresDB
	table   string
	columns []string
}

// NewFeedingRepo -.
func NewFeedingRepo(pg *postgres.PostgresDB) *FeedingRepo {
	return &FeedingRepo{
		pg,
		"feedings",
		[]string{"id", "animal_id", "meal_id", "feeding_time"}}
}

func (r *FeedingRepo) CreateFeeding(ctx context.Context, req *entity.Feeding) (*entity.Feeding, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, err
	}
	toSql, args, err := r.db.Sq.Builder.
		Insert(r.table).
		Columns("id", "animal_id", "meal_id", "feeding_time", "created_at").
		Values(req.ID, req.AnimalID, req.MaleID, req.FeedingTime, time.Now()).
		ToSql()
	if err != nil {
		tx.Rollback(ctx)
		return nil, err
	}

	_, err = r.db.Exec(ctx, toSql, args...)

	if err != nil {
		tx.Rollback(ctx)
		return nil, err
	}

	query := `UPDATE food_warehouse SET quantity = quantity - 1 WHERE id = $1;`

	if err != nil {
		tx.Rollback(ctx)
		return nil, err
	}

	_, err = r.db.Exec(ctx, query, req.MaleID)
	if err != nil {
		tx.Rollback(ctx)
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return r.GetFeeding(ctx, &entity.FieldValueReq{Field: "id", Value: req.ID})
}

func (r *FeedingRepo) GetFeeding(ctx context.Context, req *entity.FieldValueReq) (*entity.Feeding, error) {
	var res entity.Feeding
	toSql, args, err := r.db.Sq.Builder.
		Select(r.columns...).
		From(r.table).
		Where(r.db.Sq.Equal(req.Field, req.Value)).
		ToSql()
	if err != nil {
		return nil, err
	}
	row := r.db.QueryRow(ctx, toSql, args...)
	err = row.Scan(&res.ID, &res.AnimalID, &res.MaleID, &res.FeedingTime)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (r *FeedingRepo) ListFeedings(ctx context.Context, req *entity.ListReq) (*entity.ListFeeding, error) {
	var res entity.ListFeeding
	toSql := r.db.Sq.Builder.
		Select(r.columns...).
		From(r.table)

	countBuilder := r.db.Sq.Builder.Select("count(*)").From(r.table)

	if req.Limit >= 1 && req.Page >= 1 {
		toSql = toSql.
			Limit(req.Limit).
			Offset(req.Limit * (req.Page - 1))
	}
	toSqls, args, err := toSql.ToSql()
	if err != nil {
		return nil, err
	}
	row, err := r.db.Query(ctx, toSqls, args...)

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

	for row.Next() {
		var feeding entity.Feeding
		err = row.Scan(&feeding.ID,
			&feeding.AnimalID,
			&feeding.MaleID,
			&feeding.FeedingTime,
		)
		if err != nil {
			return nil, err
		}
		res.Feedings = append(res.Feedings, feeding)
	}
	return &res, nil
}

func (r *FeedingRepo) ListFeedingsByAnimalID(ctx context.Context, req *entity.ListFeedingByAnimalIDReq) (*entity.ListFeeding, error) {
	var res entity.ListFeeding
	toSql := r.db.Sq.Builder.
		Select(r.columns...).
		From(r.table).
		Where(r.db.Sq.Equal("animal_id", req.ID))

	countBuilder := r.db.Sq.Builder.Select("count(*)").From(r.table)

	if req.Limit >= 1 && req.Page >= 1 {
		toSql = toSql.
			Limit(req.Limit).
			Offset(req.Limit * (req.Page - 1))
	}
	toSqls, args, err := toSql.ToSql()
	if err != nil {
		return nil, err
	}
	row, err := r.db.Query(ctx, toSqls, args...)

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

	for row.Next() {
		var feeding entity.Feeding
		err = row.Scan(&feeding.ID,
			&feeding.AnimalID,
			&feeding.MaleID,
			&feeding.FeedingTime,
		)
		if err != nil {
			return nil, err
		}
		res.Feedings = append(res.Feedings, feeding)
	}
	return &res, nil
}

func (r *FeedingRepo) CreateWatering(ctx context.Context, req *entity.Watering) (*entity.Watering, error) {
	toSql, args, err := r.db.Sq.Builder.
		Insert(r.table).
		Columns("id", "animal_id", "watering_time", "created_at").
		Values(req.ID, req.AnimalID, req.WateringTime, time.Now()).
		ToSql()

	if err != nil {
		return nil, err
	}
	_, err = r.db.Exec(ctx, toSql, args...)
	if err != nil {
		return nil, err
	}
	return r.GetWatering(ctx, &entity.FieldValueReq{Field: "id", Value: req.ID})
}

func (r *FeedingRepo) GetWatering(ctx context.Context, req *entity.FieldValueReq) (*entity.Watering, error) {
	toSql, args, err := r.db.Sq.Builder.
		Select("id", "animal_id", "watering_time").
		From(r.table).
		Where(r.db.Sq.Equal(req.Field, req.Value)).
		ToSql()
	if err != nil {
		return nil, err
	}
	var res entity.Watering
	if err = r.db.QueryRow(ctx, toSql, args...).Scan(&res.ID, &res.AnimalID, &res.WateringTime); err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *FeedingRepo) NotWatered(ctx context.Context) (*entity.ListWatering, error) {
	var res entity.ListWatering
	toSql, args, err := r.db.Sq.Builder.
		Select("id", "animal_id", "watering_time").
		From(r.table).
		Where(r.db.Sq.Equal("writer_id", nil)).
		ToSql()

	rows, err := r.db.Query(ctx, toSql, args...)
	if err != nil {
		return nil, err
	}
	countBuilder := r.db.Sq.Builder.Select("count(*)").From(r.table)

	queryCount, _, err := countBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	err = r.db.QueryRow(ctx, queryCount).Scan(&res.Count)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var resw entity.Watering
		err = rows.Scan(&resw.ID, &resw.AnimalID, &resw.WateringTime)
		if err != nil {
			return nil, err
		}
		res.Watering = append(res.Watering, resw)
	}
	return &res, nil
}

func (r *FeedingRepo) NotFeedings(ctx context.Context) (*entity.ListFeeding, error) {
	var res entity.ListFeeding
	toSql, args, err := r.db.Sq.Builder.
		Select("id", "animal_id", "feeding_time").
		From(r.table).
		Where(r.db.Sq.Equal("feeding_time", nil)).
		ToSql()

	countBuilder := r.db.Sq.Builder.Select("count(*)").From(r.table)

	row, err := r.db.Query(ctx, toSql, args...)

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

	for row.Next() {
		var feeding entity.Feeding
		err = row.Scan(&feeding.ID,
			&feeding.AnimalID,
			&feeding.FeedingTime,
		)
		if err != nil {
			return nil, err
		}
		res.Feedings = append(res.Feedings, feeding)
	}
	return &res, nil
}
