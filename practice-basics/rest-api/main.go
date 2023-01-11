package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// if r.URL.Path != "/hello" {
	// 	http.Error(w, "404 not found.", http.StatusNotFound)
	// 	return
	// }

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello!")
}

func checkUser(w http.ResponseWriter, r *http.Request) {
	var user User
	dbPassword := "texas"
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatal("Err decoding in struct")
	}

	if user.Password == dbPassword {
		fmt.Println("Success logged in")
	}
	fmt.Fprintf(w, "Response :%v", user)
}

func getJson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader((http.StatusOK))
	switch r.Method {
	case "GET":
		w.Write([]byte(`"message:Get is called"`))
	case "POST":
		w.Write([]byte(`"message:POST is called"`))
	}
}

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Home Page")
	})
	router.HandleFunc("/hello", helloHandler) // Update this line of code
	router.HandleFunc("/json", getJson)
	router.HandleFunc("/user", checkUser)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
