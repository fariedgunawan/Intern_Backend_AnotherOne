package config

import (
	"Intern_Backend/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDataBase() *gorm.DB {
	username := "root"
	password := "Antimaling2@"    // change this to your local db
	host := "tcp(127.0.0.1:3306)" // change this to your local db
	database := "db_intern"       // change this to your local db

	dsn := fmt.Sprintf("%v:%v@%v/%v?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&models.AdminModel{}, &models.ManagerModel{}, &models.BarangModel{})

	return db
}
