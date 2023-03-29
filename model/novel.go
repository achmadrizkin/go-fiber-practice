package model

import "time"

type Novel struct {
	Id          int       `gorm:"type:int;primary_key" json:"id"`
	Name        string    `gorm:"type:varchar(50);not null" json:"name" binding:"required"`
	Author      string    `gorm:"type:varchar(50);not null" json:"author" binding:"required"`
	Description string    `gorm:"type:varchar(50);not null" json:"description" binding:"required"`
	CreatedAt   time.Time `gorm:"type:time;not null" json:"created_at"`
	UpdatedAt   time.Time `gorm:"type:time;not null" json:"updated_at"`
}
