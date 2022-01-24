package main

import (
	"api/config"
	"api/routes"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	logger := log.New(os.Stderr, "", log.Ldate|log.Ltime)
	cfg, _ := config.LoadConfig()

	// if cfgErr != nil {
	// 	logger.Fatal(cfgErr)
	// }

	srv := &http.Server{
		Addr:        fmt.Sprintf(":%s", cfg.Port),
		Handler:     routes.NewRouter(cfg),
		IdleTimeout: time.Minute,
		ReadTimeout: 10 * time.Second, WriteTimeout: 30 * time.Second,
	}
	logger.Printf("starting server on %s", srv.Addr)
	err := srv.ListenAndServe()
	logger.Fatal(err)

}
