package handler

import (
	"projec1/database"
	"projec1/model"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)



func GetAllUser(c *fiber.Ctx) error { 
	db := database.DB.Db
	var users []model.User
	db.Find(&users)
	if len(users) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "users not found", "data": users})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "users found", "data": users})
}

func GetUser(c *fiber.Ctx) error { 
	db := database.DB.Db
	username := c.Params("username")
	var user model.User
	db.Find(&user, "username = ?", username)
	if user.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "user not found", "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "user found", "data": user})

}

func Update(c *fiber.Ctx) error {

	type UpdateUser struct {
		NewUsername string `json:"username"`
		NewEmail    string `json:"email"`
		NewPassword string `json:"password"`
	}

	db := database.DB.Db
	var aktifuser model.User
	username := c.Params("username")
	db.Find(&aktifuser, "username = ?", username)
	if aktifuser.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "user not found", "data": nil})
	}
	var updateUser UpdateUser
	err := c.BodyParser(&updateUser)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "user not found", "data": nil})
	}
	aktifuser.Username = updateUser.NewUsername
	aktifuser.Email = updateUser.NewEmail
	aktifuser.Password = updateUser.NewPassword
	db.Save(&aktifuser)

	//username control
	var existingUser model.User
	result := db.Where("username = ?", aktifuser.Username).First(&existingUser)

	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "User not found", "data": err})
	}

	//email control
	var existingEmail model.User
	result = db.Where("email = ?", aktifuser.Email).First(&existingEmail)

	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "email is used by someone else", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "updated user data", "data": aktifuser})
}

func Delete(c *fiber.Ctx) error {

	db := database.DB.Db
	username := c.Params("username")
	var user model.User
	db.Find(&user, "username = ?", username)
	if user.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "user not found", "data": nil})
	}

	db.Delete(&user)
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "user deleted"})

}

func DeleteAllUser(c *fiber.Ctx) error {

	db := database.DB.Db
	var users []model.User
	db.Find(&users)
	if len(users) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "database is empty", "data": users})
	}
	db.Delete(&users)
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "users deleted"})
}

