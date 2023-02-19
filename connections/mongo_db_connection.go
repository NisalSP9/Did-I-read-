package connections

import (
	"context"
	"github.com/NisalSP9/Did-I-read/commons"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

func Connect() *mongo.Database {
	connectionString := os.Getenv("DID_I_READ_ADMIN_DB_URI")
	dbName := os.Getenv("DID_I_READ_ADMIN_DB_NAME")
	clientOptions := options.Client().ApplyURI(connectionString)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		commons.ErrorLogger.Println("Error while connecting to the DB : " + err.Error())
	}
	if client != nil {
		return client.Database(dbName)
	} else {
		commons.ErrorLogger.Println("Client is nil : ")
		return nil
	}
}

func CheckConnection(client *mongo.Client) *mongo.Client {
	err := client.Ping(context.TODO(), nil)
	if err != nil {
		commons.ErrorLogger.Println("Error while ping to the DB : " + err.Error())
	}
	log.Println("Connected to MongoDB!")
	return client
}

func Disconnect(DB *mongo.Database) {
	err := DB.Client().Disconnect(context.TODO())
	if err != nil {
		commons.ErrorLogger.Println(err.Error())
	}
}
