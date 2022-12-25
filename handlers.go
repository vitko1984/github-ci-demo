package main

import (
	"net/http"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func homePageHandler(c *gin.Context) {
	db := c.Value("database").(*gorm.DB)
	products := []Product{}
	if err := db.Find(&products).Error; err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.HTML(http.StatusOK, "pages/home.page.html", gin.H{"products": products})
}
