package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	myname := "John Doe"
	app := fiber.New()
	// This is a simple Go program that prints "Hello, World!" to the console.
	println(myname)
	log.Fatal(app.Listen(":3000"))
	// The program uses the built-in println function to output the string.
}
