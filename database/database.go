package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/riquefsouza/hefesto-fiber-gorm/models"
)

var (
	DBConn *gorm.DB
)

func InitDatabase() {
	var err error
	DBConn, err = gorm.Open("sqlite3", "dbhefesto.db")
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Database connection successfully opened")

	DBConn.AutoMigrate(&models.AdmParameterCategory{})
	fmt.Println("Database Migrated")
}
