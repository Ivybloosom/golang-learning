// 功能说明
/*
* @Author: Ivy
* @Date: 2022/7/20 10:34 PM
 */
package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//const MONGODB_URL = "mongodb+srv://Ivy:<password>@cluster0.ac69p6c.mongodb.net/?retryWrites=true&w=majority"

func main() {
	// Connect to your MongoDB
	uri := "mongodb://127.0.0.1:27017"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// Get the database and collection.
	db := client.Database("test")
	coll := db.Collection("user")

	var result bson.M
	err = coll.FindOne(context.TODO(), bson.D{{"name", "Ivy"}}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return
		}
		panic(err)
	}
	fmt.Println(result)

	//// Insert documents
	//docs := []interface{}{
	//	bson.D{{"name", "Halley's Comet"}, {"officialName", "1P/Halley"}, {"orbitalPeriod", 75}, {"radius", 3.4175}, {"mass", 2.2e14}},
	//	bson.D{{"name", "Wild2"}, {"officialName", "81P/Wild"}, {"orbitalPeriod", 6.41}, {"radius", 1.5534}, {"mass", 2.3e13}},
	//	bson.D{{"name", "Comet Hyakutake"}, {"officialName", "C/1996 B2"}, {"orbitalPeriod", 17000}, {"radius", 0.77671}, {"mass", 8.8e12}},
	//}
	//result, err := coll.InsertMany(context.TODO(), docs)
	//if err != nil {
	//	panic(err)
	//}
	//
	//for _, id := range result.InsertedIDs {
	//	fmt.Printf("\t%s\n", id)
	//}

	//// Update documents
	//filter := bson.D{{}}
	//update := bson.D{{"$mul", bson.D{{"Radius", 1.60934}}}}
	//result, err = coll.UpdateMany(context.TODO(), filter, update)
	//if err != nil {
	//	panic(err)
	//}
	////fmt.Printf("Number of documents updated: %d", result.ModifiedCount)
	//
	//// Delete Documents
	//filter := bson.D{
	//	{"$and",
	//		bson.A{
	//			bson.D{{"orbitalPeriod", bson.D{{"$gt", 5}}}},
	//			bson.D{{"orbitalPeriod", bson.D{{"$lt", 85}}}},
	//		},
	//	},
	//}
	//result, err = coll.DeleteMany(context.TODO(), filter)
	//if err != nil {
	//	panic(err)
	//}
	//// amount deleted code goes here
	////fmt.Printf("Number of documents deleted: %d", result.DeletedCount)
	//
	//// Find documents
	//cursor, err := coll.Find(context.TODO(), bson.D{})
	//if err != nil {
	//	panic(err)
	//}
	//for cursor.Next(context.TODO()) {
	//	var result bson.M
	//	if err := cursor.Decode(&result); err != nil {
	//		panic(err)
	//	}
	//	fmt.Println(result)
	//}
	//if err := cursor.Err(); err != nil {
	//	panic(err)
	//}
	//
	//// Find document with filter
	//filter = bson.D{
	//	{"$and",
	//		bson.A{
	//			bson.D{{"hasRings", false}},
	//			bson.D{{"mainAtmosphere", "Ar"}},
	//		},
	//	},
	//}
	//cursor, err = coll.Find(context.TODO(), filter)
	//if err != nil {
	//	panic(err)
	//}
	//for cursor.Next(context.TODO()) {
	//	var result bson.M
	//	if err := cursor.Decode(&result); err != nil {
	//		panic(err)
	//	}
	//	fmt.Println(result)
	//}
	//if err := cursor.Err(); err != nil {
	//	panic(err)
	//}
}
