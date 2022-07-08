package persistence

import (
	"context"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_PostService/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE_MSG   = "messages"
	COLLECTION_MSG = "message"
)

type MessageMongoDBStore struct {
	messages *mongo.Collection
}

func (store MessageMongoDBStore) Insert(messages *domain.MessageUsers) error {
	result, err := store.messages.InsertOne(context.TODO(), messages)
	if err != nil {
		return err
	}
	messages.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (store MessageMongoDBStore) DeleteAll() {
	store.messages.DeleteMany(context.TODO(), bson.D{{}})
}

func (store MessageMongoDBStore) GetByUsers(firstUsername, secondUsername string) (*domain.MessageUsers, error) {
	filter1 := bson.M{"first_user": firstUsername, "second_user": secondUsername}
	filter2 := bson.M{"first_user": secondUsername, "second_user": firstUsername}
	filterr := bson.M{
		"$or": []bson.M{
			filter1,
			filter2,
		},
	}

	return store.filterOne(filterr)
}

func (store MessageMongoDBStore) SendMessage(message *domain.Message) error {
	filter1 := bson.M{"first_user": message.MessageTo, "second_user": message.MessageFrom}
	filter2 := bson.M{"first_user": message.MessageFrom, "second_user": message.MessageTo}
	filterr := bson.M{
		"$or": []bson.M{
			filter1,
			filter2,
		},
	}
	messages, err := store.filterOne(filterr)
	if err != nil {
		return err
	}
	messages.Messages = append(messages.Messages, *message)

	_, err = store.messages.UpdateOne(
		context.TODO(),
		bson.M{"_id": messages.Id},
		bson.D{
			{"$set", bson.D{{"messages", messages.Messages}}},
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func NewMessageMongoDBStore(client *mongo.Client) domain.MessageStore {
	messages := client.Database(DATABASE_MSG).Collection(COLLECTION_MSG)
	return &MessageMongoDBStore{
		messages: messages,
	}
}

func (store *MessageMongoDBStore) filter(filter interface{}) ([]*domain.MessageUsers, error) {
	cursor, err := store.messages.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decodeMessage(cursor)
}

func (store *MessageMongoDBStore) filterOne(filter interface{}) (message *domain.MessageUsers, err error) {
	result := store.messages.FindOne(context.TODO(), filter)
	err = result.Decode(&message)
	return
}

func decodeMessage(cursor *mongo.Cursor) (messages []*domain.MessageUsers, err error) {
	for cursor.Next(context.TODO()) {
		var message domain.MessageUsers
		err = cursor.Decode(&message)
		if err != nil {
			return
		}
		messages = append(messages, &message)
	}
	err = cursor.Err()
	return
}
