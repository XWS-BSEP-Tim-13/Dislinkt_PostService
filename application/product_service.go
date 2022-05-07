package application

import (
	"errors"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_PostService/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
