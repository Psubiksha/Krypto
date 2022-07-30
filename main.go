package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"

	"github.com/gorilla/mux"
	
)
{
	"handler":"email",
	"notification":"connection-failed",
	"targets":[
		"email@somecompany.com",
		"person@somecompany.com"
	]
}
{
    "status": 200,
    "type": "notification_types",
    "notification_types": [
        {
            "id": "connection-failed",
            "name": "Connection failed",
            "description": "Receive a notification whenever a connection fails to open."
        },
        {
            "id": "product-upgrade-available",
            "name": "Upgrade available",
            "description": "Receive a notification when a new version of the software is available."
        },
        {
            "id": "license-expired",
            "name": "License has expired",
            "description": "Receive a notification when your license has expired."
        },
        {
            "id": "license-expiration-warning",
            "name": "License expiry approaching",
            "description": "Receive a notification when your license will expire soon."
        },
        {
            "id": "job-failed",
            "name": "Job failed",
            "description": "Receive a notification whenever a job fails."
        },
        {
            "id": "job-completed",
            "name": "Job completed",
            "description": "Receive a notification whenever a job meets its stop policy."
        }
    ]
}
func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/users", AllUsers).Methods("GET")
	myRouter.HandleFunc("/user/{id}/{name}", NewUser).Methods("POST")
	myRouter.HandleFunc("/likes/{id}", Likes).Methods("GET")
	myRouter.HandleFunc("/likes/{id}", Distance).Methods("GET")
	myRouter.HandleFunc("/user/{id}/{name}", Query).Methods("GET")
	myRouter.HandleFunc("/user/{id}/{name}", UpdateUser).Methods("PUT")
	
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}