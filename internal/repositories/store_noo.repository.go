package repositories

import (
	"context"
	"time"

	"gorm.io/gorm"

	"go-learning/internal/dtos"
)

type StoreRepository struct {
	db *gorm.DB
}

func NewStoreRepository(db *gorm.DB) *StoreRepository {
	return &StoreRepository{db: db}
}

func (r *StoreRepository) GetStoreNoo(
	ctx context.Context,
	userID int64,
	page, limit int,
	filter dtos.StoreNooFilter,
) (*dtos.StoreNooResult, error) {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}
	offset := (page - 1) * limit

	visitDate := time.Now().Format("2006-01-02")

	// base query
	q := r.db.WithContext(ctx).
		Table("store AS s").
		Joins("JOIN selling_order_header AS soh ON soh.order_header_id = s.last_order_header_id").
		Joins("LEFT JOIN store_type AS st ON st.type_id = s.store_type_id").
		Where("s.is_prospek = ?", 1).
		Where("s.created_by = ?", userID).
		Where("soh.order_status <> ?", "no-order")

	// filter search
	if filter.Search != "" {
		like := "%" + filter.Search + "%"
		q = q.Where("(s.store_code LIKE ? OR s.store_name LIKE ?)", like, like)
	}

	// filter status
	switch filter.Status {
	case "waiting":
		q = q.Where("s.flag_type = 'waiting' AND s.is_noo = 0 AND soh.is_approved = 0")
	case "approved":
		q = q.Where("s.flag_type = 'new' AND s.is_noo = 1 AND soh.is_approved = 1")
	case "rejected":
		q = q.Where("s.flag_type = 'rejected' AND s.is_noo = 0 AND soh.is_approved = 2")
	}

	// filter date range
	if filter.OrderDateStart != "" && filter.OrderDateEnd != "" {
		q = q.Where("soh.order_date BETWEEN ? AND ?", filter.OrderDateStart, filter.OrderDateEnd)
	}

	// hitung total (tanpa limit/offset)
	var total int64
	if err := q.
		Select("COUNT(DISTINCT s.store_id)").
		Count(&total).Error; err != nil {
		return nil, err
	}

	// query data
	var stores []dtos.StoreNoo

	err := q.
		Select(`
			s.store_id,
			s.store_code,
			s.reference_code,
			s.store_name,
			s.address,
			s.latitude,
			s.longitude,
			s.is_geofencing,
			IFNULL(s.kode_pos, '') AS kode_pos,
			s.store_type_id,
			st.type_name AS store_type_name,
			s.image_path,
			s.pic_phone_number,
			s.pic_name,
			s.created_by,
			s.is_noo,
			s.is_prospek,
			s.flag_type,
			IFNULL(soh.order_header_id, '') AS last_order_header_id,
			IFNULL(soh.order_date, '') AS last_order_date,
			IFNULL(soh.order_id, '') AS last_order_id,
			IFNULL(soh.reason_no_order, '') AS last_reason_no_order,
			IFNULL(soh.is_approved, 0) AS last_is_approved,
			DATE_FORMAT(IFNULL(soh.modified_date, soh.created_date), '%Y-%m-%d %H:%i:%s') AS last_order_modified,
			CASE
				WHEN s.flag_type = 'waiting' AND s.is_noo = 0 AND soh.is_approved = 0 THEN 'waiting'
				WHEN s.flag_type = 'new'     AND s.is_noo = 1 AND soh.is_approved = 1 THEN 'approved'
				WHEN s.flag_type = 'rejected'AND s.is_noo = 0 AND soh.is_approved = 2 THEN 'rejected'
				ELSE ''
			END AS status_noo,
			EXISTS(
				SELECT 1
				FROM visit vv
				WHERE vv.store_id = s.store_id
				  AND vv.visit_date = ?
				  AND vv.visit_type_id IN (1, 6)
				  AND vv.is_visit = 1
				  AND vv.end_datetime IS NOT NULL
			) AS is_visited_by_visit
		`, visitDate).
		Group("s.store_id").
		Order("soh.is_approved ASC").
		Order("soh.order_date ASC").
		Order("s.store_name ASC").
		Limit(limit).
		Offset(offset).
		Scan(&stores).Error
	if err != nil {
		return nil, err
	}

	totalPages := 0
	if limit > 0 {
		totalPages = int((total + int64(limit) - 1) / int64(limit))
	}

	res := &dtos.StoreNooResult{
		Pagination: dtos.Pagination{
			Total:      total,
			Page:       page,
			Limit:      limit,
			TotalPages: totalPages,
		},
		Filter: filter,
		Stores: stores,
	}

	return res, nil
}
