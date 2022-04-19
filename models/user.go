package models

import "time"

type User struct {
	ID        uint64    `gorm:"column:id;type:bigint(20);not null;autoIncrement" json:"id"`
	Name      string    `gorm:"column:name;type:varchar(45);default:null" json:"name"`
	UserName  string    `gorm:"column:user_name;type:varchar(45);default:null" json:"user_name"`
	Password  string    `gorm:"column:password;type:varchar(255);default:null" json:"password"`
	CreatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"updated_at"`
	CreatedBy uint64    `gorm:"column:created_by;type:bigint(20);not null" json:"created_by"`
	UpdatedBy uint64    `gorm:"column:updated_by;type:bigint(20);not null" json:"updated_by"`
}

type Login struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}
