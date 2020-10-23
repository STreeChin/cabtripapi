package model

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

const (
	USERNAME   = "root"
	PASSWORD   = "root"
	NETWORK    = "tcp"
	SERVER     = "@cluster0-6kp2k.mongodb.net"
	PORT       = 3306
	DATABASE   = "cabtrip"
	COLLECTION = "cab_trip_data"
)
/* collcetion struct
type CabTrip struct {
	medallion          string
	hack_license       string
	vendor_id          string
	rate_code          int
	store_and_fwd_flag string
	pickup_datetime    time.Time
	dropoff_datetime   time.Time
	passenger_count    int
	trip_time_in_secs  int
	trip_distance      float64
	pickup_longitude   float64
	pickup_latitude    float64
	dropoff_longitude  float64
	dropoff_latitude   float64
}
*/
type MongoDBClient struct {
	Client *mongo.Client
}

var MgDB *MongoDBClient

func GetMongodbClient() *MongoDBClient {
	return MgDB
}
func MongoConnectionClose() {
	err := MgDB.Client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}

func MongoDBConnect() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//"mongodb+srv://root:root@cluster0-6kp2k.mongodb.net/test?retryWrites=true&w=majority"
	url := "mongodb+srv://" + USERNAME + ":" + PASSWORD + SERVER + "/test?retryWrites=true&w=majority"
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal(err)
	}
	MgDB = new(MongoDBClient)
	MgDB.Client = client
	fmt.Println("Connected to MongoDB!")
}
func (c *MongoDBClient) GetCount(ctx context.Context, id, date string) int64 {
	timeLayout := "2006-01-02 15:04:05"
	// if we need the time location, use time.LoadLocation("Local"), ParseInLocation
	startOfDay, _ := time.Parse(timeLayout, date+" 00:00:00")
	endOfDay, _ := time.Parse(timeLayout, date+" 23:59:59")

	collection := c.Client.Database(DATABASE).Collection(COLLECTION)
	filter := bson.M{"medallion": id, "pickup_datetime": bson.M{"$gte": startOfDay, "$lte": endOfDay}}
	//get the count
	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	return count
}
