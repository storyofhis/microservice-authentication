package main

import (
	"log"
	"net"

	"github.com/storyofhis/auth-service/repositories/database"
	app "github.com/storyofhis/auth-service/repositories/proto"
	gogrpc "google.golang.org/grpc"
)

func main() {
	database.Load()

	listen, err := net.Listen("tcp", database.Get().Address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err.Error())
	}

	// Creating a New Server
	server := gogrpc.NewServer()

	// Auth registration server
	app.RegisterAuthServer(server, nil)

	// Message of success
	log.Println(database.Get().Name, "is running on port", database.Get().Address)

	// Initializing server
	if err = server.Serve(listen); err != nil {
		log.Fatalf("Failed to server: %v", err.Error())
	}
}
