package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Person struct {
	Name string
	Age  string
}

var user []Person

func handleGetPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// if r.Method == "GET" {
	// 	name := r.URL.Query()["name"][0]
	// 	for _, structs := range user {
	// 		if structs.Name == name {
	// 			err := json.NewEncoder(w).Encode(&structs)
	// 			if err != nil {
	// 				log.Fatal(err)
	// 			}
	// 		}
	// 	}
	// } else {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	w.Write([]byte("Bad Request"))
	// }
	if r.Method == "GET" {
		json.NewEncoder(w).Encode(&user)
	}
}

func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "PUT" {
		var human Person
		err := json.NewDecoder(r.Body).Decode(&human)
		if err != nil {
			log.Fatal(err)
		}

		for index, structs := range user {
			if structs.Name == human.Name {
				user = append(user[:index], user[index+1:]...)
			}
		}
		user = append(user, human)
		err = json.NewEncoder(w).Encode(&human)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
	}
}

func handleAddPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "POST" {
		w.WriteHeader(http.StatusOK)
		var human Person
		err := json.NewDecoder(r.Body).Decode(&human)
		if err != nil {
			log.Fatal(err)
		}
		user = append(user, human)
		err = json.NewEncoder(w).Encode(&human)
		if err != nil {
			log.Fatal(err)
		}

	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
	}
}

func handleDeletePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "DELETE" {
		name := r.URL.Query()["name"][0]
		indexChoice := 0
		for index, structs := range user {
			if structs.Name == name {
				indexChoice = index
			}
		}
		user = append(user[:indexChoice], user[indexChoice+1:]...)
		w.Write([]byte("Deleted It"))
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
	}
	json.NewEncoder(w).Encode(&user) // return user list
}

func main() {

	user = append(user, Person{Name: "Alex", Age: "28"})

	http.HandleFunc("/api/v1/getName", handleGetPerson)
	http.HandleFunc("/api/v1/updateName", UpdatePerson)
	http.HandleFunc("/api/v1/addName", handleAddPerson)
	http.HandleFunc("/api/v1/deleteName", handleDeletePerson)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}
