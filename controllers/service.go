package controllers

import (
	"context"
	"fmt"
	"log"
	"sync"

	pb "github.com/gabriel-henriq/klever-challenge/api/gen/proto/go/proto"
	"github.com/gabriel-henriq/klever-challenge/api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UpvoteServiceServer struct {
	Db  *mongo.Collection
	Ctx context.Context
	pb.UnimplementedUpvoteServiceServer
}

func (s *UpvoteServiceServer) Upvote(ctx context.Context, req *pb.VoteRequest) (*pb.VoteResponse, error) {

	id, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Invalid id: %v", err))
	}

	cursor := s.Db.FindOneAndUpdate(
		s.Ctx,
		bson.D{{Key: "_id", Value: id}},
		bson.D{{Key: "$inc", Value: bson.D{{Key: "likes", Value: 1}}}})

	if cursor == nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("This title does not exists in database %s", req.GetId()),
		)
	}

	var data models.Book

	cursor = s.Db.FindOne(s.Ctx, bson.D{{Key: "_id", Value: id}})

	err = cursor.Decode(&data)
	if err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("This title does not exists in database %s", req.GetId()),
		)
	}

	return &pb.VoteResponse{
		Id:     data.ID.Hex(),
		Title:  data.Title,
		Author: data.Author,
		Likes:  data.Likes,
	}, nil

}

func (s *UpvoteServiceServer) ListBooks(req *pb.ListBooksRequest, stream pb.UpvoteService_ListBooksServer) error {
	title := string(req.GetTitle())

	var waitGroup sync.WaitGroup

	// Create a pipeline to watch when a upvote is updated or inserted in the database by title
	pipeline := mongo.Pipeline{
		{
			{Key: "$match", Value: bson.D{{Key: "$and", Value: bson.A{
				bson.D{{Key: "operationType", Value: "update"}},
				bson.D{{Key: "fullDocument.title", Value: title}},
			}}}},
		},
	}

	// Setting options for enable full document projection
	opts := options.ChangeStream().SetFullDocument(options.UpdateLookup)

	// Create a change stream to listen to the books collection
	DBstream, err := s.Db.Watch(s.Ctx, pipeline, opts)
	if err != nil {
		panic(err)
	}

	// Waitgroup counter
	waitGroup.Add(1)

	// Stop listening to the database when the stream is closed
	routineCtx, cancelFn := context.WithCancel(context.Background())
	_ = cancelFn

	// Start the goroutine that will listen to the DB stream
	go listenToDBChangeStream(routineCtx, waitGroup, DBstream, s.Db, stream)

	// Wait for the worker to finish
	waitGroup.Wait()
	return nil
}

func listenToDBChangeStream(
	routineCtx context.Context,
	waitGroup sync.WaitGroup,
	stream *mongo.ChangeStream,
	collection *mongo.Collection,
	streamServer pb.UpvoteService_ListBooksServer,
) {
	// Cleanup defer functions when this function exits
	defer stream.Close(routineCtx)

	// Wrap the worker call in a closure that makes sure to tell the WaitGroup that this worker is done
	defer waitGroup.Done()

	// Whenever there is a change in the books collection, decode the change
	for stream.Next(routineCtx) {
		var event bson.M
		var data models.Book

		// Decode the change event
		if err := stream.Decode(&event); err != nil {
			panic(err)
		}

		// Find the mongodb document based on the objectID
		err := collection.FindOne(context.TODO(), event["documentKey"]).Decode(&data)
		if err != nil {
			log.Fatal(err)
		}

		// Send the book to the client
		streamServer.Send(&pb.ListBooksResponse{
			Id:     data.ID.Hex(),
			Title:  data.Title,
			Author: data.Author,
			Likes:  data.Likes,
		})

	}
}
