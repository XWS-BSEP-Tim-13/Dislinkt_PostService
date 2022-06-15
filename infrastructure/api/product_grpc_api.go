package api

import (
	"context"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_PostService/application"
	pb "github.com/XWS-BSEP-Tim-13/Dislinkt_PostService/infrastructure/grpc/proto"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_PostService/jwt"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_PostService/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/status"
)

type PostHandler struct {
	pb.UnimplementedPostServiceServer
	service *application.PostService
	logger  *logger.Logger
}

func NewPostHandler(service *application.PostService, logger *logger.Logger) *PostHandler {
	return &PostHandler{
		service: service,
		logger:  logger,
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
		handler.logger.ErrorMessage("Action: GP/ " + id)
		return nil, err
	}
	postPb := mapPostToPb(post)
	response := &pb.GetResponse{
		Post: postPb,
	}

	handler.logger.InfoMessage("Action: Get GP/ " + id)
	return response, nil
}

func (handler *PostHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	posts, err := handler.service.GetAll()
	if err != nil {
		handler.logger.ErrorMessage("Action: GP")
		return nil, err
	}
	response := &pb.GetAllResponse{
		Posts: []*pb.Post{},
	}
	for _, post := range posts {
		current := mapPostToPb(post)
		response.Posts = append(response.Posts, current)
	}

	handler.logger.InfoMessage("Action: GP")
	return response, nil
}

func (handler *PostHandler) GetByUser(ctx context.Context, request *pb.GetByUserRequest) (*pb.GetAllResponse, error) {
	username := request.Username
	posts, err := handler.service.GetByUser(username)
	if err != nil {
		handler.logger.ErrorMessage("Action: GP/" + username)
		return nil, err
	}
	response := &pb.GetAllResponse{
		Posts: []*pb.Post{},
	}
	for _, post := range posts {
		current := mapPostToPb(post)
		response.Posts = append(response.Posts, current)
	}

	handler.logger.InfoMessage("Action: GP/ " + username)
	return response, nil
}

func (handler *PostHandler) CreatePost(ctx context.Context, request *pb.NewPostRequest) (*pb.NewPost, error) {
	username, _ := jwt.ExtractUsernameFromToken(ctx)
	post := mapPostDtoPbToDomain(request.Post, username)

	newPost, err := handler.service.CreateNewPost(post)
	if err != nil {
		handler.logger.ErrorMessage("User: " + username + " | Action: CP")
		return nil, status.Error(400, err.Error())
	}

	response := &pb.NewPost{
		Post: mapPostToPb(newPost),
	}

	handler.logger.InfoMessage("User: " + username + " | Action: CP")
	return response, nil
}

func (handler *PostHandler) ReactToPost(ctx context.Context, request *pb.ReactionRequest) (*pb.ReactionResponse, error) {
	username, _ := jwt.ExtractUsernameFromToken(ctx)
	reaction := mapReactionToDomain((*request).Reaction)

	postId, err := handler.service.ReactToPost(reaction)
	if err != nil {
		return nil, err
	}

	reactionResponse := &pb.ReactionResponse{
		PostId: postId,
	}

	handler.logger.InfoMessage("User: " + username + " | Action: RoP")
	return reactionResponse, nil
}

func (handler *PostHandler) CreateCommentOnPost(ctx context.Context, request *pb.CommentRequest) (*pb.CommentResponse, error) {
	username, _ := jwt.ExtractUsernameFromToken(ctx)
	comment := mapCommentDtoToDomain(request.Comment)
	postId := (*request).PostId

	newComment, err := handler.service.CreateNewComment(comment, postId)
	if err != nil {
		handler.logger.ErrorMessage("User: " + username + " | Action: CoP")
		return nil, status.Error(400, err.Error())
	}

	response := &pb.CommentResponse{
		CommentId: newComment.Id.Hex(),
	}

	handler.logger.InfoMessage("User: " + username + " | Action: CoP")
	return response, nil
}

func (handler *PostHandler) GetFeedPosts(ctx context.Context, request *pb.FeedRequest) (*pb.FeedResponse, error) {
	principal, _ := jwt.ExtractUsernameFromToken(ctx)
	usernames := mapUsernamesToDomain(request.Usernames)
	dto, err := handler.service.GetFeedPosts(request.Page, usernames)
	if err != nil {
		handler.logger.ErrorMessage("User: " + principal + " | Action: GFP")
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
	handler.logger.InfoMessage("User: " + principal + " | Action: GFP")
	return response, nil
}

func (handler *PostHandler) GetFeedPostsAnonymous(ctx context.Context, request *pb.FeedRequestAnonymous) (*pb.FeedResponse, error) {
	dto, err := handler.service.GetFeedPostsAnonymous(request.Page)
	if err != nil {
		handler.logger.ErrorMessage("User: Anonymous | Action: GFP")
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
	handler.logger.InfoMessage("User: Anonymous | Action: GFP")
	return response, nil
}

func (handler *PostHandler) UploadImage(ctx context.Context, request *pb.ImageRequest) (*pb.ImageResponse, error) {
	principal, _ := jwt.ExtractUsernameFromToken(ctx)
	image := request.Image
	imagePath, err := handler.service.UploadImage(image)
	if err != nil {
		handler.logger.ErrorMessage("User: " + principal + " | Action: ImgU")
		return nil, err
	}
	response := &pb.ImageResponse{
		ImagePath: imagePath,
	}

	handler.logger.InfoMessage("User: " + principal + " | Action: ImgU")
	return response, nil
}

func (handler *PostHandler) GetImage(ctx context.Context, request *pb.ImageResponse) (*pb.ImageRequest, error) {
	principal, _ := jwt.ExtractUsernameFromToken(ctx)
	imagePath := request.ImagePath
	image, err := handler.service.GetImage(imagePath)
	if err != nil {
		handler.logger.ErrorMessage("User: " + principal + " | Action: GImg")
		return nil, err
	}
	response := &pb.ImageRequest{
		Image: image,
	}
	handler.logger.InfoMessage("User: " + principal + " | Action: GImg")
	return response, nil
}
