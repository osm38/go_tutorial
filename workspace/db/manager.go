package db

import (
	"sample/query"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Query(logLevel logger.LogLevel) *query.Query {
	dsn := "host=db user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("db connection error!!")
	}
	db.Logger.LogMode(logLevel)

	query := query.Use(db)
	return query
}
