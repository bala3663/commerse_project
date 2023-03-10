package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const URL = "mongodb://localhost:27017"

type Books struct {
	BookId     int
	BookName   string
	AuthorName string
	year       int
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//connect to db
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(URL))
	if err != nil {
		panic(err)
	}
	//disconnect object
	defer func() {
		err := client.Disconnect(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}()

	//use the client to print all databases list
	// resultSet, _ := client.ListDatabaseNames(ctx, bson.D{})
	// for _, db := range resultSet {
	// 	fmt.Println(db)
	// }

	// //use the client to acess databse and collection

	collections := client.Database("First_Database").Collection("First COllection")
	cursor, _ := collections.Find(ctx, bson.D{})
	for cursor.Next(ctx) {
		var results bson.D
		err := cursor.Decode(&results)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(results)

	}

	//INSERT VALUES IN DATABASES
	// books := Books{

	// 	BookId:     7,
	// 	BookName:   "HARRY POTTER-deathly hallows",
	// 	AuthorName: "J K ROLINGS",
	// 	year:       2007,
	// }

	// docs := bson.D{
	// 	{Key: "BookId", Value: books.BookId},
	// 	{Key: "BookName", Value: books.BookName},
	// 	{Key: "AuthorName", Value: books.AuthorName},
	// 	{Key: "year", Value: books.year},
	// }
	// resultSet, _ := collections.InsertOne(ctx, docs)
	// fmt.Println("Inserted record with ID: %v", resultSet.InsertedID)

	//UPDATE RECORDS
	// id, _ := primitive.ObjectIDFromHex("6406df0cb25aea28e47861a3")
	// updated, _ := collections.UpdateOne(
	// 	ctx,
	// 	bson.M{"_id": id},
	// 	bson.D{
	// 		{Key: "$set", Value: bson.M{"bookname": "WINGS OF FIRE"}},
	// 	})
	// fmt.Println("Updates %v documents!", updated.ModifiedCount)

	//DELETE DOCUMENT
	// dele, err := collections.DeleteOne(ctx, bson.M{"BookId": 1})
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("Deleted %v documents", dele.DeletedCount)
}
