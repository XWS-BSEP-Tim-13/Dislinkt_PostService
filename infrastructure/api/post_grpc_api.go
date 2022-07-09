package api

import (
	"context"
	"fmt"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_PostService/application"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_PostService/domain"
	pb "github.com/XWS-BSEP-Tim-13/Dislinkt_PostService/infrastructure/grpc/proto"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_PostService/jwt"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_PostService/logger"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_PostService/tracer"
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
	span := tracer.StartSpanFromContext(ctx, "API Get")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)

	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	post, err := handler.service.Get(ctx, objectId)
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
	span := tracer.StartSpanFromContext(ctx, "API GetAll")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)

	posts, err := handler.service.GetAll(ctx)
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
	span := tracer.StartSpanFromContext(ctx, "API GetByUser")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)

	username := request.Username
	posts, err := handler.service.GetByUser(ctx, username)
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
	span := tracer.StartSpanFromContext(ctx, "API CreatePost")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)

	//username, _ := jwt.ExtractUsernameFromToken(ctx)
	post := mapPostDtoPbToDomain(request.Post)

	newPost, err := handler.service.CreateNewPost(ctx, post)
	if err != nil {
		handler.logger.ErrorMessage("User: " + post.Username + " | Action: CP")
		return nil, status.Error(400, err.Error())
	}

	response := &pb.NewPost{
		Post: mapPostToPb(newPost),
	}

	handler.logger.InfoMessage("User: " + post.Username + " | Action: CP")
	return response, nil
}

func (handler *PostHandler) ReactToPost(ctx context.Context, request *pb.ReactionRequest) (*pb.ReactionResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "API ReactToPost")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)

	username, _ := jwt.ExtractUsernameFromToken(ctx)
	reaction := mapReactionToDomain((*request).Reaction)

	postId, err := handler.service.ReactToPost(ctx, reaction)
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
	span := tracer.StartSpanFromContext(ctx, "API CreateCommentOnPost")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)

	username, _ := jwt.ExtractUsernameFromToken(ctx)
	comment := mapCommentDtoToDomain(request.Comment)
	postId := (*request).PostId

	newComment, err := handler.service.CreateNewComment(ctx, comment, postId)
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

func (handler *PostHandler) DeletePost(ctx context.Context, requset *pb.GetRequest) (*pb.GetAllRequest, error) {
	span := tracer.StartSpanFromContext(ctx, "API DeletePost")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)

	id, _ := primitive.ObjectIDFromHex(requset.Id)
	err := handler.service.Delete(ctx, id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	response := &pb.GetAllRequest{}
	return response, nil
}

func (handler *PostHandler) GetFeedPosts(ctx context.Context, request *pb.FeedRequest) (*pb.FeedResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "API GetFeedPosts")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)

	principal, _ := jwt.ExtractUsernameFromToken(ctx)
	usernames := mapUsernamesToDomain(request.Usernames)
	var dto *domain.FeedDto
	var err error
	if len(usernames) == 0 {
		dto, err = handler.service.GetFeedPostsAnonymous(ctx, request.Page)
	} else {
		dto, err = handler.service.GetFeedPosts(ctx, request.Page, usernames)
	}
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
	span := tracer.StartSpanFromContext(ctx, "API GetFeedPostsAnonymous")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)

	dto, err := handler.service.GetFeedPostsAnonymous(ctx, request.Page)
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
	span := tracer.StartSpanFromContext(ctx, "API UploadImage")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)

	principal, _ := jwt.ExtractUsernameFromToken(ctx)
	image := request.Image
	imagePath, err := handler.service.UploadImage(ctx, image)
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
	span := tracer.StartSpanFromContext(ctx, "API GetImage")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)

	principal, _ := jwt.ExtractUsernameFromToken(ctx)
	imagePath := request.ImagePath
	image, err := handler.service.GetImage(ctx, imagePath)
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

func (handler *PostHandler) GetMessagesForUsers(ctx context.Context, request *pb.GetByUserRequest) (*pb.MessageResponse, error) {
	principal, _ := jwt.ExtractUsernameFromToken(ctx)
	messages, err := handler.service.GetMessagesByUsers(request.Username, principal)
	if err != nil {
		return nil, err
	}

	response := &pb.MessageResponse{
		Messages: mapMessagesToPb(messages),
	}
	return response, nil
}

func (handler *PostHandler) GetMessagesForUser(ctx context.Context, request *pb.GetAllRequest) (*pb.GetByUserResponse, error) {
	principal, _ := jwt.ExtractUsernameFromToken(ctx)
	messages, err := handler.service.GetMessagesByUser(principal)
	if err != nil {
		return nil, err
	}

	response := &pb.GetByUserResponse{
		Messages: []*pb.MessageUsers{},
	}

	for _, message := range messages {
		current := mapMessagesToPb(message)
		response.Messages = append(response.Messages, current)
	}

	return response, nil
}

func (handler *PostHandler) SaveMessage(ctx context.Context, request *pb.SaveMessageRequest) (*pb.GetAllRequest, error) {
	message := mapMessagePbToDomain(request.Message)
	err := handler.service.SaveMessage(message)
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllRequest{}
	return response, nil
}

func (handler *PostHandler) GetEvents(ctx context.Context, request *pb.EventRequest) (*pb.Events, error) {
	events, err := handler.service.GetAllEvents()
	if err != nil {
		return nil, err
	}
	response := &pb.Events{
		Events: []*pb.Event{},
	}

	for _, event := range events {
		current := mapEventToPB(event)
		response.Events = append(response.Events, current)
	}

	return response, nil
}
