package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db * gorm.DB
)

func Connect()  {
	// dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_DATABASE"))
	dsn := "test:password@tcp(127.0.0.1:3306)/bookstore?charset=utf8"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB  {
	return db
}