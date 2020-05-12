package main

import (
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/gofiber/websocket"
	"log"
)

func main() {
	app := fiber.New()

	app.Use(func(c *fiber.Ctx) {
		c.Locals("Hello", "World")
		c.Next()
	})

	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		fmt.Println(c.Locals("Hello")) // "World"
		// Websocket logic...
		for {
			mt, msg, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				break
			}
			log.Printf("recv: %s", msg)
			err = c.WriteMessage(mt, msg)
			if err != nil {
				log.Println("write:", err)
				break
			}
		}
	}))

	app.Get("/helloWorld", func(ctx *fiber.Ctx) {
		ctx.Send("Hello Paul")
	})

	app.Listen(7888) // ws://localhost:3000/ws
}
