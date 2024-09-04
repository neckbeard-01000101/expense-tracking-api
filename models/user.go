package models

import (
	"context"
	"errors"
	"expense-tracking/config"
	emailverifier "github.com/AfterShip/email-verifier"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserName string             `bson:"username" json:"username"`
	Password string             `bson:"password" json:"password"`
	Email    string             `bson:"email" json:"email"`
}

func (u *User) ValidUser(dbName, collectionName string) (bool, error) {
	collection := config.Client.Database(dbName).Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check if either username or email already exists
	filter := bson.M{
		"$or": []bson.M{
			{"username": u.UserName},
			{"email": u.Email},
		},
	}

	var existingUser User
	err := collection.FindOne(ctx, filter).Decode(&existingUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return true, nil
		}
		return false, err
	}

	if existingUser.UserName == u.UserName {
		return false, errors.New("username is already taken")
	}
	if existingUser.Email == u.Email {
		return false, errors.New("email is already taken")
	}

	return true, nil
}

func (u *User) ValidEmail() bool {
	verifier := emailverifier.NewVerifier()
	ret, err := verifier.Verify(u.Email)
	if err != nil {
		return false
	}
	if !ret.Syntax.Valid {
		return false
	}
	return true
}
