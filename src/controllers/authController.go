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

type Input struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *fiber.Ctx) error {
	fmt.Println("login method")
	data := new(Input)

	if err := c.BodyParser(data); err != nil {
		return err
	}

	fmt.Println("password:", []byte(data.Password))
	var user models.User
	database.DB.Where("email = ?", data.Email).First(&user)

	if user.Id == 0 {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "user not found!",
		})
	}
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data.Password)); err != nil {
		c.Status(fiber.StatusBadRequest)
		c.JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "email or password entry error, please try agin.",
		})
	}

	return c.JSON(user)
}
