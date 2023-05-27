package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kpv22/GO-React-Crud/db"
	"github.com/kpv22/GO-React-Crud/models"
)

func GetUsersHandler(c *fiber.Ctx) error {
	var users []models.User
	db.DB.Find(&users)
	return c.JSON(&users)
}

func GetUserHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	db.DB.First(&user, id)

	if user.ID == 0 {
		return c.Status(fiber.StatusNotFound).SendString("User not found!")
	}

	return c.JSON(&user)
}

func PostUserHandler(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	// Consultar el último ID de usuario en la base de datos
	var lastUserID uint
	db.DB.Table("users").Select("id").Order("id DESC").Limit(1).Scan(&lastUserID)

	// Asignar un nuevo ID al usuario en caso de que no exista ningún usuario con ID siguiente
	if lastUserID > 0 {
		user.ID = lastUserID + 1
	} else {
		user.ID = 1
	}

	createdUser := db.DB.Create(&user)
	if createdUser.Error != nil {
		return c.Status(fiber.StatusBadRequest).SendString(createdUser.Error.Error())
	}

	return c.JSON(&user)
}

func DeleteUserHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	fmt.Println("User ID:", id)

	var user models.User
	db.DB.First(&user, id)
	if user.ID == 0 {
		return c.Status(fiber.StatusNotFound).SendString("User not found!")
	}
	if err := db.DB.Unscoped().Delete(&user).Error; err != nil {
		fmt.Println("Error:", err.Error())
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to delete user")
	}

	return c.SendStatus(fiber.StatusOK)
}
