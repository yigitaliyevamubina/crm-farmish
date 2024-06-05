package repo

import (
	"context"
	"crm-farmish/internal/entity"
	"crm-farmish/internal/pkg/postgres"
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
	toSql, args, err := r.db.Sq.Builder.
		Insert(r.table).
		Columns(r.columns...).
		Values(req.ID, req.AnimalID, req.MaleID, req.FeedingTime).
		ToSql()
	if err != nil {
		return nil, err
	}
	_, err = r.db.Exec(ctx, toSql, args...)
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
