package database

import (
	"fmt"
	"sync"

	"github.com/Manizmn84/hasin_interview/bootstrap"
	"github.com/Manizmn84/hasin_interview/internal/domain/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dbInstance *gorm.DB
	dbOnce     sync.Once
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
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(fmt.Errorf("failed to connect database"))
		}

		dbInstance = db

		err = dbInstance.AutoMigrate(
			&entities.Product{},
		)

		if err != nil {
			panic(err)
		}
	})

	return dbInstance
}
