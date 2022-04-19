package models

import "time"

type Outlet struct {
	ID         uint64    `gorm:"column:id;type:bigint(20);not null;autoIncrement" json:"id"`
	MerchantID uint64    `gorm:"column:merchant_id;type:bigint(20);not null" json:"merchant_id"`
	OutletName string    `gorm:"column:outlet_name;type:varchar(40);not null" json:"outlet_name"`
	CreatedAt  time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"updated_at"`
	CreatedBy  uint64    `gorm:"column:created_by;type:bigint(20);not null" json:"created_by"`
	UpdatedBy  uint64    `gorm:"column:updated_by;type:bigint(20);not null" json:"updated_by"`
}

type OutletData struct {
	ID         uint64       `json:"id"`
	OutletName string       `json:"outlet_name"`
	DailyOmzet []DailyOmzet `json:"daily_omzet"`
}
