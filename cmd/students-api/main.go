package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/thissidemayur/1-golang-studentsApi/internal/config"
)

func main(){
	// 1. Load config
	cfg:= config.MustLoad()

	// 2 set custom/built in logger (if applicable)
	// 3. Database connection

	// 4. Setup router
	router := http.NewServeMux()

	router.HandleFunc("GET /",func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Welcome to student API"))
	})
	// 5. setup server
	server :=http.Server{
		Addr: cfg.Addr,
		Handler:router,
	}

	// 6. start server
	fmt.Printf("Starting server at %s\n",cfg.Addr)
	if err:=server.ListenAndServe(); err != nil{
		log.Fatalf("failed to start server: %v \n", err)
	}
}
/*
================ Explanation ==================
1. Load configuration using the MustLoad function from the config package.
2. (Placeholder) Set up custom or built-in logger if needed.
3. (Placeholder) Establish a database connection if applicable.
4. Set up an HTTP router using http.NewServeMux and define a simple route that responds with "Welcome to student API".
5. Configure the HTTP server with the address from the loaded configuration and the router as the handler.
6. Start the server and log any errors that occur during startup.
this is a basic skeleton for a HTTP server in Go.


but you can expand upon it by adding more routes, middleware, error handling,gracefull shutdown, and other functionalities as needed.

*/


/*
What is conccurency
what is channel 
what is synchronization in golang
What is goroutine
difference between buffered and unbuffered channel
*/