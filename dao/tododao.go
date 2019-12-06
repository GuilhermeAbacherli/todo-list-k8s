package dao

import (
	"context"
	"log"

	"github.com/GuilhermeAbacherli/todolistgo/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// SelectManyTodos return all of the TODOs from mongodb
func SelectManyTodos(client *mongo.Client, filter bson.M) []*entity.Todo {
	var todolist []*entity.Todo
	collection := client.Database("todolistgo").Collection("todolist")
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Println("Error on finding all the documents: ", err)
	}
	for cur.Next(context.TODO()) {
		var todo entity.Todo
		err = cur.Decode(&todo)
		if err != nil {
			log.Println("Error on decoding the document: ", err)
		}
		todolist = append(todolist, &todo)
	}
	return todolist
}

// SelectOneTodo return a TODO from mongodb
func SelectOneTodo(client *mongo.Client, filter bson.M) entity.Todo {
	var todo entity.Todo
	collection := client.Database("todolistgo").Collection("todolist")
	documentReturned := collection.FindOne(context.TODO(), filter)
	documentReturned.Decode(&todo)
	return todo
}

// InsertOneTodo insert a new todo in mongodb
func InsertOneTodo(client *mongo.Client, todo entity.Todo) entity.Todo {
	collection := client.Database("todolistgo").Collection("todolist")
	options := options.FindOne()
	// Sort by 'ID' field descending
	options.SetSort(bson.M{"id": -1})
	var lastTodo entity.Todo
	documentReturned := collection.FindOne(context.TODO(), bson.M{}, options)
	documentReturned.Decode(&lastTodo)
	todo.ID = lastTodo.ID + 1
	_, err := collection.InsertOne(context.TODO(), todo)
	if err != nil {
		log.Println("Error on inserting one document: ", err)
	}
	return todo
}

// UpdateOneTodo updates one existing TODO in mongodb
func UpdateOneTodo(client *mongo.Client, newTodo entity.Todo, filter bson.M) (updatedTodo entity.Todo) {
	collection := client.Database("todolistgo").Collection("todolist")
	update := bson.M{"$set": newTodo}
	err := collection.FindOneAndUpdate(context.TODO(), filter, update,
		options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&updatedTodo)
	if err != nil {
		log.Println("Error on updating one document: ", err)
	}
	return
}

// DeleteOneTodo deletes one existing TODO in mongodb
func DeleteOneTodo(client *mongo.Client, filter bson.M) (deletedTodo entity.Todo) {
	collection := client.Database("todolistgo").Collection("todolist")
	err := collection.FindOneAndDelete(context.TODO(), filter,
		options.FindOneAndDelete()).Decode(&deletedTodo)
	if err != nil {
		log.Println("Error on deleting one document: ", err)
	}
	return
}

// DeleteAllTodos deletes all existing TODOs in mongodb
func DeleteAllTodos(client *mongo.Client) int64 {
	collection := client.Database("todolistgo").Collection("todolist")
	deleteResult, err := collection.DeleteMany(context.TODO(), bson.D{})
	if err != nil {
		log.Println("Error on deleting all documents: ", err)
	}
	return deleteResult.DeletedCount
}
