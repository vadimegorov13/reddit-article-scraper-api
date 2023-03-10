package post

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	api := app.Group("/api")
	v1 := api.Group("/v1")
	posts := v1.Group("/posts")

	posts.Get("/", h.GetPosts)
	posts.Get("/:id", h.GetPost)
	posts.Post("/:subreddit", h.ScrapePosts)
	posts.Delete("/:id", h.DeletePost)
}
