package db

import (
	"jwt-auth/migration"
	"jwt-auth/util"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var log = util.NewLogger("db")

// DBConfig is a data strucute to populate database configurations.
type DBConfig struct {
	*gorm.DB
}

// New initilaizes new database configurations.
func New(config map[string]string) (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=postgres dbname=auth port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&migration.User{})

	return db, err
}

// AutoMigrate automatically migrates the schemas associates with this method.
func (dbc *DBConfig) AutoMigrate() {
	err := dbc.DB.AutoMigrate(&migration.User{})
	if err != nil {
		log.Debug(err)
		os.Exit(1)
	}
}
