package models

import "time"

type Transaction struct {
	ID         uint64    `gorm:"column:id;type:bigint(20);not null;autoIncrement" json:"id"`
	MerchantID uint64    `gorm:"column:merchant_id;type:bigint(20);not null" json:"merchant_id"`
	OutletID   uint64    `gorm:"column:outlet_id;type:bigint(20);not null" json:"outlet_id"`
	BillTotal  float64   `gorm:"column:bill_total;type:double;not null" json:"bill_total"`
	CreatedAt  time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"updated_at"`
	CreatedBy  uint64    `gorm:"column:created_by;type:bigint(20);not null" json:"created_by"`
	UpdatedBy  uint64    `gorm:"column:updated_by;type:bigint(20);not null" json:"updated_by"`
}

type DailyOmzet struct {
	Date  string `json:"date"`
	Omzet int64  `json:"omzet"`
}
