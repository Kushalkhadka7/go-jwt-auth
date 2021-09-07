package main

import (
	"jwt-auth/app"
	"jwt-auth/db"
	"jwt-auth/server"
	"log"
	"os"
)

func main() {
	// Generates database connection.
	var config map[string]string = map[string]string{
		"Port": "5432",
	}

	dbConn, err := db.New(config)
	if err != nil {
		log.Fatalf("Cannot establish database connection: %s", err)
		os.Exit(1)
	}

	server := server.New(6000, "localhost")

	listenter, err := server.GetHttpListener()
	if err != nil {
		log.Fatalf("Cannot initialize http server: %s", err)
		os.Exit(1)
	}

	grpcServer, err := server.GetGrpcServerInstance()
	if err != nil {
		log.Fatalf("Cannot initialize http server: %s", err)
		os.Exit(1)
	}

	appServer := app.NewAppServer(dbConn)
	appServer.RegisterServers(grpcServer)

	if err := grpcServer.Serve(listenter); err != nil {
		log.Fatalf("Cannot initialize http server: %s", err)
		os.Exit(1)
	}
}
