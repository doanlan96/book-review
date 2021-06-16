package controller

import (
	"fmt"

	"github.com/TechMaster/golang/08Fiber/Repository/model"
	repo "github.com/TechMaster/golang/08Fiber/Repository/repository"
	"github.com/gofiber/fiber/v2"
)

var books []model.Book

func GetAllReview(c *fiber.Ctx) error {
	return c.JSON(repo.Reviews.GetAllReviews())
}

func DeleteReviewById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	err = repo.Reviews.DeleteReviewById(int64(id))
	if err != nil {
		return c.Status(404).SendString(err.Error())
	} else {
		return c.SendString("delete successfully")
	}
}

func CreateReview(c *fiber.Ctx) error {
	review := new(model.Review)

	err := c.BodyParser(review)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	reviewId := repo.Reviews.CreateNewReview(review)
	return c.SendString(fmt.Sprintf("New review is created successfully with id = %d", reviewId))

}

func AverageRating(c *fiber.Ctx) error {
	listReview := repo.Reviews.GetAllReviews()
	for _, review := range listReview {
		book, err := repo.Books.FindBookById(review.BookId)
		if err != nil {
			return c.JSON(fiber.Map{
				"message": "Not found this book id",
			})
		}
		result := repo.Reviews.AverageRating()
		book.Rating = result[int64(review.BookId)]
		repo.Books.UpsertBook(book)
	}
	return c.SendString(fmt.Sprintf("successfully"))
}
