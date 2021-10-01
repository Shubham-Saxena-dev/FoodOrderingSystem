package main

import (
	"FoodOrderingSystem/controllers"
	"FoodOrderingSystem/repositories"
	"FoodOrderingSystem/routes"
	"FoodOrderingSystem/services"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const (
	Database   = "goMongo"
	Collection = "onlineFood"
	MongoDbUrl = "mongodb://localhost:27017"
)

var (
	collection *mongo.Collection
	repo       repositories.Repository
	serv       services.Service
	controller controllers.Controller
	ctx        context.Context
)

func main() {
	log.Info("Hi, this is food ordering system")
	initDatabase()
	createServer()
}

func createServer() {

	server := gin.Default()
	initializeLayers()

	routes.NewRoutesHandler(server, controller).RegisterHandlers()
	go heartbeat()
	err := server.Run()

	if err != nil {
		failOnError(err, "Unable to start server")
	}
}

func initializeLayers() {
	repo = repositories.NewMongoRepository(collection, ctx)
	serv = services.NewService(repo)
	controller = controllers.NewController(serv)
}

func initDatabase() {
	log.Info("Connecting to MongoDb...")
	clientOptions := options.Client().ApplyURI(MongoDbUrl)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		failOnError(err, "Unable to connect to MongoDb")
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		failOnError(err, "MongoDb Connection is not responding")
	}

	log.Info("Connected to MongoDB!")

	collection = client.Database(Database).Collection(Collection)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Error(msg)
		panic(err)
	}
}


func heartbeat() {
	for {
		timer := time.After(time.Second * 10)
		<-timer
		fmt.Println("heartbeat !")
	}
}
