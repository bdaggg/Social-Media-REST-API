package router

import (
	"projec1/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRouters(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")
	user := v1.Group("/user")
	post := v1.Group("/post")

	//users
	user.Get("/alluser", handler.GetAllUser)
	user.Get("/:username", handler.GetUser)
	user.Post("/singup", handler.SingUp)
	user.Post("/singin", handler.SingIn)
	user.Post("/logout", handler.LogOut)
	user.Patch("/update/:username", handler.Update)
	user.Delete("/delete/:username", handler.Delete)
	user.Delete("/deleteall/", handler.DeleteAllUser)

	//posts
	post.Post("/newpost", handler.CreatePost)
	post.Get("/allpost", handler.GetAllPost)
	post.Get("/:username", handler.GetPosts)
	post.Patch("/update/:username", handler.UpdatePost)

}
