package db

import (
	"jwt-auth/migration"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Conn struct {
	config map[string]string
}

func New(config map[string]string) (*gorm.DB, error) {
	// https://github.com/go-gorm/postgres
	dsn := "host=localhost user=postgres password=postgres dbname=auth port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&migration.User{})

	return db, err
}
