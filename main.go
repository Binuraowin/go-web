package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type Todo struct {
	ID        int    `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

func main() {
	myname := "John Doess"
	app := fiber.New()
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	PORT := os.Getenv("PORT")
	println(myname)

	todos := []Todo{}

	app.Get("/api/todos", func(c *fiber.Ctx) error {

		var x int = 5
		var p *int = &x

		println("Value of x:", x)
		println("Address of x:", &x)
		println("Value of p:", p)
		println("Value pointed to by p:", *p)
		println("Address of p:", &p)
		return c.Status(200).JSON(todos)
	},
	)

	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := &Todo{}

		if err := c.BodyParser(todo); err != nil {
			return err
		}
		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{
				"message": "Todo body is required",
			})
		}
		todo.ID = len(todos) + 1
		todos = append(todos, *todo)
		return c.Status(201).JSON(todo)
	})

	app.Patch("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				todo.Completed = !todo.Completed
				todos[i] = todo
				return c.Status(200).JSON(todo)
			}
		}
		return c.Status(404).JSON(fiber.Map{
			"message": "Todo not found",
		})
	})
	app.Delete("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				todos = append(todos[:i], todos[i+1:]...)
				return c.Status(200).JSON(fiber.Map{
					"message": "Todo deleted",
				})
			}
		}
		return c.Status(404).JSON(fiber.Map{
			"message": "Todo not found",
		})
	})
	log.Fatal(app.Listen(":" + PORT))
}
