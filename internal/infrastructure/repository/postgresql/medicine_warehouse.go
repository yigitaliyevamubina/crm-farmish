package repo

import (
	"context"
	sql2 "database/sql"
	"fmt"
	"time"

	"crm-farmish/internal/entity"
	"crm-farmish/internal/pkg/postgres"
)

// MedicineWarehouse -.
type MedicineWarehouse struct {
	db      *postgres.PostgresDB
	table   string
	columns []string
}

// NewMedicineWarehouseRepo -.
func NewMedicineWarehouseRepo(pg *postgres.PostgresDB) *MedicineWarehouse {
	return &MedicineWarehouse{
		pg,
		"medicine_warehouse",
		[]string{"id", "name", "quantity", "quantity_type", "created_at", "updated_at", "deleted_at"}}
}

func (r *MedicineWarehouse) CreateMedicine(ctx context.Context, req *entity.MedicineWarehouseCreate) (*entity.MedicineWarehouse, error) {
	var (
		res   entity.MedicineWarehouse
		upAt  sql2.NullTime
		delAt sql2.NullTime
	)
	sql, args, err := r.db.Sq.Builder.
		Insert(r.table).
		Columns("id", "name", "quantity", "quantity_type").
		Values(req.ID, req.Name, req.Quantity, req.QuantityType).
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

func (r *MedicineWarehouse) GetMedicine(ctx context.Context, req *entity.FieldValueReq) (*entity.MedicineWarehouse, error) {

	var (
		res   entity.MedicineWarehouse
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
		&res.Name,
		&res.Quantity,
		&res.QuantityType,
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

func (r *MedicineWarehouse) ListMedicine(ctx context.Context, req *entity.ListReq) (*entity.ListMedicineWarehouse, error) {
	var (
		res   entity.ListMedicineWarehouse
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
		var medicine entity.MedicineWarehouse

		err = rows.Scan(
			&medicine.ID,
			&medicine.Name,
			&medicine.Quantity,
			&medicine.QuantityType,
			&medicine.CreatedAt,
			&upAt,
			&delAt,
		)
		if err != nil {
			return nil, err
		}
		if upAt.Valid {
			medicine.UpdatedAt = upAt.Time
		}
		if delAt.Valid {
			medicine.DeletedAt = delAt.Time
		}
		res.MedicineWarehouses = append(res.MedicineWarehouses, medicine)
	}
	return &res, nil
}

func (r *MedicineWarehouse) UpdateMedicineType(ctx context.Context, req *entity.UpdateMedicineWarehouseReq) (*entity.MedicineWarehouse, error) {
	toSql, args, err := r.db.Sq.Builder.Update(r.table).
		SetMap(map[string]interface{}{
			"name":          req.Name,
			"quantity":      req.Quantity,
			"quantity_type": req.QuantityType,
			"updated_at":    time.Now(),
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

	return r.GetMedicine(ctx, &entity.FieldValueReq{
		Field: "id",
		Value: req.ID,
	})
}

func (r *MedicineWarehouse) DeleteMedicine(ctx context.Context, req *entity.FieldValueReq) (*entity.StatusRes, error) {
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
