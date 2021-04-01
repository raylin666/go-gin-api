package model

import (
	"gin-api/pkg/database"
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID        uint64     `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `gorm:"index"`
}

func GetLocalDB() *gorm.DB {
	return database.GetDB("local")
}
