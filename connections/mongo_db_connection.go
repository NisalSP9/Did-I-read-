package connections

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/NisalSP9/Did-I-read/commons"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func Connect() {
	connectionString := os.Getenv("DID_I_READ_ADMIN_DB_URI")
	dbName := os.Getenv("DID_I_READ_ADMIN_DB_NAME")
	clientOptions := options.Client().ApplyURI(connectionString)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		commons.ErrorLogger.Println("Error while connecting to the DB : " + err.Error())
	}
	DB = client.Database(dbName)
	log.Println("Database connection established")
}

func CheckConnection() bool {
	err := DB.Client().Ping(context.TODO(), nil)
	if err != nil {
		commons.ErrorLogger.Println("Error while ping to the DB : " + err.Error())
		panic(err)
	}
	return true
}

func Disconnect() {
	err := DB.Client().Disconnect(context.TODO())
	if err != nil {
		commons.ErrorLogger.Println(err.Error())
		panic(err)
	}
}
