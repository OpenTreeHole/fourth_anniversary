package main

import (
	"context"
	"github.com/caarlos0/env/v9"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var DB *gorm.DB

var Config struct {
	DbURL          string `env:"DB_URL"`
	SpecialHoleIDs []int  `env:"SPECIAL_HOLE_IDS"`
}

func Init() (app *gin.Engine) {
	var err error
	if err := env.Parse(&Config); err != nil {
		log.Fatal("failed to parse env: %v", err)
	}
	log.Println("Config: ", Config)

	DB, err = gorm.Open(mysql.Open(Config.DbURL), &gorm.Config{})

	if err != nil {
		log.Fatalf("cannot connect to database: %v", err)
		return nil
	}

	app = gin.Default()

	// 	app.Get("/floors/:id<int>/special", GetSpecialFloor)
	// Define a route with an integer ID parameter and an empty path (might be a subgroup)
	app.GET("/floors/:id/special", ListFloorsInASpecialHole)
	return app
}

func main() {
	app := Init()

	// Create a custom HTTP server with Gin
	srv := &http.Server{
		Addr:    ":8080",
		Handler: app,
	}

	// Run server in a goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Wait for interrupt signal (CTRL+C / SIGTERM)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	// Create a context with a 5-second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Gracefully shut down the server
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited gracefully")
}
