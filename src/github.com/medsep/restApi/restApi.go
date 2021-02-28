package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	ID        string   `json:"id,omitempty"`
	FirstName string   `json:"firstname, omitempty"`
	LastName  string   `json:"lastname, omitempty"`
	Address   *Address `json:"address, omitempty"`
    Phone     *Phone   `json:"phone, omitempty"`
	Email     string   `json:"email,omitempty"`
}

type Address struct {
	Street string `json:"street, omitempty"`
	City   string `json:"city, omitempty"`
	State  string `json:"state, omitempty"`
}

type Phone struct {
	Work string `json:"Work, omitempty"`
	Home   string `json:"home, omitempty"`
	Mobile  string `json:"mobile, omitempty"`
}

var user []User

func GetUserEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range user {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&User{})
}
func GetUsersEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(user)
}
func CreateUserEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var users User
	_ = json.NewDecoder(req.Body).Decode(&user)
	users.ID = params["id"]
	user = append(user, users)

	json.NewEncoder(w).Encode(user)
}
func DeleteUserEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range user {
		if item.ID == params["id"] {
			user = append(user[:index], user[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(user)
}

func main() {
	router := mux.NewRouter()
	user = append(user, User{ID: "1", FirstName: "Nic", LastName: "Raboy", Address: &Address{Street: "Valerie", City: "Glen", State: "VIC"} Phone: &Phone{Work: 001, Home:002, Mobile: 003}, Email: 123@yahoo.com})
	user = append(user, User{ID: "2", FirstName: "Eric", LastName: "Bannter", Address: &Address{Street: "Valerie", City: "Glen", State: "VIC"}, Email: 33421@gmail.com})
	router.HandleFunc("/user", GetUsersEndpoint).Methods("GET")
	router.HandleFunc("/User/{id}", GetUserEndpoint).Methods("GET")
	router.HandleFunc("/User/{id}", CreateUserEndpoint).Methods("POST")
	router.HandleFunc("/User/{id}", DeleteUserEndpoint).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":3130", router))
}
