Music Microservices Application (AP2-Final)

A scalable microservices-based music platform developed for the Advanced Programming 2 course at Astana IT University.

The project demonstrates modern backend engineering practices using Go, gRPC, Protocol Buffers, and PostgreSQL.
Instead of using a monolithic architecture, the system is separated into independent microservices communicating through high-performance gRPC (HTTP/2) calls.

Project Goals

This project was designed to demonstrate:

Microservices architecture principles
Service-to-service communication with gRPC
API Gateway pattern implementation
Clean backend architecture using Repository Pattern
REST ↔ gRPC request translation
PostgreSQL integration
Scalable distributed backend design
System Architecture

The application consists of four independent backend services and one shared Proto module.

                +----------------------+
                |      Frontend        |
                |   Postman / Client   |
                +----------+-----------+
                           |
                           | HTTP/JSON
                           v
                +----------------------+
                |     API Gateway      |
                |       :8080          |
                |   Gin + REST API     |
                +----------+-----------+
         +-----------------+-----------------+
         |                 |                 |
         v                 v                 v
 +--------------+ +--------------+ +--------------+
 | User Service | | CatalogServ. | | PlaylistServ.|
 |    :50051    | |    :50052    | |    :50053    |
 +------+-------+ +------+-------+ +------+-------+
        |                |                |
        v                v                v
   PostgreSQL       PostgreSQL       PostgreSQL
Services Overview
1. API Gateway (:8080)

The API Gateway acts as the single public entry point for all client requests.

Responsibilities
Accept REST HTTP requests
Parse JSON payloads
Route requests to proper gRPC services
Convert REST ↔ Protobuf messages
Return JSON responses to clients
Technologies
Gin-Gonic
gRPC Client
REST API
2. User Service (:50051)

Handles all user-related operations.

Features
User registration
User authentication
Password handling
Database interaction
Database
PostgreSQL
Architecture
Repository Pattern
gRPC Server
3. Catalog Service (:50052)

Responsible for music catalog management.

Features
Add tracks
Retrieve track metadata
Search tracks
Manage catalog storage
Example Track Metadata
Title
Artist
Duration
4. Playlist Service (:50053)

Manages playlists and track associations.

Features
Create playlists
Attach tracks to playlists
User-playlist relationships
Playlist retrieval
5. Proto Module (/proto)

Contains all shared .proto contract files.

Responsibilities
Define gRPC services
Define Protobuf message structures
Generate Go code for:
gRPC clients
gRPC servers
Tech Stack
Category	Technology
Language	Go (Golang)
API Framework	Gin-Gonic
Communication	gRPC
Serialization	Protocol Buffers v3
Database	PostgreSQL
Architecture	Microservices
Patterns	API Gateway, Repository Pattern
Protocol	HTTP/2
Suggested Project Structure
ap2-final/
│
├── api-gateway/
│   ├── cmd/
│   ├── internal/
│   └── main.go
│
├── user-service/
│   ├── cmd/
│   ├── internal/
│   └── main.go
│
├── catalog-service/
│   ├── cmd/
│   ├── internal/
│   └── main.go
│
├── playlist-service/
│   ├── cmd/
│   ├── internal/
│   └── main.go
│
├── proto/
│   ├── user.proto
│   ├── catalog.proto
│   ├── playlist.proto
│   └── generated/
│
├── docker-compose.yml
├── go.work
├── go.mod
└── README.md
Running the Project Locally

Before running the project, ensure that:

Go is installed
PostgreSQL is running
Ports 8080, 50051, 50052, 50053 are available
1. Start User Service

Open Terminal 1:

cd user-service
go run cmd/main.go
2. Start Catalog Service

Open Terminal 2:

cd catalog-service
go run cmd/main.go
3. Start Playlist Service

Open Terminal 3:

cd playlist-service
go run cmd/main.go
4. Start API Gateway

Open Terminal 4:

cd api-gateway
go run cmd/main.go
API Specification

All client requests must go through the API Gateway.

Base URL:

http://localhost:8080
Register a New User
Endpoint
POST /register
Request Body
{
  "username": "bakdaulet",
  "email": "test@aitu.kz",
  "password": "superpassword123"
}
Example Response
{
  "message": "user registered successfully"
}
Add Track to Catalog
Endpoint
POST /tracks
Request Body
{
  "title": "New Song",
  "artist": "Famous Artist",
  "duration": 180
}
Example Response
{
  "id": "track_1",
  "message": "track created successfully"
}
Get Track by ID
Endpoint
GET /tracks/:id
Example Request
GET /tracks/track_1
Example Response
{
  "id": "track_1",
  "title": "New Song",
  "artist": "Famous Artist",
  "duration": 180
}
Create Playlist
Endpoint
POST /playlists
Request Body
{
  "user_id": "1",
  "title": "My Favorite Tracks",
  "track_ids": [
    "track_1",
    "track_2",
    "track_3"
  ]
}
Example Response
{
  "playlist_id": "playlist_1",
  "message": "playlist created successfully"
}
Testing the Application

You can test the system using:

Postman
cURL
Insomnia
Frontend client applications
Example cURL Request
curl -X POST http://localhost:8080/register \
-H "Content-Type: application/json" \
-d '{
  "username":"bakdaulet",
  "email":"test@aitu.kz",
  "password":"superpassword123"
}'
Why gRPC?

This project uses gRPC because it provides:

High performance via HTTP/2
Compact binary serialization
Strongly typed contracts
Faster inter-service communication
Automatic code generation
Better scalability for distributed systems
Design Patterns Used
API Gateway Pattern

Provides a single entry point for all external requests.

Repository Pattern

Separates business logic from database access logic.

Microservices Architecture

Allows each service to scale and evolve independently.

Future Improvements

Possible future enhancements:

JWT Authentication
Docker containerization
Kubernetes deployment
Service Discovery
API Rate Limiting
Redis caching
Message Broker (Kafka/RabbitMQ)
CI/CD pipelines
Unit & Integration Testing
Monitoring with Prometheus + Grafana
