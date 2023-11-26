package main

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		res, _ := json.Marshal(Hello{Hello: "world"})
		return c.Send(res)
	})

	app.Listen(":8080")
}

type Hello struct {
	Hello string `json:"hello"`
}
