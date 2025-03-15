package main

import (
	"log"
	"os"

	"github.com/meahmadhassan/go-fiber-api/api"
	"github.com/meahmadhassan/go-fiber-api/models"
	"github.com/meahmadhassan/go-fiber-api/storage"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)



func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	db, err := storage.NewConnection(config)

	if err != nil {
		log.Fatal("could not load the database", err)
	}

	err = models.MigrateBooks(db)
	if err != nil {
		log.Fatal("could not migrate db")
	}

	
	r := &api.Repository{DB: db}

	app := fiber.New()
	r.SetupRoutes(app)

	app.Listen(":8080")
	// log.Fatal(app.Listen(":3000"))

}
