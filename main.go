package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Task model for the database
type Task struct {
	gorm.Model
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

// Initialize DB connection
func setupDB() *gorm.DB {
	// Update with your database details
	dsn := "host=db user=youruser password=yourpassword dbname=yourdbname port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// Migrate the schema
	db.AutoMigrate(&Task{})
	return db
}

// Create a new task
func createTask(c *gin.Context, db *gorm.DB) {
	var newTask Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Create(&newTask)
	c.JSON(http.StatusCreated, newTask)
}

// Get all tasks
func getTasks(c *gin.Context, db *gorm.DB) {
	var tasks []Task
	db.Find(&tasks)
	c.JSON(http.StatusOK, tasks)
}

func main() {
	r := gin.Default()
	db := setupDB()

	r.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Hello, World! try /tasks."})
	})

	// Routes
	r.POST("/tasks", func(c *gin.Context) {
		createTask(c, db)
	})

	r.GET("/tasks", func(c *gin.Context) {
		getTasks(c, db)
	})

	// Start server
	r.Run() // Listen and serve on 0.0.0.0:8080
}
