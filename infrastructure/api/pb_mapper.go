package api

import (
	"github.com/XWS-BSEP-Tim-13/Dislinkt_PostService/domain"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_PostService/domain/enum"
	pb "github.com/XWS-BSEP-Tim-13/Dislinkt_PostService/infrastructure/grpc/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func mapPost(post *domain.Post) *pb.Post {
	postPb := &pb.Post{
		Id:       post.Id.Hex(),
		Username: post.Username,
		Content:  post.Content,
		Image:    post.Image,
		Likes:    post.Likes,
		Dislikes: post.Dislikes,
	}
	for _, comment := range post.Comments {
		postPb.Comments = append(postPb.Comments, &pb.Comment{
			Id:       comment.Id.Hex(),
			Content:  comment.Content,
			Date:     comment.Date,
			Username: comment.Username,
		})
	}
	return postPb
}

func mapReactionToDomain(reactionPb *pb.Reaction) *domain.Reaction {
	postId, err := primitive.ObjectIDFromHex(reactionPb.PostId)
	if err != nil {
		return &domain.Reaction{}
	}

	reaction := &domain.Reaction{
		Username:     reactionPb.Username,
		PostId:       postId,
		ReactionType: enum.ReactionType(reactionPb.ReactionType),
	}

	return reaction
}
