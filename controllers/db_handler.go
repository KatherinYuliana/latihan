package controllers

import (
	"database/sql"
	"log"

	m "latihan/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func connect() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/db_latihan_pbp?parseTime=true&loc=Asia%2FJakarta")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

// Inisialisasi database
func initDB() {
	dsn := "user:password@tcp(localhost:3306)/db_latihan_pbp?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Auto Migrate the schema
	db.AutoMigrate(&m.User{})
}

func conn() *gorm.DB {
	return db
}
