package handlers

import (
	"context"
	"errors"

	pb "github.com/imakhlaq/userservice-grpc/proto"
	"github.com/imakhlaq/userservice-grpc/utils"
)

type Server struct {
	pb.UserServiceServer
}

// takes and id and returns a user
func (s *Server) GetUser(ctx context.Context, req *pb.UserIDRequest) (*pb.UserResponse, error) {
	// Implementation for fetching a user by ID
	for _, user := range utils.ListOfUsers {
		if user.ID == req.Id {
			return &pb.UserResponse{
				User: &pb.User{
					Id:      user.ID,
					Fname:   user.Fname,
					City:    user.City,
					Phone:   user.Phone,
					Height:  user.Height,
					Married: user.Married,
				},
			}, nil
		}
	}

	return nil, errors.New("user not found")
}

// takes a list of ids and returns a List of users
func (s *Server) GetUsers(ctx context.Context, req *pb.UserIDsRequest) (*pb.UsersResponse, error) {
	// Implementation for fetching users by a list of IDs

	var users []*pb.User
	for _, id := range req.Ids {
		for _, user := range utils.ListOfUsers {
			if user.ID == id {
				users = append(users, &pb.User{
					Id:      user.ID,
					Fname:   user.Fname,
					City:    user.City,
					Phone:   user.Phone,
					Height:  user.Height,
					Married: user.Married,
				})
				break
			}
		}
	}

	return &pb.UsersResponse{Users: users}, nil
}

// Returns all the married users
func (s *Server) SearchUsers(ctx context.Context, req *pb.SearchRequest) (*pb.UsersResponse, error) {
	// Implementation for searching users based on criteria
	var users []*pb.User
	for _, user := range utils.ListOfUsers {
		if user.Married == req.Married {
			users = append(users, &pb.User{
				Id:      user.ID,
				Fname:   user.Fname,
				City:    user.City,
				Phone:   user.Phone,
				Height:  user.Height,
				Married: user.Married,
			})
		}
	}

	return &pb.UsersResponse{Users: users}, nil
}
