package main

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/swaggo/echo-swagger/example/docs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
)

var teams = []team{
	{1, "3CRABS"},
	{2, "CRABS"},
	{3, "C"},
}

type team struct {
	Id   int    `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}

var collection *mongo.Collection
var ctx = context.TODO()

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func main() {
	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/teams", func(c echo.Context) error {
		return c.JSON(http.StatusOK, teams)
	})

	err := createTeam(&team{Id: 1, Name: "CRAB"})
	if err != nil {
		log.Fatal(err)
	}
	teams, err := getTeams()
	if err != nil {
		log.Fatal(err)
	}
	for _, t := range teams {
		fmt.Println(t.Name)
	}

	e.Logger.Fatal(e.Start(":1323"))
}

func init() {
	credential := options.Credential{
		Username: "root",
		Password: "qwerty",
	}
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/").SetAuth(credential)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	collection = client.Database("hackpoint").Collection("teams")
}

func createTeam(team *team) error {
	_, err := collection.InsertOne(ctx, team)
	return err
}

func getTeams() ([]*team, error) {
	// passing bson.D{{}} matches all documents in the collection
	filter := bson.D{{}}
	return filterTasks(filter)
}

func filterTasks(filter interface{}) ([]*team, error) {
	// A slice of tasks for storing the decoded documents
	var teams []*team

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return teams, err
	}

	for cur.Next(ctx) {
		var t team
		err := cur.Decode(&t)
		if err != nil {
			return teams, err
		}

		teams = append(teams, &t)
	}

	if err := cur.Err(); err != nil {
		return teams, err
	}

	// once exhausted, close the cursor
	cur.Close(ctx)

	if len(teams) == 0 {
		return teams, mongo.ErrNoDocuments
	}

	return teams, nil
}
