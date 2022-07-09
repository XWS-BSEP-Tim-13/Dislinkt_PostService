package application

import (
	"context"
	"errors"
	"fmt"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_PostService/domain"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_PostService/logger"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_PostService/tracer"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_PostService/util"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/status"
	"time"
)

type PostService struct {
	store        domain.PostStore
	imageStore   domain.UploadImageStore
	logger       *logger.Logger
	messageStore domain.MessageStore
	eventStore   domain.EventStore
}

func NewPostService(store domain.PostStore, imageStore domain.UploadImageStore, logger *logger.Logger, messageStore domain.MessageStore, eventStore domain.EventStore) *PostService {
	return &PostService{
		store:        store,
		imageStore:   imageStore,
		logger:       logger,
		messageStore: messageStore,
		eventStore:   eventStore,
	}
}

func (service *PostService) Get(ctx context.Context, id primitive.ObjectID) (*domain.Post, error) {
	span := tracer.StartSpanFromContext(ctx, "SERVICE Get")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.store.Get(ctx, id)
}

func (service *PostService) GetAll(ctx context.Context) ([]*domain.Post, error) {
	span := tracer.StartSpanFromContext(ctx, "SERVICE GetAll")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.store.GetAll(ctx)
}

func (service *PostService) GetByUser(ctx context.Context, username string) ([]*domain.Post, error) {
	span := tracer.StartSpanFromContext(ctx, "SERVICE GetByUser")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)

	return service.store.GetByUser(ctx, username)
}

func (service *PostService) ReactToPost(ctx context.Context, reaction *domain.Reaction) (string, error) {
	span := tracer.StartSpanFromContext(ctx, "SERVICE ReactToPost")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)

	post, err := service.Get(ctx, (*reaction).PostId)
	if err != nil {
		service.logger.ErrorMessage("User: " + post.Username + " Action: PNF")
		return "", status.Error(500, err.Error())
	}
	if post == nil {
		service.logger.ErrorMessage("User: " + post.Username + " Action: PNF")
		return "", status.Error(400, "Post not found!")
	}

	if (*reaction).ReactionType == 0 {
		foundInLikes := false
		for idx, like := range (*post).Likes {
			if like == (*reaction).Username {
				(*post).Likes = util.RemoveElement((*post).Likes, idx)
				foundInLikes = true
				break
			}
		}

		if !foundInLikes {
			for idx, dislike := range (*post).Dislikes {
				if dislike == (*reaction).Username {
					(*post).Dislikes = util.RemoveElement((*post).Dislikes, idx)
					break
				}
			}

			(*post).Likes = append((*post).Likes, (*reaction).Username)
			var event = domain.Event{Id: primitive.NewObjectID(), User: (*reaction).Username, Action: `Liked post of user ` + post.Username, Published: time.Now()}
			service.eventStore.Insert(&event)
		}
	} else if (*reaction).ReactionType == 1 {
		foundInDislikes := false
		for idx, dislike := range (*post).Dislikes {
			if dislike == (*reaction).Username {
				(*post).Dislikes = util.RemoveElement((*post).Dislikes, idx)
				foundInDislikes = true
				break
			}
		}

		if !foundInDislikes {
			for idx, like := range (*post).Likes {
				if like == (*reaction).Username {
					(*post).Likes = util.RemoveElement((*post).Likes, idx)
					break
				}
			}

			(*post).Dislikes = append((*post).Dislikes, (*reaction).Username)
			var event = domain.Event{Id: primitive.NewObjectID(), User: (*reaction).Username, Action: `Disliked post of user ` + post.Username, Published: time.Now()}
			service.eventStore.Insert(&event)
		}
	} else {
		service.logger.ErrorMessage("User: " + post.Username + " Action: RNS")
		return "", status.Error(400, "This reaction is not supported!")
	}

	postID, err := service.store.UpdateReactions(ctx, post)
	if err != nil {
		service.logger.ErrorMessage("User: " + post.Username + " Action: UPR")
		return "", status.Error(500, "Error while updating post!")
	}

	return postID, nil
}

