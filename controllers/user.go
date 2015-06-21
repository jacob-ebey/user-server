package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/julienschmidt/httprouter"

	"github.com/jacob-ebey/user-server/models"
	"github.com/jacob-ebey/user-server/validate"
)

type (
	// Represents the controller for operating on the User resource
	UserControler struct {
		session *mgo.Session
	}
)

func NewUserController(s *mgo.Session) *UserControler {
	return &UserControler{s}
}

// CreateUser creates a user resource
func (uc UserControler) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Create a user to be populated from the body
	u := models.User{}

	// Populate the user data
	json.NewDecoder(r.Body).Decode(&u)

	emailCheck, errMsg := validate.Email(u.Email)

	if !emailCheck {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		fmt.Fprintf(w, "{ err: \"%s\" }", errMsg)
		return
	}

	tmpUser := models.User{}

	// Fetch user
	if err := uc.session.DB("user-server").C("users").Find(bson.M{"email": u.Email}).One(&tmpUser); err == nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "{ err: \"%s\" }", "Email is already registered.")
		return
	}

	// Add an Id
	u.Id = bson.NewObjectId()

	// Write the user to mongo
	uc.session.DB("user-server").C("users").Insert(u)

	// Marshal the user to JSON
	jr, _ := json.Marshal(u)

	// Write content-type, statuscode and payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", jr)
}

// GetUser retrieves and individual user resource
func (uc UserControler) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id := p.ByName("id")

	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	// Grab id
	oid := bson.ObjectIdHex(id)

	// Create a user
	u := models.User{}

	// Fetch user
	if err := uc.session.DB("user-server").C("users").Find(bson.M{"id": oid}).One(&u); err != nil {
		w.WriteHeader(404)
		return
	}

	jr, _ := json.Marshal(u)

	// Write content-type, statuscode and payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", jr)
}

// RemoveUser removes a user resource
func (uc UserControler) RemoveUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id := p.ByName("id")

	// Grab id
	oid := bson.ObjectId(id)

	// Remove user
	if err := uc.session.DB("user-server").C("users").RemoveId(oid); err != nil {
		w.WriteHeader(404)
		return
	}

	// Write status
	w.WriteHeader(200)
}
