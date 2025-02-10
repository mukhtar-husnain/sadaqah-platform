package controllers

import (
	"github.com/mukhtar-husnain/sadaqah-platform/db"
	"github.com/mukhtar-husnain/sadaqah-platform/models"

	"github.com/gofiber/fiber/v2"
)

func CreateVolunteerRequest(c *fiber.Ctx) error {
	var data struct {
		UserID             uint   `json:"user_id"`
		Title              string `json:"title"`
		Description        string `json:"description"`
		Location           string `json:"location"`
		EventDate          string `json:"event_date"`
		RequiredVolunteers int    `json:"required_volunteers"`
		SkillsRequired     string `json:"skills_required"`
	}

	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	request := models.VolunteerRequest{
		UserID:             data.UserID,
		Title:              data.Title,
		Description:        data.Description,
		Location:           data.Location,
		EventDate:          data.EventDate,
		RequiredVolunteers: data.RequiredVolunteers,
		SkillsRequired:     data.SkillsRequired,
	}

	db.DB.Create(&request)

	return c.JSON(fiber.Map{"message": "Volunteer request created successfully!"})
}

func GetAllVolunteerRequests(c *fiber.Ctx) error {
	var requests []models.VolunteerRequest

	result := db.DB.Unscoped().Find(&requests)

	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
	}

	if len(requests) == 0 {
		return c.JSON(fiber.Map{"message": "No volunteer requests found."})
	}

	return c.JSON(requests)
}