func (service *PostService) CreateNewPost(ctx context.Context, post *domain.Post) (*domain.Post, error) {
	span := tracer.StartSpanFromContext(ctx, "SERVICE CreateNewPost")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)

	err := service.store.Insert(ctx, post)
	if err != nil {
		err := errors.New("error while creating new post")
		return nil, err
	}

	var event = domain.Event{Id: primitive.NewObjectID(), User: post.Username, Action: `Liked post of user ` + post.Username, Published: time.Now()}
	service.eventStore.Insert(&event)

	return post, nil
}

func (service *PostService) CreateNewComment(ctx context.Context, comment *domain.Comment, postId string) (*domain.Comment, error) {
	span := tracer.StartSpanFromContext(ctx, "SERVICE CreateNewComment")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)

	(*comment).Id = primitive.NewObjectID()
	id, err := primitive.ObjectIDFromHex(postId)
	if err != nil {
		return nil, err
	}
	post, err := service.store.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	var event = domain.Event{Id: primitive.NewObjectID(), User: (*comment).Username, Action: `Comment on post of user ` + post.Username, Published: time.Now()}
	service.eventStore.Insert(&event)
	(*post).Comments = append((*post).Comments, *comment)
	_, err = service.store.UpdateReactions(ctx, post)
	if err != nil {
		err := errors.New("error while creating new comment")
		return nil, err
	}
	return comment, nil
}

func (service *PostService) Delete(ctx context.Context, id primitive.ObjectID) error {
	span := tracer.StartSpanFromContext(ctx, "SERVICE Delete")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)

	err := service.store.Delete(ctx, id)
	return err
}

func (service *PostService) GetFeedPosts(ctx context.Context, page int64, usernames []string) (*domain.FeedDto, error) {
	span := tracer.StartSpanFromContext(ctx, "SERVICE GetFeedPosts")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)

	dto, err := service.store.GetFeed(ctx, page, usernames)
	if err != nil {
		return nil, err
	}
	return dto, err
}

func (service *PostService) GetFeedPostsAnonymous(ctx context.Context, page int64) (*domain.FeedDto, error) {
	span := tracer.StartSpanFromContext(ctx, "SERVICE GetFeedPostsAnonymous")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)

	fmt.Println("Posts anonymus ", page)
	dto, err := service.store.GetFeedAnonymous(ctx, page)
	if err != nil {
		return nil, err
	}
	return dto, err
}

func (service *PostService) UploadImage(ctx context.Context, image []byte) (string, error) {
	span := tracer.StartSpanFromContext(ctx, "SERVICE UploadImage")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)

	filename, err := service.imageStore.UploadObject(ctx, image)
	if err != nil {
		return "", nil
	}
	return filename, nil
}

func (service *PostService) GetImage(ctx context.Context, imagePath string) ([]byte, error) {
	span := tracer.StartSpanFromContext(ctx, "SERVICE GetImage")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)

	image := service.imageStore.GetObject(ctx, imagePath)
	return image, nil
}

func (service *PostService) GetMessagesByUsers(firstUsername, secondUsername string) (*domain.MessageUsers, error) {
	messages, err := service.messageStore.GetByUsers(firstUsername, secondUsername)
	if err != nil {
		messages = &domain.MessageUsers{
			Id:         primitive.NewObjectID(),
			FirstUser:  firstUsername,
			SecondUser: secondUsername,
			Messages:   []domain.Message{},
		}
		err = service.messageStore.Insert(messages)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}
	return messages, nil
}
func (service *PostService) SaveMessage(message *domain.Message) error {
	error := service.messageStore.SendMessage(message)
	var event = domain.Event{Id: primitive.NewObjectID(), User: message.MessageFrom, Action: `Message to user ` + message.MessageTo, Published: time.Now()}
	service.eventStore.Insert(&event)
	return error
}

func (service *PostService) GetMessagesByUser(username string) ([]*domain.MessageUsers, error) {
	messages, err := service.messageStore.GetByUser(username)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return messages, nil
}

func (service *PostService) GetAllEvents() ([]*domain.Event, error) {
	return service.eventStore.GetAll()
}
