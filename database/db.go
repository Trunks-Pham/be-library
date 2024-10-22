package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	//dsn := "host=localhost user=postgres password=yourpassword dbname=library port=5432 sslmode=disable"
	dsn := "postgresql://libraray_server_user:Txy8XXVjBSfoGEYJEUxDGFkMxKrbE3Y1@dpg-cs9luk08fa8c73cceucg-a.singapore-postgres.render.com/libraray_server"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	log.Println("Database connection established.")
}
