package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Serj1c/yaml-to-openmetrics/pkg/handlers"
	"github.com/Serj1c/yaml-to-openmetrics/pkg/util"
	"github.com/gorilla/mux"
	//"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	config, err := util.LoadConfig("./")
	if err != nil {
		log.Fatal("Unable to read configuration: ", err)
	}

	path := "currencies.yaml"

	metrics := handlers.NewMetrics(path)

	sm := mux.NewRouter()

	sm.HandleFunc("/metrics", metrics.Prepare).Methods("GET")

	server := &http.Server{
		Addr:         config.ServerPort,
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		fmt.Printf("Server is listening on port %s\n", server.Addr)
		err := server.ListenAndServe()
		if err != nil {
			fmt.Printf("Error starting server: %s:\n", err)
			os.Exit(1)
		}
	}()

	// gracefully shutdown
	sigChannel := make(chan os.Signal)
	signal.Notify(sigChannel, os.Interrupt)
	signal.Notify(sigChannel, os.Kill)

	// block until a signal is received
	sig := <-sigChannel
	fmt.Println("\nCommand to terminate received, server is being shutdown...", sig)

	// gracefully shutdown the server waiting 10 seconds for current operations to complete
	timeoutContext, finish := context.WithTimeout(context.Background(), 10*time.Second)
	defer finish()
	server.Shutdown(timeoutContext)
}
