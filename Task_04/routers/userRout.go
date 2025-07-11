package routers

import (
	"Task_04/config"
	"Task_04/model"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// 用户注册
func Register(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = string(hashedPassword)
	//插入数据库
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// 用户登录
func Login(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var storedUser model.User
	if err := config.DB.Where("username = ?", user.UserName).First(&storedUser).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// 生成 JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       storedUser.ID,
		"username": storedUser.UserName,
		"exp":      time.Now().Add(time.Hour * 2).Unix(),
	})
	//通过密钥签名
	tokenString, err := token.SignedString([]byte(config.SecrectKey))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	//设置 Cookie?
	//c.SetCookie("token", tokenString, 3600*24, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
	return
}

// 创建文章
func CreateArticle(context *gin.Context) {
	var article model.Post
	if err := context.ShouldBindJSON(&article); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	if err := config.DB.Create(&article).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create article"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Article created successfully"})
}

// 通过user id获取 获取文章列表
func GetArticles(context *gin.Context) {
	var articles []model.Post
	it, err := strconv.ParseInt(context.Query("userID"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	config.DB.Where("user_id = ?", it).Find(&articles)
	context.JSON(http.StatusOK, gin.H{"data": articles})
}

// 获取单个文章
func GetArticle(context *gin.Context) {
	var article model.Post
	id := context.Query("id")
	if id == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid article ID"})
		return
	}
	num, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "parse data error"})
		return
	}
	config.DB.Where("id = ?", num).First(&article)
	context.JSON(http.StatusOK, gin.H{"data": article})
	return

}

// 更新文章，只有文章的作者才能更新
func UpdateArticle(c *gin.Context) {
	var article model.Post
	c.ShouldBindJSON(&article)
	tx := config.DB.Model(&article).Where("user_id = ?", article.UserID).Updates(&article)
	if tx.Error != nil {
		c.JSON(http.StatusOK, gin.H{"message": "更新失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "更新成功"})
	return
}

// 删除文章，只有文章的作者才能删除
func DeleteArticle(c *gin.Context) {
	var article model.Post
	c.ShouldBindJSON(&article)
	tx := config.DB.Where("user_id = ?", article.UserID).Delete(&article)
	if tx.Error != nil {
		c.JSON(http.StatusOK, gin.H{"message": "删除失败"})
	}
	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}

// 评论功能的创建
func CreateComments(c *gin.Context) {
	var comment model.Comment
	c.ShouldBindJSON(&comment)
	tx := config.DB.Create(&comment)
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "创建失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "创建成功"})
}

func ReadComment(c *gin.Context) {
	var comment model.Comment
	c.ShouldBindJSON(&comment)
	tx := config.DB.Where("post_id = ?", comment.PostID).First(&comment)
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "查询失败"})
	}
	c.JSON(http.StatusOK, gin.H{"message": "查询成功", "data": comment})
}
