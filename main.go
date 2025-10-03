package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books = []Book{
	{ID: 1, Title: "The Great Gatsby", Author: "F. Scott Fitzgerald"},
	{ID: 2, Title: "To Kill a Mockingbird", Author: "F. Scott Fitzgerald"}}

func checkMiddleware(c *fiber.Ctx) error {
	start := time.Now()
	fmt.Printf(
		"URL = %s, Method = %s, Time = %s\n",
		c.OriginalURL(),
		c.Method(),
		start,
	)
	return c.Next()
}

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(checkMiddleware)

	app.Get("/books", getBooks)
	app.Get("/books/:id", getBook)
	app.Post("/books", createBook)
	app.Put("/books/:id", updateBook)
	app.Delete("/books/:id", deleteBook)

	app.Post("/upload", uploadFile)
	app.Get("/test-html", testHTML)

	app.Get("/config_path", getENV)

	app.Get("/config_file", getENVFile)

	app.Listen(":8080")

}

func uploadFile(c *fiber.Ctx) error {
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	err = c.SaveFile(file, "./uploads/"+file.Filename)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.SendString("File uploaded successfully")

}

func testHTML(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": "Hello, World!",
	})
}

func getENV(c *fiber.Ctx) error {
	if value, exists := os.LookupEnv("Path"); exists {
		return c.JSON(fiber.Map{"Path": value})
	}
	return c.JSON(fiber.Map{"Path": "Not Set"})
}

func getENVFile(c *fiber.Ctx) error {

	return c.JSON(fiber.Map{"SECRET": os.Getenv("SECRET")})
}
