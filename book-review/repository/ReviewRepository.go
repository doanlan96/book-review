package repository

import (
	"errors"
	"fmt"

	"github.com/TechMaster/golang/08Fiber/Repository/model"
)

type ReviewRepo struct {
	reviews map[int64]*model.Review
	autoID  int64
}

var Reviews ReviewRepo

func init() {
	Reviews = ReviewRepo{autoID: 0}
	Reviews.reviews = make(map[int64]*model.Review)
	// Reviews.InitData("sql:45312")
}

func (r *ReviewRepo) getAutoID() int64 {
	r.autoID += 1
	return r.autoID
}
func (r *ReviewRepo) CreateNewReview(review *model.Review) int64 {
	nextID := r.getAutoID()
	review.Id = nextID
	r.reviews[nextID] = review
	return nextID
}

// func (r *ReviewRepo) InitData(connection string) {
// 	fmt.Println("Connect to ", connection)

// 	r.CreateNewReview(&model.Review{
// 		BookId:  1,
// 		Comment: "Hay",
// 		Rating:  4})

// 	r.CreateNewReview(&model.Review{
// 		BookId:  2,
// 		Comment: "Hay",
// 		Rating:  4})

// 	r.CreateNewReview(&model.Review{
// 		BookId:  1,
// 		Comment: "Rất Hay",
// 		Rating:  5})

// 	r.CreateNewReview(&model.Review{
// 		BookId:  2,
// 		Comment: "Dở",
// 		Rating:  2})
// }

func (r *ReviewRepo) GetAllReviews() map[int64]*model.Review {
	return r.reviews
}

func (r *ReviewRepo) DeleteReviewById(Id int64) error {
	if _, ok := r.reviews[Id]; ok {
		delete(r.reviews, Id)
		return nil
	} else {
		return errors.New("review not found")
	}
}

func (r *ReviewRepo) UpdateReviewById(newReview *model.Review) error {
	if _, ok := r.reviews[newReview.Id]; ok {
		r.reviews[newReview.Id] = newReview
		return nil //tìm được
	} else {
		return errors.New("review not found")
	}
}

func (r *ReviewRepo) UpsertReview(review *model.Review) int64 {
	if _, ok := r.reviews[review.Id]; ok {
		r.reviews[review.Id] = review
		return review.Id
	} else {
		return r.CreateNewReview(review)
	}
}

func (r *ReviewRepo) AverageRating() (result map[int64]float32) {
	sumRating := make(map[int64]int)
	numberRating := make(map[int64]int)
	result = make(map[int64]float32)

	for _, value := range r.reviews {
		numberRating[value.BookId]++
		sumRating[value.BookId] += value.Rating
	}
	fmt.Println(sumRating)
	fmt.Println(numberRating)

	for key := range numberRating {
		result[key] = float32(sumRating[key]) / float32(numberRating[key])
	}
	return result
}
