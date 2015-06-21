package main

import (
	"log"
	"net/http"

	"gopkg.in/mgo.v2"

	"github.com/julienschmidt/httprouter"

	"github.com/jacob-ebey/user-server/controllers"
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
	// Instantiate a new router
	r := httprouter.New()

	// Get a UserController instance
	uc := controllers.NewUserController(getSession())

	// Get a user resource
	r.GET("/user/:key/:id", uc.GetUser)

	// Create a user resource
	r.POST("/user/:key", uc.CreateUser)

	// Remvoe a user
	r.DELETE("/user/:key/:id", uc.RemoveUser)

	// Fire up the server
	log.Fatal(http.ListenAndServe("localhost:8080", r))

	// Create user
	// curl -XPOST -H 'Content-Type: application/json' -d '{"name": "Jacob Ebey", "gender": "male", "age": 21, "email": "jacob.ebey@live.com"}' http://localhost:8080/user

	// Get user
	// curl http://localhost:8080/user/5586593e4a75d52b22000001
}
