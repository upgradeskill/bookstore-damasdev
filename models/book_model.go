package models

import "time"

type BookModel struct {
	Id        uint `gorm:"primaryKey"`
	Name      string
	CreatedAt time.Time
}
