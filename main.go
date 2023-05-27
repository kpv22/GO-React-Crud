package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/kpv22/GO-React-Crud/db"
	"github.com/kpv22/GO-React-Crud/models"
	"github.com/kpv22/GO-React-Crud/routes"
)

func main() {
	db.DBConnection()
	db.DB.AutoMigrate(&models.User{})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	app := fiber.New()

	app.Use(cors.New())
	// Agregar el manejador OPTIONS
	app.Options("*", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNoContent)
	})

	app.Get("/users", routes.GetUsersHandler)
	app.Get("/users/:id", routes.GetUserHandler)
	app.Post("/users", routes.PostUserHandler)
	app.Delete("/users/:id", routes.DeleteUserHandler)

	app.Static("/", "./client/dist")

	fmt.Println("Server on port " + port)
	app.Listen(":" + port)
}
