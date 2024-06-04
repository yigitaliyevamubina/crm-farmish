package repo

import (
	"context"
	sql2 "database/sql"
	"time"

	"crm-farmish/internal/entity"
	"crm-farmish/internal/pkg/postgres"
)

// AnimalsRepo -.
type AnimalsRepo struct {
	db      *postgres.PostgresDB
	table   string
	columns []string
}

// NewAnimalsRepo -.
func NewAnimalsRepo(pg *postgres.PostgresDB) *AnimalsRepo {
	return &AnimalsRepo{
		pg,
		"animals",
		[]string{"id", "type", "name", "gender", "weight", "last_fed_time", "last_watered_time", "disease", "created_at", "updated_at", "deleted_at"}}
}

func (r *AnimalsRepo) CreateAnimals(ctx context.Context, req *entity.AnimalsCreate) (*entity.Animals, error) {
	sql, args, err := r.db.Sq.Builder.
		Insert(r.table).
		Columns("id", "type", "name", "gender", "weight", "last_fed_time", "last_watered_time", "disease", "created_at").
		Values(req.ID, req.Type, req.Name, req.Gender, req.Weight, req.LastFedTime, req.LastWateredTime, req.Disease, time.Now()).
		ToSql()

	if err != nil {
		return nil, err
	}

	_, err = r.db.Exec(ctx, sql, args...)

	if err != nil {
		return nil, err
	}

	return r.GetAnimals(ctx, &entity.FieldValueReq{
		Field: "id",
		Value: req.ID,
	})
}

func (r *AnimalsRepo) GetAnimals(ctx context.Context, req *entity.FieldValueReq) (*entity.Animals, error) {

	var (
		res   entity.Animals
		upAt  sql2.NullTime
		delAt sql2.NullTime
	)
	toSql := r.db.Sq.Builder.
		Select(r.columns...).
		From(r.table)

	if req.Field != "" && req.Value != "" {
		toSql = toSql.
			Where(r.db.Sq.Equal(req.Field, req.Value))
	}

	if !req.DeleteStatus {
		toSql = toSql.Where(r.db.Sq.Equal("deleted_at", nil))
	}
	toSqls, args, err := toSql.ToSql()
	if err != nil {
		return nil, err
	}

	row := r.db.QueryRow(ctx, toSqls, args...)
	err = row.Scan(
		&res.ID,
		&res.Type,
		&res.Name,
		&res.Gender,
		&res.Weight,
		&res.LastFedTime,
		&res.LastWateredTime,
		&res.Disease,
		&res.CreatedAt,
		&upAt,
		&delAt,
	)

	if err != nil {
		return nil, err
	}

	if upAt.Valid {
		res.UpdatedAt = upAt.Time
	}
	if delAt.Valid {
		res.DeletedAt = delAt.Time
	}

	return &res, nil
}

func (r *AnimalsRepo) ListAnimals(ctx context.Context, req *entity.ListReq) (*entity.ListAnimalRes, error) {
	var (
		res   entity.ListAnimalRes
		upAt  sql2.NullTime
		delAt sql2.NullTime
	)

	toSql := r.db.Sq.Builder.
		Select(r.columns...).
		From(r.table)

	countBuilder := r.db.Sq.Builder.Select("count(*)").From(r.table)

	if req.Page >= 1 && req.Limit >= 1 {
		toSql = toSql.
			Limit(req.Limit).
			Offset(req.Limit * (req.Page - 1))
	}

	if req.OrderBy != "" {
		toSql = toSql.OrderBy(req.OrderBy)
	}
	if !req.DeleteStatus {
		countBuilder = countBuilder.Where(r.db.Sq.Equal("deleted_at", nil))
		toSql = toSql.Where(r.db.Sq.Equal("deleted_at", nil))
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
		var animalType entity.Animals

		err = rows.Scan(
			&animalType.ID,
			&animalType.Type,
			&animalType.Name,
			&animalType.Gender,
			&animalType.Weight,
			&animalType.LastFedTime,
			&animalType.LastWateredTime,
			&animalType.Disease,
			&animalType.CreatedAt,
			&upAt,
			&delAt,
		)
		if err != nil {
			return nil, err
		}
		if upAt.Valid {
			animalType.UpdatedAt = upAt.Time
		}
		if delAt.Valid {
			animalType.DeletedAt = delAt.Time
		}
		res.Animals = append(res.Animals, animalType)
	}
	return &res, nil
}

func (r *AnimalsRepo) UpdateAnimals(ctx context.Context, req *entity.UpdateAnimalReq) (*entity.Animals, error) {
	toSql, args, err := r.db.Sq.Builder.Update(r.table).
		SetMap(map[string]interface{}{
			"name":              req.Name,
			"gender":            req.Gender,
			"weight":            req.Weight,
			"last_fed_time":     req.LastFedTime,
			"last_watered_time": req.LastWateredTime,
			"disease":           req.Disease,
			"updated_at":        time.Now(),
		}).
		Where(r.db.Sq.Equal("id", req.ID)).
		ToSql()

	if err != nil {
		return nil, err
	}

	_, err = r.db.Exec(ctx, toSql, args...)
	if err != nil {
		return nil, err
	}

	return r.GetAnimals(ctx, &entity.FieldValueReq{
		Field: "id",
		Value: req.ID,
	})
}

func (r *AnimalsRepo) DeleteAnimals(ctx context.Context, req *entity.FieldValueReq) (*entity.StatusRes, error) {
	if !req.DeleteStatus {
		toSql, args, err := r.db.Sq.Builder.
			Update(r.table).
			Set("deleted_at", time.Now()).
			Where(r.db.Sq.EqualMany(map[string]interface{}{
				"deleted_at": nil,
				req.Field:    req.Value,
			})).
			ToSql()
		if err != nil {
			return &entity.StatusRes{Status: false}, err
		}

		resp, err := r.db.Exec(ctx, toSql, args...)
		if err != nil {
			return &entity.StatusRes{Status: false}, err
		}
		if resp.RowsAffected() > 0 {
			return &entity.StatusRes{Status: true}, nil
		}
		return &entity.StatusRes{Status: false}, nil

	} else {
		toSql, args, err := r.db.Sq.Builder.
			Delete(r.table).
			Where(r.db.Sq.Equal(req.Field, req.Value)).
			ToSql()
		if err != nil {
			return &entity.StatusRes{Status: false}, err
		}

		resp, err := r.db.Exec(ctx, toSql, args...)
		if err != nil {
			return &entity.StatusRes{Status: false}, err
		}

		if resp.RowsAffected() > 0 {
			return &entity.StatusRes{Status: true}, nil
		}
		return &entity.StatusRes{Status: false}, nil
	}

}
