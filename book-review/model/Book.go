package model

type Book struct {
	Id      int64    `json:"id"`
	Title   string   `json:"title"`
	Authors []Author `json:"authors"`
	Rating  float32  `json:"rating"`
	// Rating := controller.Rating() `json:"rating"`
}
