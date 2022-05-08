package application

import (
	"errors"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_PostService/domain"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_PostService/util"
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
		return "", status.Error(400, "This reaction is not supported!")
	}

	postID, err := service.store.UpdateReactions(post)
	if err != nil {
		return "", status.Error(500, "Error while updating post!")
	}

	return postID, nil
}

func (service *PostService) CreateNewPost(post *domain.Post) (*domain.Post, error) {
	(*post).Id = primitive.NewObjectID()
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
