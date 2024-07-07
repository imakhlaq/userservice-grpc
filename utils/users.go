package utils

import (
	"github.com/imakhlaq/userservice-grpc/models"
)

var ListOfUsers = []models.User{
	{ID: 1, Fname: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
	{ID: 2, Fname: "Bob", City: "DL", Phone: 224567890, Height: 6.1, Married: false},
	{ID: 3, Fname: "Sam", City: "NY", Phone: 888456780, Height: 5.2, Married: true},
	{ID: 4, Fname: "Eric", City: "NY", Phone: 773456890, Height: 5.6, Married: true},
	{ID: 3, Fname: "Danial", City: "LA", Phone: 554567890, Height: 5.5, Married: false},
	{ID: 4, Fname: "Ron", City: "LA", Phone: 1234997890, Height: 5.10, Married: true},
}
