package initializer

import (
	"fmt"
	"log"
	"os"

	todo "github.com/eskylake/go-todo/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB(config *Config) {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", config.DBHost, config.DBUser, config.DBPassword, config.DBName, config.DBPort)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database! \n", err.Error())
		os.Exit(1)
	}

	DB.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Running Migrations")
	DB.AutoMigrate(&todo.Todo{})

	log.Println("🚀 Connected Successfully to the Database")
}

func Ping() error {
	d, err := DB.DB()
	if err != nil {
		log.Fatal("Failed to connect to the Database! \n", err.Error())
	}

	return d.Ping()
}
