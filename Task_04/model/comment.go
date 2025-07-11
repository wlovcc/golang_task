package model

import (
	"time"
)

// 存储评论信息
type Comment struct {
	//gorm.Model
	ID        uint      `gorm:"primaryKey"`
	Content   string    `gorm:"column:content"`
	UserID    uint      `gorm:"column:user_id"`
	PostID    uint      `gorm:"column:post_id"`
	CreatedAt time.Time `gorm:"column:created_at"`
}
