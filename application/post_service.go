package application

import (
	"errors"
	"fmt"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_PostService/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/status"
)

type PostService struct {
	store domain.PostStore
}

func NewPostService(store domain.PostStore) *PostService {
	return &PostService{
		store: store,
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
		return "", status.Error(500, err.Error())
	}
	if post == nil {
		return "", status.Error(400, "Post not found!")
	}

	if (*reaction).ReactionType == 0 {
		fmt.Println("This is like")
		(*post).Likes = append((*post).Likes, (*reaction).Username)
		fmt.Println("Liked")
	} else if (*reaction).ReactionType == 1 {
		fmt.Println("This is dislike")
		(*post).Dislikes = append((*post).Dislikes, (*reaction).Username)
		fmt.Println("Disliked")
	} else {
		return "", status.Error(400, "This reaction is not supported!")
	}

	postID, err := service.store.UpdateReactions(post)
	if err != nil {
		return "", status.Error(500, "Error while updating post!")
	}

	return postID, nil
}

func (service *PostService) CreateNewPost(post *domain.Post) (*domain.Post, error) {
	//dbUser, _ := service.store.GetByUsername((*user).Username)
	//if dbUser == nil {
	//	err := errors.New("user with this username not exists")
	//	return nil, err
	//}
	(*post).Id = primitive.NewObjectID()
	err := service.store.Insert(post)
	if err != nil {
		err := errors.New("error while creating new post")
		return nil, err
	}

	return post, nil
}
