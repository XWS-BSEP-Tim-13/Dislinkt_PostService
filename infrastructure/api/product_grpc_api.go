package api

import (
	"context"
	"fmt"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_PostService/application"
	pb "github.com/XWS-BSEP-Tim-13/Dislinkt_PostService/infrastructure/grpc/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/status"
)

type PostHandler struct {
	pb.UnimplementedPostServiceServer
	service *application.PostService
}

func NewPostHandler(service *application.PostService) *PostHandler {
	return &PostHandler{
		service: service,
	}
}

func (handler *PostHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	post, err := handler.service.Get(objectId)
	if err != nil {
		return nil, err
	}
	postPb := mapPostToPb(post)
	response := &pb.GetResponse{
		Post: postPb,
	}
	return response, nil
}

func (handler *PostHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	posts, err := handler.service.GetAll()
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		Posts: []*pb.Post{},
	}
	for _, post := range posts {
		fmt.Printf("Image %s\n", post.Image)
		current := mapPostToPb(post)
		response.Posts = append(response.Posts, current)
	}
	return response, nil
}

func (handler *PostHandler) GetByUser(ctx context.Context, request *pb.GetByUserRequest) (*pb.GetAllResponse, error) {
	username := request.Username
	posts, err := handler.service.GetByUser(username)
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		Posts: []*pb.Post{},
	}
	for _, post := range posts {
		current := mapPostToPb(post)
		response.Posts = append(response.Posts, current)
	}
	return response, nil
}

func (handler *PostHandler) CreatePost(ctx context.Context, request *pb.NewPostRequest) (*pb.NewPost, error) {
	fmt.Println((*request).Post)
	post := mapPostDtoPbToDomain(request.Post, "stefanljubovic")
	fmt.Println(post)

	newPost, err := handler.service.CreateNewPost(post)
	if err != nil {
		return nil, status.Error(400, err.Error())
	}

	response := &pb.NewPost{
		Post: mapPostToPb(newPost),
	}
	return response, nil
}

func (handler *PostHandler) ReactToPost(ctx context.Context, request *pb.ReactionRequest) (*pb.ReactionResponse, error) {
	fmt.Println(request.Reaction)
	reaction := mapReactionToDomain((*request).Reaction)
	fmt.Println(reaction)

	postId, err := handler.service.ReactToPost(reaction)
	if err != nil {
		return nil, err
	}

	reactionResponse := &pb.ReactionResponse{
		PostId: postId,
	}

	return reactionResponse, nil
}

func (handler *PostHandler) CreateCommentOnPost(ctx context.Context, request *pb.CommentRequest) (*pb.CommentResponse, error) {
	fmt.Println((*request).Comment)
	comment := mapCommentDtoToDomain(request.Comment)
	fmt.Println(comment)
	postId := (*request).PostId

	newComment, err := handler.service.CreateNewComment(comment, postId)
	if err != nil {
		return nil, status.Error(400, err.Error())
	}

	response := &pb.CommentResponse{
		CommentId: newComment.Id.Hex(),
	}

	return response, nil
}

func (handler *PostHandler) GetFeedPosts(ctx context.Context, request *pb.FeedRequest) (*pb.FeedResponse, error) {
	fmt.Println("Posts microservice")
	usernames := mapUsernamesToDomain(request.Usernames)
	dto, err := handler.service.GetFeedPosts(request.Page, usernames)
	if err != nil {
		return nil, err
	}
	pbPosts := []*pb.Post{}
	for _, post := range dto.Posts {
		current := mapPostToPb(post)
		pbPosts = append(pbPosts, current)
	}
	response := &pb.FeedResponse{
		Posts:    pbPosts,
		LastPage: dto.LastPage,
		Page:     dto.Page,
	}
	return response, nil
}

func (handler *PostHandler) UploadImage(ctx context.Context, request *pb.ImageRequest) (*pb.ImageResponse, error) {
	fmt.Println("Upload slike")
	image := request.Image
	imagePath, err := handler.service.UploadImage(image)
	if err != nil {
		return nil, err
	}
	response := &pb.ImageResponse{
		ImagePath: imagePath,
	}
	return response, nil
}

func (handler *PostHandler) GetImage(ctx context.Context, request *pb.ImageResponse) (*pb.ImageRequest, error) {
	fmt.Println("Dobavljanje slike")
	imagePath := request.ImagePath
	image, err := handler.service.GetImage(imagePath)
	if err != nil {
		return nil, err
	}
	response := &pb.ImageRequest{
		Image: image,
	}
	return response, nil
}
