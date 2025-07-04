package main

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"primary key;autoIncrement"`
	Name      string `gorm:"type:varchar(16)"`
	Age       int
	Sex       int `gorm:"type:tinyint"`
	PostCount int
	Posts     []Post
	CreatedAt *time.Time `gorm:"autoCreateTime"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime"`
}

type Post struct {
	ID            uint   `gorm:"primary key;autoIncrement"`
	Title         string `gorm:"type:varchar(16)"`
	Content       string `gorm:"type:text"`
	UserID        uint
	User          User
	CommentStatus int `gorm:"type:tinyint"`
	CommentCount  int
	Comments      []Comment
	CreatedAt     *time.Time `gorm:"autoCreateTime"`
	UpdatedAt     *time.Time `gorm:"autoUpdateTime"`
}

type Comment struct {
	ID        uint   `gorm:"primary key;autoIncrement"`
	Content   string `gorm:"type:text"`
	PostID    uint
	Post      Post
	CreatedAt *time.Time `gorm:"autoCreateTime"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime"`
}

// 钩子函数
// 更新用户的文章数量
func (p *Post) BeforeCreate(tx *gorm.DB) (err error) {
	tx.Model(&User{}).Where("ID = ?", p.UserID).Update("PostCount", gorm.Expr("post_count + ?", 1))
	return
}

// 更新文章的评论数量和评论状态
func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	// 更新文章的评论数量和评论状态
	tx.Model(&Post{}).Select("CommentStatus", "CommentCount").Where("ID = ?", c.PostID).Updates(map[string]interface{}{"CommentStatus": 1, "CommentCount": gorm.Expr("comment_count + ?", 1)})
	return
}

// 更新文章的评论数量和评论状态
func (c *Comment) BeforeDelete(tx *gorm.DB) (err error) {
	tx.Model(&Post{}).Where("ID = ?", c.PostID).Update("CommentCount", gorm.Expr("comment_count - ?", 1))
	var post Post
	tx.First(&post, c.PostID)
	if post.CommentCount == 0 {
		tx.Model(&post).Select("CommentStatus").Update("comment_status", 0)
	}
	return
}

func gormTest(db *gorm.DB) {
	err := db.AutoMigrate(&User{}, &Post{}, &Comment{})
	if err != nil {
		log.Fatal(err)
	}

	users := []User{
		{Name: "张三", Age: 28, Sex: 1, PostCount: 0},
		{Name: "李四", Age: 30, Sex: 0, PostCount: 0},
		{Name: "王五", Age: 32, Sex: 1, PostCount: 0}}
	db.Create(&users)

	posts := []Post{
		{Title: "第一篇博客", Content: "这是一篇测试博客", UserID: 1, CommentStatus: 0, CommentCount: 0},
		{Title: "第二篇博客", Content: "这是第二篇测试博客", UserID: 2, CommentStatus: 0, CommentCount: 0},
		{Title: "第三篇博客", Content: "这是第三篇测试博客", UserID: 3, CommentStatus: 0, CommentCount: 0},
	}
	db.Create(&posts)
	comments := []Comment{
		{Content: "这是第条1评论", PostID: 1},
		{Content: "这是第条2评论", PostID: 2},
		{Content: "这是第条3评论", PostID: 3},
		// {Content: "这是第条4评论", PostID: 1},
		// {Content: "这是第条5评论", PostID: 2},
		// {Content: "这是第条6评论", PostID: 3},
		// {Content: "这是第条7评论", PostID: 1},
		// {Content: "这是第条8评论", PostID: 2},
		// {Content: "这是第条9评论", PostID: 3},
		// {Content: "这是第条10评论", PostID: 1},
		// {Content: "这是第条11评论", PostID: 2},
		// {Content: "这是第条12评论", PostID: 3},
		// {Content: "这是第条13评论", PostID: 1},
		// {Content: "这是第条14评论", PostID: 2},
		// {Content: "这是第条15评论", PostID: 3},
	}
	db.Create(&comments)

	// 查询某个用户发布的所有文章及其对应的评论信息。
	var user User
	db.Preload("Posts").Preload("Posts.Comments").Where("id = ?", 1).First(&user)
	fmt.Println(user)

	// 查询评论数最多的文章
	var post Post
	db.Preload("Comments").Order("comment_count desc").First(&post)
	fmt.Println("评论数最高的文章：", post)

	// 测试删除钩子
	db.Delete(&Comment{ID: 14, PostID: 8})
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/godb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("连接数据库失败:", err)
		return
	}

	// 获取底层 sql.DB 连接池
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get database: %v", err)
	}
	fmt.Println("连接数据库成功")
	defer sqlDB.Close()

	gormTest(db)
}
