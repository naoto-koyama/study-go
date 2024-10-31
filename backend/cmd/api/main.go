package main

import (
	"log"
	"path/filepath"
	"runtime"
	"study-go/internal/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func init() {
	// 環境変数の読み込み
	_, b, _, _ := runtime.Caller(0)
	rootDir := filepath.Join(filepath.Dir(b), "../..")
	if err := godotenv.Load(filepath.Join(rootDir, ".env")); err != nil {
		log.Fatal("Error loading .env file:", err)
	}
}

func setupDB() *gorm.DB {
	db := database.NewDB()
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get database instance:", err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}
	log.Println("Successfully connected to database")
	return db
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return r
}

func main() {
	setupDB()
	r := setupRouter()

	log.Println("Server starting on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
