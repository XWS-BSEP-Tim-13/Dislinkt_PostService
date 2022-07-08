package persistence

import (
	"context"
	"fmt"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_PostService/domain"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_PostService/tracer"
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

func (store *PostMongoDBStore) Get(ctx context.Context, id primitive.ObjectID) (*domain.Post, error) {
	span := tracer.StartSpanFromContext(ctx, "DB Get")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)

	filter := bson.M{"_id": id}
	return store.filterOne(ctx, filter)
}

func (store *PostMongoDBStore) GetAll(ctx context.Context) ([]*domain.Post, error) {
	span := tracer.StartSpanFromContext(ctx, "DB GetAll")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)

	filter := bson.D{{}}
	return store.filter(ctx, filter)
}

func (store *PostMongoDBStore) Insert(ctx context.Context, product *domain.Post) error {
	span := tracer.StartSpanFromContext(ctx, "DB Insert")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)

	result, err := store.posts.InsertOne(context.TODO(), product)
	if err != nil {
		return err
	}
	product.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (store *PostMongoDBStore) DeleteAll(ctx context.Context) {
	span := tracer.StartSpanFromContext(ctx, "DB DeleteAll")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)

	store.posts.DeleteMany(context.TODO(), bson.D{{}})
}

func (store *PostMongoDBStore) GetByUser(ctx context.Context, username string) ([]*domain.Post, error) {
	span := tracer.StartSpanFromContext(ctx, "DB GetByUser")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)

	filter := bson.M{"username": username}
	return store.filter(ctx, filter)
}

func (store *PostMongoDBStore) filter(ctx context.Context, filter interface{}) ([]*domain.Post, error) {
	span := tracer.StartSpanFromContext(ctx, "DB filter")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)

	cursor, err := store.posts.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (store *PostMongoDBStore) Delete(ctx context.Context, id primitive.ObjectID) error {
	span := tracer.StartSpanFromContext(ctx, "DB Delete")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)

	filter := bson.M{"_id": id}
	_, err := store.posts.DeleteOne(context.TODO(), filter)
	return err
}

func (store *PostMongoDBStore) filterOne(ctx context.Context, filter interface{}) (product *domain.Post, err error) {
	span := tracer.StartSpanFromContext(ctx, "DB filterOne")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)

	result := store.posts.FindOne(context.TODO(), filter)
	err = result.Decode(&product)
	return
}

func (store *PostMongoDBStore) UpdateReactions(ctx context.Context, post *domain.Post) (string, error) {
	span := tracer.StartSpanFromContext(ctx, "DB UpdateReactions")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)

	filter := bson.M{"_id": (*post).Id}
	replacementObj := post
	_, err := store.posts.ReplaceOne(context.TODO(), filter, replacementObj)

	fmt.Printf("Updated \n")
	if err != nil {
		return "", err
	}
	return (*post).Id.Hex(), nil
}

func (store *PostMongoDBStore) GetFeed(ctx context.Context, page int64, usernames []string) (*domain.FeedDto, error) {
	span := tracer.StartSpanFromContext(ctx, "DB GetFeed")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)

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

func (store *PostMongoDBStore) GetFeedAnonymous(ctx context.Context, page int64) (*domain.FeedDto, error) {
	span := tracer.StartSpanFromContext(ctx, "DB GetFeedAnonymous")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)

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
