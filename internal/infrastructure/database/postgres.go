package database

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/Manizmn84/hasin_interview/bootstrap"
	"github.com/Manizmn84/hasin_interview/internal/domain/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dbInstance *gorm.DB
	dbOnce     sync.Once
)

const (
	maxRetries = 10
	retryDelay = 3 * time.Second
)

func NewPostgresDataBase(dbConfig *bootstrap.Database) *gorm.DB {

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tehran",
		dbConfig.Host,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Name,
		dbConfig.Port,
	)

	dbOnce.Do(func() {
		var db *gorm.DB
		var err error

		for i := 1; i <= maxRetries; i++ {
			db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
			if err == nil {
				break
			}
			log.Printf("database connection attempt %d/%d failed: %v", i, maxRetries, err)
			if i < maxRetries {
				time.Sleep(retryDelay)
			}
		}

		if err != nil {
			panic(fmt.Errorf("failed to connect to database after %d attempts: %w", maxRetries, err))
		}

		dbInstance = db

		if err = dbInstance.AutoMigrate(&entities.Todo{}); err != nil {
			panic(err)
		}
	})

	return dbInstance
}
