package persistence

import (
	"context"
	"fmt"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_PostService/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"math"
)

const (
	DATABASE   = "posts"
	COLLECTION = "post"
)

type PostMongoDBStore struct {
	posts *mongo.Collection
}

func NewPostMongoDBStore(client *mongo.Client) domain.PostStore {
	posts := client.Database(DATABASE).Collection(COLLECTION)
	return &PostMongoDBStore{
		posts: posts,
	}
}

func (store *PostMongoDBStore) Get(id primitive.ObjectID) (*domain.Post, error) {
	filter := bson.M{"_id": id}
	return store.filterOne(filter)
}

func (store *PostMongoDBStore) GetAll() ([]*domain.Post, error) {
	filter := bson.D{{}}
	return store.filter(filter)
}

func (store *PostMongoDBStore) Insert(product *domain.Post) error {
	result, err := store.posts.InsertOne(context.TODO(), product)
	if err != nil {
		return err
	}
	product.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (store *PostMongoDBStore) DeleteAll() {
	store.posts.DeleteMany(context.TODO(), bson.D{{}})
}

func (store *PostMongoDBStore) GetByUser(username string) ([]*domain.Post, error) {
	filter := bson.M{"username": username}
	return store.filter(filter)
}

func (store *PostMongoDBStore) filter(filter interface{}) ([]*domain.Post, error) {
	cursor, err := store.posts.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (store *PostMongoDBStore) Delete(id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	_, err := store.posts.DeleteOne(context.TODO(), filter)
	return err
}

func (store *PostMongoDBStore) filterOne(filter interface{}) (product *domain.Post, err error) {
	result := store.posts.FindOne(context.TODO(), filter)
	err = result.Decode(&product)
	return
}

func (store *PostMongoDBStore) UpdateReactions(post *domain.Post) (string, error) {
	filter := bson.M{"_id": (*post).Id}
	replacementObj := post
	_, err := store.posts.ReplaceOne(context.TODO(), filter, replacementObj)

	fmt.Printf("Updated \n")
	if err != nil {
		return "", err
	}
	return (*post).Id.Hex(), nil
}

func (store *PostMongoDBStore) GetFeed(page int64, usernames []string) (*domain.FeedDto, error) {
	filter := bson.D{{"username", bson.D{{"$in", usernames}}}}
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"date", -1}})
	var perPage int64 = 5
	total, _ := store.posts.CountDocuments(context.TODO(), filter)
	findOptions.SetSkip((page - 1) * perPage)
	findOptions.SetLimit(perPage)
	cursor, err := store.posts.Find(context.TODO(), filter, findOptions)
	fmt.Printf("Total %d, %f\n", total, math.Ceil(float64(total)/float64(perPage)))
	if err != nil {
		return nil, err
	}
	posts, _ := decode(cursor)
	dto := domain.FeedDto{
		Posts:    posts,
		Page:     page,
		LastPage: int64(math.Ceil(float64(total) / float64(perPage))),
	}
	return &dto, nil
}

func (store *PostMongoDBStore) GetFeedAnonymous(page int64) (*domain.FeedDto, error) {
	filter := bson.D{}
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"likes", -1}, {"date", -1}})
	var perPage int64 = 5
	total, _ := store.posts.CountDocuments(context.TODO(), filter)
	findOptions.SetSkip((page - 1) * perPage)
	findOptions.SetLimit(perPage)
	cursor, err := store.posts.Find(context.TODO(), filter, findOptions)
	fmt.Printf("Total %d, %f\n", total, math.Ceil(float64(total)/float64(perPage)))
	if err != nil {
		return nil, err
	}
	posts, _ := decode(cursor)
	dto := domain.FeedDto{
		Posts:    posts,
		Page:     page,
		LastPage: int64(math.Ceil(float64(total) / float64(perPage))),
	}
	return &dto, nil
}

func decode(cursor *mongo.Cursor) (products []*domain.Post, err error) {
	for cursor.Next(context.TODO()) {
		var product domain.Post
		err = cursor.Decode(&product)
		if err != nil {
			return
		}
		products = append(products, &product)
	}
	err = cursor.Err()
	return
}
