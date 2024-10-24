package main

import (
	"fmt"
	"go_mongo/controllers"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()
	session, err := getSession()
	if err != nil {
		fmt.Println("Could not connect to MongoDB:", err)
		return
	}
	defer session.Close()

	uc := controllers.NewUserController(session)
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.Deleteuser)

	fmt.Println("Server running at localhost:9000")
	err = http.ListenAndServe("localhost:9000", r)

	if err != nil {
		fmt.Println("Failed to start server:", err)
	}
}

func getSession() (*mgo.Session, error) {
	s, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		return nil, fmt.Errorf("failed to dial MongoDB: %v", err)
	}

	// Verify the connection by pinging MongoDB
	err = s.Ping()
	if err != nil {
		return nil, fmt.Errorf("unable to ping MongoDB: %v", err)
	}

	fmt.Println("Connected to MongoDB")
	return s, nil
}
