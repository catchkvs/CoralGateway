package dbconnector
import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func Init() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	handleError(err)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := client.Database("testing").Collection("numbers")
	log.Println("Collection: ", collection)
	res, err := collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
	log.Println("Response: ", res)
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil { log.Fatal(err) }
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil { log.Fatal(err) }
		log.Println(result)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
}

func handleError(err error) {
	if err != nil {
		log.Println("handling error::::", err)

	}
}
