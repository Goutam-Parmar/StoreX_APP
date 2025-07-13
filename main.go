package main

import (
	"StoreXApp/database"
	"StoreXApp/routes"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const shutdownTimeout = 10 * time.Second

func main() {
	err := database.ConnectionAndMigrate()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Database connection succeeded")
	}
	router := routes.InitRoutes()
	srv := &http.Server{
		Addr:    ":8081",
		Handler: router,
	}
	//  goroutine for faster and also it didnt block main thread , now main thread can do other task
	go func() {
		fmt.Println("Server running at http://localhost:8081")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown signal received...")
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}

	if err := database.ShutDownDBN(); err != nil {
		log.Printf("Error closing DB: %v", err)
	}

	log.Println("Server stop gracefully")
}
