package models

type Course struct {
	ID          string `json:"id" bson:"_id,omitempty"`
	Image       string `json:"image" bson:"image"`
	Title       string `json:"title" bson:"title"`
	Link        string `json:"link" bson:"link"`
	IsFavourite bool   `json:"isFavourite" bson:"isFavourite"`
}
