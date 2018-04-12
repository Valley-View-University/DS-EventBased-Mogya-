//basic server
package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"net/http"
	"time"
	"log"
)

const port = "8006"

func main() {
	headers:= handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	origins:= handlers.AllowedOrigins([]string{"*"})
	methods:= handlers.AllowedMethods([]string{
		http.MethodGet,
		http.MethodPut,
		http.MethodPost,
		http.MethodDelete,
		http.MethodOptions,
	})

	//starting the API
	fmt.Println("Server API starting on port", port)
	fmt.Println("Application running on")

	//Todo: create a router, and perform routing functions

	router:= NewRouter()

	server:= &http.Server{
		Addr:fmt.Sprintf("127.0.0.1:%v", port),
		Handler:handlers.CORS(origins,headers,methods)(router),
		WriteTimeout:1*time.Second,
		ReadTimeout:1*time.Second,
	}
	log.Fatal(server.ListenAndServe())
}