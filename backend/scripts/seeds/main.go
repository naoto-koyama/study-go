package main

import (
	"log"
	"path/filepath"
	"runtime"
	"study-go/internal/database"
	"study-go/internal/models"
	customfaker "study-go/internal/test/faker"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func init() {
	_, b, _, _ := runtime.Caller(0)
	rootDir := filepath.Join(filepath.Dir(b), "../..")
	if err := godotenv.Load(filepath.Join(rootDir, ".env")); err != nil {
		log.Fatal("Error loading .env file:", err)
	}
}

func cleanupDatabase(db *gorm.DB) error {
	if err := db.Migrator().DropTable(&models.Bid{}, &models.Item{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&models.Bid{}, &models.Item{}); err != nil {
		return err
	}

	log.Println("Cleaned up existing data")
	return nil
}

func createItems(db *gorm.DB, count int) error {
	for i := 0; i < count; i++ {
		startPrice := uint(customfaker.RandomInt(1000, 100000))
		startAt := time.Now().Add(time.Duration(customfaker.RandomInt(1, 48)) * time.Hour)

		item := models.Item{
			Title:        customfaker.GenerateTitle(),
			Description:  customfaker.GenerateDescription(),
			StartPrice:   int(startPrice),
			CurrentPrice: int(startPrice),
			StartAt:      startAt,
			EndAt:        startAt.Add(time.Duration(customfaker.RandomInt(24, 168)) * time.Hour),
			Status:       "active",
			CreatedAt:    time.Now(),
		}

		if err := db.Create(&item).Error; err != nil {
			return err
		}
	}

	log.Printf("Created %d items\n", count)
	return nil
}

func createUsers(db *gorm.DB, count int) error {
	for i := 0; i < count; i++ {
		user := models.User{
			Username:  faker.Username(),
			Email:     faker.Email(),
			Password:  faker.Password(),
			CreatedAt: time.Now(),
		}

		if err := db.Create(&user).Error; err != nil {
			return err
		}
	}

	log.Printf("Created %d users\n", count)
	return nil
}

func createBids(db *gorm.DB) error {
	// Itemに対して0-2件の入札があるようにする
	items := []models.Item{}
	if err := db.Find(&items).Error; err != nil {
		return err
	}

	for _, item := range items {
		for i := 0; i < customfaker.RandomInt(0, 2); i++ {
			user := models.User{}
			if err := db.Order("RANDOM()").First(&user).Error; err != nil {
				return err
			}

			bid := models.Bid{
				ItemID:    item.ID,
				UserID:    user.ID,
				Price:     item.CurrentPrice + customfaker.RandomInt(100, 1000),
				CreatedAt: time.Now(),
			}

			if err := db.Create(&bid).Error; err != nil {
				return err
			}

			item.CurrentPrice = bid.Price
			if err := db.Save(&item).Error; err != nil {
				return err
			}
		}
	}

	log.Println("Created bids")
	return nil
}

func main() {
	db := database.NewDB()

	if err := cleanupDatabase(db); err != nil {
		log.Fatal("Failed to cleanup database:", err)
	}

	log.Println("Starting seed data creation...")

	if err := createItems(db, 50); err != nil {
		log.Fatal("Failed to create items:", err)
	}

	if err := createUsers(db, 100); err != nil {
		log.Fatal("Failed to create users:", err)
	}

	if err := createBids(db); err != nil {
		log.Fatal("Failed to create bids:", err)
	}

	log.Println("Seed data creation completed successfully")
}
