package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type TodoItem struct {
	Id          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}
type ToDoItemCreateDto struct {
	Id          int    `json:"-" gorm:"column:id;"`
	Title       string `json:"title" gorm:"column:title;"`
	Description string `json:"description" gorm:"column:description;"`
	Status      string `json:"status" gorm:"column:status;"`
}
type ToDoItemUpdate struct {
	Title       string `json:"title" gorm:"column:title;"`
	Description string `json:"description" gorm:"column:description;"`
	Status      string `json:"status" gorm:"column:status;"`
}

// @title 	Tag Service API
// @version	1.0
// @description A Tag service API in Go using Gin framework

// @host 	localhost:8080
// @BasePath /swagger
func (TodoItem) TableName() string {
	return "todo_items"
}
func (ToDoItemCreateDto) TableName() string {
	return TodoItem{}.TableName()
}
func (ToDoItemUpdate) TableName() string {
	return TodoItem{}.TableName()
}
func CreateItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data ToDoItemCreateDto
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if err := db.Create(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": data.Id,
		})

	}
}
func GetItemById(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data TodoItem
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Bad Request",
			})
		}
		if err := db.Where("id = ?", id).First(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}
}
func main() {
	godotenv.Load()
	r := gin.Default()

	dsn := os.Getenv("DB_CONN_STR")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Connect database is success", db)
	}
	v1 := r.Group("/v1")
	{
		items := v1.Group("/items")
		{
			items.POST("/create-items", CreateItem(db))
			items.GET("/:id", GetItemById(db))
		}
	}
	r.Run(":8080")
}
