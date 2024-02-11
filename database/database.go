package database

import (
	"fmt"
	"log"
	"projec1/config"
	"projec1/model"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func Connect() {
	p := config.Config("POSTGRES_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		log.Fatal(err)
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		config.Config("POSTGRES_HOST"), config.Config("POSTGRES_USER"),
		config.Config("POSTGRES_PASSWORD"), config.Config("POSTGRES_DB"), port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("connected success")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("running migrations")
	db.AutoMigrate(&model.User{}, &model.Post{}, &model.IsActive{})
	DB = Dbinstance{
		Db: db,
	}

}
