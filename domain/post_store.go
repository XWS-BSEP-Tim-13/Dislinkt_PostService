package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostStore interface {
	Get(id primitive.ObjectID) (*Post, error)
	GetAll() ([]*Post, error)
	Insert(post *Post) error
	DeleteAll()
	GetByUser(username string) ([]*Post, error)
	UpdateReactions(post *Post) (string, error)
	GetFeed(page int64, usernames []string) (*FeedDto, error)
	GetFeedAnonymous(page int64) (*FeedDto, error)
	Delete(id primitive.ObjectID) error
}
