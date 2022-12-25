package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=7 dbname=dvDB port=5432 sslmode=disable TimeZone=Europe/Moscow"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Срыв соединения с БД: %s", err)
	}
	err = setupDatabase(db)
	if err != nil {
		log.Fatalf("Ошибка установки БД: %s", err)
	}

	r := gin.Default()
	setupRouter(r, db)
	err = r.Run(":5000")
	if err != nil {
		log.Fatalf("Ошибка запуска Gin %s", err)
	}
}
