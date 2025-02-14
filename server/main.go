package main

import (
	"context" //**
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"           //**
	"go.mongodb.org/mongo-driver/bson/primitive" //**
	"go.mongodb.org/mongo-driver/mongo"          //**
	"go.mongodb.org/mongo-driver/mongo/options"  //**
)

// Config
var mongoUri string = "mongodb://localhost:27017"
var mongoDbName string = "ars_app_db"

// Database variables
var mongoClient *mongo.Client
var flightCollection *mongo.Collection

type Flight struct {
	Id          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Number      string             `json:"number" bson:"number"`
	AirlineName string             `json:"airline_name" bson:"airline_name"`
	Source      string             `json:"source" bson:"source"`
	Destination string             `json:"destination" bson:"destination"`
	Capacity    int                `json:"capacity" bson:"capacity"`
	Price       float64            `json:"price" bson:"price"`
}

// mongo connect
func connectToMongo() {
	var err error
	mongoClient, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoUri))
	if err != nil {
		fmt.Println("Mongo DB Connection Error!" + err.Error())
		return
	}
	name := "flights"
	flightCollection = mongoClient.Database(mongoDbName).Collection(name)
}

// apis
func createFlight(c *gin.Context) {
	var flight Flight
	if err := c.BindJSON(&flight); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server Error." + err.Error()})
		return
	}
	//
	flight.Id = primitive.NewObjectID()
	_, err := flightCollection.InsertOne(context.TODO(), flight)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server Error.\n" + err.Error()})
		return
	}
	//
	c.JSON(http.StatusCreated,
		gin.H{"message": "Flight Created Successfully", "flight": flight})
}

func readAllFlights(c *gin.Context) {
	//
	var flights []Flight
	cursor, err := flightCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server Error.\n" + err.Error()})
		return
	}
	defer cursor.Close(context.TODO())
	//
	flights = []Flight{}
	err = cursor.All(context.TODO(), &flights)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server Error.\n" + err.Error()})
		return
	}
	//
	c.JSON(http.StatusOK, flights)
}

func readFlightById(c *gin.Context) {
	id := c.Param("id")
	//
	flightId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID.\n" + err.Error()})
		return
	}
	//
	var flight Flight
	err = flightCollection.FindOne(context.TODO(), bson.M{"_id": flightId}).Decode(&flight)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Flight Not Found."})
		return
	}
	//
	c.JSON(http.StatusOK, flight)
}

func updateFlight(c *gin.Context) {
	id := c.Param("id")
	flightId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID.\n" + err.Error()})
		return
	}
	//
	var oldFlight Flight
	err = flightCollection.FindOne(context.TODO(), bson.M{"_id": flightId}).Decode(&oldFlight)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Flight Not Found."})
		return
	}
	//
	var jbodyFlight Flight
	err = c.BindJSON(&jbodyFlight)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server Error." + err.Error()})
		return
	}
	oldFlight.Price = jbodyFlight.Price
	//
	_, err = flightCollection.UpdateOne(context.TODO(),
		bson.M{"_id": flightId},
		bson.M{"$set": bson.M{"price": jbodyFlight.Price}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server Error." + err.Error()})
		return
	}
	//response
	c.JSON(http.StatusOK, gin.H{"message": "Flight Updated Successfully", "flight": oldFlight})
}

func deleteFlight(c *gin.Context) {
	id := c.Param("id")
	flightId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID.\n" + err.Error()})
		return
	}
	//
	var oldFlight Flight
	err = flightCollection.FindOne(context.TODO(), bson.M{"_id": flightId}).Decode(&oldFlight)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Flight Not Found."})
		return
	}
	//delete
	_, err = flightCollection.DeleteOne(context.TODO(), bson.M{"_id": flightId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server Error." + err.Error()})
		return
	}
	//response
	c.JSON(http.StatusOK, gin.H{"message": "Flight deleted successfully."})
}

func main() {
	//
	connectToMongo()
	//router
	r := gin.Default()
	//cors
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // React frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	//routes
	r.POST("/flights", createFlight)
	r.GET("/flights", readAllFlights)
	r.GET("/flights/:id", readFlightById)
	r.PUT("/flights/:id", updateFlight)
	r.DELETE("/flights/:id", deleteFlight)
	//server
	r.Run(":8080")
}