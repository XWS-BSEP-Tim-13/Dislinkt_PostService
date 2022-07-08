package application

import (
	"errors"
	"fmt"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_PostService/domain"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_PostService/logger"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_PostService/util"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/status"
)

type PostService struct {
	store        domain.PostStore
	imageStore   domain.UploadImageStore
	logger       *logger.Logger
	messageStore domain.MessageStore
}

func NewPostService(store domain.PostStore, imageStore domain.UploadImageStore, logger *logger.Logger, messageStore domain.MessageStore) *PostService {
	return &PostService{
		store:        store,
		imageStore:   imageStore,
		logger:       logger,
		messageStore: messageStore,
	}
}

func (service *PostService) Get(id primitive.ObjectID) (*domain.Post, error) {
	return service.store.Get(id)
}

func (service *PostService) GetAll() ([]*domain.Post, error) {
	return service.store.GetAll()
}

func (service *PostService) GetByUser(username string) ([]*domain.Post, error) {
	return service.store.GetByUser(username)
}

func (service *PostService) ReactToPost(reaction *domain.Reaction) (string, error) {
	post, err := service.Get((*reaction).PostId)
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
		}
	} else {
		service.logger.ErrorMessage("User: " + post.Username + " Action: RNS")
		return "", status.Error(400, "This reaction is not supported!")
	}

	postID, err := service.store.UpdateReactions(post)
	if err != nil {
		service.logger.ErrorMessage("User: " + post.Username + " Action: UPR")
		return "", status.Error(500, "Error while updating post!")
	}

	return postID, nil
}

func (service *PostService) CreateNewPost(post *domain.Post) (*domain.Post, error) {
	err := service.store.Insert(post)
	if err != nil {
		err := errors.New("error while creating new post")
		return nil, err
	}

	return post, nil
}

func (service *PostService) CreateNewComment(comment *domain.Comment, postId string) (*domain.Comment, error) {
	(*comment).Id = primitive.NewObjectID()
	id, err := primitive.ObjectIDFromHex(postId)
	if err != nil {
		return nil, err
	}
	post, err := service.store.Get(id)
	if err != nil {
		return nil, err
	}
	(*post).Comments = append((*post).Comments, *comment)
	_, err = service.store.UpdateReactions(post)
	if err != nil {
		err := errors.New("error while creating new comment")
		return nil, err
	}
	return comment, nil
}

func (service *PostService) Delete(id primitive.ObjectID) error {
	err := service.store.Delete(id)
	return err
}

func (service *PostService) GetFeedPosts(page int64, usernames []string) (*domain.FeedDto, error) {
	dto, err := service.store.GetFeed(page, usernames)
	if err != nil {
		return nil, err
	}
	return dto, err
}

func (service *PostService) GetFeedPostsAnonymous(page int64) (*domain.FeedDto, error) {
	fmt.Println("Posts anonymus ", page)
	dto, err := service.store.GetFeedAnonymous(page)
	if err != nil {
		return nil, err
	}
	return dto, err
}

func (service *PostService) UploadImage(image []byte) (string, error) {
	filename, err := service.imageStore.UploadObject(image)
	if err != nil {
		return "", nil
	}
	return filename, nil
}

func (service *PostService) GetImage(imagePath string) ([]byte, error) {
	image := service.imageStore.GetObject(imagePath)
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
	return service.messageStore.SendMessage(message)
}

func (service *PostService) GetMessagesByUser(username string) ([]*domain.MessageUsers, error) {
	messages, err := service.messageStore.GetByUser(username)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return messages, nil
}
