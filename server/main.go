package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/gabriel-henriq/klever-challenge/api/controllers"
	"github.com/gabriel-henriq/klever-challenge/api/db"
	pb "github.com/gabriel-henriq/klever-challenge/api/gen/proto/go/proto"
	"go.mongodb.org/mongo-driver/mongo"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var bookDb *mongo.Collection
var mongoCtx context.Context

func main() {

	bookDb, mongoCtx = db.Connect()

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Could not connect on port :50051: %s", err.Error())
	}

	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer)

	upvoteService := controllers.UpvoteServiceServer{
		Ctx: mongoCtx,
		Db:  bookDb,
	}

	pb.RegisterUpvoteServiceServer(grpcServer, &upvoteService)

	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()
	fmt.Printf("Server started on port :50051\n")

	// Right way to stop the server using a SHUTDOWN HOOK
	// Create a channel to receive OS signals
	ch := make(chan os.Signal, 1)

	// Relay os.Interrupt to our channel (os.Interrupt = CTRL+C)
	// Ignore other incoming signals
	signal.Notify(ch, os.Interrupt)

	// Block main routine until a signal is received
	// As long as user doesn't press CTRL+C a message is not passed and our main routine keeps running
	<-ch

	fmt.Println("\nStopping the server...")

	// Graceful shutdown of the server
	grpcServer.Stop()
	listener.Close()

	// Close the database connection
	bookDb.Database().Client().Disconnect(mongoCtx)

	fmt.Println("Server stopped.")
}
