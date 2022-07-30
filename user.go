package main

import (
	"fmt"
	 "encoding/json"
	"net/http"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"strings"
	"math"

	
)
var db *gorm.DB
var err error



	
func initialMigration() {
	db, err := gorm.Open("sqlite3", "users.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&User{})
}


func AllUsers(w http.ResponseWriter,r *http.Request){
	db, err := gorm.Open("sqlite3", "users.db")
	if err != nil {
		panic("Couldn't able to connect to the database")
	}
	defer db.Close()

	var users []User
	db.Find(&users)
	fmt.Println("{}", users)

	json.NewEncoder(w).Encode(users)
}
func NewUser(w http.ResponseWriter,r *http.Request){
	db, err := gorm.Open("sqlite3", "users.db")
	if err != nil {
		panic("Couldn't able to connect to the database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	id := vars["id"]
	name := vars["name"]
	location := vars["location"]
	gender := vars["gender"]



	fmt.Println(id)
	fmt.Println(name)
	fmt.Println(location)
	fmt.Println(gender)

	db.Create(&User{Id:id, Location: location, Name: name, Gender: gender})
	fmt.Fprintf(w, "New User Successfully Created")
}

func Likes(w http.ResponseWriter,r *http.Request){
	db, err := gorm.Open("sqlite3", "likes.db")
	if err != nil {
		panic("Couldn't able to connect to the database")
	}
	defer db.Close()
	// read our opened jsonFile as a byte array.
    byteValue, _ := ioutil.ReadAll(jsonFile)

    // we initialize our Users array
    var likes Likes
    json.Unmarshal(byteValue, &likes)

    for i := 0; i < len(likes.Likes); i++ {
		for j := i+1; j < len(likes.Likes); j++ {
			if strconv.Itoa(likes.Likes[i].who_likes)==strconv.Itoa(likes.Likes[j].who_is_liked){
				fmt.Println("User ID: " + likes.Likes[i].id)
			}

		}
		
	}

}

func Distance(w http.ResponseWriter,r *http.Request){
	var k int
	var x string 
	// read our opened jsonFile as a byte array.
     byteValue, _ := ioutil.ReadAll(jsonFile)

     // we initialize our Users array
     var users Users

      json.Unmarshal(byteValue, &users)

    for i := 0; i < len(users.Users); i++ {
		if math.abs(strconv.Itoa(users.Users[i].location))-math.abs(strconv.Itoa(users.x.location)<=k){
			json.NewEncoder(w).Encode(users)
		
		}
    
}

}

func Query(w http.ResponseWriter,r *http.Request){
	var s string
	// read our opened jsonFile as a byte array.
     byteValue, _ := ioutil.ReadAll(jsonFile)

     // we initialize our Users array
     var users Users

      json.Unmarshal(byteValue, &users)

    for i := 0; i < len(users.Users); i++ {
		if strings.Contains(users.Users[i].name,s){
			json.NewEncoder(w).Encode(users)
		}

	}

}


func UpdateUser(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "users.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	id := vars["id"]
	location := vars["location"]
	gender := vars["gender"]

	var user User
	db.Where("name = ?", name).Find(&user)

	user.id = id
	user.location = location

	db.Save(&user)
	fmt.Fprintf(w, "Successfully Updated User")
}