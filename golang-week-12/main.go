package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var baseURL = "http://localhost:8080"

type student struct {
	ID    string
	Name  string
	Grade int
}

var data = []student{
	student{"E001", "ethan", 21},
	student{"W001", "wick", 22},
	student{"B001", "bourne", 23},
	student{"B002", "bond", 23},
}

func users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		var result, err = json.Marshal(data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func user(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		var id = r.FormValue("id")
		var result []byte
		var err error

		for _, each := range data {
			if each.ID == id {
				result, err = json.Marshal(each)

				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				w.Write(result)
				return
			}
		}

		http.Error(w, "User not found!", http.StatusNotFound)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func fetchUsers() ([]student, error) {
	fmt.Println("here")
	var err error
	var client = &http.Client{}
	var data []student

	request, err := http.NewRequest("GET", baseURL+"/users", nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func main() {
	http.HandleFunc("/users", users)
	http.HandleFunc("/user", user)

	fmt.Println("Starting web server at http://localhost:8080")
	http.ListenAndServe(":8080", nil)

	var fetchs, err = fetchUsers()
	if err != nil {
		fmt.Println("Error!", err.Error())
		return
	}

	for _, each := range fetchs {
		fmt.Printf("ID: %s\t Name: %s\t Grade: %d\n", each.ID, each.Name, each.Grade)
	}
}
