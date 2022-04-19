package config

import (
	"fmt"
	"majoo-test-debidarmawan/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbConnection struct {
	MajooDB *gorm.DB
}

func InitConnection(URL string, maxOpenConnection int) *DbConnection {
	db, err := gorm.Open(mysql.Open(URL), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		fmt.Println(err)
		fmt.Println("Database connection failed")
		return nil
	}
	dbConn, err := db.DB()
	if err != nil {
		fmt.Println(err)
		fmt.Println("Database connection failed")
		return nil
	} else {
		dbConn.SetMaxOpenConns(maxOpenConnection)
	}
	return &DbConnection{db}
}

func Migrate(dbConn *gorm.DB) {
	dbConn.AutoMigrate(
		&models.User{},
		&models.Merchant{},
		&models.Outlet{},
		&models.Transaction{},
	)
}
