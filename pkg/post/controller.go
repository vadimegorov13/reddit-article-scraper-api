package post

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

/*
The handler type contains the database connection.
*/
type handler struct {
	DB *gorm.DB
}

/*
RegisterRoutes creates a new handler and registers the API routes with the given Fiber app.
@param {*fiber.App} app - The Fiber app instance.
@param {*gorm.DB} db - The GORM database connection.
*/
func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	// Define the API routes
	api := app.Group("/api")
	v1 := api.Group("/v1")
	posts := v1.Group("/posts")

	// Register the route handlers
	posts.Get("/", h.GetPosts)
	posts.Get("/:id", h.GetPost)
	posts.Post("/:subreddit", h.ScrapePosts)
	posts.Delete("/:id", h.DeletePost)
}
