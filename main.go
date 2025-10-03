package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books = []Book{
	{ID: 1, Title: "The Great Gatsby", Author: "F. Scott Fitzgerald"},
	{ID: 2, Title: "To Kill a Mockingbird", Author: "F. Scott Fitzgerald"}}

func main() {

	app := fiber.New()

	app.Get("/books", getBooks)
	app.Get("/books/:id", getBook)
	app.Post("/books", createBook)
	app.Put("/books/:id", updateBook)
	app.Delete("/books/:id", deleteBook)
	app.Listen(":8080")

}

func getBooks(c *fiber.Ctx) error {
	return c.JSON(books)
}

func getBook(c *fiber.Ctx) error {
	bookID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	for _, book := range books {
		if book.ID == bookID {
			return c.JSON(book)
		}
	}

	return c.Status(404).SendString("Book not found")
}

func createBook(c *fiber.Ctx) error {
	book := new(Book)

	if err := c.BodyParser(book); err != nil {
		return err
	}

	books = append(books, *book)

	return c.JSON(books)
}

func updateBook(c *fiber.Ctx) error {
	bookID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	bookUpdate := new(Book)

	if err := c.BodyParser(bookUpdate); err != nil {
		return err
	}

	for i, book := range books {
		if book.ID == bookID {
			books[i].Title = bookUpdate.Title
			books[i].Author = bookUpdate.Author
			return c.JSON(books[i])
		}
	}
	return c.Status(404).SendString("Book not found")
}

func deleteBook(c *fiber.Ctx) error {
	bookID, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	for _, book := range books {
		if book.ID == bookID {
			books = append(books[:bookID-1], books[bookID:]...)
			return c.SendStatus(fiber.StatusNoContent)
		}
	}
	return c.Status(404).SendString("Book not found")
}
