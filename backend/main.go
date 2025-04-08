package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := os.Getenv("CONFIG")
	if dsn == "" {
		log.Fatal("CONFIG environment variable is not set")
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("cannot connect to database: %v", err)
	}

	if err := db.AutoMigrate(&User{}); err != nil {
		log.Fatalf("automigrate failed: %v", err)
	}

	r := gin.Default()

	// 	app.Get("/floors/:id<int>/special", GetSpecialFloor)
	r.Get("")

	r.Run(":8080")
}
