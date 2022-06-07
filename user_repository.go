package main

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepo struct {
	coll *mongo.Collection
}

func NewUserRepo(db *mongo.Database) *UserRepo {
	return &UserRepo{
		db.Collection("users"),
	}
}

func (r *UserRepo) Find(UserID int64) *User {
	res := r.coll.FindOne(context.TODO(), bson.D{
		{Key: "user_id", Value: UserID},
	})

	if res.Err() != nil {
		return nil
	}

	var User User
	res.Decode(&User)

	return &User
}

func (r *UserRepo) FindOrInsert(UserID int64) (bool, *User) {
	var (
		user  = r.Find(UserID)
		insrt = false
	)

	if user == nil {
		res, err := r.coll.InsertOne(context.TODO(), User{
			UserID: UserID,
		})
		CheckIfError(err)
		insrt = true

		if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
			user = &User{ID: &oid, UserID: UserID}
		}
	}

	return insrt, user
}

func (r *UserRepo) Update(user *User) {
	_, err := r.coll.ReplaceOne(
		context.TODO(),
		bson.M{"_id": user.ID},
		user,
	)
	CheckIfError(err)
}
