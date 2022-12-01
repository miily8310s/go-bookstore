package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db * gorm.DB
)

func Connect()  {
	// TODO: dbのurl
	dsn := ""
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB  {
	return db
}