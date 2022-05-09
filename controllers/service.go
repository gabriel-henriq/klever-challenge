package controllers

import (
	"context"
	"fmt"

	"github.com/gabriel-henriq/klever-challenge/api/models"
	"github.com/gabriel-henriq/klever-challenge/api/pb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
	cursor, err := s.Db.Find(s.Ctx, bson.M{})
	if err != nil {
		return err
	}
	books := []models.Book{}
	err = cursor.All(s.Ctx, &books)
	if err != nil {
		return err
	}

	for _, book := range books {
		err := stream.Send(&pb.ListBooksResponse{
			Id:     book.ID.Hex(),
			Author: book.Author,
			Title:  book.Title,
			Likes:  book.Likes,
		})
		if err != nil {
			return err
		}
	}

	return nil
}
