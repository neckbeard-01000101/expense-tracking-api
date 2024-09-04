package controllers

import (
	"context"
	"encoding/json"
	"expense-tracking/config"
	"expense-tracking/models"
	"expense-tracking/utils"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
	"time"
)

const (
	DB_NAME         string = "expense-tracking"
	COLLECTION_NAME string = "users"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("Error decoding request body in register: %v", err)
		http.Error(w, `{"error": "Error decoding the request body"}`, http.StatusInternalServerError)
		return
	}
	isValidUser, err := user.ValidUser(DB_NAME, COLLECTION_NAME)
	if !isValidUser {
		log.Println(err)
		http.Error(w, `{"error": username or email are taken}`, http.StatusBadRequest)
		return
	}
	if !user.ValidEmail() {
		log.Println("email is invalid")
		http.Error(w, `{"error": email in invalid}`, http.StatusBadRequest)
		return
	}
	hashedPassword := utils.HashString(user.Password)
	user.Password = hashedPassword
	collection := config.Client.Database(DB_NAME).Collection(COLLECTION_NAME)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := collection.InsertOne(ctx, bson.M{
		"username": user.UserName,
		"email":    user.Email,
		"password": user.Password,
	})
	if err != nil {
		log.Printf("error registering user: %v", err)
		http.Error(w, `{"error": "error registering the user"}`, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"message": "User has been registered successfully",
		"id":      result.InsertedID,
	}
	json.NewEncoder(w).Encode(response)
	log.Printf("User registered successfully with ID: %v", result.InsertedID)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var loginData LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginData)
	if err != nil {
		log.Fatal("error decoding request data in login request")
	}
	loginData.Password = utils.HashString(loginData.Password)
	fmt.Println(loginData)
	collection := config.Client.Database(DB_NAME).Collection(COLLECTION_NAME)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.M{
		"username": loginData.Username,
		"password": loginData.Password,
	}
	err = collection.FindOne(ctx, filter).Decode(&loginData)
	if err != nil {
		log.Printf("error loging in: %v", err)
		http.Error(w, `{"error": "username or password are incorrect"}`, http.StatusUnauthorized)
		return
	}
	w.WriteHeader(http.StatusOK)
}
