package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/thissidemayur/1-golang-studentsApi/internal/config"
	"github.com/thissidemayur/1-golang-studentsApi/internal/http/handlers/student"
	"github.com/thissidemayur/1-golang-studentsApi/internal/storage/sqlite"
)

func main(){
	// 1. Load config
	cfg:= config.MustLoad()

	// 2 set custom/built in logger (if applicable)
	
	// 3. Database connection
	storage,err:=sqlite.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	slog.Info("storage initalized",slog.String("env",cfg.Env), slog.String("version","1.0.0"))
	

	// 4. Setup router
	router := http.NewServeMux()

	router.HandleFunc("POST /api/v1/students",student.New(storage))
	router.HandleFunc("GET /api/v1/students/{id}",student.GetStudentById(storage))
	// 5. setup server
	server :=http.Server{
		Addr: cfg.Addr,
		Handler:router,
	}

	// 6. start server and gracefull shutdown 
	slog.Info("Starting server at ", slog.String("address",cfg.Addr))

	done:=make(chan os.Signal ,1)

	// add signal in done channel
	signal.Notify(done,os.Interrupt,syscall.SIGINT,syscall.SIGTERM) // listen to these signals

	go func() {
		if err:=server.ListenAndServe(); err != nil{
		log.Fatalf("failed to start server: %v \n", err)
		}
	}()

	<-done

	slog.Info("Shutting down server...")

	// add gracefull shutdown logic
	ctx,cancel := context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()

	if err:=server.Shutdown(ctx); err != nil {
		slog.Error("failed to shutdown server",slog.String("error",err.Error()))

	}
	slog.Info("Server shutdown properly with gracefull manner!")
}
/*
================ Explanation ==================
1. Load configuration using the MustLoad function from the config package.
2. (Placeholder) Set up custom or built-in logger if needed.
3. (Placeholder) Establish a database connection if applicable.
4. Set up an HTTP router using http.NewServeMux and define a simple route that responds with "Welcome to student API".
5. Configure the HTTP server with the address from the loaded configuration and the router as the handler.
6. Start the server and log any errors that occur during startup.
7. Implement graceful shutdown by listening for OS interrupt signals (like SIGINT and SIGTERM) and blocking until such a signal is received.

Note:
this is a basic skeleton for a HTTP server in Go.
2. SIGINT is the signal sent when you press Ctrl+C in the terminal to interrupt a running process.
3. SIGTERM is a termination signal that can be sent to a process to request its termination. It's commonly used for graceful shutdowns.
4. OS.Interrupt is a generic signal that represents an interrupt from the operating system, which typically maps to SIGINT on Unix-like systems.

You can run this code as a starting point for a simple HTTP server in Go,
but you can expand upon it by adding more routes, middleware, error handling,gracefull shutdown, and other functionalities as needed.

*/


