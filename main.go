package main

import (
	"gopkg.in/mgo.v2"

	"github.com/jacob-ebey/user-server/controllers"
	"github.com/jacob-ebey/user-server/core"
)

func getSession() *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial("mongodb://localhost")

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}
	return s
}

func main() {
	// Get a UserController instance
	uc := controllers.NewUserController(getSession())

	// Create a new Server
	server := core.NewServer("localhost:8080")

	// Get a user resource
	server.GET("/user/:key/:id", uc.GetUser)

	// Create a user resource
	server.POST("/user/:key", uc.CreateUser)

	// Remvoe a user
	server.DELETE("/user/:key/:id", uc.RemoveUser)

	server.Run()

	// Create user
	// curl -XPOST -H 'Content-Type: application/json' -d '{"name": "Jacob Ebey", "gender": "male", "age": 21, "email": "jacob.ebey@live.com"}' http://localhost:8080/user

	// Get user
	// curl http://localhost:8080/user/5586593e4a75d52b22000001
}
