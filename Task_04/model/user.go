package model

type User struct {
	//gorm.Model
	ID       uint   `gorm:"primaryKey"`
	UserName string `gorm:"column:user_name"`
	Password string `gorm:"column:password"`
	Email    string `gorm:"column:email"`
	Posts    []Post `gorm:"foreignKey:UserID"`
}
