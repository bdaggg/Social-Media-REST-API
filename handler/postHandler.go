package handler

import (
	"projec1/database"
	"projec1/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreatePost(c *fiber.Ctx) error {

	db := database.DB.Db
	req := new(model.PostRequest)

	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "server error was able to body parse", "data": err})
	}
	var user model.User

	if err := db.Where("username = ?", req.Username).First(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "user not found", "data": err})
	}

	post := model.Post{
		Description: req.Description,
		UserID:      user.ID,
		Username:    req.Username,
	}

	err = db.Create(&post).Error

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "post created", "data": nil})
}

func GetAllPost(c *fiber.Ctx) error {
	db := database.DB.Db
	var posts []model.Post
	db.Find(&posts)

	if len(posts) == 0 {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "posts are not there or system error", "data": posts})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "posts get from database", "data": posts})
}

func GetPosts(c *fiber.Ctx) error {
	db := database.DB.Db
	username := c.Params("username")
	var posts []model.Post
	db.Find(&posts, "username = ?", username)
	if len(posts) == 0 {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "user don't have anypost or system error", "data": posts})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "posts get from database", "data": posts})

}

func UpdatePost(c *fiber.Ctx) error {
	type UPost struct {
		NewDescription string `json:"description"`
		ID             string `json:"id"`
	}
	db := database.DB.Db
	var req UPost
	c.BodyParser(&req)
	PostId := req.ID
	var newPost model.Post
	id, err := strconv.Atoi(PostId)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "success", "message": "postun uuid kismi bos", "data": nil})
	}
	db.Find(&newPost, "ID1 = ?", id)
	if newPost.UserID == uuid.Nil {

		return c.Status(500).JSON(fiber.Map{"status": "success", "message": "postun uuid kismi bos", "data": nil})
	}
	newPost.Description = req.NewDescription
	db.Save(&newPost)
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "posts updated", "data": nil})

}
