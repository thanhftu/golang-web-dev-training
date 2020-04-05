package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/thanhftu/golang-web-dev/042_mongodb/05_mongodb/01_update-user-controller/controllers"
	"gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()
	// Get a UserController instance
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}

func getSession() *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial("mongodb+srv://thanhdhnt:Donga130@contactkeeper-uo0tn.mongodb.net/golang?retryWrites=true&w=majority")

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}
	return s
}
