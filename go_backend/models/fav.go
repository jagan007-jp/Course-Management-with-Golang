package models

type Favourites struct {
	ID          string `json:"id" bson:"id"`
	Image       string `json:"image" bson:"image"`
	Title       string `json:"title" bson:"title"`
	Link        string `json:"link" bson:"link"`
	IsFavourite bool   `json:"isFavourite" bson:"isFavourite"`
	Username    string `json:"username" bson:"username"`
}
