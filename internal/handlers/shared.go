package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nt2311-vn/go-fiber/pkg/validation"
)

func ValidateFields(c *fiber.Ctx) error {
	field := c.Query("field")
	value := c.Query(field)

	var err error

	switch field {
	case "email":
		err = validation.ValidateEmail(value)
	case "password":
		err = validation.ValidatePassword(value)
	case "confirm-password":
		password := c.Query("password")
		err = validation.ValidatePasswordConfirm(password, value)
	default:
		err = nil
	}

	if err != nil {
		return c.SendString(err.Error())
	}

	return c.SendString("")
}
