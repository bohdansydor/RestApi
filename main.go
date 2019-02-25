package main

import (
	_ "RestApi_v2.0/routers"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)
type User struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Surname string `json:"surname"`
	Age int `json:"age"`
}
type Users struct {
	Users []User `json:"users"`
}



func parseJson () Users{
	jsonFile, err := os.Open("static/list.json")
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var users Users
	json.Unmarshal(byteValue, &users)

	return users

}


func getUsers(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Convert-Tipe","application/json")
	json.NewEncoder(w).Encode(parseJson())

}

func getUser(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Convert-Tipe","application/json")
	params := mux.Vars(r)
	for _, item := range parseJson().Users {
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func addUser(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Convert-Tipe","application/json")
	var newUser User

	var Users = parseJson()

	json.NewDecoder(r.Body).Decode(&newUser)
	newUser.ID = strconv.Itoa(len(Users.Users) + 1)
	Users.Users = append(Users.Users,newUser)
	json.NewEncoder(w).Encode(Users.Users)      //   масив чи json Users



}

func editUser(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Convert-Tipe","application/json")
	params := mux.Vars(r)

	var Users = parseJson()

	for index, item := range parseJson().Users{
		if item.ID == params["id"] {

			var user User

			json.NewDecoder(r.Body).Decode(&user)

			user.ID = params["id"]
			Users.Users[index] = user

			json.NewEncoder(w).Encode(Users.Users)
			return
		}
	}
	json.NewEncoder(w).Encode(Users.Users)
}

func deleteUser(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Convert-Tipe","application/json")
	params := mux.Vars(r)
	var Users = parseJson()
	for index, item := range parseJson().Users{
		if item.ID == params["id"] {

			Users.Users = append(parseJson().Users[:index], parseJson().Users[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(Users.Users)

}

func handleRequest()  {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/users", getUsers).Methods("GET")
	myRouter.HandleFunc("/users/{id}", getUser).Methods("GET")
	myRouter.HandleFunc("/users", addUser).Methods("POST")
	myRouter.HandleFunc("/users/{id}",editUser).Methods("PUT")
	myRouter.HandleFunc("/users/{id}",deleteUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":5000", myRouter))

}

func main() {
	handleRequest()
}

