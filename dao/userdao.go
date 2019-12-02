package dao

import (
	"context"
	"log"

	"github.com/GuilhermeAbacherli/todolistgo/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// SelectOneUser returns a single user registered in mongodb
func SelectOneUser(client *mongo.Client, filter bson.M, user *entity.User) {
	collection := client.Database("todolistgo").Collection("users")
	documentReturned := collection.FindOne(context.TODO(), filter)
	err := documentReturned.Decode(&user)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return
		}
	}
	return
}

// InsertOneUser insert a new user in mongodb
func InsertOneUser(client *mongo.Client, user entity.User) interface{} {
	collection := client.Database("todolistgo").Collection("users")
	insertResult, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Println("Error on inserting one user: ", err)
	}
	return insertResult.InsertedID
}
