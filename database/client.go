package database

import (
	"JobScheduler/Server/model"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

type Config struct {
	Host     string
	User     string
	Password string
	DB       string
}

var Connector *gorm.DB

func Connect(config Config) error {
	var err error

	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&multiStatements=true", config.User, config.Password, config.Host, config.DB)

	Connector, err = gorm.Open("mysql", connectionString)

	return err

}

func Migrate(table *model.Job) {
	Connector.AutoMigrate(&table)
	log.Println("Table migrated")
}
