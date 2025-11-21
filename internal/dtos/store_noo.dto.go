package dtos

type StoreNooFilter struct {
	Search         string
	Status         string
	OrderDateStart string
	OrderDateEnd   string
}

type StoreNoo struct {
	StoreID           int64   `gorm:"column:store_id" json:"store_id"`
	StoreCode         string  `gorm:"column:store_code" json:"store_code"`
	ReferenceCode     string  `gorm:"column:reference_code" json:"reference_code"`
	StoreName         string  `gorm:"column:store_name" json:"store_name"`
	Address           string  `gorm:"column:address" json:"address"`
	Latitude          float64 `gorm:"column:latitude" json:"latitude"`
	Longitude         float64 `gorm:"column:longitude" json:"longitude"`
	IsGeofencing      bool    `gorm:"column:is_geofencing" json:"is_geofencing"`
	KodePos           string  `gorm:"column:kode_pos" json:"kode_pos"`
	StoreTypeID       int64   `gorm:"column:store_type_id" json:"store_type_id"`
	StoreTypeName     string  `gorm:"column:store_type_name" json:"store_type_name"`
	StatusNoo         string  `gorm:"column:status_noo" json:"status_noo"`
	ImagePath         string  `gorm:"column:image_path" json:"image_path"`
	IsVisitedByVisit  bool    `gorm:"column:is_visited_by_visit" json:"is_visited_by_visit"`
	PicPhoneNumber    string  `gorm:"column:pic_phone_number" json:"pic_phone_number"`
	PicName           string  `gorm:"column:pic_name" json:"pic_name"`
	CreatedBy         int64   `gorm:"column:created_by" json:"created_by"`
	IsNoo             int     `gorm:"column:is_noo" json:"is_noo"`
	IsProspek         int     `gorm:"column:is_prospek" json:"is_prospek"`
	FlagType          string  `gorm:"column:flag_type" json:"flag_type"`
	LastOrderHeaderID string  `gorm:"column:last_order_header_id" json:"last_order_header_id"`
	LastOrderDate     string  `gorm:"column:last_order_date" json:"last_order_date"`
	LastOrderID       string  `gorm:"column:last_order_id" json:"last_order_id"`
	LastReasonNoOrder string  `gorm:"column:last_reason_no_order" json:"last_reason_no_order"`
	LastIsApproved    int     `gorm:"column:last_is_approved" json:"last_is_approved"`
	LastOrderModified string  `gorm:"column:last_order_modified" json:"last_order_modified"`
}

type Pagination struct {
	Total      int64
	Page       int
	Limit      int
	TotalPages int
}

type StoreNooResult struct {
	Pagination Pagination
	Filter     StoreNooFilter
	Stores     []StoreNoo
}
