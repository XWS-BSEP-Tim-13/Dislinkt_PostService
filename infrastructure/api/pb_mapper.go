package api

import (
	"github.com/XWS-BSEP-Tim-13/Dislinkt_PostService/domain"
	pb "github.com/XWS-BSEP-Tim-13/Dislinkt_PostService/infrastructure/grpc/proto"
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
