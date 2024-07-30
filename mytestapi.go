package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"mytestapi/cmd/db"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/go-sql-driver/mysql"
	"github.com/urfave/negroni"

	"mytestapi/cmd/middlewares"
	"mytestapi/cmd/mytestapi"
)

func main() {
	port := flag.String("port", "8080", "port to listen on")
	dbHost := flag.String("dbHost", "localhost", "database host")
	dbPort := flag.String("dbPort", "3306", "database port")
	dbUser := flag.String("dbUser", "", "database user")
	dbPass := flag.String("dbPass", "", "database password")
	flag.Parse()

	db := db.NewDBClient(fmt.Sprintf("%s:%s@tcp(%s:%s)/testdb", *dbUser, *dbPass, *dbHost, *dbPort))
	profileService := mytestapi.NewProfileService(db)
	authService := mytestapi.NewAuthService(db)

	server := mytestapi.Server{
		ProfileService: profileService,
	}
	accessMiddleware := middlewares.NewAuthMiddleware(authService)
	middleware := negroni.New(accessMiddleware)

	mux := http.NewServeMux()
	mux.Handle("/profile", server.GetProfile())
	middleware.UseHandler(mux)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	server.HttpServer = &http.Server{
		Addr:    fmt.Sprintf(":%s", *port),
		Handler: middleware,
	}

	go func() {
		log.Printf("Server is running on port %s\n", *port)
		if err := server.HttpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Could not listen on :%s: %v\n", *port, err)
		}
	}()

	<-quit
	log.Println("shutting down")
	db.Close()
}
