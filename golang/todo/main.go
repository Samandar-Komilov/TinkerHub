package main

import (
	"context"
	"database/sql"
	"log"
	"strconv"

	"todo/gen"

	"github.com/gofiber/fiber/v2"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	db, err := sql.Open("pgx", "postgres://postgres:password@localhost:5432/todos?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	q := gen.New(db)
	app := fiber.New()

	app.Post("/todos", func(c *fiber.Ctx) error {
		var body struct {
			Title string `json:"title"`
		}
		if err := c.BodyParser(&body); err != nil {
			return c.Status(400).SendString("Invalid body")
		}
		todo, err := q.CreateTodo(context.Background(), body.Title)
		if err != nil {
			return c.Status(500).SendString("DB error")
		}
		return c.JSON(todo)
	})

	app.Get("/todos", func(c *fiber.Ctx) error {
		todos, err := q.ListTodos(context.Background())
		if err != nil {
			return c.Status(500).SendString("DB error")
		}
		return c.JSON(todos)
	})

	app.Put("/todos/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(400).SendString("Invalid ID")
		}
		var body struct {
			Title     string `json:"title"`
			Completed bool   `json:"completed"`
		}
		if err := c.BodyParser(&body); err != nil {
			return c.Status(400).SendString("Invalid body")
		}
		todo, err := q.UpdateTodo(context.Background(), gen.UpdateTodoParams{
			ID:        int32(id),
			Title:     body.Title,
			Completed: body.Completed,
		})
		if err != nil {
			return c.Status(500).SendString("DB error")
		}
		return c.JSON(todo)
	})

	app.Delete("/todos/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(400).SendString("Invalid ID")
		}
		err = q.DeleteTodo(context.Background(), int32(id))
		if err != nil {
			return c.Status(500).SendString("DB error")
		}
		return c.SendStatus(204)
	})

	log.Fatal(app.Listen(":3000"))
}
