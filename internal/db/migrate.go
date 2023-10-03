package db

import "github.com/jenniekibiri/go-stickers/internal/models"

func SyncDb() {
	DB.AutoMigrate(&models.Sticker{})
}
