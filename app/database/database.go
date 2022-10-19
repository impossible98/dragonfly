package database

import (
	// import built-in packages
	"fmt"
	"strconv"
	// import third-party packages
	"github.com/glebarez/sqlite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	// import local packages
	"dragonfly/app/config"
	"dragonfly/app/model"
	"dragonfly/app/utils/path"
)

func SetupDatabase() *gorm.DB {
	// control flow
	if config.DatabaseType == "postgres" {
		dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=disable", config.DatabaseHost, config.DatabaseUser, config.DatabasePassword, strconv.Itoa(config.DatabasePort), config.DatabaseDatabase)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		// control flow
		if err != nil {
			panic("Failed to create a connection to database")
		}

		db.AutoMigrate(&model.Favourite{})
		fmt.Println("Database connected!")
		// return
		return db
	} else {
		path.CreateDir("./data")
		db, err := gorm.Open(sqlite.Open("./data/test.db"), &gorm.Config{})
		// control flow
		if err != nil {
			panic("failed to connect database")
		}
		db.AutoMigrate(&model.Favourite{})
		fmt.Println("Database connected!")
		// return
		return db
	}

}
