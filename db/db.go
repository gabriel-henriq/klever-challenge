package db

import (
	"context"
	"fmt"

	"github.com/gabriel-henriq/klever-challenge/api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func Connect() (*mongo.Collection, context.Context) {
	// Set client options
	mongoCtx := context.Background()

	// Create a new client and connect to the server
	client, err := mongo.Connect(mongoCtx, options.Client().SetDirect(true).ApplyURI("mongodb://mongodb:27017/?connect=direct"))
	if err != nil {
		panic(err)
	}

	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected and pinged.")

	// Get a handle to the collection
	bookDb := client.Database("klever-challenge").Collection("books")

	// Populate the database with some books
	populateWithBooks(bookDb, mongoCtx)

	return bookDb, mongoCtx
}

func populateWithBooks(bookDb *mongo.Collection, mongoCtx context.Context) {
	book, err := bookDb.FindOne(mongoCtx, bson.M{"title": "Clean Code"}).DecodeBytes()
	if book == nil && err != nil {
		fmt.Println("Books not found, inserting...")

		book1 := models.Book{
			Title:  "Clean Code",
			Author: "Robert C. Martin",
			Likes:  0,
		}
		book2 := models.Book{
			Title:  "Go Programming Language",
			Author: "John Doe",
			Likes:  0,
		}
		_, err := bookDb.InsertMany(mongoCtx, []interface{}{book1, book2})
		if err != nil {
			panic(err)
		}
	}
}
