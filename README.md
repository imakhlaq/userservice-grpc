# gRPC User Service

This project is a Golang gRPC service that manages user details and includes search functionality. The service simulates a database by maintaining a list of user details in memory and provides gRPC endpoints for fetching and searching user details.

## Features

- Fetch user details by user ID
- Retrieve a list of user details by a list of user IDs
- Search for user details based on specific criteria(marital status)

## User Model

The user model is defined as follows:

```json
{
  "id": 1,
  "fname": "Steve",
  "city": "LA",
  "phone": 1234567890,
  "height": 5.8,
  "married": true
}
```

## Setup

### Prerequisites

- Go 1.16 or later
- Protobuf compiler (protoc)
- Docker (optional for containerization)

## Installation

- Clone the repository:

```
git clone https://github.com/your-username/grpc-user-service.git

cd grpc-user-service
```

- Generate the gRPC code from the protobuf file:

```
protoc --go_out=plugins=grpc:. proto/user.proto
```

- Build and run the server:

```
go run server/main.go
```

### Docker

To run the service in a Docker container:

- Build the Docker image:

```
docker build -t grpc-user-service .
```

- Run the container:

```
docker run -p 50051:50051 grpc-user-service
```

Alternatively, you can use docker-compose:

- Build and start the service:

```
docker-compose up --build
```
