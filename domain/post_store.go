package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostStore interface {
	Get(ctx context.Context, id primitive.ObjectID) (*Post, error)
	GetAll(ctx context.Context) ([]*Post, error)
	Insert(ctx context.Context, post *Post) error
	DeleteAll(ctx context.Context)
	GetByUser(ctx context.Context, username string) ([]*Post, error)
	UpdateReactions(ctx context.Context, post *Post) (string, error)
	GetFeed(ctx context.Context, page int64, usernames []string) (*FeedDto, error)
	GetFeedAnonymous(ctx context.Context, page int64) (*FeedDto, error)
	Delete(ctx context.Context, id primitive.ObjectID) error
}
