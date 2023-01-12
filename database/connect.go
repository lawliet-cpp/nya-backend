package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", "containers-us-west-56.railway.app", "postgres","GldyzYjiQPO2IA4rp6FN" , "railway", "7332", &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	DB = db

}
