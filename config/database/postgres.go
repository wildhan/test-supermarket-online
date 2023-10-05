package database

import (
	"fmt"
	"time"

	"test-lion-superindo/lib/log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbConnection struct {
	DB *gorm.DB
}

func CreateConnection(connectString string) *DbConnection {
	log.Info("Starting connection ...")
	var db *gorm.DB
	var err error
	tryConnectTimes := 0
	// connectString := os.Getenv("DB_STRING")

	for {
		db, err = gorm.Open(postgres.Open(connectString), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		tryConnectTimes++
		if err != nil {
			log.Error(fmt.Sprintf("Try to create DB Connection %d: %v \n", tryConnectTimes, err.Error()))
			time.Sleep(3 * time.Second)
			continue
		}
		break
	}

	log.Info("DB Connection Success")
	return &DbConnection{db}
}
