package main

import (
	"fmt"
	"mongo-golang/controllers"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user/", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)

	fmt.Printf("Running at 3000")
	http.ListenAndServe("localhost:3000", nil)
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://root:123456@localhost:8081/mongo-golang")
	if err != nil {
		panic(err)
	}
	return s
}
