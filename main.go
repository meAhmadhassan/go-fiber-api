package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"

)

type Book struct {
	Author 		string		`json:"author"`
	Title 		string		`json:"title"`
	Publisher	string		`json:"publisher"`
}


type Repository struct {
	DB *gorm.DB 
}

func(r *Repository) SetupRoutes(app *fiber.App){
	api := app.Group("/api")
	api.Post("/create_books", r.CreateBook)
	api.Delete("/delete_book/:id", r.DeleteBook)
	api.Get("/get_book/:id", r.GetBookByID)
	api.Get("/books", r.GetBooks)

}


func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	db, err := storage.NewConnection(config)
	if err != nil {
		log.Fatal("could not load the database", err)
	}

	app := fiber.New()

	r := Repository(
		DB: db, 
	)

	r.SetupRoutes(app)

	app.Listen(":8000")
    // log.Fatal(app.Listen(":3000"))

}