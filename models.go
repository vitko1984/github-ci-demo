package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint
	ProductName string
	Description string
	Photo       string
	Price       int64
}

func setupDatabase(db *gorm.DB) error {
	err := db.AutoMigrate(
		&Product{},
	)
	if err != nil {
		return fmt.Errorf("Ошибка миграции ДБ: %s", err)
	}
	return nil
}

func connectDatabase(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("database", db)
	}
}
