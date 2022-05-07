package api

import (
	"github.com/XWS-BSEP-Tim-13/Dislinkt_PostService/domain"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_PostService/domain/enum"
	pb "github.com/XWS-BSEP-Tim-13/Dislinkt_PostService/infrastructure/grpc/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func mapPostToPb(post *domain.Post) *pb.Post {
	postPb := &pb.Post{
		Id:       post.Id.Hex(),
		Username: post.Username,
		Content:  post.Content,
		Image:    post.Image,
		Likes:    post.Likes,
		Dislikes: post.Dislikes,
		Date:     timestamppb.New(post.Date),
	}
	for _, comment := range post.Comments {
		postPb.Comments = append(postPb.Comments, &pb.Comment{
			Id:       comment.Id.Hex(),
			Content:  comment.Content,
			Date:     timestamppb.New(comment.Date),
			Username: comment.Username,
		})
	}
	return postPb
}

func mapPostPbToDomain(postPb *pb.Post) *domain.Post {
	post := &domain.Post{
		Username: (*postPb).Username,
		Content:  (*postPb).Content,
		Image:    (*postPb).Image,
		Date:     (*((*postPb).Date)).AsTime(),
	}

	post.Likes = []string{}
	for _, like := range (*postPb).Likes {
		post.Likes = append(post.Likes, like)
	}

	post.Dislikes = []string{}
	for _, dislike := range (*postPb).Dislikes {
		post.Dislikes = append(post.Dislikes, dislike)
	}

	post.Comments = []domain.Comment{}
	for _, comment := range postPb.Comments {
		id, err := primitive.ObjectIDFromHex(comment.Id)
		if err != nil {
			continue
		}
		post.Comments = append(post.Comments, domain.Comment{
			Id:       id,
			Content:  comment.Content,
			Date:     comment.Date.AsTime(),
			Username: comment.Username,
		})
	}
	return post
}

func mapReactionToDomain(reactionPb *pb.Reaction) *domain.Reaction {
	postId, err := primitive.ObjectIDFromHex((*reactionPb).PostId)
	if err != nil {
		return &domain.Reaction{}
	}

	reaction := &domain.Reaction{
		Username:     (*reactionPb).Username,
		PostId:       postId,
		ReactionType: enum.ReactionType((*reactionPb).ReactionType),
	}

	return reaction
}
