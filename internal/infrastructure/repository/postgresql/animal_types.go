package repo

import (
	"context"
	sql2 "database/sql"
	"time"

	"crm-farmish/internal/entity"
	"crm-farmish/internal/pkg/postgres"
)

// AnimalTypesRepo -.
type AnimalTypesRepo struct {
	db      *postgres.PostgresDB
	table   string
	columns []string
}

// NewAnimalTypesRepo -.
func NewAnimalTypesRepo(pg *postgres.PostgresDB) *AnimalTypesRepo {
	return &AnimalTypesRepo{
		pg,
		"animal_types",
		[]string{"id", "type", "feeding_interval", "watering_interval", "created_at", "updated_at", "deleted_at"}}
}

func (r *AnimalTypesRepo) CreateAnimalTypes(ctx context.Context, req *entity.AnimalTypeCreate) (*entity.AnimalType, error) {
	sql, args, err := r.db.Sq.Builder.
		Insert(r.table).
		Columns("id", "type", "feeding_interval", "watering_interval", "created_at").
		Values(req.ID, req.Type, req.FeedingInterval, req.WateringInterval, time.Now()).
		ToSql()

	if err != nil {
		return nil, err
	}

	_, err = r.db.Exec(ctx, sql, args...)

	return r.GetAnimalTypes(ctx, &entity.FieldValueReq{
		Field: "id",
		Value: req.ID,
	})
}

func (r *AnimalTypesRepo) GetAnimalTypes(ctx context.Context, req *entity.FieldValueReq) (*entity.AnimalType, error) {

	var (
		res   entity.AnimalType
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
		&res.FeedingInterval,
		&res.WateringInterval,
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

func (r *AnimalTypesRepo) ListAnimalTypes(ctx context.Context, req *entity.ListReq) (*entity.ListAnimalTypeRes, error) {
	var (
		res   entity.ListAnimalTypeRes
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
		var animalType entity.AnimalType

		err = rows.Scan(
			&animalType.ID,
			&animalType.Type,
			&animalType.FeedingInterval,
			&animalType.WateringInterval,
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
		res.AnimalTypes = append(res.AnimalTypes, animalType)
	}
	return &res, nil
}

func (r *AnimalTypesRepo) UpdateAnimalsType(ctx context.Context, req *entity.UpdateAnimalTypeReq) (*entity.AnimalType, error) {
	toSql, args, err := r.db.Sq.Builder.Update(r.table).
		SetMap(map[string]interface{}{
			"type":              req.Type,
			"feeding_interval":  req.FeedingInterval,
			"watering_interval": req.WateringInterval,
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

	return r.GetAnimalTypes(ctx, &entity.FieldValueReq{
		Field: "id",
		Value: req.ID,
	})
}

func (r *AnimalTypesRepo) DeleteAnimalTypes(ctx context.Context, req *entity.FieldValueReq) (*entity.StatusRes, error) {
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
