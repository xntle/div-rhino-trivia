package databse

import (
	"fmt"
	"log"
	"os"

	"github.com/xntle/div-rhino-trivia/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func ConnectDb() {

	// %s is environment variable
	dsn := fmt.Sprintf("host=db user=%s password=%s dbname=%s port = 5432 sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("failed to connect to database", err)
		os.Exit(2)
	}

	log.Println("Connected to database")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Running Migrations....")
	db.AutoMigrate(&models.Fact{})


	// set our global bd value to the database we setted up
	DB = Dbinstance{
		Db: db,
	}
}