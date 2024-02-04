package handler

import (
	"log"
	"projec1/database"
	"projec1/model"

	"github.com/gofiber/fiber/v2"
)

func SingUp(c *fiber.Ctx) error {

	db := database.DB.Db
	user := new(model.User)
	err := c.BodyParser(user)
	if err != nil {
		log.Fatal(err)
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "server error was able to body parse", "data": err})
	}

	//username control
	var existingUser model.User
	result := db.Where("username = ?", user.Username).First(&existingUser)

	if result.Error == nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Username is used by someone else", "data": err})
	}

	//email control
	var existingEmail model.User
	result = db.Where("email = ?", user.Email).First(&existingEmail)

	if result.Error == nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "email is used by someone else", "data": err})
	}

	err = db.Create(&user).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "server error could not create new user", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "user created", "data": user})
}

func SingIn(c *fiber.Ctx) error {

	db := database.DB.Db
	var loginData model.LogIn
	c.BodyParser(&loginData)
	username := loginData.Username
	password := loginData.Password
	var user model.User
	result := db.Where("username = ? and password =?", username, password).Find(&user)
	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"status": "fail", "message": "username or password is fail", "data": result.Error})
	}

	var activeTrue model.IsActive
	activeTrue.UserID = user.ID
	activeTrue.IsActive = true
	db.Save(&activeTrue)

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "login succesed", "data": user})
}

func LogOut(c *fiber.Ctx) error {
	db := database.DB.Db
	var logOut model.LogOut
	c.BodyParser(&logOut)
	username := logOut.Username
	if len(username) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "fail", "message": "user not found", "data": nil})
	}
	var activefalse model.IsActive
	var user model.User
	db.Where("username = ?", username).First(&user)
	id := user.ID
	db.Where("user_id = ?", id).First(&activefalse)
	db.Delete(&activefalse)

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "logout succesed"})
}
