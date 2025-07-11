package model

import (
	"time"
)

// post表存储博客文章信息
type Post struct {
	//gorm.Model
	ID        uint      `gorm:"primaryKey"`
	Title     string    `gorm:"column:title"`
	Content   string    `gorm:"column:content"`
	UserID    uint      `gorm:"column:user_id"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	Comments  []Comment `gorm:"foreignKey:PostID"`
}
