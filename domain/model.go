package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Comment struct {
	Id       primitive.ObjectID `bson:"_id"`
	Content  string             `bson:"content"`
	Date     time.Time          `bson:"date"`
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
	Date     time.Time          `bson:"date"`
}

type Message struct {
	MessageFrom string    `bson:"message_from"`
	MessageTo   string    `bson:"message_to"`
	Content     string    `bson:"content"`
	Date        time.Time `bson:"date"`
}

type MessageUsers struct {
	Id         primitive.ObjectID `bson:"_id"`
	FirstUser  string             `bson:"first_user"`
	SecondUser string             `bson:"second_user"`
	Messages   []Message          `bson:"messages"`
}
