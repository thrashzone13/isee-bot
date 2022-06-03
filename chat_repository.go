package main

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Chat struct {
	ID     *primitive.ObjectID `json:"_id" bson:"_id"`
	ChatID int64               `json:"chat_id" bson:"chat_id"`
}

type ChatRepo struct {
	coll *mongo.Collection
}

func NewChatRepo(db *mongo.Database) *ChatRepo {
	return &ChatRepo{
		db.Collection("chats"),
	}
}

func (r *ChatRepo) Find(id int64) *Chat {
	res, err := r.coll.Distinct(context.TODO(), "chat_id", bson.D{
		{Key: "id", Value: id},
	})
	LogIfError(err)

	var chat Chat
	UnMarshalDBResult(res, &chat)

	return &chat
}

func (r *ChatRepo) Create(chatID int64) {
	_, err := r.coll.InsertOne(context.TODO(), bson.D{
		{"chat_id", chatID},
	})
	LogIfError(err)
}
