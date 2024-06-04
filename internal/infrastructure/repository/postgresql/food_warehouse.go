package repo

import (
	"context"
	"fmt"
	"time"

	"crm-farmish/internal/entity"
	"crm-farmish/internal/pkg/postgres"
)

// FoodWarehouseRepo -.
type FoodWarehouseRepo struct {
	db      *postgres.PostgresDB
	table   string
	columns []string
}

// NewFoodWarehouseRepo -.
func NewFoodWarehouseRepo(pg *postgres.PostgresDB) *FoodWarehouseRepo {
	return &FoodWarehouseRepo{
		pg,
		"food_warehouse",
		[]string{"id", "name", "quantity", "quantity_type", "animal_id", "animal_type", "group_feeding", "created_at", "updated_at", "deleted_at"}}
}

func (r *FoodWarehouseRepo) CreateFoodWarehouse(ctx context.Context, req *entity.FoodWarehouseCreate) (*entity.FoodWarehouse, error) {
	var res entity.FoodWarehouse
	sql, args, err := r.db.Sq.Builder.
		Insert(r.table).
		Columns("id", "name", "quantity", "quantity_type", "animal_id", "animal_type", "group_feeding").
		Values(req.ID, req.Name, req.Quantity, req.QuantityType, req.AnimalID, req.AnimalType, req.GroupFeeding).
		Suffix(fmt.Sprintf("RETURNING %s", r.columns)).
		ToSql()

	if err != nil {
		return nil, err
	}

	row := r.db.QueryRow(ctx, sql, args...)

	err = row.Scan(
		&res.ID,
		&res.Name,
		&res.Quantity,
		&res.QuantityType,
		&res.AnimalID,
		&res.AnimalType,
		&res.GroupFeeding,
	)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *FoodWarehouseRepo) GetFoodWarehouse(ctx context.Context, req *entity.FieldValueReq) (*entity.FoodWarehouse, error) {
	var res entity.FoodWarehouse

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
		&res.Name,
		&res.Quantity,
		&res.QuantityType,
		&res.AnimalID,
		&res.AnimalType,
		&res.GroupFeeding,
	)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *FoodWarehouseRepo) ListFoodWarehouse(ctx context.Context, req *entity.ListReq) (*entity.ListFoodWarehouse, error) {
	var res entity.ListFoodWarehouse

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
		var foodWarehouse entity.FoodWarehouse

		err = rows.Scan(
			&foodWarehouse.ID,
			&foodWarehouse.Name,
			&foodWarehouse.Quantity,
			&foodWarehouse.QuantityType,
			&foodWarehouse.AnimalID,
			&foodWarehouse.AnimalType,
			&foodWarehouse.GroupFeeding,
		)
		if err != nil {
			return nil, err
		}

		res.FoodWarehouses = append(res.FoodWarehouses, foodWarehouse)
	}
	return &res, nil
}

func (r *FoodWarehouseRepo) UpdateFoodWarehouseType(ctx context.Context, req *entity.UpdateFoodWarehouseReq) (*entity.FoodWarehouse, error) {
	toSql, args, err := r.db.Sq.Builder.Update(r.table).
		SetMap(map[string]interface{}{
			"name":         req.Name,
			"quantity":     req.Quantity,
			"quantityType": req.QuantityType,
			"animalID":     req.AnimalID,
			"animalType":   req.AnimalType,
			"groupFeeding": req.GroupFeeding,
			"updated_at":   time.Now(),
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

	return r.GetFoodWarehouse(ctx, &entity.FieldValueReq{
		Field: "id",
		Value: req.ID,
	})
}

func (r *FoodWarehouseRepo) DeleteFoodWarehouse(ctx context.Context, req *entity.FieldValueReq) (*entity.StatusRes, error) {
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
