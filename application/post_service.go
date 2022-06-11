package application

import (
	"errors"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_PostService/domain"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_PostService/logger"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_PostService/util"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/status"
)

type PostService struct {
	store      domain.PostStore
	imageStore domain.UploadImageStore
	logger     *logger.Logger
}

func NewPostService(store domain.PostStore, imageStore domain.UploadImageStore, logger *logger.Logger) *PostService {
	return &PostService{
		store:      store,
		imageStore: imageStore,
		logger:     logger,
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
		service.logger.ErrorMessage("User: " + post.Username + " Action: Post not found")
		return "", status.Error(500, err.Error())
	}
	if post == nil {
		service.logger.ErrorMessage("User: " + post.Username + " Action: Post not found")
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
		service.logger.ErrorMessage("User: " + post.Username + " Action: Reaction not supported")
		return "", status.Error(400, "This reaction is not supported!")
	}

	postID, err := service.store.UpdateReactions(post)
	if err != nil {
		service.logger.ErrorMessage("User: " + post.Username + " Action: Update post reaction")
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

func (service *PostService) GetFeedPosts(page int64, usernames []string) (*domain.FeedDto, error) {
	dto, err := service.store.GetFeed(page, usernames)
	if err != nil {
		return nil, err
	}
	return dto, err
}

func (service *PostService) GetFeedPostsAnonymous(page int64) (*domain.FeedDto, error) {
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
