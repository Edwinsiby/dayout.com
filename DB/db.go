package DB

import (
	"context"
	"fmt"
	"log"
	"main/models"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
	"gorm.io/gorm"
)

// psql
var Db *gorm.DB

// mongodb
var Ctx context.Context
var Cancel context.CancelFunc
var UserList []models.User
var err error

var Client *mongo.Client
var collection *mongo.Collection

func init() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	Client = client
	collection = client.Database("sample").Collection("user")
}

func InsertOne(dataBase, col string, doc models.User) (*mongo.InsertOneResult, error) {

	result, err := collection.InsertOne(Ctx, doc)
	if err != nil {
		fmt.Println("error in mongo")
		fmt.Println(err)
	}
	return result, nil
}

func FindUser(query bson.M, field bson.M, dataBase, col string) (*models.User, error) {
	var user models.User
	var Client *mongo.Client
	collection := Client.Database(dataBase).Collection(col)
	err := collection.FindOne(context.Background(), query, options.FindOne().SetProjection(field)).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}
		return nil, err
	}
	return &user, nil
}

func FindAllUsers(query bson.M, field bson.M, databse, col string) ([]models.User, error) {
	users := []models.User{}
	var Client *mongo.Client
	collection := Client.Database(databse).Collection(col)
	user, err := collection.Find(context.Background(), query, options.Find().SetProjection(field))
	if err != nil {
		return nil, err
	}
	defer user.Close(context.Background())
	var userData models.User
	for user.Next(context.Background()) {
		err := user.Decode(&userData)
		if err != nil {
			return nil, err
		}
		users = append(users, userData)
	}
	return users, nil
}

func DeleteUser(query bson.M, database, col string) error {
	var Client *mongo.Client
	collection := Client.Database(database).Collection(col)
	_, err := collection.DeleteOne(context.Background(), query)
	if err != nil {
		return err
	}
	return nil
}
