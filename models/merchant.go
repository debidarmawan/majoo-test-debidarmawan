package models

import "time"

type Merchant struct {
	ID           uint64    `gorm:"column:id;type:bigint(20);not null;autoIncrement" json:"id"`
	UserID       uint64    `gorm:"column:user_id;type:bigint(20);not null" json:"user_id"`
	MerchantName string    `gorm:"column:merchant_name;type:varchar(40);not null" json:"merchant_name"`
	CreatedAt    time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"updated_at"`
	CreatedBy    uint64    `gorm:"column:created_by;type:bigint(20);not null" json:"created_by"`
	UpdatedBy    uint64    `gorm:"column:updated_by;type:bigint(20);not null" json:"updated_by"`
}

type MerchantOmzet struct {
	UserID uint64 `query:"user_id"`
	Period string `query:"period"`
	Limit  int    `query:"limit"`
	Page   int    `query:"page"`
}

type MerchantOmzetResponse struct {
	MerchantID   uint64       `json:"merchant_id"`
	MerchantName string       `json:"merchant_name"`
	DailyOmzet   []DailyOmzet `json:"daily_omzet"`
}

type MerchantOutletOmzet struct {
	UserID   uint64 `query:"user_id"`
	OutletID uint64 `query:"outlet_id"`
	Period   string `query:"period"`
	Limit    int    `query:"limit"`
	Page     int    `query:"page"`
}

type MerchantOutletOmzetResponse struct {
	MerchantID   uint64       `json:"merchant_id"`
	MerchantName string       `json:"merchant_name"`
	OutletsData  []OutletData `json:"outlets"`
}
