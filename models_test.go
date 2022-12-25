package main

import (
	"testing"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func freshDb(t *testing.T, path ...string) *gorm.DB {
	var dbUri string
	t.Helper()

	// Примечание: path(путь) можно указать в отдельном тесте для отладки
	// цели -- чтобы файл db можно было проверить после запуска теста.
	// Обычно его следует опустить, чтобы была действительно свежая база данных памяти.
	// используется каждый раз.

	if len(path) == 0 {
		//dbUri = ":memory:"
		dbUri = "host=localhost user=postgres password=7 dbname=dvDB port=5432 sslmode=disable TimeZone=Europe/Moscow"
	} else {
		dbUri = path[0]
	}

	db, err := gorm.Open(postgres.Open(dbUri), &gorm.Config{})
	if err != nil {
		t.Fatalf("Ошибка открытия БД памяти: %s", err)
	}
	if err := setupDatabase(db); err != nil {
		t.Fatalf("Ошибка при настройке БД: %s", err)
	}
	return db
}

func TestProductEmpty(t *testing.T) {
	//t.Parallel()
	db := freshDb(t)
	products := []Product{}
	if err := db.Find(&products).Error; err != nil {
		t.Fatalf("Ошибка запроса товаров из тeстовой БД: %s", err)
	}
	got := len(products)
	if got != 0 {
		t.Errorf("Ожидалось 0 товаров, получено %d", got)
	}
}
