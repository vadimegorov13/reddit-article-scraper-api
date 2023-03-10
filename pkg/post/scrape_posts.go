package post

import (
	"github.com/gofiber/fiber/v2"
)

func (h handler) ScrapePosts(c *fiber.Ctx) error {
	return c.JSON("")

}
