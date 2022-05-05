package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Comment struct {
	Id       primitive.ObjectID `bson:"_id"`
	Content  string             `bson:"content"`
	Date     string             `bson:"date"`
	Username string             `bson:"username"`
}

type Post struct {
	Id       primitive.ObjectID `bson:"_id"`
	Username string             `bson:"username"`
	Content  string             `bson:"content"`
	Image    string             `bson:"image"`
	Likes    []string           `bson:"likes"`
	Dislikes []string           `bson:"dislikes"`
	Comments []Comment          `bson:"comments"`
}
