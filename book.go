package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

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
