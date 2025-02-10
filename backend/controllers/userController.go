package controllers

import (
	"github.com/mukhtar-husnain/sadaqah-platform/db"
	"github.com/mukhtar-husnain/sadaqah-platform/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c *fiber.Ctx) error {
	var data struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
		UserType string `json:"user_type"`
	}

	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Hash the password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)

	// Create user
	user := models.User{
		Name:     data.Name,
		Email:    data.Email,
		Password: string(hashedPassword),
		UserType: data.UserType,
	}

	db.DB.Create(&user)

	return c.JSON(fiber.Map{"message": "User registered successfully!"})
}
