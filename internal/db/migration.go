package db

import (
	"log"
	"os"

	"gorm.io/gorm"
)

func Migrations(db *gorm.DB) {
	sqlFile, err := os.ReadFile("../migrations/001_init.sql")
	if err != nil {
		log.Fatalf("Ошибка чтения SQL файла: %v", err)
	}

	sqlScript := string(sqlFile)
	result := db.Exec(sqlScript)

	if result.Error != nil {
		log.Fatalf("Ошибка выполнения SQL файла: %v", result.Error)
	}

	log.Println("Иницилизация завершена табилца успешно создана")
}
