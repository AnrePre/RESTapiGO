package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type user struct {
	Name string `json:"name"`
	Age  string `json:"age"`
	City string `json:"city"`
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint(w, "Hello Root Route")
	w.Write([]byte("Hello Root Route")) //response to when the handler is called
	fmt.Println("Hello Root Route")
}

func teachersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Println(r.URL.Path)
		path := strings.TrimPrefix(r.URL.Path, "/teachers/")
		userID := strings.TrimSuffix(path, "/")
		fmt.Println("The Id is: ", userID)

		fmt.Println("Query Params : ", r.URL.Query())
		queryParams := r.URL.Query()
		sortby := queryParams.Get("sortby")
		key := queryParams.Get("key")
		sortOrder := queryParams.Get("sortorder")
		fmt.Printf("Sortby: %v, Sort order: %v, Key: %v", sortby, sortOrder, key)
		if sortOrder == "" {
			sortOrder = "DESC"
		}

		//w.Write([]byte("Hello GET Method on Teachers Route"))
		//fmt.Println("Hello GET Method on Teachers Route")
	case http.MethodPost:
		w.Write([]byte("Hello Post Method on Teachers Route"))
		fmt.Println("Hello Post Method on Teachers Route")
	case http.MethodPut:
		w.Write([]byte("Hello Put Method on Teachers Route"))
		fmt.Println("Hello Put Method on Teachers Route")
	case http.MethodPatch:
		w.Write([]byte("Hello Patch Method on Teachers Route"))
		fmt.Println("Hello Patch Method on Teachers Route")
	case http.MethodDelete:
		w.Write([]byte("Hello Delete Method on Teachers Route"))
		fmt.Println("Hello Delete Method on Teachers Route")
	}
	//w.Write([]byte("Hello Teachers Route")) //response to when the handler is called
	//fmt.Println("Hello Teachers Route")
}

func studentsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello GET Method on Students Route"))
		fmt.Println("Hello GET Method on Students Route")
	case http.MethodPost:
		w.Write([]byte("Hello Post Method on Students Route"))
		fmt.Println("Hello Post Method on Students Route")
	case http.MethodPut:
		w.Write([]byte("Hello Put Method on Students Route"))
		fmt.Println("Hello Put Method on Students Route")
	case http.MethodPatch:
		w.Write([]byte("Hello Patch Method on Students Route"))
		fmt.Println("Hello Patch Method on Students Route")
	case http.MethodDelete:
		w.Write([]byte("Hello Delete Method on Students Route"))
		fmt.Println("Hello Delete Method on Students Route")
	}
	w.Write([]byte("Hello Students Route")) //response to when the handler is called
	fmt.Println("Hello Students Route")
}

func execsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello GET Method on Execs Route"))
		fmt.Println("Hello GET Method on Execs Route")
	case http.MethodPost:
		w.Write([]byte("Hello Post Method on Execs Route"))
		fmt.Println("Hello Post Method on Execs Route")
	case http.MethodPut:
		w.Write([]byte("Hello Put Method on Execs Route"))
		fmt.Println("Hello Put Method on Execs Route")
	case http.MethodPatch:
		w.Write([]byte("Hello Patch Method on Execs Route"))
		fmt.Println("Hello Patch Method on Execs Route")
	case http.MethodDelete:
		w.Write([]byte("Hello Delete Method on Route"))
		fmt.Println("Hello Delete Method on Route")
	}
	w.Write([]byte("Hello Execs Route")) //response to when the handler is called
	fmt.Println("Hello Execs Route")
}

func main() {
	port := ":3000"

	http.HandleFunc("/", rootHandler)

	http.HandleFunc("/teachers/", teachersHandler)

	http.HandleFunc("/students/", studentsHandler)

	http.HandleFunc("/execs/", execsHandler)

	fmt.Println("Server is running on port", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalln("Error starting the server", err)
	}

}
