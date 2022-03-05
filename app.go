package main

import (
    "fmt"
    // "log"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/template/html"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
    // Initialize standard Go html template engine
    engine := html.New("./views", ".html")
    app := fiber.New(fiber.Config{
        Views: engine,
    })
    app.Use(logger.New())
	
     

    app.Get("/", func(c *fiber.Ctx) error {
        // Render index template
        return c.Render("index", fiber.Map{
            "INTERPRETAR": "",
        })
    })

	app.Post("/interpretar", execute)

	/*
		Se ejecuta el servidor y en caso de fallar, muetra log.Fatal con el error.
	*/
    _ = app.Listen(":3000")
}


type getInput struct {
    Input string `form:"editor"`
}   

func execute(c *fiber.Ctx) error {
    data := new(getInput)
    if err := c.BodyParser(data); err != nil {
        return err
    }
    fmt.Println("sii")
    fmt.Println(data)
    return c.Render("index", fiber.Map{
        "INTERPRETAR": data.Input,
    })
}