package routes

import (
	"github.com/TechMaster/golang/08Fiber/Repository/controller"
	"github.com/gofiber/fiber/v2"
)

func ConfigBookRouter(router *fiber.Router) {
	//Return all books
	(*router).Get("/", controller.GetAllBook)

	(*router).Get("/:id", controller.GetBookById)

	(*router).Delete("/:id", controller.DeleteBookById)

	(*router).Patch("", controller.UpdateBookById)

	(*router).Post("", controller.CreateBook)

	(*router).Put("", controller.UpsertBook)
}

func ConfigReviewRouter(router *fiber.Router) {
	(*router).Get("/", controller.GetAllReview)
	(*router).Post("", controller.CreateReview)
	(*router).Delete("/:id", controller.DeleteReviewById)
	(*router).Get("/average/:id", controller.AverageRating)
}
