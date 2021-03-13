package model

type Model struct {
	ID        uint64 `gorm:"primary_key" json:"id"`
	CreatedAt int    `json:"created_at"`
	UpdatedAt int    `json:"updated_at"`
	DeletedAt int    `gorm:"index"`
}