package controllers

import (
	"ambassador/src/database"
	"ambassador/src/models"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	FirstName       string `json:"first_name" form:"first_name"`
	LastName        string `json:"last_name" form:"last_name"`
	Email           string `json:"email" form:"email"`
	Password        string `json:"password" form:"password"`
	PasswordConfirm string `json:"password_confirm" form:"password_confirm"`
}

func Register(c *fiber.Ctx) error {
	data := new(User)

	if err := c.BodyParser(data); err != nil {
		return err
	}

	fmt.Println("data:", data)
	fmt.Println("password:", data.Password)
	fmt.Println("password confirm:", data.PasswordConfirm)

	if data.Password != data.PasswordConfirm {
		c.Status(400)
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "password do not match.",
		})
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data.Password), 12)

	user := models.User{
		FirstName:    data.FirstName,
		LastName:     data.LastName,
		Email:        data.Email,
		Password:     password,
		IsAmbassador: false,
	}

	database.DB.Create(&user)

	return c.JSON(fiber.Map{
		"code":    http.StatusOK,
		"message": "register success",
		"data":    &user,
	})
}
