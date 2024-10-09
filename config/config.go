package config

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// InitializeDatabase initializes and returns a database connection with connection pooling
func InitializeDatabase() (*gorm.DB, error) {
	// DSN (Data Source Name) untuk MySQL
	dsn := "root:Insider2816.@tcp(127.0.0.1:3306)/travel_db?charset=utf8mb4&parseTime=True&loc=Local"

	// Opsi tambahan untuk GORM (misalnya, menggunakan logger)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Atur log mode sesuai kebutuhan
	})
	if err != nil {
		log.Println("Error connecting to the database: ", err)
		return nil, err
	}

	// Mendapatkan koneksi *sql.DB dari GORM untuk mengatur connection pooling
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// Set the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100) // Jumlah maksimal koneksi yang aktif

	// Set the maximum number of idle connections to the database.
	sqlDB.SetMaxIdleConns(10) // Jumlah koneksi idle yang dibiarkan terbuka

	// Set the maximum lifetime of a connection.
	sqlDB.SetConnMaxLifetime(30 * time.Minute) // Waktu maksimal koneksi bisa digunakan

	return db, nil
}
