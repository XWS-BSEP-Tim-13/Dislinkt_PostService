package domain

import (
	"github.com/XWS-BSEP-Tim-13/Dislinkt_PostService/domain/enum"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Reaction struct {
	PostId       primitive.ObjectID
	Username     string
	ReactionType enum.ReactionType
}

type FeedDto struct {
	Posts    []*Post
	Page     int64
	LastPage int64
}
