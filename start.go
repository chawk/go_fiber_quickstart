package main

import (
	"encoding/json"

	"github.com/gofiber/fiber"
)

type FruitInterface struct {
	Page   int
	Fruits []string
}

func main() {
	app := fiber.New()

	// Returns plain text.
	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello, Fiber!")
		// navigate to => http://localhost:3000/
	})

	// Load static files like CSS, Images & JavaScript
	app.Static("/public", "./public")

	// Returns a local HTML file.
	app.Get("/hello", func(c *fiber.Ctx) {
		if err := c.SendFile("./templates/hello.html"); err != nil {
			c.Next(err)
		}
		// navigate to => http://localhost:3000/hello
	})

	// Use parameters
	app.Get("/parameter/:value", func(c *fiber.Ctx) {
		c.Send("Get request with value: " + c.Params("value"))
		// navigate to => http://localhost:3000/parameter/this_is_the_parameter
	})

	// Use wildcards to design your API.
	app.Get("/api/*", func(c *fiber.Ctx) {
		c.Send("API path: " + c.Params("*") + " -> do lookups with these values")
		// navigate to => http://localhost:3000/api/user/chris

		// return serialized JSON.
		if c.Params("*") == "fruits" {

			response := FruitInterface{
				Page:   1,
				Fruits: []string{"apple", "peach", "pear"},
			}

			responseJson, _ := json.Marshal(response)

			c.Send(responseJson)

			// navigate to => http://localhost:3000/api/fruits
		}
	})

	app.Listen(3000)
}
